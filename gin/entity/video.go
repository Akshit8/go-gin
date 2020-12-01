package entity

// Person struct
// not exporting
type person struct {
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	Age       int8   `json:"age" binding:"gte=1,lte=100"`
	Email     string `json:"email" binding:"required,email"`
}

// Video struct
type Video struct {
	Title       string `json:"title" binding:"min=2,max=10" validate:"titleCool"`
	Description string `json:"description" binding:"max=20"`
	URL         string `json:"url" binding:"required,url"`
	Author      person `json:"author" binding:"required"`
}
