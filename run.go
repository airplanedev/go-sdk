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

// Run wraps provides your task with context-cancellation and
// error-handling logic. Any errors returned, or bubbled up via
// a panic, will be logged as Airplane outputs.
func Run(f func(ctx context.Context) error) {
	// Handles task cancellation by cancelling the provided context.
	ctx := trap.Context()

	// Handle uncaught panics.
	defer func() {
		if r := recover(); r != nil {
			MustNamedOutput("error", fmt.Sprintf("%+v", r))
			os.Exit(1)
		}
	}()

	if err := f(ctx); err != nil {
		MustNamedOutput("error", fmt.Sprintf("%+v", err))
		os.Exit(1)
	}
}
