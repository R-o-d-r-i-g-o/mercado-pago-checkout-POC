package main

import (
	"code-space-backend-api/api"
	"code-space-backend-api/app/config"
	"code-space-backend-api/common/cache"
	"code-space-backend-api/common/process"
	"code-space-backend-api/env"
	"code-space-backend-api/infra/database"
	"context"
	"time"
)

func init() {
	env.Load()
	database.Load()
	cache.InitializeCacheManager(time.Minute * 5)
}

func main() {
	var container = config.NewContainer()

	defer container.Log.SyncLogs()
	defer database.CloseDB(container.Infrastructure.DB)

	var server = api.NewServer(container)

	go server.Run()

	process.GracefulShutdown(func(ctx context.Context) {
		server.Shutdown(ctx)
		container.Tracing.Provider.Shutdown(ctx)
	})
}
