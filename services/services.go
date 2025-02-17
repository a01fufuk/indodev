package services

import (
	"database/sql"
	"indodev/repositories"

	"github.com/go-redis/redis"
)

type UsecaseService struct {
	RepoDB       *sql.DB
	RedisClient  *redis.Client
	DivisionRepo repositories.DivisionRepository
}

func NewUsecaseService(repoDB *sql.DB,
	RedisClient *redis.Client,
	DivisionRepo repositories.DivisionRepository,
) UsecaseService {
	return UsecaseService{
		RepoDB:       repoDB,
		RedisClient:  RedisClient,
		DivisionRepo: DivisionRepo,
	}
}
