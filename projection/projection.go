package main

import (
	"fmt"
	"os"

	"github.com/lukeshiner/raytrace/canvas"
	"github.com/lukeshiner/raytrace/colour"
	"github.com/lukeshiner/raytrace/material"
	"github.com/lukeshiner/raytrace/matrix"
	"github.com/lukeshiner/raytrace/object"
	"github.com/lukeshiner/raytrace/ray"
	"github.com/lukeshiner/raytrace/vector"
)

var drawCol = colour.New(1, 0, 0)
var background = colour.New(0, 0, 0)
var canvasSize = 100
var wallSize = 7.0
var pixelSize = wallSize / float64(canvasSize)
var halfWall = wallSize / 2
var rayOrigin = vector.NewPoint(0, 0, -5)
var wallZ = 10.0

func main() {
	var col colour.Colour
	var worldX, worldY float64
	c := canvas.New(canvasSize, canvasSize)
	s := getObject()
	for y := 0; y < c.Height; y++ {
		worldY = halfWall - pixelSize*float64(y)
		for x := 0; x < c.Width; x++ {
			worldX = -halfWall + pixelSize*float64(x)
			col = calculatePixelColour(x, y, worldX, worldY, s)
			c.WritePixel(x, y, col)
		}
	}
	writePPM(c)
}

func getObject() object.Sphere {
	s := object.NewSphere()
	s.SetTransform(getTransform())
	s.SetMaterial(getMaterial())
	return s
}

func getTransform() matrix.Matrix {
	return matrix.IdentityMatrix(4)
}

func getMaterial() material.Material {
	return material.New()
}

func calculatePixelColour(x, y int, worldX, worldY float64, s object.Sphere) colour.Colour {
	position := vector.NewPoint(worldX, worldY, wallZ)
	rayDirection := vector.Subtract(position, rayOrigin)
	rayDirection = rayDirection.Normalize()
	r := ray.New(rayOrigin, rayDirection)
	xs := ray.Intersect(s, r)
	_, err := xs.Hit()
	if err == nil {
		return drawCol
	}
	return background
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
