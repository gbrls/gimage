package cmd

import (
	"fmt"
	"time"

	"github.com/gabrielSchneider100/gimage/img"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Downloads images",
	Long:  `It does what you expect.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("please indicate your search term and the folder")
		} else if len(args) < 2 {
			return fmt.Errorf("please indicate the destination folder")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		past := time.Now()

		img.GetImagesURLs(args[0], args[1])
		fmt.Printf("Downloaded your images in %v, with tag: %v\n", time.Since(past), args[0])

	},
}
