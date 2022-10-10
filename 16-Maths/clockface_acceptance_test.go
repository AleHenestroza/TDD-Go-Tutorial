package clockface_test

import (
	clockface "go-tdd-tutorial/16-Maths"
	"strings"
	"testing"
	"time"
)

func TestSecondHandAtMidnight(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

	var b strings.Builder
	clockface.SvgWriter(&b, tm)
	got := b.String()

	want := `<line x1="150" y1="150" x2="150.000" y2="60.000"`

	if !strings.Contains(got, want) {
		t.Errorf("Expected to find the second hand %v, in the SVG output %v", want, got)
	}
}

func TestSecondHandAt30Seconds(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 30, 0, time.UTC)

	var b strings.Builder
	clockface.SvgWriter(&b, tm)
	got := b.String()

	want := `<line x1="150" y1="150" x2="150.000" y2="240.000"`

	if !strings.Contains(got, want) {
		t.Errorf("Expected to find the second hand %v, in the SVG output %v", want, got)
	}
}
