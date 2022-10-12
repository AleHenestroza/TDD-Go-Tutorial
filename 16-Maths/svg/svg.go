package svg

import (
	"fmt"
	"io"
	"time"

	cf "go-tdd-tutorial/16-Maths"
)

const (
	secondsHandLength = 90
	minutesHandLength = 80
	hoursHandLength   = 50
	clockCentreX      = 150
	clockCentreY      = 150
)
const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`
const bezel = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`
const svgEnd = `</svg>`

func SVGWriter(w io.Writer, t time.Time) {
	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)
	writeSecondsHand(w, t)
	writeMinutesHand(w, t)
	writeHoursHand(w, t)
	io.WriteString(w, svgEnd)
}

// SecondsHand is the unit vector of the seconds hand of an analogue clock at time 't', represented as a Point
func writeSecondsHand(w io.Writer, t time.Time) {
	p := makeHand(cf.SecondsHandPoint(t), secondsHandLength)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
}

func writeMinutesHand(w io.Writer, t time.Time) {
	p := makeHand(cf.MinutesHandPoint(t), minutesHandLength)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;"/>`, p.X, p.Y)
}

func writeHoursHand(w io.Writer, t time.Time) {
	p := makeHand(cf.HoursHandPoint(t), hoursHandLength)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;"/>`, p.X, p.Y)
}

func makeHand(p cf.Point, length float64) cf.Point {
	p = cf.Point{X: p.X * length, Y: p.Y * length}             // scale
	p = cf.Point{X: p.X, Y: -p.Y}                              // flip
	p = cf.Point{X: p.X + clockCentreX, Y: p.Y + clockCentreY} // translate
	return cf.Point{X: p.X, Y: p.Y}
}
