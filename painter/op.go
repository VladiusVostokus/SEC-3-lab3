package painter

import (
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/image/draw"
	"image"
	"image/color"
)

// Operation змінює вхідну текстуру.
type Operation interface {
	// Do виконує зміну операції, повертаючи true, якщо текстура вважається готовою для відображення.
	Do(t screen.Texture) (ready bool)
}

type BackGroundRect struct {
	X1 int
	Y1 int
	X2 int
	Y2 int
}

func (bgr *BackGroundRect) Do(t screen.Texture) (ready bool) {
	i := image.Rect(bgr.X1, bgr.Y1, bgr.X2, bgr.Y2)
	t.Fill(i.Bounds(), color.Black, screen.Src)
	return false
}

// OperationList групує список операції в одну.
type OperationList []Operation

func (ol OperationList) Do(t screen.Texture) (ready bool) {
	for _, o := range ol {
		ready = o.Do(t) || ready
	}
	return
}

// UpdateOp операція, яка не змінює текстуру, але сигналізує, що текстуру потрібно розглядати як готову.
var UpdateOp = updateOp{}

type updateOp struct{}

func (op updateOp) Do(t screen.Texture) bool { return true }

// OperationFunc використовується для перетворення функції оновлення текстури в Operation.
type OperationFunc func(t screen.Texture)

func (f OperationFunc) Do(t screen.Texture) bool {
	f(t)
	return false
}

// WhiteFill зафарбовує тестуру у білий колір. Може бути викоистана як Operation через OperationFunc(WhiteFill).
func WhiteFill(t screen.Texture) {
	t.Fill(t.Bounds(), color.White, screen.Src)
}

// GreenFill зафарбовує тестуру у зелений колір. Може бути викоистана як Operation через OperationFunc(GreenFill).
func GreenFill(t screen.Texture) {
	t.Fill(t.Bounds(), color.RGBA{G: 0xff, A: 0xff}, screen.Src)
}

func Reset(t screen.Texture) {
	t.Fill(t.Bounds(), color.Black, screen.Src)
}

type Cross struct {
	X int
	Y int
}

func (c *Cross) Do(t screen.Texture) (ready bool) {
	blue := color.RGBA{0, 0, 255, 0}
	x1 := c.X - 200
	y1 := c.Y - 100
	x2 := x1 + 400
	y2 := y1 + 200
	i := image.Rect(x1, y1, x2, y2)
	t.Fill(i.Bounds(), blue, draw.Src)

	x1 = x1 + 100
	y1 = y1 - 100
	x2 = x1 + 200
	y2 = y1 + 400
	i = image.Rect(x1, y1, x2, y2)
	t.Fill(i.Bounds(), blue, draw.Src)
	return false
}
