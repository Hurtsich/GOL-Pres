package organism

import (
	"GOL-Pres/cell"
	"fmt"
	"image"
	"image/color"
	"math/rand"
	"sync"
)

type Organism struct {
	matrice [][]*cell.Cell
}

var (
	organism Organism

	white = color.RGBA{R: uint8(255), G: uint8(255), B: uint8(255), A: uint8(255)}
	black = color.RGBA{R: uint8(0), G: uint8(0), B: uint8(0), A: uint8(255)}
)

func (o *Organism) Breathe() {
	wg := new(sync.WaitGroup)
	for _, col := range o.matrice {
		for _, blob := range col {
			wg.Add(1)
			go blob.Live(wg)
		}
	}
	wg.Wait()
}

func (o *Organism) Move() *image.Paletted {
	palette := []color.Color{
		white, black,
	}
	topLeft := image.Point{0, 0}
	bottomRight := image.Point{len(o.matrice[0]), len(o.matrice)}
	img := image.NewPaletted(image.Rectangle{topLeft, bottomRight}, palette)
	for x, col := range o.matrice {
		for y, blob := range col {
			if blob.Alive {
				img.Set(x, y, white)
			} else {
				img.Set(x, y, black)
			}
		}
	}
	return img
}

func NewOrganism(lenght int) Organism {
	organism = Organism{matrice: make([][]*cell.Cell, lenght)}
	for x := range organism.matrice {
		organism.matrice[x] = make([]*cell.Cell, lenght)
	}
	for x := 0; x < lenght; x++ {
		for y := 0; y < lenght; y++ {
			var blob *cell.Cell
			if organism.matrice[x][y] == nil {
				blob = cell.NewCell(randomStatus())
				organism.matrice[x][y] = blob
			} else {
				blob = organism.matrice[x][y]
			}
			newNeighbor(mod(x, lenght), mod(y-1, lenght), Up, blob.Up)
			newNeighbor(mod(x+1, lenght), mod(y-1, lenght), UpRight, blob.UpRight)
			newNeighbor(mod(x+1, lenght), mod(y, lenght), Right, blob.Right)
			newNeighbor(mod(x+1, lenght), mod(y+1, lenght), DownRight, blob.DownRight)
			newNeighbor(mod(x, lenght), mod(y+1, lenght), Down, blob.Down)
			newNeighbor(mod(x-1, lenght), mod(y+1, lenght), DownLeft, blob.DownLeft)
			newNeighbor(mod(x-1, lenght), mod(y, lenght), Left, blob.Left)
			newNeighbor(mod(x-1, lenght), mod(y-1, lenght), UpLeft, blob.UpLeft)
		}
		fmt.Println("...")
	}
	for _, col := range organism.matrice {
		for _, blob := range col {
			blob.ChitChat()
		}
	}
	return organism
}

func newNeighbor(x, y int, side Side, membrane cell.Membrane) {
	if organism.matrice[x][y] == nil {
		blob := cell.NewCell(randomStatus())
		organism.matrice[x][y] = blob
		registerNeighbor(blob, side, membrane)
	} else {
		blob := organism.matrice[x][y]
		registerNeighbor(blob, side, membrane)
	}
}

func registerNeighbor(blob *cell.Cell, side Side, membrane cell.Membrane) {
	switch side {
	case Up:
		blob.Down.In = membrane.Out
		blob.Down.Out = membrane.In
	case UpLeft:
		blob.DownRight.In = membrane.Out
		blob.DownRight.Out = membrane.In
	case Left:
		blob.Right.In = membrane.Out
		blob.Right.Out = membrane.In
	case DownLeft:
		blob.UpRight.In = membrane.Out
		blob.UpRight.Out = membrane.In
	case Down:
		blob.Up.In = membrane.Out
		blob.Up.Out = membrane.In
	case DownRight:
		blob.UpLeft.In = membrane.Out
		blob.UpLeft.Out = membrane.In
	case Right:
		blob.Left.In = membrane.Out
		blob.Left.Out = membrane.In
	case UpRight:
		blob.DownLeft.In = membrane.Out
		blob.DownLeft.Out = membrane.In
	}
}

func randomStatus() bool {
	return rand.Intn(100) <= 12
}

func mod(a, b int) int {
	return (a%b + b) % b
}

type Side int

const (
	Up Side = iota
	UpRight
	Right
	DownRight
	Down
	DownLeft
	Left
	UpLeft
)
