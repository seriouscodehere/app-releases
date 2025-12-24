package onboarding_routes

import (
	"net/http"
	onboarding_controller "sraraa/reciever_src/controllers/auth/signup/onboarding_controllers"
)

func RegisterOnboardingRoutes() {
	http.HandleFunc("/api/onboarding/username", onboarding_controller.SetUsernameHandler)
	http.HandleFunc("/api/onboarding/fullname", onboarding_controller.SetFullnameHandler)
	http.HandleFunc("/api/onboarding/password", onboarding_controller.SetPasswordHandler)
}
