package main

import(
  "time"
  "math/rand"
  "github.com/nsf/termbox-go"
)

func nextTetrimino() Block {
  // 次に落ちてくるブロック
  rand.Seed(time.Now().UnixNano())
  i := rand.Intn(len(blockList))
  return blockList[i]

}

func drawBlock(x, y int, block Block) {
  termbox.Clear(coldef, coldef)
  point := block.Point
  display := []rune(block.Display)[0]
  width := 2

  for r := 0; r < len(point); r++ {
    for c := 0; c < len(point[r]); c++ {
      if point[r][c] {
        termbox.SetCell((x+r)*width, y+c, display, coldef, coldef)
      }
    }
  }
  termbox.Flush()
}


/*********  ブロック定義 *********/

type Block struct {
  Display string
  Point [][]bool
}

var blockList = []Block {
  square,
}

var square = Block {
  Display: "正",
  Point: [][]bool {
    { true, true },
    { true, true },
  },
}

