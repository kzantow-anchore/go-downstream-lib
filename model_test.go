package downstream_test

import (
	"github.com/kzantow-anchore/go-downstream-lib"
	"testing"
)

func Test_v2(t *testing.T) {
	test := downstream.Model{
		Name2: "Name2",
	}
	t.Log(test)
}
