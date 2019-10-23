/*
Copyright © 2019 Kyle J <kjcodeact+dsc AT gmail DOT com>

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
	errors "dsc/fancy_errors"
	"dsc/printer"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a working dsc directory.",
	Long: `'init' will initialize and create a working .dsc subdirectory to store compressed commits,
logging information, history, server configurations, etc.
`,
	Run: func(cmd *cobra.Command, args []string) {
		InitializeWorkingDirectory()
	},
}

var workingDir string

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().String("dir", "", "specify working directory ")
	workingDir, _ = initCmd.Flags().GetString("dir")
}

func InitializeWorkingDirectory() {

	if workingDir == "" {
		localDir, err := os.Getwd()
		if err != nil {
			printer.Fatalln(errors.Wrap(err))
		}

		workingDir = localDir
	}

	workingDir = filepath.Join(workingDir, ".dsc")

	err := createWorkingDirectory(workingDir)
	if err != nil {
		printer.Fatalln(errors.Wrap(err))
	}

	printer.Println("initialized empty dsc working directory in %s", workingDir)
}

func createWorkingDirectory(workingDir string) error {

	_, err := os.Stat(workingDir)
	if os.IsNotExist(err) {
		os.Mkdir(workingDir, 0644)
		// copy dsc.db from internalPackageDir
		_, err := os.Create("index.db")
		if err != nil {
			printer.Fatalln(errors.Wrap(err))
		}
	} else {
		return errors.New("dsc working directory already exsits")
	}

	return nil
}
