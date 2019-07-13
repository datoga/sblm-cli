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

var newLevel string
var loggers string

// editCmd represents the set command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "It edits the logger level of a logger or set of loggers",
	Long:  `It allows to edit the logger level of different loggers`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("set called")

		n, err := client.EditLoggers(loggers, newLevel)

		if err != nil {
			fmt.Println(err)
		}

		if n == 0 {
			fmt.Println("None logger has been modified with name", loggers)
		} else if n == 1 {
			fmt.Println("Logger", loggers, "has been set with level", newLevel, "successfully")
		} else {
			fmt.Println("Logger group", loggers, "has been set with level", newLevel, "successfully")
		}
	},
}

func init() {
	rootCmd.AddCommand(editCmd)

	// Here you will define your flags and configuration settings.
	editCmd.PersistentFlags().StringVarP(&newLevel, "level", "l", "INFO", "New level to set the logger or group of loggers")

	editCmd.MarkFlagRequired("level")

	editCmd.PersistentFlags().StringVarP(&loggers, "name", "n", "ROOT", "Name of the logger or group of loggers to apply the new level configuration")

	editCmd.MarkFlagRequired("logger")

}
