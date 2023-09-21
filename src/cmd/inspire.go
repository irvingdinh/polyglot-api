package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var inspireCmd = &cobra.Command{
	Use:   "inspire",
	Short: "Lorem ipsum dolor sit amet",
	Run: func(cmd *cobra.Command, args []string) {
		//ctx := context.Background()
		//uri := "mongodb://127.0.0.1:27017"
		//
		//client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
		//if err != nil {
		//	log.Fatalf("Failed to connect to MongoDB with error: %s", err.Error())
		//}
		//
		//defer func() {
		//	if err := client.Disconnect(ctx); err != nil {
		//		log.Fatalf("Failed to disconnect from MongoDB with error: %s", err.Error())
		//	}
		//}()
		//
		//repositoryObj := repository.New(client)
		//
		//agencyId, _ := primitive.ObjectIDFromHex("6507ff6ccdaea839a0407471")
		//
		//repositoryObj.UserRepository().InsertOne(ctx, model.User{
		//	AgencyID: agencyId,
		//	Email:    "john.doe@example.com",
		//})

		fmt.Println("Lorem ipsum dolor sit amet")
	},
}
