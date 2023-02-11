/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/m4salah/tri/todo"
	"github.com/spf13/cobra"
)

func runList(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(datafile)
	if err != nil {
		log.Printf("%v", err)
	}
	for index, item := range items {
		fmt.Printf("%d. %s\n", index+1, item.Text)
	}
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all todo items",
	Long:  ``,
	Run:   runList,
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
}
