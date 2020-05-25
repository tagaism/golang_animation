package main

import (
	"github.com/h8gi/canvas"
	"math/rand"
)

type Vector struct {
	x float64
	y float64
}

type Ball struct {
	location Vector
	velocity Vector
	radius float64
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

	balls := []Ball{}
	for i := 0; i < 10; i++ {
		r := randFloat(0, 80)
		b := Ball{
			location : Vector{randFloat(r, WIDTH-r), randFloat(r, HEIGHT-r)},
			velocity : Vector{randFloat(-5, 5), randFloat(-5, 5)},
			radius : r,
		}
		balls = append(balls, b)
	}
	c.Draw(func(ctx *canvas.Context) {
		ctx.SetRGB(1, 1, 1)
		ctx.Clear()
		ctx.Push()
		for i := 0; i < len(balls); i++ {
			balls[i].DrawBall(ctx)
			balls[i].Move()
			balls[i].Edge()
		}
		ctx.Pop()
	})
}

// Vectors logic
func (a *Vector) Add(b Vector) {
	(*a).x += b.x
	(*a).y += b.y
}

// Balls logic
func randFloat(min, max float64) float64 {
	return min + rand.Float64()*(max - min)
}

func (b *Ball) DrawBall(ctx *canvas.Context) {
	ctx.SetRGB(0, 0, 0)
	ctx.DrawCircle(b.location.x, b.location.y, b.radius)
	ctx.Stroke()
	ctx.SetLineWidth(1)
}

func (b *Ball) Move() {
	b.location.Add(b.velocity)
}

func (b *Ball) Edge() {
	if b.location.x < b.radius || b.location.x > WIDTH - b.radius {
		b.velocity.x *= -1
	}
	if b.location.y < b.radius || b.location.y > HEIGHT - b.radius {
		b.velocity.y *= -1
	}
}