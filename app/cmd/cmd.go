package cmd

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/cobra"

	"github.com/TaichiOhmi/jwt-auth-with-goreact/app/database"
	"github.com/TaichiOhmi/jwt-auth-with-goreact/app/routes"
)

//func NewRootCmd() *cobra.Command {
//	return &cobra.Command{
//		Use:   "server",
//		Short: "run server",
//		Long:  `This is server run command`,
//		RunE:  serverRun,
//	}
//}

func RegisterComannd(rootCmd *cobra.Command) {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:  "server",
	RunE: serverRun,
}

func serverRun(_ *cobra.Command, _ []string) error {
	database.Connect()

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	routes.Setup(app)

	app.Listen("localhost:8080")

	return nil
}
