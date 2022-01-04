package app

import (
	"github.com/Sanjungliu/assesment-user-service/internal/auth"
	"github.com/Sanjungliu/assesment-user-service/internal/user"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	Storage *mongo.Client
	Auth    auth.Service
	User    user.Service
}
