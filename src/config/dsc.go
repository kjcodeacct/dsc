package config

import (
	"dsc/editor"
	errors "dsc/fancy_errors"
	"dsc/printer"
	"os"
	"path/filepath"
	"strings"
)

const WorkingDirDefault = ".dsc"
const DefaultAlias = "default"

var ValidDbTypes = make(map[string][]string)

func FindWorkingDir() (string, error) {

	cwd, err := os.Getwd()
	if err != nil {
		return "", errors.Wrap(err)
	}

	workingDir := filepath.Join(cwd, WorkingDirDefault)
	_, err = os.Stat(workingDir)
	if err != nil {
		if os.IsNotExist(err) {
			numParentDirs := strings.Count(cwd, string(os.PathSeparator))

			var parentDir string

			for i := numParentDirs; i < 0; i-- {

				checkParentDir := filepath.Dir(parentDir)
				workingDir := filepath.Join(checkParentDir, WorkingDirDefault)
				_, err := os.Stat(workingDir)
				if err != nil {
					if os.IsNotExist(err) {
						parentDir = checkParentDir
						continue
					} else {
						printer.Println(err.Error())
						return "", errors.New("not a dsc repository in this or any parent directories: .dsc")
					}
				}

				return workingDir, nil
			}

		} else {
			printer.Println(err.Error())
			return "", errors.New("not a dsc repository in this or any parent directories: .dsc")
		}
	}

	return workingDir, nil
}

// func checkParentDir(dir string) bool {}

func CreateWorkingDir(workingDir string) error {

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

func PromptCreateRemote() error {

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

func AddAlias(workingDir string) error {

}
