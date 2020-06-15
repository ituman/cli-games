package seabattle

type Fleet []*Ship

func NewFleet() Fleet {
	return Fleet{
		{length: 4, helth: 4},
		{length: 3, helth: 3},
		{length: 3, helth: 3},
		{length: 2, helth: 2},
		{length: 2, helth: 2},
		{length: 2, helth: 2},
		{length: 1, helth: 1},
		{length: 1, helth: 1},
		{length: 1, helth: 1},
		{length: 1, helth: 1},
	}
}

func (f Fleet) StillAlive() int {
	var alive int
	for _, s := range f {
		if s.helth > 0 {
			alive++
		}
	}
	return alive
}
