package access_auth_routes

import (
	"database/sql"
	"net/http"
	access_auth_controller "sraraa/reciever_src/controllers/auth/access_auth"
)

func RegisterAccessAuthRoutes(db *sql.DB) {
	http.HandleFunc("/api/user/info", access_auth_controller.CheckUserCompleteHandler(db))
	http.HandleFunc("/api/user/check-username", access_auth_controller.CheckUsernameHandler(db))
}
