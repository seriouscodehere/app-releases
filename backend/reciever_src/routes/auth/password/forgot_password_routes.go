package forgot_password_routes

import (
	"net/http"
	forgot_password_controller "sraraa/reciever_src/controllers/auth/password"
)

func RegisterForgotPasswordRoutes() {
	http.HandleFunc("/auth/forgot-password/send-otp", forgot_password_controller.SendResetOTP)
	http.HandleFunc("/auth/forgot-password/verify-otp", forgot_password_controller.VerifyResetOTP)
}
