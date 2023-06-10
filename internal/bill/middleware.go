package bill

import (
	"github.com/gofiber/fiber/v2"
	"gopkg.in/go-playground/validator.v9"
)

var validate = validator.New()

func (h *BillHandler) checkGetBillsParamsMiddleware(c *fiber.Ctx) error {
	year, _ := c.ParamsInt("year", 0)
	month, _ := c.ParamsInt("month", 0)
	day, _ := c.ParamsInt("day", 0)
	date := BDate{year, month, day}

	err := validate.Struct(date)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": "Please provide valid params",
		})
	}

	c.Locals("bDate", date)
	return c.Next()
}
