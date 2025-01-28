/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"main/pkg/config"
	"main/pkg/server"
	"net/http"

	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the core server module",
	Long: `The server command starts the core server module, which includes
connecting to the database, Minio, and Firebase.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

		log.Println("Starting the server module...")

		log.Println("Reading configuration...")
		configKey, err := config.GetConfig()
		if err != nil {
			panic(err)
		}

		log.Println("Connecting to database")
		database, err := config.DBConfig(configKey.DBHost, configKey.DBUser, configKey.DBPassword, configKey.DBName, configKey.DBPort, configKey.DBTimeZone, configKey.DBSSLMode)
		if err != nil {
			panic(err)
		}

		log.Println("Initializing firebase")
		firebaseClient, err := config.FirebaseConfig()
		if err != nil {
			panic(err)
		}

		log.Println("Initializing minio")
		minioClient, err := config.MinioConfig(configKey.MinioEndpoint, configKey.MinioAccess, configKey.MinioSecret, configKey.MinioBucket, configKey.MinioSSLMode)
		if err != nil {
			panic(err)
		}

		log.Println("Initializing sentry")
		sentryHandler, err := config.SentryConfig(configKey.SentryDSN)
		if err != nil {
			panic(err)
		}

		log.Printf("Starting server at %s", configKey.AppPort)
		server := server.Router(firebaseClient, minioClient, database)
		err = http.ListenAndServe(":"+configKey.AppPort, sentryHandler.Handle(server))
		if err != nil {
			panic(err)
		}

		log.Println("Server module started successfully.")
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
