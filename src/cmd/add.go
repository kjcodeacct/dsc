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
	errors "dsc/fancy_errors"
	"dsc/printer"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a file or files to be commited.",
	Long: `'add' will add file(s) to be tracked and commited by dsc.
Please see the docs for more information on adding files in order to apply changes correctly.
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)

		if len(args) > 1 {
			// TODO prompt user to correctly order files
			for _, filePath := range args {
				err := addFile(filePath)
				if err != nil {
					printer.Fatalln(errors.Wrap(err).Error())
				}
			}
		} else {
			for _, filePath := range args {
				err := addFile(filePath)
				if err != nil {
					printer.Fatalln(errors.Wrap(err).Error())
				}
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func addFile(filePath string) error {

	err := checkFile(filePath)
	if err != nil {
		return err
	}

	return nil
}

func checkFile(filePath string) error {

	fileName := filePath

	_, err := os.Stat(filePath)
	if err != nil {

		if os.IsNotExist(err) {

			cwd, err := os.Getwd()
			if err != nil {
				return fancy_errors.Wrap(err)
			}

			filePath = filepath.Join(cwd, filePath)

			_, err = os.Stat(filePath)
			if err != nil {

				if os.IsNotExist(err) {
					errMsg := fmt.Sprintf("file '%s' does not exist", fileName)
					return fancy_errors.New(errMsg)
				} else {
					return fancy_errors.Wrap(err)
				}

			}

		} else {
			return fancy_errors.Wrap(err)
		}
	}

	return nil
}
