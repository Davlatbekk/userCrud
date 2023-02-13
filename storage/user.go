package storage

import (
	"app/models"
	"encoding/json"
	"io/ioutil"
	"os"
)

type userRepo struct {
	fileName string
	file     *os.File
}

// Constructor
func NewUserRepo(fileName string, file *os.File) *userRepo {
	return &userRepo{
		fileName: fileName,
		file:     file,
	}
}

func (u *userRepo) Create(req *models.CreateUser) (id int, err error) {

	var users []*models.User
	err = json.NewDecoder(u.file).Decode(&users)
	if err != nil {
		return 0, err
	}

	if len(users) > 0 {
		id = users[len(users)-1].Id + 1
		users = append(users, &models.User{
			Id:      id,
			Name:    req.Name,
			Surname: req.Surname,
		})
	} else {
		id = 1
		users = append(users, &models.User{
			Id:      id,
			Name:    req.Name,
			Surname: req.Surname,
		})
	}

	body, err := json.MarshalIndent(users, "", "   ")

	err = ioutil.WriteFile(u.fileName, body, os.ModePerm)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (u *userRepo) GitListUser(req *models.GetListRequest) (*models.GetListResponse, error) {

	response := models.GetListResponse{
		Users: make([]*models.User, 0),
	}

	err := json.NewDecoder(u.file).Decode(&response.Users)

	if err != nil {
		return nil, err
	}
	return &response, nil

}
