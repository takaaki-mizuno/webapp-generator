/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/omiselabs/opn-generator/internal"
	"github.com/omiselabs/opn-generator/internal/handlers"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create new project",
	Long:  ``,
	RunE: func(command *cobra.Command, args []string) error {
		currentPath, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
			return err
		}
		apiDefinitionPath, err := command.Flags().GetString("api")
		if err != nil {
			log.Fatal(err)
			return err
		}
		container := internal.BuildContainer()
		var newHandler *handlers.NewHandler

		if err := container.Invoke(func(
			_newHandler *handlers.NewHandler,
		) {
			newHandler = _newHandler
		}); err != nil {
			log.Fatal(err)
			return err
		}

		return newHandler.Execute(args[0], currentPath, apiDefinitionPath)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
	newCmd.Flags().String("api", "", "OPN spec file ( yaml or json ) for user api")
}
