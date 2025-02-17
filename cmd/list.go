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

var All bool

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(All)
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

		var reqList []db.Request
		if All {
			reqList, err = queries.ListAllRequests(ctx)
		} else {
			reqList, err = queries.ListProjectRequests(ctx, currId)
		}

		for _, r := range reqList {
			// just print for now, we'll add bubble tea component later
			fmt.Println(r)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	listCmd.Flags().BoolVarP(&All, "all", "a", false, "Lists all requests across ALL projects")
}
