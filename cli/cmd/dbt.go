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
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const DOCKER_IMAGE = "flip:latest"

// dbtCmd represents the dbt command
var dbtCmd = &cobra.Command{
	Use:   "dbt",
	Short: "Model blockchain data using the DBT SQL framework.",
	// Run: func(cmd *cobra.Command, args []string) {
	// 	runConsole()
	// },
}

func newConsoleCommand() *cobra.Command {
	cmd := &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			runConsole()
		},
		Use:   `console`,
		Short: "Enter into a DBT environment for your project.",
		Long:  `Creates a DBT environment where you can run any DBT commands against the 'sql_models' directory.`,
	}

	return cmd
}

func runConsole() {
	fmt.Println("⌨️ Spinning up an interactive DBT console")
	path, _ := os.Getwd()
	mount := filepath.Join(path, "sql_models")
	username := viper.GetString("username")
	docker_args := []string{"run", "-it", "--env", "FLIP_USERNAME=" + username, "-v", mount + ":/sql_models", DOCKER_IMAGE, "/support/dbt_console.sh"}
	cmdh := exec.Command("docker", docker_args...)
	// stdout, err := cmdh.Output()
	cmdh.Stdout = os.Stdout
	cmdh.Stdin = os.Stdin
	cmdh.Stderr = os.Stderr
	cmdh.Run()
}

func newDocsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			port, _ := cmd.Flags().GetString("port")
			if port == "" {
				port = "8085"
			}
			path, _ := os.Getwd()
			mount := filepath.Join(path, "sql_models")
			username := viper.GetString("username")
			docker_args := []string{"run", "-it", "-p", port + ":" + port, "--env", "FLIP_USERNAME=" + username, "--env", "DBT_DOCS_PORT=" + port, "-v", mount + ":/sql_models", DOCKER_IMAGE, "/support/dbt_docs.sh"}
			cmdh := exec.Command("docker", docker_args...)
			// stdout, err := cmdh.Output()
			cmdh.Stdout = os.Stdout
			cmdh.Stdin = os.Stdin
			cmdh.Stderr = os.Stderr
			cmdh.Run()
		},
		Use:   `docs`,
		Short: "Launch DBT docs",
	}

	return cmd
}

func init() {
	docsCmd := newDocsCommand()
	var Port string
	docsCmd.Flags().StringVarP(&Port, "port", "p", "", "Local port to run dbt docs server on.")

	dbtCmd.AddCommand(docsCmd)
	dbtCmd.AddCommand(newConsoleCommand())

	rootCmd.AddCommand(dbtCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dbtCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dbtCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
