package user

type Formatter struct {
	Id 		  int `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email 	  string `json:"email"`
	Token 	  string `json:"token"`
}

func FormatUser(user User, token string) Formatter {
	formatter := Formatter{
		Id:        user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Token:     token,
	}
	return formatter
}