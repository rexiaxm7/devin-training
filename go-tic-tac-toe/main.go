package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Board struct {
	cells [9]string
}

func NewBoard() *Board {
	return &Board{cells: [9]string{" ", " ", " ", " ", " ", " ", " ", " ", " "}}
}

func (b *Board) Display() {
	fmt.Println()
	fmt.Printf(" %s | %s | %s \n", b.cells[0], b.cells[1], b.cells[2])
	fmt.Println("-----------")
	fmt.Printf(" %s | %s | %s \n", b.cells[3], b.cells[4], b.cells[5])
	fmt.Println("-----------")
	fmt.Printf(" %s | %s | %s \n", b.cells[6], b.cells[7], b.cells[8])
	fmt.Println()
}

func (b *Board) DisplayWithNumbers() {
	fmt.Println("マス目の番号は以下のように配置されています：")
	fmt.Println("0 | 1 | 2")
	fmt.Println("---------")
	fmt.Println("3 | 4 | 5")
	fmt.Println("---------")
	fmt.Println("6 | 7 | 8")
	fmt.Println()
}

func (b *Board) IsCellEmpty(position int) bool {
	return b.cells[position] == " "
}

func (b *Board) PlaceMark(position int, mark string) bool {
	if position < 0 || position > 8 {
		return false
	}
	if !b.IsCellEmpty(position) {
		return false
	}
	b.cells[position] = mark
	return true
}

func (b *Board) IsFull() bool {
	for _, cell := range b.cells {
		if cell == " " {
			return false
		}
	}
	return true
}

func (b *Board) CheckWin(mark string) bool {
	for i := 0; i < 9; i += 3 {
		if b.cells[i] == mark && b.cells[i+1] == mark && b.cells[i+2] == mark {
			return true
		}
	}
	for i := 0; i < 3; i++ {
		if b.cells[i] == mark && b.cells[i+3] == mark && b.cells[i+6] == mark {
			return true
		}
	}
	if b.cells[0] == mark && b.cells[4] == mark && b.cells[8] == mark {
		return true
	}
	if b.cells[2] == mark && b.cells[4] == mark && b.cells[6] == mark {
		return true
	}
	return false
}

type Game struct {
	board       *Board
	currentMark string
}

func NewGame() *Game {
	return &Game{
		board:       NewBoard(),
		currentMark: "○",
	}
}

func (g *Game) SwitchPlayer() {
	if g.currentMark == "○" {
		g.currentMark = "×"
	} else {
		g.currentMark = "○"
	}
}

func (g *Game) Play() {
	scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Println("まるばつゲーム（Tic-Tac-Toe）へようこそ！")
	g.board.DisplayWithNumbers()
	
	for {
		fmt.Printf("現在の盤面:\n")
		g.board.Display()
		
		fmt.Printf("プレイヤー %s の番です。0-8の番号を入力してください: ", g.currentMark)
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		
		position, err := strconv.Atoi(input)
		if err != nil || position < 0 || position > 8 {
			fmt.Println("無効な入力です。0から8までの数字を入力してください。")
			continue
		}
		
		if !g.board.PlaceMark(position, g.currentMark) {
			fmt.Println("そのマスはすでに埋まっています。別のマスを選んでください。")
			continue
		}
		
		if g.board.CheckWin(g.currentMark) {
			g.board.Display()
			fmt.Printf("%s の勝ち！\n", g.currentMark)
			break
		}
		
		if g.board.IsFull() {
			g.board.Display()
			fmt.Println("引き分け！")
			break
		}
		
		g.SwitchPlayer()
	}
	
	fmt.Print("もう一度プレイしますか？ (y/n): ")
	scanner.Scan()
	replay := strings.ToLower(strings.TrimSpace(scanner.Text()))
	if replay == "y" {
		g = NewGame()
		g.Play()
	} else {
		fmt.Println("ゲームを終了します。ありがとうございました！")
	}
}

func main() {
	game := NewGame()
	game.Play()
}
