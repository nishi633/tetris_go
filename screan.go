package main

import(
  "github.com/nsf/termbox-go"
)

var (
  frame = map[string]int { "top": 2, "botton": 2, "rignt": 2, "left": 2 }
  wall = "⬜"
  screan [DisplayY][DisplayX] rune
)

// 表示枠
func mainScrean() {
  wallRune := []rune(wall)[0]

  for r := 0; r < DisplayY; r++ {
    for c := 0; c < frame["top"]; c++ {
      screan[r][c] = wallRune
    }
  }

  for r := 0; r < DisplayY; r++ {
    for c := DisplayX - frame["botton"]; c < DisplayX; c++ {
      screan[r][c] = wallRune
    }
  }

  for r := 0; r < frame["left"]; r++ {
    for c := 0; c < DisplayX; c++ {
      screan[r][c] = wallRune
    }
  }

  for r := DisplayY - frame["left"]; r < DisplayY; r++ {
    for c := 0; c < DisplayX; c++ {
      screan[r][c] = wallRune
    }
  }
  drawScrean()
}

func drawScrean() {
  for r := 0; r < DisplayY; r++ {
    for c := 0; c < DisplayX; c++ {
      termbox.SetCell(c*StrWidth, r, screan[r][c], Coldef, Coldef)
    }
  }
}

func captureBlock() {
  for r := 0; r < len(currentBlock); r++ {
    for c := 0; c < len(currentBlock[r]); c++ {
      screan[r][c] = currentBlock[r][c]
    }
  }
}
