/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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
  "os"

  "github.com/spf13/cobra"

  homedir "github.com/mitchellh/go-homedir"
  "github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
  Use:   "dsc",
  Short: "dsc (database schema control)",
  Long: `
usage: dsc [--version][--help][--webpage][--doc] [<command> [<args>]

These are common dsc commands:

Start a working directory:
    init      Create a working directory for dsc.

Work on a scema change:
    add       Add file to be tracked by local dsc index.
    rm        Remove file from local dsc index.

Examine history and state of schema changes:
    log       Show previously applied changes.
    status    Show working status.
    diff      Show diff between two databases.

Collaborate:
    pull      Fetch a copy of a schema from a remote database.
    populate  Fetch a copy of a schema from a remote database and it's entries.
    push      Apply schema to a remote database.
    dryrun    Test schema changes against a remote database (in an uncommited transaction).
`,
  // Uncomment the following line if your bare application
  // has an action associated with it:
  //  Run: func(cmd *cobra.Command, args []string) { },
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

  rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.dsc.yaml)")

  // Cobra also supports local flags, which will only run
  // when this action is called directly.
  rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
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

    // Search config in home directory with name ".dsc" (without extension).
    viper.AddConfigPath(home)
    viper.SetConfigName(".dsc")
  }

  viper.AutomaticEnv() // read in environment variables that match

  // If a config file is found, read it in.
  if err := viper.ReadInConfig(); err == nil {
    fmt.Println("Using config file:", viper.ConfigFileUsed())
  }
}
