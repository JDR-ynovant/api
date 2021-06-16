package routes

import (
	"github.com/JDR-ynovant/api/internal/repository"
	"github.com/gofiber/fiber/v2"
	"log"
)

type GridsRouteHandler struct{}

func (GridsRouteHandler) Register(api fiber.Router) {
	api.Get("/grids/:id", handleGetGrid)

	log.Println("Registered push api group.")
}

// handleGetGrid godoc
// @Summary Get Grid
// @Description Fetch grid based on its ID
// @Tags grids
// @Accept  json
// @Produce  json
// @Param id path string true "Grid ID"
// @Success 200 {object} models.Grid
// @Router /grids/:id [get]
func handleGetGrid(c *fiber.Ctx) error {
	gr := repository.NewGridRepository()

	grid, err := gr.FindOneById(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(grid)
}
