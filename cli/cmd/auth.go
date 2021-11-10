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
	"flip/api"
	"flip/utils"
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// authCMD represents the auth command
var authCMD = &cobra.Command{
	Use:   "auth",
	Short: "Authentication",
	Long:  `Authenticate to your Flipside account, logout and view the currently logged in user :)`,
}

// loginCmd represents the generate command
func loginCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "login",
		Short: "Login to your Flipside account",
		Long:  `Use your username and password to authenticate with your Flipside account. Authentication is required to run most commands.`,
		Run: func(cmd *cobra.Command, args []string) {
			username, usernameErr := utils.GetUsername()
			if usernameErr != nil {
				fmt.Printf("Invalid username %v\n", usernameErr)
				return
			}

			pwd, pwdErr := utils.GetPwd()
			if pwdErr != nil {
				fmt.Printf("Invalid password %v\n", pwdErr)
				return
			}

			config := api.Config{BaseURL: "http://localhost:3000"}
			client, _ := api.NewClient(config)
			response, authErr := client.Authenticate(username, pwd)
			if authErr != nil {
				fmt.Println("⛔️ Failed to authenticate. Invalid username/password. Are you sure you're using your Flipside creds?")
				return
			}

			viper.Set("JWT", response.JWT)
			viper.WriteConfig()

			fmt.Println(fmt.Sprintf("🥳 Welcome %s! You have successfully authenticated ✅", username))
		},
	}
}

func logoutCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "logout",
		Short: "Logout of your Flipside account",
		Long:  `Logout of your Flipside account.`,
		Run: func(cmd *cobra.Command, args []string) {
			jwt := viper.Get("jwt")
			if jwt == nil || jwt == "" {
				fmt.Println("No user is currently logged-in.")
				return
			}
			viper.Set("JWT", "")
			viper.WriteConfig()
			fmt.Println(fmt.Sprintf("Peace out ✌️. You've been successfully logged out."))
		},
	}
}

// meCmd represents the generate command
func meCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "me",
		Short: "Get the currently logged-in user",
		Long:  `Get the currently logged-in user.`,
		Run: func(cmd *cobra.Command, args []string) {
			jwt := viper.Get("jwt")
			if jwt == nil || jwt == "" {
				fmt.Println("No user is currently logged-in. Run `flip auth login` to login to your account!")
				return
			}

			config := api.Config{BaseURL: "http://localhost:3000", JWT: jwt.(string)}
			client, _ := api.NewClient(config)
			me, err2 := client.GetMe()
			if err2 != nil {
				fmt.Println(err2)
				return
			}

			t := table.NewWriter()
			t.SetOutputMirror(os.Stdout)
			fmt.Println()
			fmt.Println(fmt.Sprintf("Sup, you're currently logged-in as `%s`", me.UserName))
			t.AppendRow([]interface{}{"User ID", me.ID})
			t.AppendRow([]interface{}{"Username", me.UserName})
			t.AppendRow([]interface{}{"Org", me.OrganizationName})
			t.AppendRow([]interface{}{"Org ID", me.OrganizationID})
			t.SetStyle(table.StyleLight)
			t.Render()
		},
	}
}

func init() {
	authCMD.AddCommand(loginCommand())
	authCMD.AddCommand(logoutCommand())
	authCMD.AddCommand(meCommand())
	rootCmd.AddCommand(authCMD)
}
