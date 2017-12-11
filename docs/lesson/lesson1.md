1. 環境構築（golangのサーバをdockerを使って立てよう)

    1.1 dockerのインストールをする  
        [こちらから](https://docs.docker.com/docker-for-mac/)

    1.1 gitからサンプルをforkして、cloneしてください  
        [こちらから](https://github.com/agato/go-exercise)

    1.2. dockerを使って、goのサーバを立てます   
    DockerFile内にbuild内容を書いてますので、詳細はそちらをご覧ください  
    golangは1.9を使いました。
    ginはhotdeploy用になります。
   
    ```:DockerFile
    #DockerFileを使い、dockerコンテナにgolangをbuildする
    #cloneしたディレクリで下記を実行する
    docker build --rm -t go-exercise-app .
     
    #ビルドしたコンテナを実行する
    docker run -p 3000:3000 -v `pwd`:/go/apps/lesson1 --name="go-exercise-app" -d go-exercise-app
    ```

    1.3.  サーバが立ち上がったか確認する  
     [http://localhost:3000/](http://localhost:3000/)
     
     
     
     
#####今回の学び  

  - golangの導入は簡単にできる
  - dockerにふれて興味がでる  
  [コマンドはこの辺りを参照](https://qiita.com/curseoff/items/a9e64ad01d673abb6866)  
  [仕組みなど](http://morizyun.github.io/docker/about-docker-command.html)
    



