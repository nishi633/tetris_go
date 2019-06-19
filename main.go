package main

import (
  "os"
  "time"
  "github.com/nsf/termbox-go"
)

const(
  displayRow int = 30
  displayColumn int = 12
  coldef = termbox.ColorDefault
)

var screan [displayRow]string
var count int = 0

func main() {
  if err := termbox.Init(); err != nil {
  }
  defer termbox.Close()
  drawBlock(displayColumn/2, 0, square) // 初期表示

  t := time.NewTicker(1 * time.Second)
  defer t.Stop()

	timerCh := make(chan bool)
	keyCh := make(chan termbox.Key)

  // 処理を待ち受けるために用意する
  go TimerLoop(timerCh)
  go KeyEventLoop(keyCh)

  // メインの処理ループ
  // 該当チャンネルになにか来ればなにかするし、来なければbreak
  for {
    select {
    case <-keyCh:
      break
    case <-timerCh:
      break
    default:
      break
    }
  }
}

//タイマーイベント
// ブロックを下に落とす
func TimerLoop(tch chan bool) {
	for {
		tch <- true
    count += 1
    drawBlock(displayColumn/2, count, square)
		time.Sleep(1 * time.Second)
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
