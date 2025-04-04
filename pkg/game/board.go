package game

import (
    "fmt"
)

// Board represents the tic-tac-toe game board
type Board struct {
    Cells [25]string
    Size  int
}

// NewBoard creates a new empty board
func NewBoard() *Board {
    cells := [25]string{}
    for i := 0; i < 25; i++ {
        cells[i] = " "
    }
    return &Board{
        Cells: cells,
        Size:  5,
    }
}

// IsValidMove checks if the move is valid
func (b *Board) IsValidMove(position int) bool {
    if position < 0 || position > 24 {
        return false
    }
    return b.Cells[position] == " "
}

// MakeMove places a symbol at the specified position
func (b *Board) MakeMove(position int, symbol string) {
    b.Cells[position] = symbol
}

// IsFull checks if the board is full
func (b *Board) IsFull() bool {
    for _, cell := range b.Cells {
        if cell == " " {
            return false
        }
    }
    return true
}

// Display prints the current state of the board
func (b *Board) Display() {
    fmt.Println()
    for i := 0; i < b.Size; i++ {
        fmt.Printf(" %s | %s | %s | %s | %s \n", 
            b.Cells[i*5], b.Cells[i*5+1], b.Cells[i*5+2], b.Cells[i*5+3], b.Cells[i*5+4])
        if i < b.Size-1 {
            fmt.Println("-------------------")
        }
    }
    fmt.Println()
}

// DisplayWithPositions prints the board with position numbers
func (b *Board) DisplayWithPositions() {
    fmt.Println("Positions:")
    for i := 0; i < b.Size; i++ {
        fmt.Printf(" %2d | %2d | %2d | %2d | %2d \n", 
            i*5, i*5+1, i*5+2, i*5+3, i*5+4)
        if i < b.Size-1 {
            fmt.Println("-------------------")
        }
    }
    fmt.Println()
}
