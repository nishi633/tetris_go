package main

import (
  "os"
  "time"
  "github.com/nsf/termbox-go"
)

const(
  DisplayX = 20
  DisplayY = 30
  Coldef = termbox.ColorDefault
  FallSpan = 300 * time.Millisecond
  StrWidth = 2 //文字幅
)

var (
  block Block
  x int
  y int
)

func main() {
  if err := termbox.Init(); err != nil {
  }
  defer termbox.Close()

	timerCh := make(chan bool)
	keyCh := make(chan termbox.Key)
	blockCh := make(chan bool)

  // 処理を待ち受けるために用意する
  go TimerLoop(timerCh, blockCh)
  go KeyEventLoop(keyCh)

  mainScrean()
  // メインの処理ループ
  // 該当チャンネルになにか来ればなにかするし、来なければbreak
  for {
    // 各落下ブロックの初期設定
    block = nextTetrimino()
    x = DisplayX/2 - len(block.Point[0])/2 - 1
    y = frame["top"]

LOOP:
    for {
      select {
      case <-keyCh:
        break
      case <-timerCh:
        break
      case <- blockCh:
        captureBlock()
        break LOOP
      default:
        break
      }
    }
  }
}

//タイマーイベント
// ブロックを下に落とす
func TimerLoop(tch, bch chan bool) {
	for {
		tch <- true
    shift := 1

    // ブロックがぶつからないか確認してから落下
    if canFall(shift) {
      y += shift
      termbox.Clear(Coldef, Coldef)
      drawScrean()
      drawBlock(x, y, block)
      termbox.Flush()

		  time.Sleep(FallSpan)
    } else {
      bch <- true
    }
	}
}


func canFall(shift int) bool {
  isMove := true
  for r := 0; r < len(block.Point); r++ {
    for c := 0; c < len(block.Point[r]); c++ {
      if screan[y + r + shift][x + c] != 0 {
        isMove = false
        break
      }
    }
  }
  return isMove
}

//キーイベント
// exit処理
func KeyEventLoop(kch chan termbox.Key) {
	for {
      switch ev := termbox.PollEvent(); ev.Type {
      case termbox.EventKey:
          switch ev.Key {
          case termbox.KeyEsc:
            os.Exit(0)
          case termbox.KeyArrowLeft:
            //左キーを押された時の処理
            if x > frame["left"] {
              x -= 1
            }
          case termbox.KeyArrowRight:
            //右キーを押された時の処理
            if x < DisplayX - frame["right"] - len(block.Point) - 1{
              x += 1
            }
          case termbox.KeyArrowDown:
            //下キーを押された時の処理
            shift := 2
            if canFall(shift) {
              y += shift
            }
          default:
          }
      default:
      }
	}
}
