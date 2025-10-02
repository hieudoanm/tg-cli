// Package cmd ...
package cmd

import (
	"fmt"
	"tg/utils"

	"github.com/spf13/cobra"
)

// webhookCmd represents the telegramWebhook command
var webhookCmd = &cobra.Command{
	Use:   "webhook",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.LogProgramName()
		// Get subcommands
		fmt.Println("delete   - Delete Webhook")
		fmt.Println("get-info - Get Info of Webhook")
		fmt.Println("set      - Set Webhook")
	},
}

func init() {
	rootCmd.AddCommand(webhookCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// webhookCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// webhookCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
