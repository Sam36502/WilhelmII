package wilhelm

type Player struct {
	position  Coords
	inventory []Item
}

func NewPlayer(game Game) *Player {
	return &Player{
		position:  game.StartCoords,
		inventory: make([]Item, 0),
	}
}

func (p *Player) GetCoords() Coords { return p.position }
