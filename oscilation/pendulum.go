package main

import (
	"github.com/h8gi/canvas"
	. "math"
	// "golang.org/x/image/colornames"
	// "fmt"
)

const (
	WIDTH = 700.0
	HEIGHT = 400.0
	FRAME_RATE = 60
)



func main() {
	c := canvas.NewCanvas(&canvas.CanvasConfig{
		Width:     WIDTH,
		Height:    HEIGHT,
		FrameRate: FRAME_RATE,
	})

	c.Setup(func(ctx *canvas.Context) {
		// ctx.SetRGB255(0, 0, 0)
	})

	amplitude := 200.0
	period := 100.0
	frame_count := 0.0

	c.Draw(func(ctx *canvas.Context) {
		// first push
		ctx.Clear()
		ctx.Push()
		frame_count += 1.0
		x := amplitude * Cos((frame_count/period) * Pi * 2)
		ctx.SetRGB(1, 1, 1)
		ctx.Translate(WIDTH/2, HEIGHT/2)
		ctx.DrawCircle(0, x, 20)
		ctx.Fill()
		ctx.DrawLine(0, 0, 0, x)
		ctx.Stroke()
		ctx.Pop() // first pop
	})
}
