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
	"fmt"
	"os"
	"os/exec"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

const DBT_PROJECT_TEMPLATE_URL = "https://github.com/jfmyers/flip-starter-project"
const DBT_PROJECT_TEMPLATE_BRANCH = "main"

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a new FLIP project.",
	Long:  `Generates a brand new flip project.`,
	Run: func(cmd *cobra.Command, args []string) {
		projectName, projectNameErr := getProjectName()
		if projectNameErr != nil {
			fmt.Printf("❌ Failed to create project %v\n", projectNameErr)
			return
		}

		if _, err := os.Stat("./" + projectName); !os.IsNotExist(err) {
			fmt.Println(fmt.Sprintf("🙅 Project `%s` already exists in this directory!", projectName))
			return
		}

		git_clone_args := []string{"clone", "--depth=1", "--branch", DBT_PROJECT_TEMPLATE_BRANCH, DBT_PROJECT_TEMPLATE_URL, projectName}
		cmdh := exec.Command("git", git_clone_args...)
		cloneStdout, cloneErr := cmdh.Output()

		if cloneErr != nil {
			fmt.Println("❌ Error cloning default project: ")
			fmt.Println(cloneStdout)
			fmt.Println(cloneErr.Error())
			return
		}

		rmCmd := exec.Command("rm", "-rf", "./"+projectName+"/.git")
		_, rmErr := rmCmd.Output()
		if rmErr != nil {
			fmt.Println("❌ Error removing .git config", rmErr.Error())
			return
		}

		fmt.Println(fmt.Sprintf("🥳 Your project was succesfully initialized! Run `cd ./%s` to navigate to your project's directory.", projectName))
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
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
