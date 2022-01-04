package main

import (
	"context"

	"github.com/Sanjungliu/assesment-user-service/config"
	"github.com/Sanjungliu/assesment-user-service/database"
	"github.com/Sanjungliu/assesment-user-service/internal/app"
	"github.com/Sanjungliu/assesment-user-service/internal/auth"
	"github.com/Sanjungliu/assesment-user-service/internal/user"
)

func buildApp(ctx context.Context, cfg *config.Config) *app.App {
	storage := database.DBinstance(ctx, cfg)
	auth := auth.NewService(cfg.JWTSecretKey())
	DBCollection := database.OpenCollection(storage, cfg.DBName(), cfg.DBCollection())
	user := user.NewService(DBCollection)

	return &app.App{
		Storage: storage,
		Auth:    auth,
		User:    user,
	}
}
