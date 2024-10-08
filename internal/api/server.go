package api

import (
	"fmt"

	"github.com/MIcQo/maptonic/config"
	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humafiber"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/mikhail-bigun/fiberlogrus"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Debug bool
	Host  string
	Port  uint
}

func NewServer(c *Config) error {
	// Create your router.
	router := fiber.New(fiber.Config{
		AppName:               "MapTonic",
		DisableStartupMessage: !c.Debug,
		EnablePrintRoutes:     c.Debug,
	})

	// Register your middleware.
	router.Use(
		registerLogger(),
		registerHealthCheck(),
		registerPrometheus(router),
	)

	// Wrap the router with Huma to create an API instance.
	var api = humafiber.New(router, humaConfig())

	// Register your operations with the API.
	huma.Register(api, huma.Operation{Path: "/reverse", Method: "GET"}, ReverseGeocodingHandler)
	huma.Register(api, huma.Operation{Path: "/search", Method: "GET"}, SearchGeocodingHandler)

	logrus.Info("Server started at ", fmt.Sprintf("%s:%d", c.Host, c.Port))
	// Start the server!
	return router.Listen(fmt.Sprintf("%s:%d", c.Host, c.Port))
}

func registerHealthCheck() fiber.Handler {
	return healthcheck.New(healthcheck.Config{
		LivenessProbe: func(c *fiber.Ctx) bool {
			return true
		},
		ReadinessProbe: func(c *fiber.Ctx) bool {
			return true
		},
	})
}

func registerLogger() fiber.Handler {
	return fiberlogrus.New()
}

func registerPrometheus(router *fiber.App) func(ctx *fiber.Ctx) error {
	prometheus := fiberprometheus.New("maptonic")
	prometheus.RegisterAt(router, "/metrics")
	prometheus.SetSkipPaths([]string{"/readyz", "livez", "/openapi.yaml"})
	return prometheus.Middleware
}

func humaConfig() huma.Config {
	var cfg = huma.DefaultConfig(
		"MapTonic",
		config.Version,
	)
	cfg.Info.Description = `Provides endpoints to retrieve information about countries name and/or address of POI.`
	return cfg
}
