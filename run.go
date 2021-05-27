package airplane

import (
	"context"
	"fmt"
	"os"
	"time"

	// TODO: move this into a separate Go package.
	"github.com/airplanedev/cli/pkg/trap"
)

func init() {
	// Use the same SIGINT-timeout as Airplane agents will provide.
	trap.Timeout = 10 * time.Second
}

// Run provides your task with context-cancellation and
// error-handling. Any errors returned, or bubbled up via
// a panic, will be logged as Airplane outputs.
func Run(f func(ctx context.Context) error) {
	// Handles task cancellation by cancelling the provided context.
	ctx := trap.Context()

	// Handle uncaught panics.
	errMsg := fmt.Sprintf("View the task logs for details: app.airplane.dev/runs/%s", os.Getenv("AIRPLANE_RUN_ID"))
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "%+v\n", r)
			if err, ok := r.(error); ok {
				MustNamedOutput("error", err.Error())
			} else {
				MustNamedOutput("error", fmt.Sprintf("%s", r))
			}
			MustNamedOutput("error", errMsg)
			os.Exit(1)
		}
	}()

	if err := f(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		MustNamedOutput("error", err.Error())
		MustNamedOutput("error", errMsg)
		os.Exit(1)
	}
}
