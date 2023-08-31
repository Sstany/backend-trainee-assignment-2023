package models

type User struct {
	ID       int    `json:"userId" uri:"id" binding:"required"`
	Username string `json:"username" uri:"name"`
}
