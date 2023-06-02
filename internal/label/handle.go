package label

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type LabelHandler struct {
	labelService LabelService
}

func NewLabelHandler(labelRoute fiber.Router, ls LabelService) {
	handler := &LabelHandler{
		labelService: ls,
	}

	labelRoute.Get("", handler.getLabels)
	labelRoute.Post("", handler.createLabel)

	labelRoute.Get("/:labelID", handler.getLabelById)
	labelRoute.Put("/:labelID", handler.checkIfLabelExistsMiddleware, handler.updateLabel)
	labelRoute.Delete("/:labelID", handler.checkIfLabelExistsMiddleware, handler.deleteLabel)
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

func (h *LabelHandler) createLabel(c *fiber.Ctx) error {
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	label := &Label{}

	if err := c.BodyParser(label); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	err := h.labelService.CreateLabel(customContext, label)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Label has been created successfully!",
	})
}

func (h *LabelHandler) getLabelById(c *fiber.Ctx) error {
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	targetLabelId, err := c.ParamsInt("labelID")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": "Please provide a valid label id",
		})
	}
	label, err := h.labelService.GetLabelById(customContext, targetLabelId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   label,
	})
}

func (h *LabelHandler) deleteLabel(c *fiber.Ctx) error {
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	targetLabelId := c.Locals("labelID").(int)
	err := h.labelService.DeleteLabel(customContext, targetLabelId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *LabelHandler) updateLabel(c *fiber.Ctx) error {
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	targetLabelId := c.Locals("labelID").(int)
	label := &Label{Model: gorm.Model{ID: uint(targetLabelId)}}

	if err := c.BodyParser(label); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	err := h.labelService.UpdateLabel(customContext, label)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Label has been updated successfully!",
	})
}
