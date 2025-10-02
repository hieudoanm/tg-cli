// Package cmd ...
package cmd

import (
	"encoding/json"
	"fmt"
	"tg/utils"

	"github.com/spf13/cobra"
)

// GetInfoResponse ...
type GetInfoResponse struct {
	Ok bool `json:"ok"`
}

// webhookInfoCmd represents the telegramWebhookGetInfo command
var webhookInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.LogProgramName()
		// Get Telegram Token
		fmt.Print("Telegram Token: ")
		var token string
		fmt.Scanln(&token)
		// Get Webhook Info
		var url string = fmt.Sprintf("https://api.telegram.org/bot%s/getWebhookInfo", token)
		responseByte, postError := utils.Post(url, utils.Options{})
		if postError != nil {
			fmt.Println("Error: ", postError)
			return
		}
		// Parse response
		var getInfoResponse GetInfoResponse
		jsonError := json.Unmarshal(responseByte, &getInfoResponse)
		if jsonError != nil {
			fmt.Println("Error: ", jsonError)
			return
		}
		fmt.Println("Success")
	},
}

func init() {
	webhookCmd.AddCommand(webhookInfoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// webhookInfoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// webhookInfoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
