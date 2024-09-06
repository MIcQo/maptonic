package api

import (
	"fmt"
	"github.com/MIcQo/maptonic/config"
	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humafiber"
	"github.com/gofiber/fiber/v2"
	"github.com/mikhail-bigun/fiberlogrus"
)

type Config struct {
	Debug bool
	Host  string
	Port  uint
}

func NewServer(c *Config) error {
	// Create your router.
	router := fiber.New(fiber.Config{
		AppName:           "MapTonic",
		EnablePrintRoutes: c.Debug,
	})

	registerLogger(router)
	registerPrometheus(router)

	// Wrap the router with Huma to create an API instance.
	_ = humafiber.New(router, humaConfig())

	// Register your operations with the API.
	// ...

	// Start the server!
	return router.Listen(fmt.Sprintf("%s:%d", c.Host, c.Port))
}

func registerLogger(router *fiber.App) {
	router.Use(fiberlogrus.New())
}

func registerPrometheus(router *fiber.App) {
	prometheus := fiberprometheus.New("maptonic")
	prometheus.RegisterAt(router, "/metrics")
	prometheus.SetSkipPaths([]string{"/health", "/openapi.yaml"})
	router.Use(prometheus.Middleware)
}

func humaConfig() huma.Config {
	var cfg = huma.DefaultConfig(
		"MapTonic",
		config.Version,
	)
	cfg.Info.Description = `Provides endpoints to retrieve information about countries name and/or address of POI.`
	return cfg
}
