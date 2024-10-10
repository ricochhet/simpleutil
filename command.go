package simpleutil

import (
	"errors"
	"fmt"
	"slices"
	"strings"
)

const BinSize int = 2

var ErrNoFunctionName = errors.New("no function name")

func NewCommand(args []string, name string, argCount int) ([]string, error) {
	var nameArgs []string

	if slices.Contains(args, name) {
		ind := slices.Index(args, name)

		if ind != -1 {
			if len(args) >= argCount+BinSize {
				nameArgs = append(nameArgs, args[ind+1:]...)
				return nameArgs, nil
			}

			return nil, errExpectedArgs(argCount, len(args)-BinSize)
		}

		return nil, fmt.Errorf("%s not found", name) //nolint:err113 // wontfix
	}

	return nil, ErrNoFunctionName
}

func SplitArguments(input string) []string {
	var parts []string

	var part strings.Builder

	inQuote := false

	for _, char := range input {
		switch {
		case char == '"':
			inQuote = !inQuote
		case char == ' ' && !inQuote:
			parts = append(parts, part.String())
			part.Reset()
		default:
			part.WriteRune(char)
		}
	}

	parts = append(parts, part.String())

	return parts
}

func CheckArgumentCount(args []string, expected int) error {
	if len(args) != expected {
		return errExpectedArgs(expected, len(args))
	}

	return nil
}

func errExpectedArgs(expected, got int) error {
	return fmt.Errorf("expected: %d arguments but got %d", expected, got) //nolint:err113 // wontfix
}
