package label

import (
	"context"
	"github.com/gofiber/fiber/v2"
)

func (h *LabelHandler) checkIfLabelExistsMiddleware(c *fiber.Ctx) error {
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	targetLabelID, err := c.ParamsInt("labelID")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": "Please specify a valid label ID!",
		})
	}

	searchLabel, err := h.labelService.GetLabelById(customContext, targetLabelID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}
	if searchLabel == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "fail",
			"message": "These is no label with this ID!",
		})
	}

	c.Locals("labelID", targetLabelID)
	return c.Next()

}
