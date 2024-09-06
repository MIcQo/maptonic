package api

import (
	"fmt"
	"github.com/MIcQo/maptonic/config"
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humafiber"
	"github.com/gofiber/fiber/v2"
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

	// Wrap the router with Huma to create an API instance.
	_ = humafiber.New(router, humaConfig())

	// Register your operations with the API.
	// ...

	// Start the server!
	return router.Listen(fmt.Sprintf("%s:%d", c.Host, c.Port))
}

func humaConfig() huma.Config {
	var cfg = huma.DefaultConfig(
		"MapTonic",
		config.Version,
	)
	cfg.Info.Description = `Provides endpoints to retrieve information about countries name and/or address of POI.`
	return cfg
}
