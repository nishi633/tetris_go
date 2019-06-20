package main

import(
  //"fmt"
  "github.com/nsf/termbox-go"
)

var (
  frame = map[string]int { "top": 2, "botton": 2, "rignt": 2, "left": 2 }
  wall = "⬜"
)

// 表示枠
func mainScrean() {
  wallRune := []rune(wall)[0]

  for r := 0; r < displayRow; r++ {
    for c := 0; c < frame["top"]; c++ {
      termbox.SetCell(r*strWidth, c, wallRune, coldef, coldef)
    }
  }

  for r := 0; r < displayRow; r++ {
    for c := displayColumn - frame["botton"]; c < displayColumn; c++ {
      termbox.SetCell(r*strWidth, c, wallRune, coldef, coldef)
    }
  }

  for r := 0; r < frame["left"]; r++ {
    for c := 0; c < displayColumn; c++ {
      termbox.SetCell(r*strWidth, c, wallRune, coldef, coldef)
    }
  }

  for r := displayRow - frame["left"]; r < displayRow; r++ {
    for c := 0; c < displayColumn; c++ {
      termbox.SetCell(r*strWidth, c, wallRune, coldef, coldef)
    }
  }
}

// 画面の状況を記録
func captureScrean() {
}
