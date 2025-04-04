package game

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

// Game represents the tic-tac-toe game
type Game struct {
    Board        *Board
    Players      [2]*Player
    CurrentIndex int
    Scanner      *bufio.Scanner
}

// NewGame creates a new game
func NewGame() *Game {
    return &Game{
        Board:        NewBoard(),
        Players:      [2]*Player{NewPlayer("○", "Player 1"), NewPlayer("×", "Player 2")},
        CurrentIndex: 0,
        Scanner:      bufio.NewScanner(os.Stdin),
    }
}

// CurrentPlayer returns the current player
func (g *Game) CurrentPlayer() *Player {
    return g.Players[g.CurrentIndex]
}

// SwitchPlayer switches to the next player
func (g *Game) SwitchPlayer() {
    g.CurrentIndex = (g.CurrentIndex + 1) % 2
}

// GetPlayerMove gets the player's move
func (g *Game) GetPlayerMove() int {
    for {
        fmt.Printf("%s (%s)の番です。0-24の数字を入力してください: ", 
            g.CurrentPlayer().Name, g.CurrentPlayer().Symbol)
        
        g.Scanner.Scan()
        input := strings.TrimSpace(g.Scanner.Text())
        
        position, err := strconv.Atoi(input)
        if err != nil {
            fmt.Println("無効な入力です。0-24の数字を入力してください。")
            continue
        }
        
        if !g.Board.IsValidMove(position) {
            if position < 0 || position > 24 {
                fmt.Println("無効な入力です。0-24の数字を入力してください。")
            } else {
                fmt.Println("そのマスはすでに埋まっています。別のマスを選んでください。")
            }
            continue
        }
        
        return position
    }
}

// CheckWin checks if the current player has won
func (g *Game) CheckWin() bool {
    symbol := g.CurrentPlayer().Symbol
    cells := g.Board.Cells
    size := g.Board.Size
    
    // Check rows
    for row := 0; row < size; row++ {
        for col := 0; col <= size-5; col++ {
            idx := row*size + col
            if cells[idx] == symbol && 
               cells[idx+1] == symbol && 
               cells[idx+2] == symbol && 
               cells[idx+3] == symbol && 
               cells[idx+4] == symbol {
                return true
            }
        }
    }
    
    // Check columns
    for col := 0; col < size; col++ {
        for row := 0; row <= size-5; row++ {
            idx := row*size + col
            if cells[idx] == symbol && 
               cells[idx+size] == symbol && 
               cells[idx+size*2] == symbol && 
               cells[idx+size*3] == symbol && 
               cells[idx+size*4] == symbol {
                return true
            }
        }
    }
    
    // Check diagonals (top-left to bottom-right)
    for row := 0; row <= size-5; row++ {
        for col := 0; col <= size-5; col++ {
            idx := row*size + col
            if cells[idx] == symbol && 
               cells[idx+size+1] == symbol && 
               cells[idx+size*2+2] == symbol && 
               cells[idx+size*3+3] == symbol && 
               cells[idx+size*4+4] == symbol {
                return true
            }
        }
    }
    
    // Check diagonals (top-right to bottom-left)
    for row := 0; row <= size-5; row++ {
        for col := 4; col < size; col++ {
            idx := row*size + col
            if cells[idx] == symbol && 
               cells[idx+size-1] == symbol && 
               cells[idx+size*2-2] == symbol && 
               cells[idx+size*3-3] == symbol && 
               cells[idx+size*4-4] == symbol {
                return true
            }
        }
    }
    
    return false
}

// Play starts the game
func (g *Game) Play() {
    fmt.Println("○×ゲームを開始します！")
    g.Board.DisplayWithPositions()
    
    for {
        g.Board.Display()
        position := g.GetPlayerMove()
        g.Board.MakeMove(position, g.CurrentPlayer().Symbol)
        
        if g.CheckWin() {
            g.Board.Display()
            fmt.Printf("%s (%s)の勝ち！\n", 
                g.CurrentPlayer().Name, g.CurrentPlayer().Symbol)
            break
        }
        
        if g.Board.IsFull() {
            g.Board.Display()
            fmt.Println("引き分け！")
            break
        }
        
        g.SwitchPlayer()
    }
    
    fmt.Println("ゲーム終了")
}
