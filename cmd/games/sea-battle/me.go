package seabattle

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Me struct {
	name     string
	field    *Field
	fleet    Fleet
	enemy    Player
	attempts []string
}

func (m *Me) initField() {
	fmt.Println("Расставь свои корабли на поле. Для этого вводи координату носа корабля")
	for _, s := range m.fleet {
		fmt.Println(m.field.Print(true))
		fmt.Println("Размер: ", s.length)
		pos := readPos()
		m.field.SetShip(pos[0], pos[1], true, s)
		fmt.Println(m.field.Print(true))
	}
}

func (m *Me) shot() bool {
	fmt.Println("Введи координаты выстрела")
	pos := readPos()
	isHit := m.getEnemy().getField().Shoot(pos[0], pos[1])
	return isHit
}

func (m *Me) isLose() bool {
	if m.fleet.StillAlive() > 0 {
		return false
	} else {
		return true
	}
}

func readPos() []int {
	reader := bufio.NewReader(os.Stdin)
	inp, _ := reader.ReadString('\n')
	nums := strings.Split(inp, "")
	x, _ := strconv.Atoi(nums[0])
	y, _ := strconv.Atoi(nums[1])
	return []int{x, y}
}

func (m *Me) getEnemy() Player {
	return m.enemy
}

func (m *Me) getField() *Field {
	return m.field
}

func (m *Me) getFleet() Fleet {
	return m.fleet
}

func (m *Me) getName() string {
	return m.name
}
