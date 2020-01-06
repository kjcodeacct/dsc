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
	"dsc/editor"
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

	checkWorkingDir, _ := initCmd.Flags().GetString("dir")

	if checkWorkingDir != "" {
		workingDir, _ = initCmd.Flags().GetString("dir")
	} else {

		cwd, err := os.Getwd()
		if err != nil {
			printer.Fatalln(errors.Wrap(err).Error())
		}

		workingDir = cwd
	}
}

func InitializeWorkingDirectory() {

	if workingDir == "" {
		localDir, err := os.Getwd()
		if err != nil {
			printer.Fatalln(errors.Wrap(err).Error())
		}

		workingDir = localDir
	}

	workingDir = filepath.Join(workingDir, ".dsc")

	err := createWorkingDirectory(workingDir)
	if err != nil {
		printer.Fatalln(errors.Wrap(err).Error())
	}

	printer.Println("initialized empty dsc working directory in %s", workingDir)

	setupConfig := editor.PromptBool("do you want to setup a remote host")

	if setupConfig {
	}
}

const DefaultAlias = "default"

func createAliases() error {

	defaultSet := false

	// aliasList := []config.Alias
AliasPrompt:
	for {

		host := editor.Prompt("hostname")
		port := editor.Prompt("port")

		var isDefault bool
		var alias string

		if !defaultSet {

			isDefault = editor.PromptBool("is this your default host")
		}

		if !isDefault {

			alias = editor.Prompt("remote alias")
			if alias == DefaultAlias && defaultSet {
				printer.Red("default remote alias is already set, please reconfigure this host alias")
				continue AliasPrompt
			}

		} else {

			alias = DefaultAlias
		}

	}

}

func createWorkingDirectory(workingDir string) error {

	_, err := os.Stat(workingDir)
	if os.IsNotExist(err) {

		os.Mkdir(workingDir, 0644)
		// copy dsc.db from internalPackageDir
		_, err := os.Create("index.db")
		if err != nil {
			printer.Fatalln(errors.Wrap(err).Error())
		}

	} else {

		if err != nil {
			return errors.Wrap(err)
		}

		return errors.New("dsc working directory already exsits")
	}

	return nil
}
