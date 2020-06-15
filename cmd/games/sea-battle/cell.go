package seabattle

type CellStatus int

const (
	EMPTY CellStatus = iota
	SHIP
	MISS
	HIT
	KILL
)

type Cell struct {
	status CellStatus
	ship   *Ship
}

func (c Cell) String(own bool) string {
	switch c.status {
	case EMPTY:
		return "~"
	case SHIP:
		if own {
			return "#"
		} else {
			return "~"
		}
	case MISS:
		return "*"
	case HIT:
		return "x"
	case KILL:
		return "X"
	default:
		return "?"
	}
}

func (c *Cell) hasShip() bool {
	if c.ship != nil {
		return true
	} else {
		return false
	}
}
