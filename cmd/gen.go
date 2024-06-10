/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/saucon/sulaptomap/converter"
	"strings"

	"github.com/spf13/cobra"
)

var structName string
var sourcePath string
var option string

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var opt converter.Options
		switch option {
		default:
			opt = converter.DefaultOptions
		}

		// split struct name if more than one
		structNames := strings.Split(structName, ",")
		for _, name := range structNames {
			err := converter.ConvertToMap(name, sourcePath, &opt)
			if err != nil {
				return
			}
		}

		// if struct name is only one
		if len(structNames) == 0 {
			err := converter.ConvertToMap(structName, sourcePath, &opt)
			if err != nil {
				return
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(genCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	genCmd.Flags().StringVarP(&structName, "struct", "s", "", "Struct name in go. Comma separated if more than one")
	err := rootCmd.MarkFlagRequired("struct")

	genCmd.Flags().StringVarP(&sourcePath, "path", "p", "", "Path to your struct file")
	err = rootCmd.MarkFlagRequired("path")

	genCmd.Flags().StringVarP(&option, "option", "o", "", "Option, if not set will be default")

	_ = err
}
