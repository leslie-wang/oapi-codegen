package externalref

import (
	"testing"

	"github.com/deepmap/oapi-codegen/v2/internal/test/externalref/packageA"
	"github.com/deepmap/oapi-codegen/v2/internal/test/externalref/packageB"
)

func TestParameters(t *testing.T) {
	b := &packageB.ObjectB{}
	_ = Container{
		ObjectA: &packageA.ObjectA{ObjectB: b},
		ObjectB: b,
	}
}
