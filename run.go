package airplane

import (
	"context"
	"fmt"
	"os"
	"runtime/debug"
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
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "%+v\n%s\n", r, debug.Stack())
			if err, ok := r.(error); ok {
				MustSetOutput(err.Error(), "error")
			} else {
				MustSetOutput(fmt.Sprintf("%s", r), "error")
			}
			os.Exit(1)
		}
	}()

	if err := f(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		MustSetOutput(err.Error(), "error")
		os.Exit(1)
	}
}
