package verify_session_routes

import (
	"net/http"
	verify_session_controller "sraraa/reciever_src/controllers/auth/verify_session"
)

func VerifySessionRoutes() {
	http.HandleFunc("/api/auth/verify_session", verify_session_controller.VerifySessionHandler)
}
