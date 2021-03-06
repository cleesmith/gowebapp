package model

import (
	"time"

	"gowebapp/shared/database"
)

type User struct {
	Email      string    `json:"email"`
	First_name string    `json:"first_name"`
	Last_name  string    `json:"last_name"`
	Password   string    `json:"password"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

// UserByEmail gets user information from email
func UserByEmail(email string) (User, error) {
	result := User{}
	// err := database.DB.Get(&result, "SELECT id, password, status_id, first_name FROM user WHERE email = ? LIMIT 1", email)
	var err error
	return result, err
}

// UserIdByEmail gets user id from email
func UserIdByEmail(email string) (User, error) {
	result := User{}
	// err := database.DB.Get(&result, "SELECT id FROM user WHERE email = ? LIMIT 1", email)
	var err error
	return result, err
}

func UserCreate(first_name, last_name, email, password string) error {
	var err error
	now := time.Now() // time.RFC3339 ???
	// if err = bucket.Put([]byte(user.Created_at.Format(time.RFC3339)), encoded); err != nil {
	// 	return err
	// }
	// if err = bucket.Put([]byte(user.Updated_at.Format(time.RFC3339)), encoded); err != nil {
	// 	return err
	// }
	user := &User{
		First_name: first_name,
		Last_name:  last_name,
		Email:      email,
		Password:   password,
		Created_at: now,
		Updated_at: now,
	}
	//              bucket,  key,        data
	database.Update("users", user.Email, &user)
	return err
}
