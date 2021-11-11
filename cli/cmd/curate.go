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
	"fmt"

	"github.com/spf13/cobra"
)

// curateCMD
var curateCMD = &cobra.Command{
	Use:   "curate",
	Short: "Manage a Flip Data Curation project",
	Long:  `Manage a curated data project built on Flipside's data exchange.`,
}

func initCurationProjectCMD() *cobra.Command {
	return &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("🤫 This feature is coming soon...")
		},
		Use:   `init`,
		Short: "Create a new Flip Data Curation project",
		Long:  `Generates a brand new Data Curation project.`,
	}
}

func dbtConsoleCMD() *cobra.Command {
	return &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("🤫 This feature is coming soon...")
		},
		Use:   `dbt-console`,
		Short: "Enter into a DBT environment for your project",
		Long:  `Creates a DBT environment where you can run any DBT commands against the 'sql_models' directory.`,
	}
}

func dbtDocsCMD() *cobra.Command {
	return &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("🤫 This feature is coming soon...")
		},
		Use:   `dbt-docs`,
		Short: "Launch DBT docs for your project",
		Long:  `Launch DBT docs for your Data Curation project.`,
	}
}

func init() {
	curateCMD.AddCommand(initCurationProjectCMD())
	curateCMD.AddCommand(dbtConsoleCMD())
	curateCMD.AddCommand(dbtDocsCMD())
	rootCmd.AddCommand(curateCMD)
}
