package anwser

import (
	"github.com/gofiber/fiber/v2"
)

func LogAnswer(c *fiber.Ctx, arrayLogVideo map[string]interface{}, page int, nextPage int, previousPage int, limit int, totalResults int) error {
	totalPages := (totalResults + limit - 1) / limit
	answer := fiber.Map{
		"data":          arrayLogVideo,
		"page":          page,
		"nextPage":      nextPage,
		"previousPage":  previousPage,
		"totalPages":    totalPages,
		"totalResults":  totalResults,
	}
	return c.Status(fiber.StatusOK).JSON(answer)
}