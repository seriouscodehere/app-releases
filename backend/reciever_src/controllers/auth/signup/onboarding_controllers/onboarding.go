package onboarding_controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"sraraa/db"
	"sraraa/reciever_src/controllers/auth/uniqueid"
	user_models "sraraa/reciever_src/models/user"
	auth_utils "sraraa/reciever_src/utils/auth"
)

// --- Username ---
func SetUsernameHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Limit body size
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)

	type request struct {
		Email    string `json:"email"`
		Username string `json:"username"`
	}

	var body request
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&body); err != nil || body.Email == "" || body.Username == "" {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if !isValidEmail(body.Email) {
		http.Error(w, "Invalid email format", http.StatusBadRequest)
		return
	}

	DB := db.DB
	if DB == nil {
		log.Println("database not initialized")
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	verified, err := user_models.IsVerified(DB, body.Email)
	if err != nil || !verified {
		http.Error(w, "Email not verified", http.StatusBadRequest)
		return
	}

	if err := user_models.SetUsername(DB, body.Email, body.Username); err != nil {
		log.Println("SetUsername error:", err)
		http.Error(w, fmt.Sprintf("Failed to set username: %s", err.Error()), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Username set successfully")
}

// --- Fullname ---
func SetFullnameHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)

	type request struct {
		Email    string `json:"email"`
		Fullname string `json:"fullname"`
	}

	var body request
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&body); err != nil || body.Email == "" || body.Fullname == "" {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if !isValidEmail(body.Email) {
		http.Error(w, "Invalid email format", http.StatusBadRequest)
		return
	}

	DB := db.DB
	if DB == nil {
		log.Println("database not initialized")
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	verified, err := user_models.IsVerified(DB, body.Email)
	if err != nil || !verified {
		http.Error(w, "Email not verified", http.StatusBadRequest)
		return
	}

	if err := user_models.SetFullname(DB, body.Email, body.Fullname); err != nil {
		log.Println("SetFullname error:", err)
		http.Error(w, fmt.Sprintf("Failed to set fullname: %s", err.Error()), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Fullname set successfully")
}

// Password
func SetPasswordHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)

	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var body request
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&body); err != nil || body.Email == "" || body.Password == "" {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if !isValidEmail(body.Email) {
		http.Error(w, "Invalid email format", http.StatusBadRequest)
		return
	}

	DB := db.DB
	if DB == nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	verified, err := user_models.IsVerified(DB, body.Email)
	if err != nil || !verified {
		http.Error(w, "Email not verified", http.StatusBadRequest)
		return
	}

	hasUID, err := user_models.HasUID(DB, body.Email)
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	if err := user_models.SetPassword(DB, body.Email, body.Password); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !hasUID {
		for {
			uid, err := uniqueid.Generate()
			if err != nil {
				http.Error(w, "Server error", http.StatusInternalServerError)
				return
			}

			if err := auth_utils.ValidateUniqueID(uid); err != nil {
				continue
			}

			exists, _ := user_models.UniqueIDExists(DB, uid)
			if !exists {
				if err := user_models.SetUniqueID(DB, body.Email, uid); err != nil {
					http.Error(w, "UID creation failed", http.StatusInternalServerError)
					return
				}
				break
			}
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Password set and UID created"))
}

// local email validation used across handlers
func isValidEmail(e string) bool {
	var re = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(e)
}
