/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var partitionsCmd = &cobra.Command{
	Use:   "partitions",
	Short: "list kafka partitions",
	Long:  `List partition details for available topics from configured kafka brokers `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("partitions command ,  topics=", args)

		partitions, err := conn.ReadPartitions(args...)
		if err != nil {
			panic(err.Error())
		}

		m := map[string]struct{}{}

		for _, p := range partitions {
			m[p.Topic] = struct{}{}
		}
		for k := range m {
			fmt.Println(k)
		}
	},
}

func init() {
	rootCmd.AddCommand(partitionsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// topicsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// topicsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
