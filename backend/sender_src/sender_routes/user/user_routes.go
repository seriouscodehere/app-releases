package user_info_sender_routes

import (
	"net/http"
	user_info_sender_controller "sraraa/sender_src/controller/user"
	"strings"
)

func RegisterUserSenderRoutes() {
	http.HandleFunc("/api/user/", func(w http.ResponseWriter, r *http.Request) {
		// Only allow GET requests
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		path := strings.TrimPrefix(r.URL.Path, "/api/user/")
		path = strings.TrimSuffix(path, "/")

		switch path {
		case "email":
			user_info_sender_controller.GetEmailHandler(w, r)
		case "username":
			user_info_sender_controller.GetUsernameHandler(w, r)
		case "fullname":
			user_info_sender_controller.GetFullnameHandler(w, r)
		case "uid":
			user_info_sender_controller.GetUIDHandler(w, r)
		case "verified":
			user_info_sender_controller.GetVerifiedHandler(w, r)
		case "user_id":
			user_info_sender_controller.GetUserIDHandler(w, r)
		case "password":
			user_info_sender_controller.GetPasswordHandler(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}
