package misc

import (
	"github.com/gofiber/fiber/v2"
	"os"
	"time"
)

// Create a handler. Leave this empty, as we have no domains nor use-cases.
type MiscHandler struct{}

// Represents a new handler.
func NewMiscHandler(miscRoute fiber.Router) {
	handler := &MiscHandler{}

	// Declare routing.
	miscRoute.Get("", handler.healthCheck)
	miscRoute.Get("/shutdown", handler.shutdown)
}

// Check for the health of the API.
func (h *MiscHandler) healthCheck(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Hello World!",
	})
}

func (h *MiscHandler) shutdown(ctx *fiber.Ctx) error {
	go func() {
		time.Sleep(1 * time.Second)
		os.Exit(0)
	}()
	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "server shutdown",
	})
}
