package shell

import (
	"fmt"
	"path/filepath"
)

// Shell is a enum to define command-line interpreter.
type Shell int

const (
	Sh         Shell = 1 << iota // https://en.wikipedia.org/wiki/Bourne_shell
	Bash                         // https://en.wikipedia.org/wiki/Bash_(Unix_shell)
	Zsh                          // https://en.wikipedia.org/wiki/Z_shell
	PowerShell                   // https://en.wikipedia.org/wiki/PowerShell

	Unknown = "unknown"
)

// String returns the string representation of the enum.
func (sh Shell) String() string {
	switch sh {
	case Sh:
		return "sh"
	case Bash:
		return "bash"
	case Zsh:
		return "zsh"
	case PowerShell:
		return "powershell"
	}
	return Unknown
}

// Operation is a enum to define type of command-line interpreter operation.
type Operation int

const (
	Assign Operation = 1 << iota
	Completion
	Exec
	Print

	All = Assign | Completion | Exec | Print
)

// Classify tries to classify passed binary as a command-line interpreter.
func Classify(bin string, operations ...Operation) (sh Shell, err error) {
	if bin == "" {
		panic("shell: cannot classify shell by empty binary name")
	}

	// naive classification
	switch filepath.Base(bin) {
	case "sh":
		sh = Sh
	case "bash":
		sh = Bash
	case "zsh":
		sh = Zsh
	case "powershell", "powershell.exe", "pwsh.exe":
		sh = PowerShell
	default:
		err = fmt.Errorf("shell: cannot classify shell by %q", bin)
	}

	var op Operation
	for _, operation := range operations {
		op |= operation
	}
	// matrix classification

	return
}
