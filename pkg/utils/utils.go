package utils

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func ParseUintParam(c *fiber.Ctx, param string) (uint, error) {
	idStr := c.Params(param)
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}
