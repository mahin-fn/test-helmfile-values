package cmd

import (
	"fmt"
	"os"

	"github.com/mokhlesurr031/sre-project/backend/internal/envs"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "root",
	Short: "go-starter Management CLI",
	Long:  `A practice golang microservice project`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	envs.Init()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
