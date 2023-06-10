package bill

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"log"
)

type BillHandler struct {
	billService BillService
}

func NewBillHandler(billRoute fiber.Router, bs BillService) {
	handler := &BillHandler{
		billService: bs,
	}

	billRoute.Get("/:year?/:month?/:day?", handler.checkGetBillsParamsMiddleware, handler.getBills)
	billRoute.Post("/", handler.createBill)
	billRoute.Put("/", handler.updateBill)
	billRoute.Delete("/", handler.deleteBill)
}

func (h *BillHandler) getBills(c *fiber.Ctx) error {
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	bDate := c.Locals("bDate").(BDate)

	bills, err := h.billService.GetBills(customContext, bDate)
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

	bill := &Bill{}
	if err := c.BodyParser(bill); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	log.Printf("%v\n", bill)

	err := h.billService.CreateBill(customContext, bill)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Bill has been created successfully!",
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
