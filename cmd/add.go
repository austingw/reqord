/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/austingw/reqord/db"
	utils "github.com/austingw/reqord/internal"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"cut", "save"},
	Short:   "Save a curl request to your current project",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		reqName, reqCurl := args[0], strings.Join(args[1:], " ")
		ctx := context.Background()

		queries, err := db.GetQueries()
		if err != nil {
			fmt.Println(err)
			return
		}

		currId, err := queries.GetSelectedProject(ctx)
		if err != nil {
			fmt.Println(err)
			return
		}

		// check for existing cmd name
		request, err := queries.GetRequest(ctx, db.GetRequestParams{
			ProjectID: currId,
			Name:      args[0],
		})
		if err == nil {
			fmt.Printf("A request named %s already exists for your current project\n", request.Name)
			return
		}

		parsedReq, err := utils.ParseCurl(reqCurl)
		if err != nil {
			fmt.Println(err)
			return
		}

		insertedReq, err := queries.CreateRequest(ctx, db.CreateRequestParams{
			ProjectID: 1, // will be replaced w/ selected project val
			Name:      reqName,
			Curl:      reqCurl,
			Method:    parsedReq.Method,
			Url:       parsedReq.Url,
			// Headers:   parsedReq.Headers, NOTE: need to fix type issue
			// Body:      parsedReq.Body,
		})
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Sucessfully created request: " + insertedReq.Name)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
