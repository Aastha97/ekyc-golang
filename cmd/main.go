package main

import (
	"os"

	"github.com/spf13/cobra"
)

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "ekyc",
	Short: "e-KYC",
	Long:  `e-KYC to perform face matching between two image and OCR`,
}
