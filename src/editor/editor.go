package editor

import (
	"dsc/fancy_errors"
	"io/ioutil"
	"os"
	"os/exec"
)

const DefaultEditor = "vim"

func OpenInEditor(filename string) error {

	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = DefaultEditor
	}

	executable, err := exec.LookPath(editor)
	if err != nil {
		return err
	}

	cmd := exec.Command(executable, filename)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
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

	file, err := ioutil.TempFile(os.TempDir(), "*")
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
