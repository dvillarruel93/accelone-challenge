package model

// Contact struct is the representation of the contact in memoryDB
// json tags are the name who are going to show in json response
// validate tags are to validate with received bodies, more info in README.md
type Contact struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name" validate:"required,min=0,max=100"`
	LastName  string `json:"last_name" validate:"required,min=0,max=100"`
	Address   string `json:"address" validate:"required,min=0,max=300"`
}
