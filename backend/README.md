# backend

## 必要な環境

- Go 1.20 以上
- Docker
- [Task](https://taskfile.dev/)

`task up`を一回実行したら、Go のプログラムを書き換えると自動的に反映されます。`go run main.go`などを実行する必要はありません。

## Task コマンド

詳しくは [Taskfile.yml](./Taskfile.yml)

- `task up`
  - GoのプログラムとMySQL、Adminerを立ち上げる。http://localhost:8081 にアクセスするとAdminerにアクセスできます。Goのサーバーはポート8080、MySQLは3306で立ち上がります。
- `task down`
  - GoとMySQLを閉じる。
- `task db`
  - MySQLに入ってSQLで操作する。
  - `exit`で抜けられる。
- `task log`
  - Go のログを見る。
- `task reset`
  - データベースを初期化する。

### Task の導入

Go が入っていることが前提

```sh
go install github.com/go-task/task/v3/cmd/task@latest
```

asdf で Go を入れた場合はその後 `asdf reshim golang`
直接入れた場合はパスを通す。
