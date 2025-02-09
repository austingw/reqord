/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/austingw/reqord/db"
	"github.com/spf13/cobra"
)

// playCmd represents the play command
var playCmd = &cobra.Command{
	Use:     "play",
	Aliases: []string{"run"},
	Short:   "A brief description of your command",
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
		fmt.Println(request)
		fields := strings.Fields(request.Curl)
		curlCmd := exec.Command(fields[0], fields[1:]...)
		curlCmd.Stdout = os.Stdout
		curlCmd.Stderr = os.Stderr
		err = curlCmd.Run()
		if err != nil {
			fmt.Println("Curl request failed with %s\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(playCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// playCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// playCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
