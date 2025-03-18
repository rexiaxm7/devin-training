package game

import (
    "fmt"
)

// Board represents the tic-tac-toe game board
type Board struct {
    Cells [9]string
}

// NewBoard creates a new empty board
func NewBoard() *Board {
    return &Board{
        Cells: [9]string{" ", " ", " ", " ", " ", " ", " ", " ", " "},
    }
}

// IsValidMove checks if the move is valid
func (b *Board) IsValidMove(position int) bool {
    if position < 0 || position > 8 {
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
    fmt.Printf(" %s | %s | %s \n", b.Cells[0], b.Cells[1], b.Cells[2])
    fmt.Println("-----------")
    fmt.Printf(" %s | %s | %s \n", b.Cells[3], b.Cells[4], b.Cells[5])
    fmt.Println("-----------")
    fmt.Printf(" %s | %s | %s \n", b.Cells[6], b.Cells[7], b.Cells[8])
    fmt.Println()
}

// DisplayWithPositions prints the board with position numbers
func (b *Board) DisplayWithPositions() {
    fmt.Println("Positions:")
    fmt.Println(" 0 | 1 | 2 ")
    fmt.Println("-----------")
    fmt.Println(" 3 | 4 | 5 ")
    fmt.Println("-----------")
    fmt.Println(" 6 | 7 | 8 ")
    fmt.Println()
}
