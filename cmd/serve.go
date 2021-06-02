package cmd

import (
	_ "github.com/JDR-ynovant/api/docs"
	"github.com/JDR-ynovant/api/internal/routes"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
	"log"
)

// @title Candy Fight API
// @version 1.0
// @description Swagger documentation for Candy-Fight Game API
// @contact.name API Support
// @contact.email fiber@swagger.io
// @host localhost:3000
// @BasePath /
func executeServeCommand() {
	handlers := []routes.RouteHandler{routes.UserRouteHandler{}}

	log.Println("Serving candy-fight API...")
	app := fiber.New()

	app.Get("/swagger/*", swagger.Handler)

	for _, handler := range handlers {
		handler.Register(app)
	}

	app.Listen(":3000")
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "HTTP REST API for Candy-Fight game.",
	Run: func(cmd *cobra.Command, args []string) {
		executeServeCommand()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}