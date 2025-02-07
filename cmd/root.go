/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"pokeShowcase-api/configs"
	"pokeShowcase-api/internal/http"

	"github.com/spf13/cobra"
)

var logo = `
Ordine API
=================================================================================================================================================================
                      $$\                  $$$$$$\  $$\                                                                                                 $$\ 
                      $$ |                $$  __$$\ $$ |                                                                                                \__|
   $$$$$$\   $$$$$$\  $$ |  $$\  $$$$$$\  $$ /  \__|$$$$$$$\   $$$$$$\  $$\  $$\  $$\  $$$$$$$\ $$$$$$\   $$$$$$$\  $$$$$$\          $$$$$$\   $$$$$$\  $$\ 
  $$  __$$\ $$  __$$\ $$ | $$  |$$  __$$\ \$$$$$$\  $$  __$$\ $$  __$$\ $$ | $$ | $$ |$$  _____|\____$$\ $$  _____|$$  __$$\ $$$$$$\ \____$$\ $$  __$$\ $$ |
  $$ /  $$ |$$ /  $$ |$$$$$$  / $$$$$$$$ | \____$$\ $$ |  $$ |$$ /  $$ |$$ | $$ | $$ |$$ /      $$$$$$$ |\$$$$$$\  $$$$$$$$ |\______|$$$$$$$ |$$ /  $$ |$$ |
  $$ |  $$ |$$ |  $$ |$$  _$$<  $$   ____|$$\   $$ |$$ |  $$ |$$ |  $$ |$$ | $$ | $$ |$$ |     $$  __$$ | \____$$\ $$   ____|       $$  __$$ |$$ |  $$ |$$ |
  $$$$$$$  |\$$$$$$  |$$ | \$$\ \$$$$$$$\ \$$$$$$  |$$ |  $$ |\$$$$$$  |\$$$$$\$$$$  |\$$$$$$$\\$$$$$$$ |$$$$$$$  |\$$$$$$$\        \$$$$$$$ |$$$$$$$  |$$ |
  $$  ____/  \______/ \__|  \__| \_______| \______/ \__|  \__| \______/  \_____\____/  \_______|\_______|\_______/  \_______|        \_______|$$  ____/ \__|
  $$ |                                                                                                                                        $$ |          
  $$ |                                                                                                                                        $$ |          
  \__|                                                                                                                                        \__|                                                                   
=================================================================================================================================================================
🚀 Powered by brasilis for Monsters! 🚀
`

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pokeShowcase-api",
	Short: "A brief description of your application",
	Long:  logo + `This application default action is start the api`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(logo+"Service http started, listening port: %v and log_level: %v\n", configs.Envs.PORT, configs.Envs.LOG_LEVEL)
		http.GetServer().Run()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
