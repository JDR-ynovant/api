package cmd

import (
	"github.com/SherClockHolmes/webpush-go"
	"github.com/spf13/cobra"
	"log"
)

var vapidCmd = &cobra.Command{
	Use:   "generate-vapid",
	Short: "Generate Vapid keypair",
	Run: func(cmd *cobra.Command, args []string) {
		executeVapidGenerateCommand()
	},
}

func executeVapidGenerateCommand() {
	privateKey, publicKey, err := webpush.GenerateVAPIDKeys()

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("Private Key : %s\n", privateKey)
	log.Printf("Public Key : %s\n", publicKey)
}

func init() {
	rootCmd.AddCommand(vapidCmd)
}