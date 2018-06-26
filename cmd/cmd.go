package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gimage",
	Short: "Gimage downloads images for you",
	Long:  `A image downloader that uses google images to quickly download lightweight, low resolution images for you, written in Go.`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, a := range args {
			fmt.Println(a)
		}
	},
}

//Execute start CLi
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
