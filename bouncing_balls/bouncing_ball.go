package main

import (
	"github.com/tagaism/vectorgo"
	"github.com/h8gi/canvas"
	"math/rand"
	"github.com/faiface/pixel/pixelgl"
	"fmt"
)

type Ball struct {
	location vector.Vector
	velocity vector.Vector
	acceleration vector.Vector
	radius float64
	mass float64
}

const (
	WIDTH = 700.0
	HEIGHT = 400.0
	FRAME_RATE = 60.0
)

func main() {
	c := canvas.NewCanvas(&canvas.CanvasConfig{
		Width: WIDTH,
		Height: HEIGHT,
		FrameRate: FRAME_RATE,
	})

	balls := make([]Ball, 1)
	for _, b := range balls {
		r := randFloat(0, 80)
		b.location = vector.Create(WIDTH/2, HEIGHT-r)
		b.velocity = vector.Create(0, -2)
		b.acceleration = vector.Create(0, 0)
		b.radius = r
		b.mass = r
		balls = append(balls, b)
	}

	m := 47.0 // coeficient of air friction

	// wind := vector.Create(-0.4, 0)
	gravitation := vector.Create(0.0, -5)
	friction := gravitation.Copy()

	friction.Normal()
	friction.Mult(-1)
	friction.Mult(m)

	fmt.Println(friction)

	c.Draw(func(ctx *canvas.Context) {
		ctx.SetRGB(1, 1, 1)
		ctx.Clear()
		for i := 0; i < len(balls); i++ {
			balls[i].DrawBall(ctx)
			balls[i].Move()
			balls[i].Edge()
			if ctx.IsKeyPressed(pixelgl.MouseButtonLeft) {
				// balls[i].applyForce(wind)
				fmt.Println("PRESSED")
				balls[i].applyForce(friction)
			}
			balls[i].applyForce(gravitation)
		}
	})
}

func randFloat(min, max float64) float64 {
	return min + rand.Float64()*(max - min)
}

func (b *Ball) DrawBall(ctx *canvas.Context) {
	ctx.SetRGB(0, 0, 0)
	ctx.DrawCircle(b.location.X, b.location.Y, b.radius)
	ctx.Stroke()
	ctx.SetLineWidth(1)
}

func (b *Ball) Move() {
	b.velocity.Add(b.acceleration)
	b.location.Add(b.velocity)
	b.acceleration.Mult(0)
}

func (b *Ball) Edge() {
	if b.location.X < b.radius || b.location.X > WIDTH - b.radius {
		b.velocity.X *= -1
	}
	if b.location.Y < b.radius || b.location.Y > HEIGHT - b.radius {
		b.velocity.Y *= -1
	}
}

// Newton's 2nd law with mass
func (b *Ball) applyForce(f vector.Vector) {
	f.Div(b.mass)
	b.acceleration.Add(f)
}
