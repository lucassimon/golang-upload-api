/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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
	"fmt"
	"log"

	"github.com/lucassimon/golang-upload-api/internal/adapters/http/chii"
	"github.com/spf13/cobra"
)

// httpCmd represents the server command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "starts http server",
	Long:  `starts http server with chi router`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("starts create a webserver")
		server := chii.MakeNewWebserver()
		server.Serve()
		fmt.Println("Server started")

	},
}

func init() {
	rootCmd.AddCommand(httpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// httpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// httpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
