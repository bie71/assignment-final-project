package models

import "time"

type UsersModels struct {
	UserId    string    `dbq:"user_id"`
	Name      string    `dbq:"name"`
	Username  string    `dbq:"username"`
	Password  string    `dbq:"password"`
	UserType  string    `dbq:"user_type"`
	CreatedAt time.Time `dbq:"created_at"`
}

func TableNameUsers() string {
	return "users"
}

func UsersFieldName() []string {
	return []string{
		"user_id",
		"name",
		"username",
		"password",
		"user_type",
		"created_at",
	}
}
