package app

import (
	"context"
	"flag"
	"fmt"
	"github.com/YANcomp/platform_common/pkg/closer"
	"github.com/YANcomp/yanbackend/stories_service/internal/config"
	desc "github.com/YANcomp/yanbackend/stories_service/pkg/stories_v1"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"sync"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config_path", "", "stories_service config path")
}

type App struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()

		err := a.runGRPCServer()
		if err != nil {
			log.Fatalf("failed to run GRPC server: %v", err)
		}
	}()

	wg.Wait()

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initGRPCServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	getwd, err := os.Getwd()
	if err != nil {
		return errors.Wrap(err, "os.Getwd")
	}
	var configFullPath string

	if configPath == "" {
		configPathFromEnv := os.Getenv("config_path")
		if configPathFromEnv != "" {
			configFullPath = fmt.Sprintf("%s/%s", getwd, configPathFromEnv)
		} else {
			configFullPath = fmt.Sprintf("%s/local.env", getwd)
		}
	}
	configFullPath = fmt.Sprintf("%s/%s", getwd, configPath)
	log.Printf(configPath)
	log.Printf(configFullPath)
	err = config.Load(configFullPath)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initGRPCServer(ctx context.Context) error {
	a.grpcServer = grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
	)

	reflection.Register(a.grpcServer)

	desc.RegisterStoriesV1Server(a.grpcServer, a.serviceProvider.StoriesImpl(ctx))

	return nil
}

func (a *App) runGRPCServer() error {
	log.Printf("GRPC server is running on %s", a.serviceProvider.GRPCConfig().Address())

	list, err := net.Listen("tcp", a.serviceProvider.GRPCConfig().Address())
	if err != nil {
		return err
	}

	err = a.grpcServer.Serve(list)
	if err != nil {
		return err
	}

	return nil
}
