package seabattle

type Ship struct {
	length    int
	helth     int     // Отсавшиеся палубы
	shipCells []*Cell // Клетки корабля
}

func (s *Ship) Kill() {
	for _, c := range s.shipCells {
		c.status = KILL
	}
}
