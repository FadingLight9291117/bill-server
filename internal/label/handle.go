package label

import (
	"context"
	"github.com/gofiber/fiber/v2"
)

type LabelHandler struct {
	labelService LabelService
}

func NewLabelHandler(labelRoute fiber.Router, ls LabelService) {
	handler := &LabelHandler{
		labelService: ls,
	}

	labelRoute.Get("", handler.getLabels)
}

func (h *LabelHandler) getLabels(c *fiber.Ctx) error {
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()
	labels, err := h.labelService.GetLabels(customContext)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   labels,
	})
}
