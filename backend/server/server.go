package server

import (
	"rnav2022/goqr/model"
	"strconv"

	b64 "encoding/base64"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/skip2/go-qrcode"
)

// GetAllGoqrs gets all goqrs
func GetAllGoqrs(c *fiber.Ctx) error {
	goqrs, err := model.GetAllGoqrs()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error getting all goqr links " + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(goqrs)
}

// CreateGoqr creates a new goqr
func CreateGoqr(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var goqr model.Goqr
	err := c.BodyParser(&goqr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error parsing JSON " + err.Error(),
		})
	}

	var qr []byte
	qr, err = qrcode.Encode(goqr.URL, qrcode.Medium, 256)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error creating QR " + err.Error(),
		})
	}

	goqr.QRCode = b64.StdEncoding.EncodeToString(qr)
	err = model.CreateGoqr(goqr)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error could not create goqr " + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(goqr)
}

// UpgdateGoqr updates a goqr
func UpdateGoqr(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var goqr model.Goqr
	err := c.BodyParser(&goqr)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error parsing JSON " + err.Error(),
		})
	}
	var qr []byte
	qr, err = qrcode.Encode(goqr.URL, qrcode.Medium, 256)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error creating QR " + err.Error(),
		})
	}

	goqr.QRCode = b64.StdEncoding.EncodeToString(qr)
	err = model.UpdateGoqr(goqr)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error could not update goqr " + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(goqr)
}

// DeleteGoqr deletes a goqr
func DeleteGoqr(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error could not parse id " + err.Error(),
		})
	}
	err = model.DeleteGoqr(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error deleting goqr " + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "goqr deleted",
	})
}

func SetupAndListen() {
	router := fiber.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	router.Get("/goqr", GetAllGoqrs)
	router.Post("/goqr", CreateGoqr)
	router.Patch("/goqr", UpdateGoqr)
	router.Delete("/goqr/:id", DeleteGoqr)
	router.Listen(":3000")
}
