package formatters

import "silih_a3/entities"

type UserDataFormatter struct {
	UserId     int    `json:"user_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	AvatarPath string `json:"avatar_path"`
}

type SignInFormatter struct {
	UserId    int    `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	UserToken string `json:"token"`
}

func UserDataFormat(userData entities.User) UserDataFormatter {
	userFormatter := UserDataFormatter{
		UserId:     userData.Id,
		FirstName:  userData.FirstName,
		LastName:   userData.LastName,
		Email:      userData.Email,
		AvatarPath: userData.AvatarPath,
	}

	return userFormatter
}

func Format(userData entities.User, userToken string) SignInFormatter {
	userFormatter := SignInFormatter{
		UserId:    userData.Id,
		FirstName: userData.FirstName,
		LastName:  userData.LastName,
		Email:     userData.Email,
		UserToken: userToken,
	}

	return userFormatter
}
