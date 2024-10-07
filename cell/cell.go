package cell

import "sync"

type Cell struct {
	Alive bool

	Up        Membrane
	UpRight   Membrane
	Right     Membrane
	DownRight Membrane
	Down      Membrane
	DownLeft  Membrane
	Left      Membrane
	UpLeft    Membrane
}

func NewCell(status bool) *Cell {
	return &Cell{
		Alive:     status,
		Up:        NewMembrane(),
		UpRight:   NewMembrane(),
		Right:     NewMembrane(),
		DownRight: NewMembrane(),
		Down:      NewMembrane(),
		DownLeft:  NewMembrane(),
		Left:      NewMembrane(),
		UpLeft:    NewMembrane(),
	}
}

func (c *Cell) Live(wg *sync.WaitGroup) {
	defer wg.Done()
	aliveNeighbors := c.listen()
	if c.Alive &&
		(aliveNeighbors == 2 || aliveNeighbors == 3) {
		c.Alive = true
	} else if !c.Alive && aliveNeighbors == 3 {
		c.Alive = true
	} else {
		c.Alive = false
	}
	c.ChitChat()
}

func (c *Cell) listen() int {
	handshake := 0
	handshake += isAlive(c.Up.In)
	handshake += isAlive(c.UpRight.In)
	handshake += isAlive(c.Right.In)
	handshake += isAlive(c.DownRight.In)
	handshake += isAlive(c.Down.In)
	handshake += isAlive(c.DownLeft.In)
	handshake += isAlive(c.Left.In)
	handshake += isAlive(c.UpLeft.In)
	return handshake
}

func (c *Cell) ChitChat() {
	c.Up.Out <- c.Alive
	c.UpRight.Out <- c.Alive
	c.Right.Out <- c.Alive
	c.DownRight.Out <- c.Alive
	c.Down.Out <- c.Alive
	c.DownLeft.Out <- c.Alive
	c.Left.Out <- c.Alive
	c.UpLeft.Out <- c.Alive
}

func isAlive(neighbor chan bool) int {
	if <-neighbor {
		return 1
	} else {
		return 0
	}
}
