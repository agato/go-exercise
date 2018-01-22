# チャットアプリケーションを作ろう  
 
  
## 準備
  
  lesson3をベースディレクトリにして起動してください
  必要なライブラリはglideなりを使って取得してください。  
  sampleの起動は、下記のようなコマンドを実行してください  
  ```
  PORT=8081 go run main.go
  ```
  
## ソースを見てみよう

#### クライアントの型を定義を見てみよう  
　　  
他のオブジェクト指向の言語のクラスを使って表現することはGoでは型を使って表現することに該当します。
   
説明
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
    
2.1 クライアントの型を定義します　　  
    他のオブジェクト指向の言語のクラスを使って表現することはGoでは型を使って表現することに
    該当します。
     
     
     
     
##### 今回の学び  

  - golangの導入は簡単にできる
  - dockerにふれて興味がでる  
  [コマンドはこの辺りを参照](https://qiita.com/curseoff/items/a9e64ad01d673abb6866)  
  [仕組みなど](http://morizyun.github.io/docker/about-docker-command.html)
    



