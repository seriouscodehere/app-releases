package verify_session_controller

import (
	"encoding/json"
	"net/http"
	"sraraa/db"
	user_models "sraraa/reciever_src/models/user"
)

// VerifySessionHandler checks if the session token from frontend is valid
func VerifySessionHandler(w http.ResponseWriter, r *http.Request) {
	// token from frontend IndexedDB should be sent in Authorization header
	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "No session token provided", http.StatusBadRequest)
		return
	}

	// validate JWT token
	claims, err := user_models.ValidateSessionToken(token)
	if err != nil {
		http.Error(w, "Invalid or expired session token", http.StatusUnauthorized)
		return
	}

	// check DB to make sure session exists
	DB := db.DB
	var exists bool
	err = DB.QueryRow(`SELECT EXISTS(SELECT 1 FROM sessions WHERE session_token=?)`, token).Scan(&exists)
	if err != nil || !exists {
		http.Error(w, "Session not found", http.StatusUnauthorized)
		return
	}

	// success: return user info including UID
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"user_id":  claims.UserID,
		"email":    claims.Email,
		"username": claims.Username,
		"fullname": claims.Fullname,
		"verified": claims.Verified,
		"uid":      claims.UID, // ✅ Added UID to response
	})
}
