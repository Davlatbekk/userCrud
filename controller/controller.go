package controller

import (
	"app/config"
	"app/storage"
)

type Controller struct {
	cfg   *config.Config
	store *storage.Store
}

// Controller Constructor
func NewController(cfg *config.Config, store *storage.Store) *Controller {
	return &Controller{
		cfg:   cfg,
		store: store,
	}
}
