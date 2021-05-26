package formatters

import "silih_a3/entities"

type UserDataFormatter struct {
	UserId    int    `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
}

type Formatter struct {
	UserId    int    `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
	Token     string `json:"token"`
}

func UserFormat(user entities.User) UserDataFormatter {
	userFormatter := UserDataFormatter{
		UserId:    user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Avatar:    user.Avatar,
	}

	return userFormatter
}

func Format(user entities.User, token string) Formatter {
	userFormatter := Formatter{
		UserId:    user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Avatar:    user.Avatar,
		Token:     token,
	}

	return userFormatter
}
