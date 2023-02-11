package main

import (
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/TaichiOhmi/jwt-auth-with-goreact/app/cmd"
)

func main() {
	c := &cobra.Command{Use: "select mode"}

	cmd.RegisterComannd(c)
	err := c.Execute()
	if err != nil {
		log.Printf(err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}
