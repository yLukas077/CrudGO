package models

type User struct {
    ID uint32 `json:"id"`
    Name string `json:"name"`
    Email string `json:"email"`
}
