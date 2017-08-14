# 脆弱なウェブアプリケーションのサンプル for Go

Go 言語による、基本的なウェブアプリケーションセキュリティの学習用サンプルアプリケーションです。

## 動作要件

* Go 1.6.x 以上
* GCC
  * 依存している cgo パッケージのインストールのため

これらの動作要件を満たせない場合、 Docker イメージを用いることもできます。その場合は Docker をインストールしている必要があります。

## インストール、実行手順

以下のどちらかの手順を選択してください。

### 本アプリを直接ローカルで実行する場合

以下のようにしてインストールをおこなってください。

```
$ go get github.com/co3k/go-webvuln
$ cd `go env GOPATH`/src/github.com/co3k/go-webvuln
$ make install
```

その後、以下のコマンドを実行してください。

```
$ make server
```

`Starting up the server` と表示されれば成功です。ブラウザ (Edge, Firefox, Chrome, Safari の最新安定版を推奨します) で http://localhost:8000/ を開いてください。

### 本アプリを Docker 経由で実行する場合

Go 公式の Docker イメージを利用しています。

以下のようにしてインストールをおこなってください。

```
$ go get github.com/co3k/go-webvuln
$ cd `go env GOPATH`/src/github.com/co3k/go-webvuln
$ docker build -t go-webvuln .
```

その後、以下のコマンドを実行してください。

```
$ docker-compose up
```

`Starting up the server` と表示されれば成功です。ブラウザ (Edge, Firefox, Chrome, Safari の最新安定版を推奨します) で http://localhost:8000/ を開いてください。
