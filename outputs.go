package airplane

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// SetOutput sets the Airplane task output, writing it to stdout. Outputs are
// separated from your logs and are used to provide structured context to
// end-users of your task.
//
// Docs: https://docs.airplane.dev/reference/outputs
func SetOutput(value interface{}) error {
	return SetOutputWithPath(value, "")
}

// MustSetOutput sets the Airplane task output, writing it to stdout. Outputs
// are separated from your logs and are used to provide structured context to
// end-users of your task.
//
// If an error is produced, MustSetOutput will panic.
//
// Docs: https://docs.airplane.dev/reference/outputs
func MustSetOutput(value interface{}) {
	if err := SetOutput(value); err != nil {
		panic(errors.Wrap(err, "setting output"))
	}
}

// SetOutputWithPath sets part of the Airplane task output at a given JSON
// path, writing it to stdout. Outputs are separated from your logs and are
// used to provide structured context to end-users of your task.
//
// Refer to the docs for a description of the JSON path format.
//
// Docs: https://docs.airplane.dev/reference/outputs
func SetOutputWithPath(value interface{}, path string) error {
	return writeOutputCommandWithPath("airplane_output_set", value, path)
}

// MustSetOutputWithPath sets part of the Airplane task output at a given JSON
// path, writing it to stdout. Outputs are separated from your logs and are
// used to provide structured context to end-users of your task.
//
// If an error is produced, MustSetOutputWithPath will panic.
//
// Refer to the docs for a description of the JSON path format.
//
// Docs: https://docs.airplane.dev/reference/outputs
func MustSetOutputWithPath(value interface{}, path string) {
	if err := SetOutputWithPath(value, path); err != nil {
		panic(errors.Wrapf(err, "setting output %s", path))
	}
}

// AppendOutput appends to the Airplane task output, writing it to stdout.
// Outputs are separated from your logs and are used to provide structured
// context to end-users of your task.
//
// Docs: https://docs.airplane.dev/reference/outputs
func AppendOutput(value interface{}) error {
	return AppendOutputWithPath(value, "")
}

// MustAppendOutput appends to the Airplane task output, writing it to stdout.
// Outputs are separated from your logs and are used to provide structured
// context to end-users of your task.
//
// If an error is produced, MustAppendOutput will panic.
//
// Docs: https://docs.airplane.dev/reference/outputs
func MustAppendOutput(value interface{}) {
	if err := AppendOutput(value); err != nil {
		panic(errors.Wrap(err, "setting output"))
	}
}

// AppendOutputWithPath appends to part of the Airplane task output at a given
// JSON path, writing it to stdout. Outputs are separated from your logs and are
// used to provide structured context to end-users of your task.
//
// Refer to the docs for a description of the JSON path format.
//
// Docs: https://docs.airplane.dev/reference/outputs
func AppendOutputWithPath(value interface{}, path string) error {
	return writeOutputCommandWithPath("airplane_output_append", value, path)
}

// MustAppendOutputWithPath appends to part of the Airplane task output at a
// given JSON path, writing it to stdout. Outputs are separated from your logs
// and are used to provide structured context to end-users of your task.
//
// If an error is produced, MustAppendOutputWithPath will panic.
//
// Refer to the docs for a description of the JSON path format.
//
// Docs: https://docs.airplane.dev/reference/outputs
func MustAppendOutputWithPath(value interface{}, path string) {
	if err := AppendOutputWithPath(value, path); err != nil {
		panic(errors.Wrapf(err, "setting output %s", path))
	}
}

func writeOutput(command string, value interface{}, path string) error {
	header := command
	if path != "" {
		header += fmt.Sprintf(`:%s`, path)
	}

	out, err := encodeOutput(value)
	if err != nil {
		return err
	}

	writeChunkedOutput(fmt.Sprintf("%s %s", header, out))
	return nil
}

func writeChunkedOutput(output string) {
	chunkSize := 8192
	if len(output) <= chunkSize {
		fmt.Printf("%s\n", output)
	} else {
		chunkKey := uuid.NewString()
		for i := 0; i < len(output); i += chunkSize {
			endIdx := i + chunkSize
			if endIdx > len(output) {
				endIdx = len(output)
			}
			fmt.Printf("airplane_chunk:%s %s\n", chunkKey, output[i:endIdx])
		}
		fmt.Printf("airplane_chunk_end:%s\n", chunkKey)
	}
}

// Output writes `value` as an Airplane output to stdout. Outputs are
// separated from your logs and used to provide context structured
// context to end-users of your task.
//
// Deprecated: Please use SetOutput or AppendOutput instead.
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
// Deprecated: Please use MustSetOutput or MustAppendOutput instead.
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
// Deprecated: Please use SetOutputWithPath or AppendOutputWithPath instead.
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
// Deprecated: Please use MustSetOutputWithPath or MustAppendOutputWithPath
// instead.
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
