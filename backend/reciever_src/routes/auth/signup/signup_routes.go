package signup_routes

import (
	"net/http"
	signup_controller "sraraa/reciever_src/controllers/auth/signup"
)

func RegisterSignupRoutes() {
	http.HandleFunc("/api/signup/send-otp", signup_controller.SendOTPHandler)
	http.HandleFunc("/api/signup/verify-otp", signup_controller.VerifyOTPHandler)
}
