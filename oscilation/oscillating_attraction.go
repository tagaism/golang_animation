package main

import(
	"fmt"
	"github.com/tagaism/vectorgo"
	"github.com/h8gi/canvas"
	"math/rand"
	"math"
	"time"
)

type Attractor struct {
	Location vector.Vector
	Radius float64
	Mass float64
}

type Oscillating struct {
	Location vector.Vector
	Velocity vector.Vector
	Acceleration vector.Vector
	Mass float64
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

	attractor := Attractor{
		Location: vector.Create(WIDTH/2, HEIGHT/2),
		Radius: 50,
		Mass: 10,
	}

	var oscillatings []Oscillating

	for i:=0; i < 7; i++ {
		osc := Oscillating{
			Location: vector.Create(randFloat(0, WIDTH), randFloat(0, HEIGHT)),
			Velocity: vector.Create(randFloat(0, 1), randFloat(0, 1)),
			Mass: 3,
		}
		oscillatings = append(oscillatings, osc)
	}

	fmt.Println(oscillatings)
	c.Setup(func(ctx *canvas.Context) {
		// ctx.SetRGB255(0, 0, 0)
	})

	c.Draw(func(ctx *canvas.Context) {
		ctx.Clear()
		ctx.SetRGB255(255, 255, 255)

		for i:=0; i < len(oscillatings); i++ {
			gravity := attractor.Attract(oscillatings[i])
			oscillatings[i].ApplyForce(gravity)
			oscillatings[i].Move(ctx)
			oscillatings[i].Edge()
			attractor.DrawAttractor(ctx)

			dir := vector.Sub(oscillatings[i].Location, attractor.Location)
			dist := dir.Mag()
			oscillatings[i].DrawOscillating(dist, ctx)
		}

	})

}

func (osc *Oscillating) DrawOscillating(dist float64, ctx *canvas.Context) {
	//find angle of direction
	rad := math.Atan2(osc.Velocity.Y, osc.Velocity.X)

	ctx.Push()
	ctx.Translate(osc.Location.X, osc.Location.Y)
	ctx.Rotate(rad)
	
	//make'm all oscillate
	amplitude := 5.0
	period := dist/70.0
	x := amplitude * math.Sin((period) * math.Pi * 2)
	y := amplitude * math.Cos((period) * math.Pi * 2)
	ctx.DrawEllipse(0, 0, x + 20, y + 13)

	ctx.SetRGB(0, 0, 0)
	// ctx.DrawLine(0, 0, 30, 0)
	ctx.Stroke()
	ctx.DrawCircle(x+13, 0, 3)
	ctx.Fill()
	ctx.Pop()
}

func (osc *Oscillating) ApplyForce(f vector.Vector) {
	force := f.Copy()
	force.Div(osc.Mass)
	osc.Acceleration.Add(force)
}

func (att *Attractor) Attract(osc Oscillating) vector.Vector {
	//Direction of the force
	force := att.Location.Copy()
	force.Sub(osc.Location)
	dist := force.Mag()
	dist = Constrain(dist, 5, 30)
	force.Normal()

	//Magnitude of the force
	strength := (G * att.Mass * osc.Mass) / (math.Pow(dist, 2))

	//Putting magnitude and direction together
	force.Mult(strength)
	return force
}

func (osc *Oscillating) Edge() {
	if osc.Location.X > WIDTH || osc.Location.X < 0 {
		osc.Velocity.X *= -1
	}
	if osc.Location.Y > HEIGHT || osc.Location.Y < 0 {
		osc.Velocity.Y *= -1
	}
}

func (osc *Oscillating) Move(ctx *canvas.Context) {
	osc.Velocity.Add(osc.Acceleration)
	osc.Location.Add(osc.Velocity)
	osc.Acceleration.Mult(0) //To avoid hyperacceleration of oscillatings
	// fmt.Println(osc.Velocity)
}


func (att *Attractor) DrawAttractor(ctx *canvas.Context) {
	ctx.Push()
	ctx.SetRGBA255(0, 0, 0, 10)
	ctx.DrawCircle(att.Location.X, att.Location.Y, att.Radius)
	ctx.Fill()
	ctx.Pop()
}

func Constrain(nam, min, max float64) float64 {
	//Limits naminal between min and max values
	if nam < min {
		return min
	}
	if nam > max {
		return max
	}
	return nam
}

func randFloat(min, max float64) float64 {
	rand.Seed(time.Now().UnixNano())// varying sequence
	return min + rand.Float64()*(max - min)
}
