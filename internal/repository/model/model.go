package model

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type IssuedPunishment struct {
	Id *primitive.ObjectID `bson:"_id"`

	LadderId    string `bson:"ladderId"`
	LadderCount uint8  `bson:"ladderCount"`

	PunisherId *uuid.UUID `bson:"punisherId"`
	Comments   string     `bson:"comments"`
}

type DurationIssuedPunishment struct {
	*IssuedPunishment
	ExpiresAt *time.Time `bson:"expiresAt"`
}

type Player struct {
	Id             uuid.UUID           `bson:"_id"`
	Punishments    []*IssuedPunishment `bson:"punishments"`
	LadderProgress map[string]uint8    `bson:"ladderProgress"` // ladderId -> punishment count in that ladder
}
