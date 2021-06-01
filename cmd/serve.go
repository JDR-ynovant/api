package cmd

import (
	"fmt"
	"github.com/JDR-ynovant/api/internal/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
)

func executeServeCommand() {
	handlers := []routes.RouteHandler{routes.UserRouteHandler{}}

	fmt.Println("Serving candy-fight API...")
	app := fiber.New()

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