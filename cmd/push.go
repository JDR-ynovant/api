package cmd

import (
	"github.com/JDR-ynovant/api/internal/repository"
	"github.com/JDR-ynovant/api/internal/service"
	"github.com/spf13/cobra"
	"log"
)

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Send push notification to given user id",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			log.Fatal("missing user id")
		}

		executePushCommand(args[0])
	},
}

func executePushCommand(userId string) {
	ur := repository.NewUserRepository()
	user, err := ur.FindOneById(userId)
	if err != nil {
		log.Fatal(err)
	}

	if user.Subscription.Endpoint == "" {
		log.Fatal("missing user subscription")
	}

	resp, err := service.SendNotification(&user.Subscription)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Status)
}

func init() {
	rootCmd.AddCommand(pushCmd)
}