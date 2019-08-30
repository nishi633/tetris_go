package main

import(
  "github.com/nsf/termbox-go"
)

var (
  frame = map[string]int { "top": 2, "botton": 3, "right": 2, "left": 2 }
  wall = "⬜"
  screan [DisplayY][DisplayX] rune
  piledBlock [DisplayY][DisplayX] rune
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
    for c := DisplayX - frame["right"]; c < DisplayX; c++ {
      screan[r][c] = wallRune
    }
  }

  for r := 0; r < frame["left"]; r++ {
    for c := 0; c < DisplayX; c++ {
      screan[r][c] = wallRune
    }
  }

  for r := DisplayY - frame["botton"]; r < DisplayY; r++ {
    for c := 0; c < DisplayX; c++ {
      screan[r][c] = wallRune
    }
  }
}

func drawScrean() {
  for r := 0; r < DisplayY; r++ {
    for c := 0; c < DisplayX; c++ {
      termbox.SetCell(c*StrWidth, r, screan[r][c], Coldef, Coldef)
      if piledBlock[r][c] != 0 {
        termbox.SetCell(c*StrWidth, r, piledBlock[r][c], Coldef, Coldef)
      }
    }
  }
}

// pile 現在のブロックをpiledBlockに保存する
func pile() {
  for r := 0; r < DisplayY; r++ {
    for c := 0; c < DisplayX; c++ {
      if currentBlock[r][c] != 0 {
        piledBlock[r][c] = currentBlock[r][c]
      }
    }
  }
}

// captureBlock 現在のブロック位置をscrean反映
func captureBlock() {
  for r := 0; r < DisplayY; r++ {
    for c := 0; c < DisplayX; c++ {
      if currentBlock[r][c] != 0 {
        screan[r][c] = currentBlock[r][c]
      }
    }
  }
}

// deleteLine 埋まった行を削除する
func deleteLine() {
}

