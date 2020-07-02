package main

import (
	"github.com/tagaism/vectorgo"
	"github.com/h8gi/canvas"
	"math/rand"
	"time"
)

type Ball struct {
	location vector.Vector
	velocity vector.Vector
	acceleration vector.Vector
	radius float64
	mass float64
}

type Attractor struct {
	location vector.Vector
	radius float64
	mass float64
}

const (
	WIDTH = 700.0
	HEIGHT = 400.0
	FRAME_RATE = 60.0
	G = 1 //global gravitational const
)

func main() {
	c := canvas.NewCanvas(&canvas.CanvasConfig{
		Width: WIDTH,
		Height: HEIGHT,
		FrameRate: FRAME_RATE,
	})

	//creating balls
	balls := []Ball{}
	for i := 0; i < 5; i++ {
		// r := randFloat(10, 30)
		m := randFloat(0.1, 1)
		r := m*20
		ball := Ball{
			location: vector.Create(randFloat(r, WIDTH), randFloat(r, HEIGHT-r)),
			// location: vector.Create(300, 400),
			velocity: vector.Create(0, 0),
			acceleration: vector.Create(0, 0),
			radius: r,
			mass: m,
		}
		balls = append(balls, ball)
	}

	attractor := Ball{
		location: vector.Create(WIDTH/2, HEIGHT/2),
		velocity: vector.Create(0, 0),
		acceleration: vector.Create(0, 0),
		radius: 40,
		mass: 5,
	}
	balls = append(balls, attractor)

	c.Draw(func(ctx *canvas.Context) {
		ctx.SetRGB(1, 1, 1)
		ctx.Clear()
		ctx.Push()
			
		for i := 0; i < len(balls); i++ {
			for j := 0; j < len(balls); j++ {
				if &balls[i] != &balls[j] {// Check if not the same ball
					m_force := balls[i].MutualAttract(balls[j])
					balls[j].applyForce(m_force)
				}
			}
			balls[i].Draw(ctx)
			balls[i].Move()
			balls[i].Edge()
		}
		ctx.Pop()
	})
}

func randFloat(min, max float64) float64 {
	rand.Seed(time.Now().UnixNano())// varying sequence
	return min + rand.Float64()*(max - min)
}

func (b *Ball) Draw(ctx *canvas.Context) {
	ctx.DrawCircle(b.location.X, b.location.Y, b.radius)
	if b.mass < 4 {
		ctx.SetRGBA255(100, 100, 100, 100)
	} else {
		ctx.SetRGBA255(80, 80, 80, 200)
	}
	ctx.Fill()
	ctx.Stroke()
	ctx.SetLineWidth(1)
}

func (b *Ball) Move() {
	b.velocity.Add(b.acceleration)
	b.location.Add(b.velocity)
	b.acceleration.Mult(0)
}

func (b *Ball) Edge() {
	if b.location.X < 0|| b.location.X > WIDTH {
		b.velocity.X *= -1
	}
	if b.location.Y < 0 || b.location.Y > HEIGHT {
		b.velocity.Y *= -1
	}
}

func (b1 *Ball) MutualAttract(b2 Ball) vector.Vector {
	//Direction of the force
	force := b1.location.Copy()
	force.Sub(b2.location)
	dist := force.Mag()
	dist = Constrain(dist, 5, 30)
	force.Normal()

	//Magnitude of the force
	strength := (G * b1.mass * b2.mass )/(dist * dist)

	//Putting magnitude and direction together
	force.Mult(strength)
	return force
}

// Newton's 2nd law with mass
func (b *Ball) applyForce(force vector.Vector) {
	c_force := force.Copy()
	c_force.Div(b.mass)
	b.acceleration.Add(c_force)
}

func Constrain(nam, min, max float64) float64 {
	if nam < min {
		return min
	}
	if nam > max {
		return max
	}
	return nam
}
