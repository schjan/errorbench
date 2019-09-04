package errorbench

import (
	"errors"
	"fmt"
	"testing"

	"github.com/giantswarm/microerror"
	pkgerrors "github.com/pkg/errors"
	"golang.org/x/xerrors"
)

var result bool


func BenchmarkErrors(b *testing.B) {
	var ReallyBadError = errors.New("something bad happened")
	err := fmt.Errorf("really bad!: %w", ReallyBadError)

	var res bool
	for i := 0; i < b.N; i++ {
		res = errors.Is(err, ReallyBadError)
	}

	result = res
}

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
	var ReallyBadError = pkgerrors.New("Something bad happened")
	err := pkgerrors.Wrap(ReallyBadError, "Really bad!")

	var res bool
	for i := 0; i < b.N; i++ {
		res = pkgerrors.Cause(err) == ReallyBadError
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
