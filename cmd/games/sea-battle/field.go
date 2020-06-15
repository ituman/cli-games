package seabattle

import (
	"bytes"
	"fmt"
)

const SIZE = 10

// Игровое поле размера SIZE*SIZE
type Field [SIZE][SIZE]*Cell

// NewField - Контруктор
func NewField() *Field {
	var f = new(Field)
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			f[i][j] = &Cell{status: EMPTY}
		}
	}
	return f
}

func (f *Field) SetShip(x, y int, vertical bool, s *Ship) {
	// TODO: сделать проверку условия расставления кораблей

	var shipCells []*Cell
	for i := 0; i < s.length; i++ {
		tmpCell := Cell{status: SHIP, ship: s}
		if vertical {
			f[x+i][y] = &tmpCell
		} else {
			f[x][y+i] = &tmpCell
		}
		shipCells = append(shipCells, &tmpCell)
	}
	s.shipCells = shipCells
}

func (f *Field) Shoot(x, y int) bool {
	cell := f[x][y]
	if cell.hasShip() {
		cell.ship.helth--
		if cell.ship.helth == 0 {
			cell.ship.Kill()
			fmt.Println("Убил!")
		} else {
			cell.status = HIT
			fmt.Println("Попал!")
		}
		return true
	} else {
		cell.status = MISS
		fmt.Println("Мимо!")
		return false
	}
}

func (f Field) Print(own bool) string {
	var buffer bytes.Buffer
	for _, line := range f {
		for _, cell := range line {
			buffer.WriteString(" " + cell.String(own) + " ")
		}
		buffer.WriteString("\n")
	}
	return buffer.String()
}
