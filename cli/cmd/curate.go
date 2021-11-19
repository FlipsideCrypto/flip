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
	"errors"
	"flip/services"
	"fmt"

	"github.com/manifoldco/promptui"
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
			user, err := services.NewUser()
			if err != nil {
				fmt.Println(err)
				return
			}
			_, err = user.GetMe()
			if err != nil {
				fmt.Println(err)
				return
			}

			projectName, projectNameErr := getProjectName()

			if projectNameErr != nil {
				fmt.Printf("❌ Failed to create project %v\n", projectNameErr)
				return
			}

			project, err := services.NewProject(projectName)
			if err != nil {
				fmt.Println(err)

				return
			}

			err = project.Init()
			if err != nil {
				fmt.Println(fmt.Sprintf("🥳 Your project was succesfully initialized! Run `cd ./%s` to navigate to your project's directory.", projectName))
				return
			}

		},
		Use:   `init`,
		Short: "Create a new Flip Data Curation project",
		Long:  `Generates a brand new Data Curation project.`,
	}
}

func dbtConsoleCMD() *cobra.Command {
	return &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			user, err := services.NewUser()
			if err != nil {
				fmt.Println(err)
				return
			}
			datax, err := user.GetDatax()
			if err != nil {
				fmt.Println(err)
				return
			}

			dbt, err := services.NewDBT(datax, "8000")
			if err != nil {
				fmt.Println(err)
				return
			}
			err = dbt.Console()
			if err != nil {
				fmt.Println(err)
				return
			}
		},
		Use:   `dbt-console`,
		Short: "Enter into a DBT environment for your project",
		Long:  `Creates a DBT environment where you can run any DBT commands against the 'sql_models' directory.`,
	}
}

func dbtDocsCMD() *cobra.Command {
	return &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			port, _ := cmd.Flags().GetString("port")
			if port == "" {
				port = "8085"
			}

			user, err := services.NewUser()
			if err != nil {
				fmt.Println(err)
				return
			}

			datax, err := user.GetDatax()
			if err != nil {
				fmt.Println(err)
				return
			}

			dbt, err := services.NewDBT(datax, port)
			if err != nil {
				fmt.Println(err)
				return
			}

			err = dbt.Docs()
			if err != nil {
				fmt.Println(err)
				return
			}
		},
		Use:   `dbt-docs`,
		Short: "Launch DBT docs for your project",
		Long:  `Launch DBT docs for your Data Curation project.`,
	}
}

func dbtResetEnv() *cobra.Command {
	return &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			user, err := services.NewUser()
			if err != nil {
				fmt.Println(err)
				return
			}

			datax, err := user.GetDatax()
			if err != nil {
				fmt.Println(err)
				return
			}

			dbt, err := services.NewDBT(datax, "8000")
			if err != nil {
				fmt.Println(err)
				return
			}

			err = dbt.ResetEnv()
			if err != nil {
				fmt.Println(err)
				return
			}
		},
		Use:   `reset-env`,
		Short: "Reset your local env (remove docker image, etc.)",
		Long:  `Reset your local env (remove docker image, etc.)`,
	}
}

func init() {
	curateCMD.AddCommand(initCurationProjectCMD())
	curateCMD.AddCommand(dbtConsoleCMD())

	var Port string
	dbtDocs := dbtDocsCMD()
	dbtDocs.Flags().StringVarP(&Port, "port", "p", "", "Local port to run dbt docs server on.")
	curateCMD.AddCommand(dbtDocs)

	curateCMD.AddCommand(dbtResetEnv())
	rootCmd.AddCommand(curateCMD)
}

func getProjectName() (string, error) {
	validate := func(input string) error {
		if input == "" {
			return errors.New("🙅 A project name is required!")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "What's the name of your project?",
		Validate: validate,
	}

	result, err := prompt.Run()

	return result, err
}
