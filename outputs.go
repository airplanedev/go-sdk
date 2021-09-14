package airplane

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/pkg/errors"
)

// Output writes `value` as an Airplane output to stdout. Outputs are
// separated from your logs and used to provide context structured
// context to end-users of your task.
//
// Docs: https://docs.airplane.dev/reference/outputs
func Output(value interface{}) error {
	return NamedOutput("", value)
}

// MustOutput writes `value` to stdout as an Airplane output. Outputs are
// separated from your logs and used to provide context structured
// context to end-users of your task.
//
// If an error is produced, MustOutput will panic.
//
// Docs: https://docs.airplane.dev/reference/outputs
func MustOutput(value interface{}) {
	if err := Output(value); err != nil {
		panic(errors.Wrap(err, "writing output"))
	}
}

// NamedOutput writes `value` to stdout as an Airplane output. Outputs are
// separated from your logs and used to provide context structured
// context to end-users of your task. Unlike Output, NamedOutput accepts
// a `name` which will be used to group separate streams of outputs together.
//
// Docs: https://docs.airplane.dev/reference/outputs
func NamedOutput(name string, value interface{}) error {
	if err := validateOutputName(name); err != nil {
		return err
	}

	header := "airplane_output"
	if name != "" {
		header += fmt.Sprintf(`:"%s"`, name)
	}

	out, err := encodeOutput(value)
	if err != nil {
		return err
	}

	fmt.Printf("%s %s\n", header, out)
	return nil
}

func encodeOutput(value interface{}) (string, error) {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	// We don't want to (or need to) escape e.g. & to \u0026:
	enc.SetEscapeHTML(false)
	if err := enc.Encode(value); err != nil {
		return "", errors.Wrap(err, "marshalling output to JSON")
	}
	// Get rid of trailing newline: https://github.com/golang/go/issues/37083
	buf.Truncate(buf.Len() - 1)
	return buf.String(), nil
}

// MustNamedOutput writes `value` to stdout as an Airplane output. Outputs are
// separated from your logs and used to provide context structured
// context to end-users of your task. Unlike Output, NamedOutput accepts
// a `name` which will be used to group separate streams of outputs together.
//
// If an error is produced, MustNamedOutput will panic.
//
// Docs: https://docs.airplane.dev/reference/outputs
func MustNamedOutput(name string, value interface{}) {
	if err := NamedOutput(name, value); err != nil {
		panic(errors.Wrapf(err, "writing output %s", name))
	}
}

var (
	outputNameRegex      = regexp.MustCompile("^[a-zA-Z_0-9]*$")
	errInvalidOutputName = errors.Errorf("expected name to match %s", outputNameRegex.String())
)

func validateOutputName(name string) error {
	if name != "" && !outputNameRegex.Match([]byte(name)) {
		return errors.Wrapf(errInvalidOutputName, "invalid output name (%s)", name)
	}

	return nil
}
