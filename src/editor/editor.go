package editor

import (
	"dsc/config"
	"dsc/fancy_errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
)

const DefaultEditor = "vim"

func OpenInEditor(fileDir string) error {

	userConfig, err := config.Get()
	if err != nil {
		return err
	}

	if runtime.GOOS == "windows" {

		var editor string
		if userConfig.Editor != "" {
			editor = userConfig.Editor
		} else {
			editor = os.Getenv("EDITOR")
			if editor == "" {
				editor = "Notepad"
			}
		}

		runCmd := fmt.Sprintf(`%s "%s"`, editor, fileDir)
		cmd := exec.Command(runCmd, fileDir)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err = cmd.Run()
		if err != nil {
			return fancy_errors.Wrap(err)
		}

	} else {

		var editor string
		if userConfig.Editor != "" {
			editor = userConfig.Editor
		} else {
			editor = os.Getenv("EDITOR")
			if editor == "" {
				editor = DefaultEditor
			} else {
				execPath, err := exec.LookPath(editor)
				if err != nil {
					return fancy_errors.Wrap(err)
				}

				editor = execPath
			}

		}

		cmd := exec.Command(editor, fileDir)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err = cmd.Run()
		if err != nil {
			return fancy_errors.Wrap(err)
		}
	}

	return nil
}

func GetEditorInput() ([]byte, error) {

	data := []byte{}

	file, err := ioutil.TempFile(os.TempDir(), "*")
	if err != nil {
		return data, fancy_errors.Wrap(err)
	}

	filename := file.Name()

	// wait to remove file at completion of the read below
	defer os.Remove(filename)

	err = file.Close()
	if err != nil {
		return data, fancy_errors.Wrap(err)
	}

	err = OpenInEditor(filename)
	if err != nil {
		return data, fancy_errors.Wrap(err)
	}

	data, err = ioutil.ReadFile(filename)
	if err != nil {
		return data, fancy_errors.Wrap(err)
	}

	return data, nil
}

// EditTemplate
// This takes in a templated data file (cfg, sql, yml, etc) temporarily saves it, allows the user
// 		to edit, and then returns the file with the edited changes
func EditTemplate(data []byte) ([]byte, error) {

	file, err := ioutil.TempFile(os.TempDir(), "dsc_")
	if err != nil {
		return data, fancy_errors.Wrap(err)
	}

	filename := file.Name()

	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return data, fancy_errors.Wrap(err)
	}

	// wait to remove file at completion of the read below
	defer os.Remove(filename)

	err = file.Close()
	if err != nil {
		return data, fancy_errors.Wrap(err)
	}

	err = OpenInEditor(filename)
	if err != nil {
		return data, fancy_errors.Wrap(err)
	}

	data, err = ioutil.ReadFile(filename)
	if err != nil {
		return data, fancy_errors.Wrap(err)
	}

	return data, nil
}
