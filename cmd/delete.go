/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/austingw/reqord/db"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		queries, err := db.GetQueries()
		if err != nil {
			fmt.Println(err)
			return
		}
		request, err := queries.GetRequest(ctx, db.GetRequestParams{
			ProjectID: 1,
			Name:      args[0],
		})
		if err != nil {
			fmt.Println(err)
			return
		}
		err = queries.DeleteRequest(ctx, request.ID)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Succesfully delete request:", request.Name)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
