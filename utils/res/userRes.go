package res

import (
	"masaku/models"
	"masaku/models/web"
)

func ConvertIndex(users []models.User) []web.UserReponse {
	var results []web.UserReponse
	for _, user := range users {
		userResponse := web.UserReponse{
			Id:    int(user.ID),
			Name:  user.Name,
			Email: user.Email,
		}
		results = append(results, userResponse)
	}

	return results
}

func ConvertGeneral(user *models.User) web.UserReponse {
	return web.UserReponse{
		Id:    int(user.ID),
		Name:  user.Name,
		Email: user.Email,
	}
}
