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
	"net/url"
	"os"
	"strings"

	"github.com/datoga/sblm_cli/core"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var verbose bool
var server string
var app string
var pattern string
var actuatorURL *url.URL
var client *core.Client

var log = logrus.New()

func init() {
	log.Out = os.Stdout
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sblm_cli",
	Short: "CLI to manage Spring Boot loggers",
	Long:  `This CLI allows to users to list or edit the Spring Boot logger from any online Spring Boot application.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if verbose {
			log.SetLevel(logrus.TraceLevel)
		} else {
			log.SetLevel(logrus.ErrorLevel)
		}

		log.Debug("Inside rootCmd PreRun with args ", args)

		log.Debug("Server: ", server)

		log.Debug("Application: ", app)

		log.Debug("Pattern: ", pattern)

		patternS := strings.Replace(pattern, "$server", server, 1)
		patternSA := strings.Replace(patternS, "$app", app, 1)

		var err error

		actuatorURL, err = url.Parse(patternSA)

		if err != nil {
			panic("Error decoding format for url " + patternSA + "Error: " + err.Error())
		}

		log.Debug("Actuator URL: ", actuatorURL)

		client = core.NewClient(actuatorURL, verbose)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.sblm_cli.yaml)")

	rootCmd.PersistentFlags().StringVarP(&server, "server", "s", "http://localhost:8080", "Server to get information")

	rootCmd.PersistentFlags().StringVarP(&app, "app", "a", "HelloWorld", "Application to get information")

	rootCmd.PersistentFlags().StringVarP(&pattern, "pattern", "p", "$server/$app/actuator/loggers", "Pattern to be applied")

	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".sblm_cli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".sblm_cli")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Debug("Using config file:", viper.ConfigFileUsed())
	}
}
