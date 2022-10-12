package clockface

import (
	"fmt"
	"io"
	"time"
)

const (
	secondsHandLength = 90
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
	io.WriteString(w, svgEnd)
}

// SecondsHand is the unit vector of the seconds hand of an analogue clock at time 't', represented as a Point
func writeSecondsHand(w io.Writer, t time.Time) {
	p := secondsHandPoint(t)                                    // convert
	p = Point{p.X * secondsHandLength, p.Y * secondsHandLength} // scale
	p = Point{p.X, -p.Y}                                        // flip
	p = Point{p.X + clockCentreX, p.Y + clockCentreY}           // translate
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
}
