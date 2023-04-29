package main

import (
	"senao/cmd/db"
	"senao/cmd/server"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "senao",
	Short: "start senao",
	Long:  "",
}

func init() {
	RootCmd.AddCommand(db.Cmd)
	RootCmd.AddCommand(server.Cmd)
}

var (
	VERSION = "v1"
)

func main() {
	server.VERSION = VERSION
	RootCmd.Execute()
	// server.Cmd.Execute()
}
