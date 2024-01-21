package app

import (
	"context"
	"github.com/YANcomp/platform_common/pkg/closer"
	"github.com/YANcomp/platform_common/pkg/db"
	"github.com/YANcomp/platform_common/pkg/db/pg"
	"github.com/YANcomp/platform_common/pkg/db/transaction"
	"github.com/YANcomp/yanbackend/stories_service/internal/api/stories"
	"github.com/YANcomp/yanbackend/stories_service/internal/config"
	"github.com/YANcomp/yanbackend/stories_service/internal/repository"
	storiesRepository "github.com/YANcomp/yanbackend/stories_service/internal/repository/stories"
	"github.com/YANcomp/yanbackend/stories_service/internal/service"
	storiesService "github.com/YANcomp/yanbackend/stories_service/internal/service/stories"
	"log"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig

	dbClient          db.Client
	txManager         db.TxManager
	storiesRepository repository.StoriesRepository

	storiesService service.StoriesService

	storiesImpl *stories.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

func (s *serviceProvider) StoriesRepository(ctx context.Context) repository.StoriesRepository {
	if s.storiesRepository == nil {
		s.storiesRepository = storiesRepository.NewRepository(s.DBClient(ctx))
	}

	return s.storiesRepository
}

func (s *serviceProvider) StoriesService(ctx context.Context) service.StoriesService {
	if s.storiesService == nil {
		s.storiesService = storiesService.NewService(
			s.StoriesRepository(ctx),
		)
	}

	return s.storiesService
}

func (s *serviceProvider) StoriesImpl(ctx context.Context) *stories.Implementation {
	if s.storiesImpl == nil {
		s.storiesImpl = stories.NewImplementation(s.StoriesService(ctx))
	}

	return s.storiesImpl
}
