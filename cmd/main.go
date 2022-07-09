package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"memoryCache"
	"memoryCache/pkg/cache"
	"memoryCache/pkg/handler"
	"memoryCache/pkg/service"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	caches := cache.NewCache(viper.GetDuration("cache.defaultExpiration"),
		viper.GetDuration("cache.cleanupInterval"))
	services := service.NewService(caches)
	handlers := handler.NewHandler(services, handler.Config{
		Duration: viper.GetDuration("cache.itemDuration"),
	})

	srv := new(memoryCache.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
