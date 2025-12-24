package forgot_password_controller

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/smtp"
	"os"
	"time"

	"sraraa/db"
	user_models "sraraa/reciever_src/models/user"
)

type requestPayload struct {
	Email string `json:"email"`
}

type verifyPayload struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

func SendResetOTP(w http.ResponseWriter, r *http.Request) {
	var payload requestPayload
	json.NewDecoder(r.Body).Decode(&payload)

	if payload.Email == "" {
		http.Error(w, "email required", http.StatusBadRequest)
		return
	}

	exists, err := user_models.EmailExists(db.DB, payload.Email)
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}
	if !exists {
		http.Error(w, "user not found", http.StatusBadRequest)
		return
	}

	cooldown, _ := user_models.GetPasswordResetCooldown(db.DB, payload.Email)
	if time.Now().Before(cooldown) {
		http.Error(w, "cooldown active", http.StatusTooManyRequests)
		return
	}

	count, _ := user_models.CountPasswordResetRequestsLastHour(db.DB, payload.Email)
	if count >= 5 {
		user_models.SetPasswordResetCooldown(db.DB, payload.Email, time.Now().Add(30*time.Minute))
		http.Error(w, "too many requests", http.StatusTooManyRequests)
		return
	}

	code := rand.Intn(900000) + 100000

	err = SendOTPEmail(payload.Email, fmt.Sprint(code))
	if err != nil {
		http.Error(w, "email service unavailable", http.StatusServiceUnavailable)
		return
	}

	user_models.SavePasswordResetOTP(db.DB, payload.Email, fmt.Sprint(code))
	user_models.AddPasswordResetRequest(db.DB, payload.Email)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"otp sent"}`))
}

func VerifyResetOTP(w http.ResponseWriter, r *http.Request) {
	var payload verifyPayload
	json.NewDecoder(r.Body).Decode(&payload)

	code, created, err := user_models.GetPasswordResetOTP(db.DB, payload.Email)
	if err != nil {
		http.Error(w, "otp not found", http.StatusBadRequest)
		return
	}

	if time.Since(created) > 10*time.Minute {
		user_models.DeletePasswordResetOTP(db.DB, payload.Email)
		http.Error(w, "otp expired", http.StatusBadRequest)
		return
	}

	if payload.Code != code {
		http.Error(w, "invalid otp", http.StatusBadRequest)
		return
	}

	user_models.DeletePasswordResetOTP(db.DB, payload.Email)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"otp verified"}`))
}

func SendOTPEmail(to, code string) error {
	host := os.Getenv("SMTP_HOST")
	port := os.Getenv("SMTP_PORT")
	from := os.Getenv("FROM_EMAIL")
	password := os.Getenv("SMTP_PASSWORD")

	auth := smtp.PlainAuth("", from, password, host)

	msg := []byte(
		"To: " + to + "\r\n" +
			"Subject: Password Reset Code\r\n\r\n" +
			"Your password reset code is: " + code + "\nThis code expires in 10 minutes.",
	)

	return smtp.SendMail(host+":"+port, auth, from, []string{to}, msg)
}
