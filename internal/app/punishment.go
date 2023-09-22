package app

import (
	"context"
	"go.uber.org/zap"
	"os/signal"
	"punishments/internal/config"
	"punishments/internal/repository"
	"sync"
	"syscall"
)

func Run(cfg *config.Config, logger *zap.SugaredLogger) {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()
	wg := &sync.WaitGroup{}

	punishmentCfg, err := config.LoadPunishmentConfig()
	if err != nil {
		logger.Fatalw("failed to load punishment config", err)
	}
	logger.Infow("loaded punishment config", "ladderCount", punishmentCfg.Ladders)

	repoWg := &sync.WaitGroup{}
	repoCtx, repoCancel := context.WithCancel(ctx)

	repo, err := repository.NewMongoRepository(repoCtx, logger, repoWg, cfg.MongoDB)
	if err != nil {
		logger.Panicw("Failed to create repository", "error", err)
	}

	wg.Wait()
	logger.Info("shutting down")

	logger.Info("shutting down repository")
	repoCancel()
	repoWg.Wait()
}
