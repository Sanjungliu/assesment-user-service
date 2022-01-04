package user

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	GetUserByID(userID string) (User, error)
}

type service struct {
	userCollection *mongo.Collection
}

func NewService(repository *mongo.Collection) Service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	count, err := s.userCollection.CountDocuments(ctx, bson.M{"email": input.Email})
	if err != nil {
		return User{}, err
	}

	if count > 0 {
		return User{}, errors.New("email already registered")
	}

	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
	if err != nil {
		return User{}, errors.New("error encrypt password")
	}

	user := User{
		ID:        primitive.NewObjectID(),
		Fullname:  input.Fullname,
		Email:     input.Email,
		Password:  string(password),
		Role:      input.Role,
		UserID:    primitive.NewObjectID().Hex(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, insertErr := s.userCollection.InsertOne(ctx, user)
	if insertErr != nil {
		return User{}, err
	}

	return user, nil
}

func (s *service) Login(input LoginInput) (User, error) {
	return User{}, nil
}

func (s *service) GetUserByID(userID string) (User, error) {
	return User{}, nil
}
