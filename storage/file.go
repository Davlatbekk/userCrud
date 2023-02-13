package storage

import (
	"app/config"
	"os"
)

type Store struct {
	User *userRepo
}

func NewFileJson(cfg *config.Config) (*Store, error) {

	// if doesFileExist(cfg.Path + cfg.UserFileName) {
	// 	_, err := os.Create(cfg.Path + cfg.UserFileName)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }

	userFile, err := os.Open(cfg.Path + cfg.UserFileName)
	if err != nil {
		return nil, err
	}

	return &Store{
		User: NewUserRepo(cfg.Path+cfg.UserFileName, userFile),
	}, nil
}

// function to check if file exists
func doesFileExist(fileName string) bool {
	_, error := os.Stat(fileName)
	return os.IsNotExist(error)
}
