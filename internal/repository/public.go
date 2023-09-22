package repository

import (
	"context"
	"github.com/google/uuid"
	"punishments/internal/repository/model"
)

type Repository interface {
	GetPlayer(ctx context.Context, uuid uuid.UUID) (*model.Player, error)

	GetActivePunishments(ctx context.Context, uuid uuid.UUID) ([]model.IssuedPunishment, error)

	GetAllPunishments(ctx context.Context, uuid uuid.UUID) ([]model.IssuedPunishment, error)
}
