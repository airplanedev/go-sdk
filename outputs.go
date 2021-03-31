package airplane

import (
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/pkg/errors"
)

func Output(value interface{}) error {
	return NamedOutput("", value)
}

func MustOutput(value interface{}) {
	if err := Output(value); err != nil {
		panic(errors.Wrap(err, "writing output"))
	}
}

func NamedOutput(name string, value interface{}) error {
	if err := validateOutputName(name); err != nil {
		return err
	}

	header := "airplane_output"
	if name != "" {
		header += ":" + name
	}

	out, err := json.Marshal(value)
	if err != nil {
		return errors.Wrap(err, "marshalling output to JSON")
	}

	fmt.Printf("%s %s\n", header, string(out))

	return nil
}

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
