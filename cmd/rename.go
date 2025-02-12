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

// TODO: Add flag for renaming project? Probably --proj -p?

// renameCmd represents the rename command
var renameCmd = &cobra.Command{
	Use:   "rename",
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
			Name:      args[0],
			ProjectID: 1,
		})
		if err != nil {
			if err.Error() == "sql: no rows in result set" {
				fmt.Printf("Command %s not found\n", args[0])
			}
			return
		}
		err = queries.UpdateRequest(ctx, db.UpdateRequestParams{
			ID:   request.ID,
			Name: args[1],
		})
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Sucessfully renamed %s to %s\n", args[0], args[1])
	},
}

func init() {
	rootCmd.AddCommand(renameCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// renameCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// renameCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
