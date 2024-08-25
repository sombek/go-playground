package clockface_test

import (
	"bytes"
	"encoding/xml"
	clockface "go-playground/math/clockface"
	"os"
	"testing"
	"time"
)

func TestSVGWriterAtMidnight(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)
	b := bytes.Buffer{}

	clockface.SVGWriter(&b, tm)

	svg := clockface.SVG{}

	xml.Unmarshal(b.Bytes(), &svg)

	want := clockface.Line{150, 150, 150, 60}

	for _, line := range svg.Line {
		if line == want {
			return
		}
	}

	t.Errorf("Expected to find the second hand line %+v, in the SVG lines %+v", want, svg.Line)
}

func TestSVGWriterMinuteHand(t *testing.T) {
	cases := []struct {
		time time.Time
		line clockface.Line
	}{
		{
			simpleTime(0, 0, 0),
			clockface.Line{150, 150, 150, 70},
		},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			b := bytes.Buffer{}
			clockface.SVGWriter(&b, c.time)

			svg := clockface.SVG{}
			xml.Unmarshal(b.Bytes(), &svg)

			if !containsLine(c.line, svg.Line) {
				t.Errorf("Expected to find the minute hand line %+v, in the SVG lines %+v", c.line, svg.Line)
			}
		})
	}
}

func TestSVGWriterHourHand(t *testing.T) {
	cases := []struct {
		time time.Time
		line clockface.Line
	}{
		{
			simpleTime(6, 0, 0),
			clockface.Line{150, 150, 150, 200},
		},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			b := bytes.Buffer{}
			clockface.SVGWriter(&b, c.time)

			svg := clockface.SVG{}
			xml.Unmarshal(b.Bytes(), &svg)

			if !containsLine(c.line, svg.Line) {
				t.Errorf("Expected to find the hour hand line %+v, in the SVG lines %+v", c.line, svg.Line)
			}
		})
	}
}

// full test coverage
func TestSVGWriter(t *testing.T) {
	b := bytes.Buffer{}
	clockface.SVGWriter(&b, simpleTime(10, 10, 12))

	svg := clockface.SVG{}
	xml.Unmarshal(b.Bytes(), &svg)
	// export the SVG to a file to see the result
	os.WriteFile("clockface.svg", b.Bytes(), 0644)

	if len(svg.Line) != 3 {
		t.Fatalf("Expected 3 lines in the SVG, but got %d", len(svg.Line))
	}
}

func containsLine(l clockface.Line, ls []clockface.Line) bool {
	for _, line := range ls {
		if line == l {
			return true
		}
	}
	return false
}

func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(312, time.October, 28, hours, minutes, seconds, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}
