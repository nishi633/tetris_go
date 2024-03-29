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
  FallSpan = 200 * time.Millisecond
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
  go timerLoop(timerCh, blockCh)
  go keyEventLoop()

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
func timerLoop(tch, bch chan bool) {
	for {
		tch <- true //これは必要ある？
    shift := 1

    // ブロックがぶつからないか確認してから落下
    if canFall(shift) {
      deleteLine()
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

// canFall ブロックが落下できるか判定
func canFall(shift int) bool {
  isMove := true
  for r := 0; r < len(block.Point); r++ {
    for c := 0; c < len(block.Point[r]); c++ {
      // TODO: shiftが1以上のときには落とせるだけ落とすようにする
      if block.Point[r][c] && screan[y + r + shift][x + c] != 0 {
        isMove = false
        pile()
        break
      }
    }
  }
  return isMove
}

// keyEventLoop キーイベント設定
func keyEventLoop() {
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
            if x + len(block.Point[0]) < DisplayX - frame["right"] {
              x += 1
            }
          case termbox.KeyArrowDown:
            //下キーを押された時の処理
            shift := 2
            if canFall(shift) {
              y += shift
            }
          case termbox.KeySpace:
            // Spaceキーを押された時の処理
            block = turn(block)
          default:
          }
      default:
      }
	}
}
