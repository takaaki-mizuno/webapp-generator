package main

import (
	"fmt"
	"github.com/omiselabs/gin-boilerplate/cmd"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var rootCmd = &cobra.Command{
	Use:   "cli",
	Short: "CLI for maintenance tasks",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {
}

func createDBConnection() (*gorm.DB, error) {
	container := cmd.BuildContainer()
	var db *gorm.DB

	err := container.Invoke(func(_db *gorm.DB) {
		db = _db
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	db, err := createDBConnection()
	if err != nil {
		panic(err)
	}
	viper.Set("database", db)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
