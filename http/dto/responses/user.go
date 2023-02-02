package responses

import "gedebook.com/api/domain"

type UserLoginResponse struct {
	ID          int64  `json:"id"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	AccessToken string `json:"access_token"`
}

type UserProfileResponse struct {
	ID             int64   `json:"id"`
	Email          string  `json:"email"`
	Name           string  `json:"name"`
	Profile        *string `json:"profile"`
	ProfilePicture *string `json:"profile_pcture"`
	LifePoint      int32   `json:"life_point"`
}

func AssignedUserProfile(src domain.User) (res UserProfileResponse) {
	res.ID = src.ID
	res.Email = src.Email
	res.Name = src.Name
	res.Profile = src.Profile
	res.ProfilePicture = src.ProfilePicture
	res.LifePoint = src.LifePoint
	return
}
