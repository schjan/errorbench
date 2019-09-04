package errorbench

import (
	"errors"
	"fmt"
	"testing"

	"github.com/giantswarm/microerror"
	pkgerrors "github.com/pkg/errors"
)

var resultErr error

func BenchmarkWrapError(b *testing.B) {
	var ReallyBadError = errors.New("something bad happened")
	var err error

	for i := 0; i < b.N; i++ {
		err = fmt.Errorf("really bad!: %w", ReallyBadError)
	}

	resultErr = err
}

func BenchmarkWrapMicroerror(b *testing.B) {
	var ReallyBadError = &microerror.Error{
		Kind: "ReallyBadError",
	}
	var err error

	for i := 0; i < b.N; i++ {
		err = microerror.Maskf(ReallyBadError, "Really bad!")
	}

	resultErr = err
}

func BenchmarkWrapPkgErrors(b *testing.B) {
	var ReallyBadError = pkgerrors.New("Something bad happened")
	var err error

	for i := 0; i < b.N; i++ {
		err = pkgerrors.Wrap(ReallyBadError, "Really bad!")
	}

	resultErr = err
}
