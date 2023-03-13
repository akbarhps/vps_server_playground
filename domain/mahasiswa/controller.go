package mahasiswa

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"vps_server_playground/util"
)

type Controller interface {
	CreateMahasiswa(ctx *fiber.Ctx) error
}

type controllerImpl struct {
	service Service
}

func NewController(service Service) Controller {
	return &controllerImpl{service}
}

func (c *controllerImpl) CreateMahasiswa(ctx *fiber.Ctx) error {
	req := new(createRequest)
	if err := ctx.BodyParser(req); err != nil {
		return err
	}

	log.Println("req: ", req)

	serviceCtx := context.Background()

	mhs, err := c.service.Create(serviceCtx, req)
	if err != nil {
		return err
	}

	return ctx.Status(http.StatusCreated).JSON(&util.WebResponse{
		StatusCode: http.StatusCreated,
		Status:     "create mahasiswa success",
		Data:       mhs,
	})
}