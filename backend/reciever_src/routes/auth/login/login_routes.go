package login_routes

import (
	"net/http"
	login_controller "sraraa/reciever_src/controllers/auth/login"
)

func LoginRoutes() {
	// Step 1: Request OTP (validates email + password, sends OTP)
	http.HandleFunc("/api/auth/login/request-otp", login_controller.RequestLoginOTPHandler)

	// Step 2: Verify OTP and get session token
	http.HandleFunc("/api/auth/login/verify-otp", login_controller.VerifyLoginOTPHandler)

	// Session management
	http.HandleFunc("/api/auth/logout", login_controller.LogoutHandler)
	http.HandleFunc("/api/auth/logout_all", login_controller.LogoutAllHandler)
	http.HandleFunc("/api/auth/validate_session", login_controller.ValidateSessionHandler)
}
