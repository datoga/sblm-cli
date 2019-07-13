// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var raw bool
var loggerWidth int
var filter string

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all loggers from the APP",
	Long:  `This command requests information about the list of the available loggers and their levels`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("list called")

		loggers, err := client.ListLoggers(filter)

		if err != nil {
			fmt.Println(err)
		}

		if raw {
			fmt.Println(loggers)
		} else {
			loggers.PrettyPrint(loggerWidth)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.
	listCmd.PersistentFlags().StringVarP(&filter, "filter", "f", "", "Filter to be applied")

	listCmd.PersistentFlags().BoolVarP(&raw, "raw", "r", false, "If the response is raw text or pretty")

	listCmd.PersistentFlags().IntVarP(&loggerWidth, "width", "w", 128, "Logger name width")

}
