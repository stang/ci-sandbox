package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/stang/ci-sandbox/internal/lib"
)

var (
	version string
	lang    string
	rootCmd = &cobra.Command{
		Use:   "greet [name]",
		Short: "A simple tool to give a sign of welcome when meeting someone",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var name string
			if len(args) == 1 {
				name = args[0]
			}
			greetings, err := lib.Greetings(lang, name)
			if err != nil {
				panic(err)
			}
			fmt.Println(greetings)
		},
	}
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("greeting %s\n", version)
		},
	}
)

func init() {
	rootCmd.Flags().StringVarP(&lang, "lang", "l", "en", "Language to use for the greeting")
	rootCmd.AddCommand(versionCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
