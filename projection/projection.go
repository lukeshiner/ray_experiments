package main

import (
	"fmt"
	"os"

	"github.com/lukeshiner/raytrace/canvas"
	"github.com/lukeshiner/raytrace/colour"
	"github.com/lukeshiner/raytrace/ray"
	"github.com/lukeshiner/raytrace/vector"
)

func main() {
	var r ray.Ray
	var worldX, worldY float64
	var rayDirection, position vector.Vector
	var xs ray.Intersections
	col := colour.Colour{Red: 1, Green: 0, Blue: 0}
	canvasSize := 100
	wallSize := 7.0
	pixelSize := wallSize / float64(canvasSize)
	halfWall := wallSize / 2
	rayOrigin := vector.NewPoint(0, 0, -5)
	wallZ := 10.0
	c := canvas.New(canvasSize, canvasSize)
	s := ray.NewSphere()
	for y := 0; y < c.Height; y++ {
		worldY = halfWall - pixelSize*float64(y)
		for x := 0; x < c.Width; x++ {
			worldX = -halfWall + pixelSize*float64(x)
			position = vector.NewPoint(worldX, worldY, wallZ)
			rayDirection = vector.Subtract(position, rayOrigin)
			rayDirection = rayDirection.Normalize()
			r = ray.New(rayOrigin, rayDirection)
			xs = ray.Intersect(s, r)
			_, err := xs.Hit()
			if err == nil {
				c.WritePixel(x, y, col)
			}
		}
	}
	writePPM(c)
}

func writePPM(c canvas.Canvas) {
	output := c.ToPPM()
	f, err := os.Create("projection.ppm")
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
