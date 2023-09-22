package repository

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"punishments/internal/config"
	"punishments/internal/repository/model"
	"punishments/internal/repository/registrytypes"
	"sync"
)

const (
	databaseName = "punishments"

	playerCollectionName = "player"
)

type mongoRepository struct {
	database *mongo.Database

	playerCollection *mongo.Collection
}

func NewMongoRepository(ctx context.Context, logger *zap.SugaredLogger, wg *sync.WaitGroup, cfg *config.MongoDBConfig) (Repository, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.URI).SetRegistry(createCodecRegistry()))
	if err != nil {
		return nil, err
	}

	database := client.Database(databaseName)
	repo := &mongoRepository{
		database:         database,
		playerCollection: database.Collection(playerCollectionName),
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		if err := client.Disconnect(ctx); err != nil {
			logger.Errorw("failed to disconnect from mongo", err)
		}
	}()

	repo.createIndexes(ctx)
	logger.Infow("created mongo indexes")

	return repo, nil
}

func (m *mongoRepository) GetPlayer(ctx context.Context, uuid uuid.UUID) (*model.Player, error) {
	return nil, nil
}

func (m *mongoRepository) GetActivePunishments(ctx context.Context, uuid uuid.UUID) ([]model.IssuedPunishment, error) {
	return nil, nil
}

func createCodecRegistry() *bsoncodec.Registry {
	r := bson.NewRegistry()

	r.RegisterTypeEncoder(registrytypes.UUIDType, bsoncodec.ValueEncoderFunc(registrytypes.UuidEncodeValue))
	r.RegisterTypeDecoder(registrytypes.UUIDType, bsoncodec.ValueDecoderFunc(registrytypes.UuidDecodeValue))

	return r
}
