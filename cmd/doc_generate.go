/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
//	"io/ioutil"

	"github.com/spf13/cobra"
	log "github.com/sirupsen/logrus"
)

// docGenerateCmd represents the generate command
var docGenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate project documentation",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Generating documentation")
		genRoutesDoc()
	},
}

func init() {
	docCmd.AddCommand(docGenerateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// docGenerateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// docGenerateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func genRoutesDoc() {
	log.Info("Generating routes markdown file: ")
	/*if err := ioutil.WriteFile("routes.md", []byte(md), 0644); err != nil {
		log.Println(err)
		return
	}*/
	log.Info("Documentation generated")
}
