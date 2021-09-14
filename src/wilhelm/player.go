package wilhelm

type Player struct {
	position  Coords
	inventory []Item
}

func (p *Player) GetCoords() Coords { return p.position }
