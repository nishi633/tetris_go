package main

import (
	"github.com/nsf/termbox-go"
	"math/rand"
	"time"
)

var currentBlock [DisplayY][DisplayX]rune

// 次に落ちてくるブロック
func nextTetrimino() Block {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(blockList))
	return blockList[i]
}

func drawBlock(x, y int, block Block) {
	point := block.Point
	display := []rune(block.Display)[0]

	clearCurrentBlock()

	for r := 0; r < len(point); r++ {
		for c := 0; c < len(point[r]); c++ {
			if point[r][c] {
				currentBlock[y+r][x+c] = display
				termbox.SetCell((x+c)*StrWidth, y+r, display, Coldef, Coldef)
			}
		}
	}
}

func clearCurrentBlock() {
	for r := 0; r < len(currentBlock); r++ {
		for c := 0; c < len(currentBlock[r]); c++ {
			currentBlock[r][c] = 0
		}
	}
}

func turn(input Block) Block {
	point := input.Point
	newRow := len(point[0])
	newColumn := len(point)
	block.Point = make([][]bool, newRow)
	// 初期化
	for r := 0; r < newRow; r++ {
		block.Point[r] = make([]bool, newColumn)
	}
	// ブロック設定
	for r := 0; r < newColumn; r++ {
		for c := 0; c < newRow; c++ {
			block.Point[c][r] = point[r][c]
		}
	}
	return block
}

/*********  ブロック定義 *********/

type Block struct {
	Display string
	Point   [][]bool
}

var blockList = []Block{
	square,
	rectangulare,
	t,
	z,
	s,
	l,
	j,
}

var square = Block{
	Display: "正",
	Point: [][]bool{
		{true, true},
		{true, true},
	},
}

var rectangulare = Block{
	Display: "長",
	Point: [][]bool{
		{true, true, true, true},
	},
}

var t = Block{
	Display: "ト",
	Point: [][]bool{
		{false, true, false},
		{true, true, true},
	},
}

var z = Block{
	Display: "Ｚ",
	Point: [][]bool{
		{true, true, false},
		{false, true, true},
	},
}

var s = Block{
	Display: "Ｓ",
	Point: [][]bool{
		{false, true, true},
		{true, true, false},
	},
}

var l = Block{
	Display: "Ｌ",
	Point: [][]bool{
		{true, false},
		{true, false},
		{true, true},
	},
}

var j = Block{
	Display: "Ｌ",
	Point: [][]bool{
		{false, true},
		{false, true},
		{true, true},
	},
}
