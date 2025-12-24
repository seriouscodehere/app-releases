package user_models

import (
	"database/sql"
	auth_models "sraraa/reciever_src/models/user/auth"
	login_models "sraraa/reciever_src/models/user/login"
	onboard_models "sraraa/reciever_src/models/user/onboard"
	session_models "sraraa/reciever_src/models/user/sessions"
	signup_models "sraraa/reciever_src/models/user/signup"
	user_images_models "sraraa/reciever_src/models/user/user_images"
	user_info_getter_models "sraraa/reciever_src/models/user/user_info_getters"
	"time"
)

// Re-export all models for easy access

// JWT Secret
var JwtSecret = session_models.JwtSecret

// Session Claims type
type SessionClaims = session_models.SessionClaims

// Auth models
func SetPassword(db *sql.DB, email, password string) error {
	return auth_models.SetPassword(db, email, password)
}

func SetUsername(db *sql.DB, email, username string) error {
	return auth_models.SetUsername(db, email, username)
}

func SetFullname(db *sql.DB, email, fullname string) error {
	return auth_models.SetFullname(db, email, fullname)
}

func UsernameExists(db *sql.DB, username string) (bool, error) {
	return auth_models.UsernameExists(db, username)
}

func EmailExists(db *sql.DB, email string) (bool, error) {
	return auth_models.EmailExists(db, email)
}

func HasAllRequiredFields(db *sql.DB, email string) (bool, time.Time, error) {
	return auth_models.HasAllRequiredFields(db, email)
}

func HasAllRequiredFieldsForLogin(db *sql.DB, email string) (bool, error) {
	return auth_models.HasAllRequiredFieldsForLogin(db, email)
}

// Onboarding models
func CreateUser(db *sql.DB, email string) error {
	return onboard_models.CreateUser(db, email)
}

func SetUniqueID(db *sql.DB, email, uid string) error {
	return onboard_models.SetUniqueID(db, email, uid)
}

func HasUID(db *sql.DB, email string) (bool, error) {
	return onboard_models.HasUID(db, email)
}

func UniqueIDExists(db *sql.DB, uid string) (bool, error) {
	return onboard_models.UniqueIDExists(db, uid)
}

// Signup models
func IsVerified(db *sql.DB, email string) (bool, error) {
	return signup_models.IsVerified(db, email)
}

func MarkVerified(db *sql.DB, email string) error {
	return signup_models.MarkVerified(db, email)
}

func SaveOTP(db *sql.DB, email, code string) error {
	return signup_models.SaveOTP(db, email, code)
}

func GetOTP(db *sql.DB, email string) (string, time.Time, error) {
	return signup_models.GetOTP(db, email)
}

func DeleteOTP(db *sql.DB, email string) error {
	return signup_models.DeleteOTP(db, email)
}

func AddOTPRequest(db *sql.DB, email string) error {
	return signup_models.AddOTPRequest(db, email)
}

func CountRequestsLastHour(db *sql.DB, email string) (int, error) {
	return signup_models.CountRequestsLastHour(db, email)
}

func SetCooldown(db *sql.DB, email string, until time.Time) error {
	return signup_models.SetCooldown(db, email, until)
}

func GetCooldown(db *sql.DB, email string) (time.Time, error) {
	return signup_models.GetCooldown(db, email)
}

// Login models
func GetUserIDByEmailOrUsername(db *sql.DB, login string) (int, error) {
	return login_models.GetUserIDByEmailOrUsername(db, login)
}

func GetStoredPasswordByEmail(db *sql.DB, email string) (string, error) {
	return login_models.GetStoredPasswordByEmail(db, email)
}

func SaveLoginOTP(db *sql.DB, email, code string) error {
	return login_models.SaveLoginOTP(db, email, code)
}

func GetLoginOTP(db *sql.DB, email string) (string, time.Time, error) {
	return login_models.GetLoginOTP(db, email)
}

func DeleteLoginOTP(db *sql.DB, email string) error {
	return login_models.DeleteLoginOTP(db, email)
}

func AddLoginOTPRequest(db *sql.DB, email string) error {
	return login_models.AddLoginOTPRequest(db, email)
}

func CountLoginRequestsLastHour(db *sql.DB, email string) (int, error) {
	return login_models.CountLoginRequestsLastHour(db, email)
}

func SetLoginCooldown(db *sql.DB, email string, until time.Time) error {
	return login_models.SetLoginCooldown(db, email, until)
}

func GetLoginCooldown(db *sql.DB, email string) (time.Time, error) {
	return login_models.GetLoginCooldown(db, email)
}

// Session models
func CreateSession(db *sql.DB, userID int, duration time.Duration, userAgent, ip string) (string, error) {
	return session_models.CreateSession(db, userID, duration, userAgent, ip)
}

func DeleteSession(db *sql.DB, token string) error {
	return session_models.DeleteSession(db, token)
}

func DeleteAllSessions(db *sql.DB, userID int) error {
	return session_models.DeleteAllSessions(db, userID)
}

func DeleteAllSessionsByUID(db *sql.DB, uid string) error {
	return session_models.DeleteAllSessionsByUID(db, uid)
}

func GetSessionsByUID(db *sql.DB, uid string) ([]map[string]interface{}, error) {
	return session_models.GetSessionsByUID(db, uid)
}

func CountActiveSessions(db *sql.DB, uid string) (int, error) {
	return session_models.CountActiveSessions(db, uid)
}

func ValidateSessionToken(tokenStr string) (*SessionClaims, error) {
	return session_models.ValidateSessionToken(tokenStr)
}

// Image models
func SaveUserImage(db *sql.DB, uid, username, imageType, imageURL string) error {
	return user_images_models.SaveUserImage(db, uid, username, imageType, imageURL)
}

func GetUserImage(db *sql.DB, uid, username, imageType string) (map[string]interface{}, error) {
	return user_images_models.GetUserImage(db, uid, username, imageType)
}

func GetAllUserImages(db *sql.DB, uid, username string) ([]map[string]interface{}, error) {
	return user_images_models.GetAllUserImages(db, uid, username)
}

func DeleteUserImage(db *sql.DB, uid, imageType string) error {
	return user_images_models.DeleteUserImage(db, uid, imageType)
}

// Getter models
func GetUserEmailByUID(uid string) (string, error) {
	return user_info_getter_models.GetUserEmailByUID(uid)
}

func GetUsernameByUID(uid string) (string, error) {
	return user_info_getter_models.GetUsernameByUID(uid)
}

func GetFullnameByUID(uid string) (string, error) {
	return user_info_getter_models.GetFullnameByUID(uid)
}

func GetPasswordByUID(uid string) (string, error) {
	return user_info_getter_models.GetPasswordByUID(uid)
}

func GetUserVerifiedByUID(uid string) (bool, error) {
	return user_info_getter_models.GetUserVerifiedByUID(uid)
}

func GetUserIDByUID(uid string) (int, error) {
	return user_info_getter_models.GetUserIDByUID(uid)
}

// Password reset models
func SavePasswordResetOTP(db *sql.DB, email, code string) error {
	return user_info_getter_models.SavePasswordResetOTP(db, email, code)
}

func GetPasswordResetOTP(db *sql.DB, email string) (string, time.Time, error) {
	return user_info_getter_models.GetPasswordResetOTP(db, email)
}

func DeletePasswordResetOTP(db *sql.DB, email string) error {
	return user_info_getter_models.DeletePasswordResetOTP(db, email)
}

func AddPasswordResetRequest(db *sql.DB, email string) error {
	return user_info_getter_models.AddPasswordResetRequest(db, email)
}

func CountPasswordResetRequestsLastHour(db *sql.DB, email string) (int, error) {
	return user_info_getter_models.CountPasswordResetRequestsLastHour(db, email)
}

func SetPasswordResetCooldown(db *sql.DB, email string, until time.Time) error {
	return user_info_getter_models.SetPasswordResetCooldown(db, email, until)
}

func GetPasswordResetCooldown(db *sql.DB, email string) (time.Time, error) {
	return user_info_getter_models.GetPasswordResetCooldown(db, email)
}
