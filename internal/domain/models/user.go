package models

import "time"

type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Login_Request struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type User_Details struct {
	ID              string `json:"id"`
	User_id         string `json:"user_id"`
	First_name      string `json:"first_name"`
	Last_name       string `json:"last_name"`
	Profile_picture string `json:"profile_picture"`
}
