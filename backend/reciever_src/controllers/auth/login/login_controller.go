package login_controller

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"net/smtp"
	"os"
	"time"

	"sraraa/db"
	user_models "sraraa/reciever_src/models/user"
)

// sendOTPEmail sends an OTP code via email using SMTP
func sendOTPEmail(to, code string) error {
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpPassword := os.Getenv("SMTP_PASSWORD")
	fromEmail := os.Getenv("FROM_EMAIL")

	if smtpHost == "" || smtpPort == "" || smtpPassword == "" || fromEmail == "" {
		return fmt.Errorf("SMTP configuration missing in environment variables")
	}

	auth := smtp.PlainAuth("", fromEmail, smtpPassword, smtpHost)

	subject := "Your Login OTP Code"
	body := fmt.Sprintf(`
Hello,

Your login verification code is: %s

This code will expire in 10 minutes.

If you did not request this code, please ignore this email.

Best regards,
Your App Team
`, code)

	message := fmt.Sprintf("From: %s\r\n"+
		"To: %s\r\n"+
		"Subject: %s\r\n"+
		"\r\n"+
		"%s", fromEmail, to, subject, body)

	addr := smtpHost + ":" + smtpPort
	err := smtp.SendMail(addr, auth, fromEmail, []string{to}, []byte(message))
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}

// RequestLoginOTPHandler - Step 1: Send OTP to user's email
func RequestLoginOTPHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var body request
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&body); err != nil || body.Email == "" || body.Password == "" {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	DB := db.DB
	if DB == nil {
		log.Println("database not initialized")
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	// Check if user exists
	var userID int
	err := DB.QueryRow(`SELECT id FROM users WHERE email=?`, body.Email).Scan(&userID)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Verify password
	storedPassword, err := user_models.GetStoredPasswordByEmail(DB, body.Email)
	if err != nil || storedPassword != body.Password {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Check if user is verified
	verified, err := user_models.IsVerified(DB, body.Email)
	if err != nil || !verified {
		http.Error(w, "Email not verified", http.StatusForbidden)
		return
	}

	// Check cooldown
	cooldownUntil, err := user_models.GetLoginCooldown(DB, body.Email)
	if err == nil && time.Now().Before(cooldownUntil) {
		remaining := int(time.Until(cooldownUntil).Minutes())
		http.Error(w, fmt.Sprintf("Too many requests. Try again in %d minutes", remaining), http.StatusTooManyRequests)
		return
	}

	// Check rate limiting (max 5 requests per hour)
	count, _ := user_models.CountLoginRequestsLastHour(DB, body.Email)
	if count >= 5 {
		cooldownUntil := time.Now().Add(1 * time.Hour)
		_ = user_models.SetLoginCooldown(DB, body.Email, cooldownUntil)
		http.Error(w, "Too many OTP requests. Try again later", http.StatusTooManyRequests)
		return
	}

	// Generate 6-digit OTP
	code, err := generateOTP(6)
	if err != nil {
		log.Println("OTP generation error:", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	// Save OTP
	if err := user_models.SaveLoginOTP(DB, body.Email, code); err != nil {
		log.Println("SaveLoginOTP error:", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	// Track request
	_ = user_models.AddLoginOTPRequest(DB, body.Email)

	// Send OTP via email
	if err := sendOTPEmail(body.Email, code); err != nil {
		log.Printf("Failed to send OTP email to %s: %v", body.Email, err)
		http.Error(w, "Failed to send OTP email", http.StatusInternalServerError)
		return
	}

	log.Printf("Login OTP sent to %s", body.Email)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "OTP sent to your email",
	})
}

// VerifyLoginOTPHandler - Step 2: Verify OTP and create session
func VerifyLoginOTPHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	type request struct {
		Email string `json:"email"`
		OTP   string `json:"otp"`
	}
	var body request
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&body); err != nil || body.Email == "" || body.OTP == "" {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	DB := db.DB
	if DB == nil {
		log.Println("database not initialized")
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	// Get stored OTP
	storedOTP, createdAt, err := user_models.GetLoginOTP(DB, body.Email)
	if err != nil {
		http.Error(w, "Invalid or expired OTP", http.StatusUnauthorized)
		return
	}

	// Check if OTP is expired (valid for 10 minutes)
	if time.Since(createdAt) > 10*time.Minute {
		_ = user_models.DeleteLoginOTP(DB, body.Email)
		http.Error(w, "OTP has expired", http.StatusUnauthorized)
		return
	}

	// Verify OTP
	if storedOTP != body.OTP {
		http.Error(w, "Invalid OTP", http.StatusUnauthorized)
		return
	}

	// Delete OTP after successful verification
	_ = user_models.DeleteLoginOTP(DB, body.Email)

	// Get user ID
	userID, err := user_models.GetUserIDByEmailOrUsername(DB, body.Email)
	if err != nil {
		log.Println("GetUserIDByEmailOrUsername error:", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	// Create session
	userAgent := r.UserAgent()
	ip := r.RemoteAddr
	token, err := user_models.CreateSession(DB, userID, 7*24*time.Hour, userAgent, ip)
	if err != nil {
		log.Println("CreateSession error:", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	// Return session token
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"session_token": token,
		"message":       "Login successful",
	})
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "No session token provided", http.StatusBadRequest)
		return
	}
	DB := db.DB
	_ = user_models.DeleteSession(DB, token)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Logged out successfully")
}

func LogoutAllHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	DB := db.DB
	claims, err := user_models.ValidateSessionToken(token)
	if err != nil {
		http.Error(w, "Invalid session", http.StatusUnauthorized)
		return
	}
	_ = user_models.DeleteAllSessions(DB, claims.UserID)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Logged out from all sessions")
}

func ValidateSessionHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "No session token provided", http.StatusBadRequest)
		return
	}

	claims, err := user_models.ValidateSessionToken(token)
	if err != nil {
		http.Error(w, "Invalid or expired session token", http.StatusUnauthorized)
		return
	}

	DB := db.DB
	var exists bool
	err = DB.QueryRow(`SELECT EXISTS(SELECT 1 FROM sessions WHERE session_token=?)`, token).Scan(&exists)
	if err != nil || !exists {
		http.Error(w, "Session token not found", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"user_id":  claims.UserID,
		"email":    claims.Email,
		"username": claims.Username,
		"fullname": claims.Fullname,
		"verified": claims.Verified,
		"uid":      claims.UID,
	})
}

// Helper function to generate OTP
func generateOTP(length int) (string, error) {
	const digits = "0123456789"
	otp := make([]byte, length)
	for i := range otp {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
		if err != nil {
			return "", err
		}
		otp[i] = digits[num.Int64()]
	}
	return string(otp), nil
}
