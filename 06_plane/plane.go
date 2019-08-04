package main

import (
	"fmt"
	"math"
	"os"

	"github.com/lukeshiner/raytrace/camera"
	"github.com/lukeshiner/raytrace/canvas"
	"github.com/lukeshiner/raytrace/colour"
	"github.com/lukeshiner/raytrace/light"
	"github.com/lukeshiner/raytrace/material"
	"github.com/lukeshiner/raytrace/matrix"
	"github.com/lukeshiner/raytrace/shape"
	"github.com/lukeshiner/raytrace/vector"
	"github.com/lukeshiner/raytrace/world"
)

func main() {
	w := getWorld()
	c := getCamera()
	img := camera.Render(c, w)
	writePPM(img)
}

func getWorld() world.World {
	world := world.New()
	world.Objects = getObjects()
	world.Lights = getLights()
	return world
}

func getObjects() []shape.Shape {
	return []shape.Shape{
		getFloor(), getLargeSphere(), getMiddleSphere(),
		getSmallSphere(),
	}
}

func getFloorMaterial() material.Material {
	m := material.New()
	m.Colour = colour.New(1, 0.9, 0.9)
	m.Specular = 0
	return m
}

func getFloor() shape.Shape {
	floor := shape.NewPlane()
	return floor
}

func getLargeSphere() shape.Shape {
	s := shape.NewSphere()
	s.SetTransform(matrix.TranslationMatrix(-0.5, 1, 0.5))
	m := material.New()
	m.Colour = colour.New(0.1, 1, 0.5)
	m.Diffuse = 0.7
	m.Specular = 0.3
	s.SetMaterial(m)
	return s
}

func getSmallSphere() shape.Shape {
	s := shape.NewSphere()
	s.SetTransform(matrix.Multiply(
		matrix.TranslationMatrix(1.5, 0.5, -0.5), matrix.ScalingMatrix(0.5, 0.5, 0.5)))
	m := material.New()
	m.Colour = colour.New(0.5, 1, 0.1)
	m.Diffuse = 0.7
	m.Specular = 0.3
	s.SetMaterial(m)
	return s
}

func getMiddleSphere() shape.Shape {
	s := shape.NewSphere()
	s.SetTransform(matrix.Multiply(
		matrix.TranslationMatrix(-1.5, 0.33, -0.75), matrix.ScalingMatrix(0.33, 0.33, 0.33)))
	m := material.New()
	m.Colour = colour.New(1, 0.8, 0.1)
	m.Diffuse = 0.7
	m.Specular = 0.3
	s.SetMaterial(m)
	return s
}

func getLights() []light.Light {
	l := light.NewPoint(colour.New(1, 1, 1), vector.NewPoint(-10, 10, -10))
	return []light.Light{l}
}

func getCamera() camera.Camera {
	c := camera.New(100, 50, math.Pi/3)
	t := camera.ViewTransform(
		vector.NewPoint(0, 1.5, -5), vector.NewPoint(0, 1, 0), vector.NewVector(0, 1, 0))
	c.SetTransform(t)
	return c
}

func writePPM(c canvas.Canvas) {
	output := c.ToPPM()
	f, err := os.Create("plane.ppm")
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
