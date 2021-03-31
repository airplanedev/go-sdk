package airplane

import (
	"encoding/json"
	"os"

	"github.com/pkg/errors"
)

// Parameters parses the CLI arguments provided to the task
// and sets them into `parameters` by matching JSON struct tags
// to parameter slugs.
//
// This expects that the first CLI argument to your task is a JSON document
// containing parameter slugs mapped to parameter values. This is done by
// setting the `arguments` field of the task to `{{.JSON}}`.
//
// Docs: https://docs.airplane.dev/reference/parameters
func Parameters(parameters interface{}) error {
	args := os.Args[1:]
	if len(args) != 1 {
		// TODO: if we don't receive exactly one JSON argument, use the reflect package
		// on parameters to parse the rest as flags and drop their values into the
		// relevant field in parameters.
		return errors.Errorf("expected a single arg, got %d args", len(args))
	}

	if err := json.Unmarshal([]byte(args[0]), parameters); err != nil {
		return errors.Wrap(err, "unmarshalling parameters")
	}

	return nil
}
