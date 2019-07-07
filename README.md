# TetrisGo

go言語の練習にテトリスを作成  

https://github.com/kurehajime/kuzusi/blob/master/kuzusi.go
https://qiita.com/kurehajime/items/a02eaae3f0c17750c07e

# 設計

```
main.go
tetrimino.go // 落ちてくるブロックの定義
```

# 設計
## ルール
複数種類のブロックがランダムに選択され
ブロックは1秒毎に1マス落下する  
横の列すべてが埋まったらその列のブロックをすべて消す  


## ブロック

```
正正
正正
```

```
長長長長
```

```
  ト
トトト
```

```
ＺＺ
  ＺＺ
```

```
  ＳＳ
ＳＳ
```

```
Ｌ
Ｌ
ＬＬ
```

```
  Ｊ
  Ｊ
ＪＪ
```

# termbox-goの使い方

[termbox-go の基本的な使い方](https://qiita.com/zenwerk/items/98f4db3285777a582fd0)

## 初期設定
```
if err := termbox.Init(); err != nil {
  panic
}
defer termbox.Close()
```

## 表示の設定
```
termbox.Clear(coldef, coldef)
termbox.SetCell(x, y, '┏', coldef, coldef)
termbox.Flush() // これを呼ぶことで描画される
```

## メインループを作る
これを設定しないと一瞬で処理が終了するため表示されない

## イベントハンドリング
Escキーを押したらループを抜けるなど、exit処理を入れる

## termbox-goのよく使う関数

### イベント実行してなにか帰ってくるまで待つ
ずっと待ってしまうので、if文のようには使えない

```
termbox.PollEvent()
```

### キーイベントを取得
```
termbox.EventKey
```

### Escキーを押したときのイベント
ループのexitに使うことが多そう

```
termbox.KeyEsc
``


# 作り方
「１秒毎にブロックを１マスづつ落下させる」というタイムイベントと、「ESCキーが押されたら処理を止める」というキーイベントを並行して走らせるためにはそれぞれのイベントを受け取るチャンネルを作る必要がある
それらのチャンネルをgoroutineで並行稼動させる

```
func main() {
  keyCh := make(chan termbox.Key) // termboxのKeyイベント取得関数を仕様
  timeCh := make(chan bool)
  
  go KeyEventLoop(keyCh)
  go TimeEventLoop(timeCh)

}

for {
  select {
    case <- keyCh:
    case <- timeCh:
    default:
  }
}
```

## チャンネルとは(復習)

## Todo
matrixパッケージを使って配列の部分を書き換えてみる

## 問題点
行全部埋まっても消えない
回転しない
キーの長押しするとバグる

