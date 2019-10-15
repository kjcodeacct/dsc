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
	"dsc/fancy_errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a working dsc directory.",
	Long: `Init will initialize and create a working .dsc subdirectory to store compressed commits,
logging information, history, server configurations, etc.
`,
	Run: func(cmd *cobra.Command, args []string) {
		InitializeWorkingDirectory()
	},
}

var internalPackageDir packr.Box

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().String("dir", "", "specify working directory ")

	internalPackageDir := packr.NewBox("../templates")

}

func InitializeWorkingDirectory() {

	workingDir, _ := initCmd.Flags().GetString("dir")
	if workingDir == "" {
		localDir, err := os.Getwd()
		if err != nil {
			fmt.Fatalln(err)
		}

		workingDir = localDir
	}

	err := createWorkingDirectory(workingDir)
	if err != nil {
		fmt.Fatalln()
	}

	fmt.Printf("initialized empty dsc working directory in %s", workingDir)
}

func createWorkingDirectory(workingDir string) error {

	_, err := os.Stat(workingDir)
	if os.IsNotExist(err) {
		os.Mkdir(workingDir, 0644)
		// copy dsc.db from internalPackageDir
	} else {
		return fancy_errors.New("dsc working directory already exsits")
	}

}
