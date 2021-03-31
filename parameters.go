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
// It expects that your task has been configured with a CLI arguments
// template value of `{{.JSON}}`. It will error otherwise.
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
