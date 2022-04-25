/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// convertCmd represents the convert command
var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert one format of adventure json to another format",
	Long: `A utility that accepts an input json file and format and converts it to a destination json file and format.

The formats can be the same (i.e. 5etools to 5etools) or different (i.e. 5etools to eplus).

Valid formats:
* 5etools
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		source, _ := cmd.Flags().GetString("source")
		// format, _ := cmd.Flags().GetString("format")
		destination, _ := cmd.Flags().GetString("destination")
		// target, _ := cmd.Flags().GetString("target")

		content, err := readFile(source)

		if err != nil {
			fmt.Printf("err: %v\n", err)
		}

		writeFile(destination, content)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(convertCmd)

	convertCmd.Flags().StringP("source", "s", "", "path to the input file")
	convertCmd.Flags().StringP("format", "f", "5etools", "format of the input file")
	convertCmd.Flags().StringP("destination", "d", "", "path to the destination file")
	convertCmd.Flags().StringP("target", "t", "5etools", "format of the destination file")

	convertCmd.MarkFlagRequired("source")
	// convertCmd.MarkFlagRequired("format")
	convertCmd.MarkFlagRequired("destination")
	// convertCmd.MarkFlagRequired("target")
}

func readFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func writeFile(path string, data []byte) error {
	os.Create(path)
	return os.WriteFile(path, data, 644)
}
