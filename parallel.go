package util

import (
	"context"
	"errors"
	"runtime/debug"

	errors2 "github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func cancellable(ctx context.Context, errChan chan error, f func() error) {
	internalErrChan := make(chan error, 1)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				logrus.Errorf("Got panic in goroutine, Panic: %+v, StackTrace: %s", r, string(debug.Stack()))
				internalErrChan <- errors2.Errorf("got panic: %v", r)
			}
		}()
		internalErrChan <- f()

	}()

	select {
	case <-ctx.Done():
		deadline, deadlineSet := ctx.Deadline()
		logrus.Error(ctx, "context done", map[string]interface{}{
			"deadline":    deadline.String(),
			"deadlineSet": deadlineSet,
		})
		errChan <- errors.New("请求超时")
		return
	case err := <-internalErrChan:
		errChan <- err
		return
	}
}

// Parallel run functions(fs) in parallel, return on first error,
// or all fs done. Results should be returned by closure's captured variables
// NOTE: you should be very careful about the captured variables, make sure THEY
// ALL HAVE A NON-NIL INITIAL VALUE, otherwise grpc marshaller would panic and
// crash the whole program
func Parallel(ctx context.Context, fs []func() error) error {
	deadline, deadlineSet := ctx.Deadline()
	logrus.WithContext(ctx).Info("%v %t", deadline, deadlineSet)
	if len(fs) == 0 {
		return nil
	}
	contextWithCancel, cancel := context.WithCancel(ctx)
	defer cancel()
	errChan := make(chan error, len(fs))
	count := 0
	for _, f := range fs {
		if f == nil { // skip nil functions
			count++
			continue
		}
		go cancellable(contextWithCancel, errChan, f)
	}

	for err := range errChan {
		if err != nil {
			cancel()
			return err
		}
		count++
		if count >= len(fs) {
			return nil
		}
	}
	return nil
}
