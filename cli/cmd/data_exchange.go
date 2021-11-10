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
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// dataExchangeCMD represents the snowflake command
var dataExchangeCMD = &cobra.Command{
	Use:   "datax",
	Short: "Direct database (snowflake) access to Flipside's Data Exchange.",
	Long:  `Interact with a Snowflake account that will give you full access Flipside's Data Exchange.`,
}

func credsCommand() *cobra.Command {
	return &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			jwt := viper.Get("jwt")
			if jwt == nil || jwt == "" {
				fmt.Println("No user is currently logged-in. Run `flip auth login` to login to your account!")
				return
			}

			config := api.Config{BaseURL: "http://localhost:3000", JWT: jwt.(string)}
			client, _ := api.NewClient(config)
			creds, err2 := client.GetDataExchangeCreds()
			if err2 != nil {
				fmt.Println("⛔️ An error occurred fetching your credentials. Please try again!")
				return
			}
			fmt.Println("")
			fmt.Println("👋 Access the Data Exchange Web UI here:")
			fmt.Println("")
			fmt.Println(fmt.Sprintf("   🔗 %s", creds.Url))
			fmt.Println("")
			fmt.Println("")
			fmt.Println("Find your account credentials below that you can use to login at the above link:")
			fmt.Println("")
			t := table.NewWriter()
			t.SetOutputMirror(os.Stdout)
			t.AppendRow([]interface{}{"Account", creds.Account})
			t.AppendRow([]interface{}{"Username", creds.Username})
			t.AppendRow([]interface{}{"Password", creds.Password})
			t.AppendRow([]interface{}{"Database", creds.Database})
			t.AppendRow([]interface{}{"Warehouse", creds.Warehouse})
			t.SetStyle(table.StyleLight)
			t.Render()
		},
		Use:   `creds`,
		Short: "Print your Data Exchange credentials.",
		Long:  `Print your Data Exchange credentials.`,
	}

}

func init() {
	dataExchangeCMD.AddCommand(credsCommand())
	rootCmd.AddCommand(dataExchangeCMD)
}
