package signup_controller

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"net/smtp"
	"os"
	"regexp"
	"time"

	"sraraa/db"
	user_models "sraraa/reciever_src/models/user"
)

// Use the shared DB instance that main initializes (db.DB). Do not init DB at package load.
// SendOTPHandler sends OTP to email with cooldowns and limits
func SendOTPHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Limit body size to avoid abuse
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20) // 1 MB

	type requestBody struct {
		Email string `json:"email"`
	}
	var body requestBody
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&body); err != nil || body.Email == "" {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// basic email validation
	if !isValidEmail(body.Email) {
		http.Error(w, "Invalid email format", http.StatusBadRequest)
		return
	}

	DB := db.DB
	if DB == nil {
		log.Println("database not initialized")
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	// Create user if not exists
	if err := user_models.CreateUser(DB, body.Email); err != nil {
		log.Println("CreateUser error:", err)
		// continue, but return server error
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	// Check if already verified
	verified, err := user_models.IsVerified(DB, body.Email)
	if err != nil && !errors.Is(err, sqlErrNoRows()) {
		log.Println("IsVerified error:", err)
		http.Error(w, "Failed to check verification status", http.StatusInternalServerError)
		return
	}
	if verified {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"Email already verified"}`))
		return
	}

	// Check cooldown
	cooldown, err := user_models.GetCooldown(DB, body.Email)
	if err == nil {
		if time.Now().Before(cooldown) {
			http.Error(w, fmt.Sprintf("Email is on cooldown until %s", cooldown.Format(time.RFC3339)), http.StatusTooManyRequests)
			return
		}
	}

	// Check last request time for 1 minute rule
	_, lastCreated, err := user_models.GetOTP(DB, body.Email)
	if err == nil && time.Since(lastCreated) < 1*time.Minute {
		http.Error(w, "You can request a new OTP after 1 minute", http.StatusTooManyRequests)
		return
	}

	// Count requests in last hour
	count, err := user_models.CountRequestsLastHour(DB, body.Email)
	if err == nil && count >= 7 {
		// set 6-hour cooldown
		_ = user_models.SetCooldown(DB, body.Email, time.Now().Add(6*time.Hour))
		http.Error(w, "Too many OTP requests, cooldown 6 hours applied", http.StatusTooManyRequests)
		return
	}

	otp, genErr := generateOTP()
	if genErr != nil {
		log.Println("OTP generation failed:", genErr)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	// Save OTP (handle error)
	if err := user_models.SaveOTP(DB, body.Email, otp); err != nil {
		log.Println("SaveOTP error:", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	// Track request
	if err := user_models.AddOTPRequest(DB, body.Email); err != nil {
		log.Println("AddOTPRequest error:", err)
		// continue; not fatal for user
	}

	// Send email. In development if SMTP env not set, skip sending and log.
	if err := sendEmail(body.Email, otp); err != nil {
		log.Println("Failed to send email:", err)
		// For dev we do not fail hard on email send; return success but log.
		// If you want to enforce sending even in dev, return an error here.
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OTP sent to %s", body.Email)
}

// VerifyOTPHandler verifies OTP and marks user as verified
func VerifyOTPHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Limit body size
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)

	type requestBody struct {
		Email string `json:"email"`
		OTP   string `json:"otp"`
	}
	var body requestBody
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&body); err != nil || body.Email == "" || body.OTP == "" {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if !isValidEmail(body.Email) {
		http.Error(w, "Invalid email format", http.StatusBadRequest)
		return
	}

	DB := db.DB
	if DB == nil {
		log.Println("database not initialized")
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	code, created, err := user_models.GetOTP(DB, body.Email)
	if err != nil {
		log.Println("GetOTP error:", err)
		http.Error(w, "OTP not found", http.StatusNotFound)
		return
	}

	if time.Since(created) > 10*time.Minute {
		_ = user_models.DeleteOTP(DB, body.Email)
		http.Error(w, "OTP expired", http.StatusBadRequest)
		return
	}

	if code != body.OTP {
		http.Error(w, "Invalid OTP", http.StatusUnauthorized)
		return
	}

	// Mark user verified
	if err := user_models.MarkVerified(DB, body.Email); err != nil {
		log.Println("MarkVerified error:", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	// Delete OTP immediately (no background goroutines)
	if err := user_models.DeleteOTP(DB, body.Email); err != nil {
		log.Println("DeleteOTP error:", err)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OTP verified successfully")
}

// generateOTP generates a 6-digit code using crypto/rand
func generateOTP() (string, error) {
	max := big.NewInt(1000000) // 0..999999
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%06d", n.Int64()), nil
}

// sendEmail sends OTP email. In development (no SMTP env) it will be skipped.
func sendEmail(to, otp string) error {
	host := os.Getenv("SMTP_HOST")
	port := os.Getenv("SMTP_PORT")
	from := os.Getenv("FROM_EMAIL")
	password := os.Getenv("SMTP_PASSWORD")

	// If SMTP not configured, assume development and skip sending
	if host == "" || port == "" || from == "" || password == "" {
		log.Printf("SMTP not configured; skipping send to %s. OTP: %s\n", to, otp)
		return nil
	}

	auth := smtp.PlainAuth("", from, password, host)
	msg := []byte(fmt.Sprintf(
		"To: %s\r\n"+
			"Subject: Your OTP Code\r\n"+
			"\r\n"+
			"Your 6-digit OTP code is: %s\r\n", to, otp))

	addr := fmt.Sprintf("%s:%s", host, port)
	// smtp.SendMail will block; keep it but return any error
	return smtp.SendMail(addr, auth, from, []string{to}, msg)
}

// simple email regex (not perfect but good enough for dev)
func isValidEmail(e string) bool {
	var re = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(e)
}

// small shim to compare sqlite no rows error without importing sqlite constants here
func sqlErrNoRows() error {
	return errors.New("sql: no rows in result set")
}
