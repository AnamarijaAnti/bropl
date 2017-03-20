package emulatormanager

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/bitrise-io/depman/pathutil"
	"github.com/bitrise-io/go-utils/command"
)

// Model ...
type Model struct {
	binPth string
}

// New ...
func New(androidHome string) (*Model, error) {
	binPth := filepath.Join(androidHome, "tools", "emulator")
	if runtime.GOOS == "linux" {
		binPth = filepath.Join(androidHome, "tools", "emulator64-arm")
	}

	if exist, err := pathutil.IsPathExists(binPth); err != nil {
		return nil, fmt.Errorf("failed to check if emulator exist, error: %s", err)
	} else if !exist {
		return nil, fmt.Errorf("emulator not exist at: %s", binPth)
	}

	return &Model{
		binPth: binPth,
	}, nil
}

// StartEmulatorCommand ...
func (model Model) StartEmulatorCommand(name, skin string, options ...string) *command.Model {
	args := []string{model.binPth, "-avd", name}
	
	if len(skin) == 0 {
		args = append(args, "-noskin")
	} else {
		args = append(args, "-skin", skin)
	}

	args = append(args, options...)

	return command.New(args[0], args[1:]...)
}
