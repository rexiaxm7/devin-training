package game

// Player represents a player in the game
type Player struct {
    Symbol string
    Name   string
}

// NewPlayer creates a new player with the given symbol
func NewPlayer(symbol, name string) *Player {
    return &Player{
        Symbol: symbol,
        Name:   name,
    }
}
