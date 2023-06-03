package bill

import (
	"context"
	"github.com/gofiber/fiber/v2"
)

type BillHandler struct {
	billService BillService
}

func NewBillHandler(billRoute fiber.Router, bs BillService) {
	handler := &BillHandler{
		billService: bs,
	}

	billRoute.Get("/", handler.getBills)
	billRoute.Post("/", handler.createBill)
	billRoute.Put("/", handler.updateBill)
	billRoute.Delete("/", handler.deleteBill)
}

func (h *BillHandler) getBills(c *fiber.Ctx) error {
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()
	// TODO()
	bills, err := h.billService.GetBills(customContext, BDate{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   bills,
	})

}

func (h *BillHandler) createBill(c *fiber.Ctx) error {
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()
	// TODO(): create bill
	err := h.billService.CreateBill(customContext, &Bill{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
	})
}

func (h *BillHandler) updateBill(c *fiber.Ctx) error {
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()
	// TODO(): update data
	err := h.billService.UpdateBill(customContext, &Bill{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
	})
}

func (h *BillHandler) deleteBill(c *fiber.Ctx) error {
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()
	// TODO(): delete data id
	err := h.billService.DeleteBill(customContext, 0)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
	})
}
