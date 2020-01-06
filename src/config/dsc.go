package config

import (
	errors "dsc/fancy_errors"
	"dsc/printer"
	"os"
	"path/filepath"
	"strings"
)

const WorkingDirDefault = ".dsc"

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
