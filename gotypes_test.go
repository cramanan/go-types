package gotypes

import (
	"testing"

	"github.com/cramanan/go-types/booleans"
	"github.com/cramanan/go-types/constants"
	"github.com/cramanan/go-types/maps"
	"github.com/cramanan/go-types/slices"
	"github.com/cramanan/go-types/strings"
)

func TestIntercompatibility(t *testing.T) {
	type File struct {
		Name strings.String
		Size float64
	}

	sl := slices.New(File{"main.go", 2 * constants.KB}, File{"go.mod", constants.B})

	isGoFile := func(f File, _ int) bool { return f.Name.Contains("go") }

	got := booleans.From(sl.Every(isGoFile))
	want := booleans.True

	if got != want {
		t.Errorf("Got: %t, want %t", got, want)
	}

	m := maps.New[int, File]()

	addToM := func(f File, idx int) { m.Set(idx, f) }

	sl.ForEach(addToM)

	if m.Size() != 2 {
		t.Errorf("Got: %d, want %d", m.Size(), 2)
	}
}
