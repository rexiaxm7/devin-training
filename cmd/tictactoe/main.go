package main

import (
    "fmt"
    "github.com/user/devin-training/pkg/game"
)

func main() {
    fmt.Println("○×ゲーム（Tic-Tac-Toe）")
    g := game.NewGame()
    g.Play()
}
