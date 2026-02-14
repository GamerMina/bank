package service

import "bank/repository"

type Authentication struct {
	Repository *repository.Repository
}

func Auth(a *repository.Repository) *Authentication {
	return &Authentication{Repository: a}
}
