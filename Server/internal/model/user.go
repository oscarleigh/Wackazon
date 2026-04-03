package model

type User struct {
	ID            string  `json:"id"`
	Email         string  `json:"email"`
	FirstName     string  `json:"firstName"`
	LastName      string  `json:"lastName"`
	JoinDate      Date    `json:"joinDate"`
	PwdHash       string  `json:"-"`
	StreetAddress *string `json:"streetAddress"`
	City          *string `json:"city"`
	Country       *string `json:"country"`
	PostCode      *string `json:"postCode"`
}
