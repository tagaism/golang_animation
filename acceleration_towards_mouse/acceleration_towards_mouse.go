package main

import (
	"github.com/h8gi/canvas"
	"math/rand"
	"math"
)

type Vector struct {
	x float64
	y float64
}

type Ball struct {
	location Vector
	velocity Vector
	acceleration Vector
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
	for i := 0; i < 3; i++ {
		r := randFloat(0, 80)
		b := Ball{
			location : Vector{randFloat(r, WIDTH-r), randFloat(r, HEIGHT-r)},
			velocity : Vector{0, 0},
			acceleration : Vector{0, 0},
			radius : r,
		}
		balls = append(balls, b)
	}
	c.Draw(func(ctx *canvas.Context) {
		mouse := Vector{ctx.Mouse.X, ctx.Mouse.Y}
		ctx.SetRGB(1, 1, 1)
		ctx.Clear()
		ctx.Push()
		for i := 0; i < len(balls); i++ {
			balls[i].DrawBall(ctx)
			balls[i].Move(mouse)
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

func (a *Vector) Sub(b Vector) {
	(*a).x -= b.x
	(*a).y -= b.y
}

func (a *Vector)Limit(max float64) {
	if (*a).x > max {
		(*a).x = max
	}
	if (*a).y > max {
		(*a).y = max
	}
}

func (a *Vector) Normalize() {
	mag := math.Sqrt(a.x * a.x + a.y * a.y)
	(*a).x = (*a).x / mag
	(*a).y = (*a).y / mag
}

func (a *Vector) Mult(v float64) {
	(*a).x *= v
	(*a).y *= v
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

func (b *Ball) Move(point Vector) {
	point.Sub(b.location)

	//set point vector magnitude
	point.Normalize()
	point.Mult(0.5)

	b.acceleration = point
	b.velocity.Add(b.acceleration)
	b.velocity.Limit(5)
	b.location.Add(b.velocity)
}

func (b *Ball) Edge() {
	if b.location.x > WIDTH {
		b.location.x = 0
	} else if b.location.x < 0 {
		b.location.x = WIDTH
	}

	if b.location.y > HEIGHT {
		b.location.y = 0
	} else if b.location.y < 0 {
		b.location.y = HEIGHT
	}
}
