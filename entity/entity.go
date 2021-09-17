package entity

// Entity is the most innermost circle so in clean code isn't affected with any external factors
// Business is the only factor


type User struct {
	Address string `json:"address"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Password string `json:"password"`
	Username string `json:"username"`
}