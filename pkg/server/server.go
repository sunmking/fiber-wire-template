package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type FiberServer struct {
	App *fiber.App
}

func NewFiberServer(app *fiber.App) *FiberServer {
	return &FiberServer{
		App: app,
	}
}

func (f *FiberServer) Start(port int) error {
	// @todo: add to env vars
	if err := f.App.Listen(":" + fmt.Sprintf("%d", port)); err != nil {
		return err
	}
	return nil
}
