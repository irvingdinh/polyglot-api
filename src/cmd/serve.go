package cmd

import (
	"context"
	"github.com/irvingdinh/polyglot-api/src/repository"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"

	"github.com/spf13/cobra"

	"github.com/irvingdinh/polyglot-api/src/http/handler"
	"github.com/irvingdinh/polyglot-api/src/http/server"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start HTTP server on the predefined port",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		uri := "mongodb://127.0.0.1:27017"

		client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
		if err != nil {
			log.Fatalf("Failed to connect to MongoDB with error: %s", err.Error())
		}

		defer func() {
			if err := client.Disconnect(ctx); err != nil {
				log.Fatalf("Failed to disconnect from MongoDB with error: %s", err.Error())
			}
		}()

		repositoryObj := repository.New(client)

		serverObj := server.New(
			handler.New(repositoryObj),
		)

		if err := serverObj.Start(); err != nil {
			log.Fatalf("Failed to start HTTP server with error: %s", err.Error())
		}
	},
}
