package errorbench

import (
	"github.com/giantswarm/microerror"
	"github.com/pkg/errors"
	"golang.org/x/xerrors"
	"testing"
)

var result bool

func BenchmarkXerrors(b *testing.B) {
	var ReallyBadError = xerrors.New("Something bad happened")
	err := xerrors.Errorf("Really bad!: %w", ReallyBadError)

	var res bool
	for i := 0; i < b.N; i++ {
		res = xerrors.Is(err, ReallyBadError)
	}

	result = res
}

func BenchmarkPkgErrors(b *testing.B) {
	var ReallyBadError = errors.New("Something bad happened")
	err := errors.Wrap(ReallyBadError, "Really bad!")

	var res bool
	for i := 0; i < b.N; i++ {
		res = errors.Cause(err) == ReallyBadError
	}

	result = res
}

func BenchmarkMicroerror(b *testing.B) {
	var ReallyBadError = &microerror.Error{
		Kind: "ReallyBadError",
	}
	err := microerror.Maskf(ReallyBadError, "Really bad!")

	var res bool
	for i := 0; i < b.N; i++ {
		res = microerror.Cause(err) == ReallyBadError
	}

	result = res
}
