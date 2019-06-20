package main

import (
  "os"
  "time"
  "github.com/nsf/termbox-go"
)

const(
  displayRow = 20
  displayColumn = 30
  coldef = termbox.ColorDefault
  fallSpan = 100 * time.Millisecond
  strWidth = 2
)

var (
  screan [displayRow]string
  block Block
  count int = 0
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

  // メインの処理ループ
  // 該当チャンネルになにか来ればなにかするし、来なければbreak
  for {
      block = nextTetrimino()
    for {
      select {
      case <-keyCh:
        break
      case <-timerCh:
        break
      case <- blockCh:
        // この状態を次のターンに引き継ぐ
        captureScrean()
        break
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
    count += 1

    fallHeight := displayColumn - frame["botton"] - 1

    if count < fallHeight {
      termbox.Clear(coldef, coldef)
      mainScrean()
      drawBlock(displayRow/2 - 1, count, block)
      termbox.Flush()

		  time.Sleep(fallSpan)
    } else {
      bch <- true
    }
	}
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
          case termbox.KeyArrowRight:
              //右キーを押された時の処理
          default:
          }
      default:
      }
	}
}
