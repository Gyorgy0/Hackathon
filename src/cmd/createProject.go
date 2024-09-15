/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// createProjectCmd represents the createProject command
var createProjectCmd = &cobra.Command{
	Use:   "create-project",
	Short: "A brief description of your command",
	Long:  `A command that makes a project`,
	Run: func(cmd *cobra.Command, args []string) {
		viper.ReadInConfig()
		println("asd" + viper.GetString("APIKEY"))
	},
}

func init() {
	rootCmd.AddCommand(createProjectCmd)
	/*createProjectCmd.Flags().String("name", "", "specifies the name of the project")
	createProjectCmd.Flags().String("tags", "", "specifies the tags associated with the project")
	createProjectCmd.Flags().String("description", "", "specifies the description of the project")
	createProjectCmd.Flags().Bool("default", false, "specifies whether the project is the default project or not")*/

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createProjectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createProjectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
