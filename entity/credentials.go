// Package entity impls api resource entities
package entity

// Credentials struct
type Credentials struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

// JWT struct
type JWT struct {
	Token string `json:"token"`
}
