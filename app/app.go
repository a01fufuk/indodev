package app

import (
	"database/sql"
	"indodev/repositories"
	"indodev/repositories/divisionRepository"
	"indodev/services"

	"github.com/go-redis/redis"
)

func SetupApp(DB *sql.DB, repo repositories.Repository, redis *redis.Client) services.UsecaseService {

	// Repository
	divisionRepo := divisionRepository.NewDivisionRepository(repo)

	// Services
	usecaseSvc := services.NewUsecaseService(DB, redis, divisionRepo)

	return usecaseSvc
}
