// Package entity impls api resource entities
package entity

// Person struct defines fields in person resource
type Person struct {
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	Age       uint8  `json:"age" binding:"gte=0,lte=130"`
	Email     string `json:"email" validate:"required,email"`
}

// Video struct defines fields in video resource
type Video struct {
	Title       string `json:"title" binding:"min=2,max=10" validate:"cool"`
	Description string `json:"description" binding:"max=20"`
	URL         string `json:"url" binding:"required,url"`
	Actors      int8   `json:"actors" binding:"gt=2,lt=20"`
	Author      Person `json:"author"`
}
