package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis range (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 5         // angle of x, y axes
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, _, ok := corner(i+1, j)
			if !ok {
				continue
			}
			bx, by, _, ok := corner(i, j)
			if !ok {
				continue
			}
			cx, cy, _, ok := corner(i, j+1)
			if !ok {
				continue
			}
			dx, dy, color, ok := corner(i+1, j+1)
			if !ok {
				continue
			}
			fmt.Printf("<polygon style='stroke:%s' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				color, ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, string, bool) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z, ok := f(x, y)
	if !ok {
		return 0, 0, "#ffffff", false
	}
	var color string
	if z > 0 {
		color = "#ff0000"
	} else {
		color = "#0000ff"
	}

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, color, true
}

func f(x, y float64) (float64, bool) {
	r := math.Hypot(x, y)
	if math.IsNaN(r) || math.IsInf(r, 0) {
		return 0, false
	}
	return math.Sin(r) / r, true
}
