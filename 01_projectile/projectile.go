package main

import (
	"fmt"
	"os"

	"github.com/lukeshiner/raytrace/canvas"
	"github.com/lukeshiner/raytrace/colour"
	"github.com/lukeshiner/raytrace/vector"
)

func main() {
	canvas := canvas.New(900, 550)
	v := vector.NewVector(1, 1.0, 0)
	v = v.Normalize()
	v = v.ScalarMultiply(11.25)
	p := projectile{vector.NewPoint(0, 1, 0), v}
	p.velocity = p.velocity.Normalize()
	env := environment{vector.NewPoint(0, -0.1, 0), vector.NewVector(-0.01, 0, 0)}
	ticks := 0
	for {
		tick(&env, &p)
		draw(&canvas, &p)
		ticks++
		if p.position.Y <= 0 {
			break
		}
	}
	writePPM(canvas)
}

type projectile struct {
	position vector.Vector
	velocity vector.Vector
}

type environment struct {
	gravity vector.Vector
	wind    vector.Vector
}

func tick(env *environment, p *projectile) {
	p.position = vector.Add(p.position, p.velocity)
	p.velocity = vector.Add(p.velocity, env.gravity)
	p.velocity = vector.Add(p.velocity, env.wind)
	p.velocity = p.velocity.Normalize()
	p.velocity = p.velocity.ScalarMultiply(11.25)
}

func draw(c *canvas.Canvas, p *projectile) {
	colour := colour.Colour{Red: 1, Green: 1, Blue: 1}
	x := int(p.position.X * 10)
	y := c.Height - int(p.position.Y*10)
	fmt.Printf("(%d,%d)\n", x, y)
	if x > 0 && x < c.Width && y > 0 && y < c.Height {
		c.WritePixel(x, y, colour)
	}
}

func writePPM(c canvas.Canvas) {
	output := c.ToPPM()
	f, err := os.Create("projectile.ppm")
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
