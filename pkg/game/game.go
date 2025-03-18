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
        fmt.Printf("%s (%s)の番です。0-8の数字を入力してください: ", 
            g.CurrentPlayer().Name, g.CurrentPlayer().Symbol)
        
        g.Scanner.Scan()
        input := strings.TrimSpace(g.Scanner.Text())
        
        position, err := strconv.Atoi(input)
        if err != nil {
            fmt.Println("無効な入力です。0-8の数字を入力してください。")
            continue
        }
        
        if !g.Board.IsValidMove(position) {
            if position < 0 || position > 8 {
                fmt.Println("無効な入力です。0-8の数字を入力してください。")
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
    
    // Check rows
    for i := 0; i <= 6; i += 3 {
        if cells[i] == symbol && cells[i+1] == symbol && cells[i+2] == symbol {
            return true
        }
    }
    
    // Check columns
    for i := 0; i <= 2; i++ {
        if cells[i] == symbol && cells[i+3] == symbol && cells[i+6] == symbol {
            return true
        }
    }
    
    // Check diagonals
    if cells[0] == symbol && cells[4] == symbol && cells[8] == symbol {
        return true
    }
    if cells[2] == symbol && cells[4] == symbol && cells[6] == symbol {
        return true
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
