package main

import (
	"errors"
	"fmt"
	"github.com/omiselabs/opn-generator/cmd/cli/generator"
	"github.com/spf13/cobra"
	"os"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create new project",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires your project name")
		}
		return nil
	},
	RunE: generator.New,
}

var rootCmd = &cobra.Command{
	Use:   "cli",
	Short: "CLI for OPN code generator",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
