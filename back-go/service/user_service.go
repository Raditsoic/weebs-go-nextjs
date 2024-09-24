package service

import (
	"fmt"
	"log"

	"github.com/Raditsoic/anime-go/database/repository"
	"github.com/Raditsoic/anime-go/graph/model"
	"github.com/Raditsoic/anime-go/utils"
)

type UserService struct {
	UserRepo repository.UserRepo
}

func NewUserService(repo repository.UserRepo) *UserService {
	return &UserService{UserRepo: repo}
}

func (s *UserService) Register(input model.RegisterUser) (*model.AuthPayload, error) {
	existingEmail, err := s.UserRepo.GetByEmail(input.Email)
	if err != nil {
		log.Printf("Failed to register user: %v", err)
		return nil, fmt.Errorf("Failed to register user.")
	}
	if existingEmail != nil {
		return nil, fmt.Errorf("Email is already used.")
	}

	exitingUsername, err := s.UserRepo.GetByUsername(input.Username)
	if err != nil {
		log.Printf("Failed to register user: %v", err)
		return nil, fmt.Errorf("Failed to register user.")
	}
	if exitingUsername != nil {
		return nil, fmt.Errorf("Username is already used.")
	}

	salt, err := utils.GenerateSalt(16)
	if err != nil {
		log.Printf("Error generating salt: %v", err)
		return nil, err
	}

	hashedPassword, err := utils.HashPassword(input.Password, salt)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return nil, err
	}

	user := &model.User{
		Username: input.Username,
		Email:    input.Email,
		Role:     "user",
	}
	result, err := s.UserRepo.Register(user, salt, hashedPassword)
	if err != nil {
		log.Printf("Error registering user: %v", err)
		return nil, err
	}

	token, err := utils.GenerateJWT(result.ID, user.Role)
	if err != nil {
		log.Printf("Error generating JWT: %v", err)
		return nil, err
	}

	return &model.AuthPayload{
		Token: token,
	}, nil
}

func (s *UserService) Login(input model.LoginUser) (*model.AuthPayload, error) {
	user, err := s.UserRepo.GetByUsername(input.Username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, fmt.Errorf("Invalid username or password.")
	}

	err = utils.ComparePassword(user.Hash, input.Password, user.Salt)
	if err != nil {
		return nil, fmt.Errorf("Invalid username or password.")
	}

	token, err := utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		return nil, err
	}

	return &model.AuthPayload{
		Token: token,
	}, nil
}
