package objects

import "time"

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

type UserLoginObjectRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserStoreObjectRequest struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}

type UserStoreObjectResponse struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserFindObjectResponse struct {
	ID        uint       `json:"id"`
	Username  string     `json:"username"`
	Phone     string     `json:"phone"`
	Email     string     `json:"email"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type UserLoginObjectResponse struct {
	ID          uint   `json:"id"`
	AccessToken string `json:"access_token"`
	ExpiredIn   int    `json:"expired_in"`
	Username    string `json:"username"`
}

type UserUpdateObjectRequest struct {
	Email string `json:"email"`
	Phone string `json:"phone"`
}
