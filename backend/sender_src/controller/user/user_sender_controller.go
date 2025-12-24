package user_info_sender_controller

import (
	"encoding/json"
	"errors"
	"net/http"
	user_models "sraraa/reciever_src/models/user"
)

// getUID extracts UID from query parameter
func getUID(r *http.Request) (string, error) {
	uid := r.URL.Query().Get("uid")
	if uid == "" {
		return "", errors.New("missing uid parameter")
	}
	return uid, nil
}

func GetEmailHandler(w http.ResponseWriter, r *http.Request) {
	uid, err := getUID(r)
	if err != nil {
		http.Error(w, "Missing UID", http.StatusBadRequest)
		return
	}

	email, err := user_models.GetUserEmailByUID(uid)
	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"email": email})
}

func GetUsernameHandler(w http.ResponseWriter, r *http.Request) {
	uid, err := getUID(r)
	if err != nil {
		http.Error(w, "Missing UID", http.StatusBadRequest)
		return
	}

	username, err := user_models.GetUsernameByUID(uid)
	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"username": username})
}

func GetFullnameHandler(w http.ResponseWriter, r *http.Request) {
	uid, err := getUID(r)
	if err != nil {
		http.Error(w, "Missing UID", http.StatusBadRequest)
		return
	}

	fullname, err := user_models.GetFullnameByUID(uid)
	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"fullname": fullname})
}

func GetVerifiedHandler(w http.ResponseWriter, r *http.Request) {
	uid, err := getUID(r)
	if err != nil {
		http.Error(w, "Missing UID", http.StatusBadRequest)
		return
	}

	verified, err := user_models.GetUserVerifiedByUID(uid)
	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"verified": verified})
}

func GetUIDHandler(w http.ResponseWriter, r *http.Request) {
	uid, err := getUID(r)
	if err != nil {
		http.Error(w, "Missing UID", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"uid": uid})
}

func GetUserIDHandler(w http.ResponseWriter, r *http.Request) {
	uid, err := getUID(r)
	if err != nil {
		http.Error(w, "Missing UID", http.StatusBadRequest)
		return
	}

	userID, err := user_models.GetUserIDByUID(uid)
	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"user_id": userID})
}

func GetPasswordHandler(w http.ResponseWriter, r *http.Request) {
	uid, err := getUID(r)
	if err != nil {
		http.Error(w, "Missing UID", http.StatusBadRequest)
		return
	}

	password, err := user_models.GetPasswordByUID(uid)
	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"Password": password})
}
