package main

import (
	"fmt"
	"math"
	"os"

	"github.com/lukeshiner/raytrace/canvas"
	"github.com/lukeshiner/raytrace/colour"
	"github.com/lukeshiner/raytrace/vector"
)

func main() {
	white := colour.Colour{Red: 1, Green: 1, Blue: 1}
	c := canvas.New(100, 100)
	point := vector.NewPoint(40, 0, 0)
	for i := 0; i < 12; i++ {
		fmt.Println(i, point.X, point.Y)
		point = point.RotateZ((2 * math.Pi) / 12)
		c.WritePixel(int(point.X)+50, int(point.Y)+50, white)
	}
	writePPM(c)
}

func writePPM(c canvas.Canvas) {
	output := c.ToPPM()
	f, err := os.Create("clock.ppm")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString(output)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
