/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/austingw/reqord/db"
	"github.com/spf13/cobra"
)

// projectCmd represents the project command
var projectCmd = &cobra.Command{
	Use:   "project",
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

		// check for existing project name
		proj, err := queries.GetProject(ctx, args[0])

		// if there's no project we make a new one and set it to current
		if errors.Is(err, sql.ErrNoRows) {
			newProj, err := queries.CreateProject(ctx, args[0])
			if err != nil {
				fmt.Println(err)
				return
			}
			err = queries.SetSelectedProject(ctx, newProj.ID)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("Successfully created and switched to new project: %s", newProj.Name)
			return
		} else if err != nil {
			fmt.Println(err)
			return
		}

		// otherwise we just switch :)
		err = queries.SetSelectedProject(ctx, proj.ID)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Successfully switched to project: %s", proj.Name)
	},
}

func init() {
	rootCmd.AddCommand(projectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// projectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// projectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
