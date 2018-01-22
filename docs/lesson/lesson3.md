# チャットアプリケーションを作ろう  
 
  
## 準備
  
  lesson3をベースディレクトリにして起動してください
  必要なライブラリはglideなりを使って取得してください。  
  sampleの起動は、下記のようなコマンドを実行してください  
  ```
  PORT=8081 go run main.go
  ```
  
## ポイント


### ServeHTTPを使う

既存のstructや型に対して、ServeHTTPメソッドを用意することで
   http.Handleに登録出来るようにする
   
```
package main

import (
  "fmt"
  "net/http"
)

type String string

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, s)
}

func main() {
  http.Handle("/", String("Hello World."))
  http.ListenAndServe("localhost:8000", nil)
}
```


#### クライアントの型を定義を見てみよう  
　　  
他のオブジェクト指向の言語のクラスを使って表現することはGoでは型を使って表現することに該当します。
   

* このコードではsocketはクライアントと通信を行うwebsocketへの参照を持ちます
* sendはフィールドのバッファ付きのチャンネルです。ここでは、受信したメッセージが
     待ち行列のように蓄積され、websocketを通じて、ブラウザへ送られるのを待機します。
* roomはクライアントがチャットを行なっているチャットルームへの参照が保持されます
    
    
```
package main
        
import  "github.com/gorilla/websocket"
        
//clientはチャットを行う1人のユーザを表します
type client struct {
    socket *websocket.Conn
    send chan []byte
    room *room
}        
```
    
#### forwardは他のクライアントに転送するためのメッセージを保持するチャネルです。  

ここで使っているチャネルはバッファのあるチャネルで、メッセージのための待ち行列
のようなものです。メモリー上に配置され、スレッドセーフな性質を備えています。
複数の送信者や受信者に同時に書き込みができ、ブロックされることはありません  
※バッファに空きがない場合、ブロックされ送受信の同期が行われます。


```
type room struct {
	// forwardは他のクライアントに転送するためのメッセージを保持するチャネルです。
	forward chan []byte
	  :
	  :
```


#### ヘルパを使って複雑さを下げる

本来なら、6行かかる部分を、1行で表せ
`r = newRoom()`のように呼び出すことができます。
なので他の人が実装するときにroomの構造を知らなくても実装できます。
```$xslt
// newRoomはすぐに利用できるチャットルームを生成して返します。
func NewRoom() *room {
	return &room{
		forward: make(chan []byte),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
		Tracer:  trace.Off(),
	}
}
```

     
     
     
     
##### 今回の学び  

  - サーバの起動方法
  - 定義方法の工夫
  - websocketの実装
    



