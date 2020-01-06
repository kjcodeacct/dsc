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
	"dsc/config"
	"dsc/editor"
	errors "dsc/fancy_errors"
	"dsc/printer"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a file or files to be staged for commitment.",
	Long: `'add' will stage file(s) to be tracked and commited by dsc.
Please see the docs for more information on adding files in order to apply changes correctly.
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
		files, err := parseInput(args)
		if err != nil {
			printer.Fatalln(err.Error())
		}

		err = add(files)
		if err != nil {
			printer.Fatalln(err.Error())
		}

	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	workingDir, _ = config.FindWorkingDir()
}

func parseInput(argList []string) ([]string, error) {

	var fileDirList []string

	for _, arg := range argList {

		checkDir := arg

		if !strings.Contains(arg, workingDir) {
			checkDir = filepath.Join(workingDir, checkDir)
		}

		_, err := os.Stat(checkDir)
		if err != nil {

			if os.IsNotExist(err) {
				errMsg := fmt.Sprintf("%s is not in working sub directory", arg)
				return nil, errors.New(errMsg)
			}

			return nil, err
		}

		fileDirList = append(fileDirList, arg)
	}

	return fileDirList, nil
}

func add(fileDirList []string) error {

	addFilename := filepath.Join(workingDir, ".add")
	var addFile *os.File
	var err error

	addFile, err = os.OpenFile(addFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	for _, fileDir := range fileDirList {

		_, err = io.WriteString(addFile, fileDir)
		if err != nil {
			return err
		}
	}

	err = addFile.Sync()
	if err != nil {
		return err
	}

	err = addFile.Close()
	if err != nil {
		return err
	}

	if len(fileDirList) > 1 {
		fmt.Println("opening staged file list for re ordering")
		err = editor.OpenInEditor(addFilename)
		if err != nil {
			return errors.Wrap(err)
		}
	}

	return nil
}
