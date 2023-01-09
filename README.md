# backend-guchitter-app
## 実行
- `$ cd backend-guchitter-app`
- `$ go run .`
### Swagger UI の表示
- `$ go run .`のあと下にアクセス
- http://localhost:8080/swagger/index.html
  - `interface/handler`のメソッド上にGo Docで記載したAPI仕様が表示されるよ。
  - Go Docを修正したら`swag init`のあと`go run .`で反映されます。
### Go Doc の表示
- `$ godoc -http=:{port}`のあと、`localhost:{port}`にアクセス
  - `Third party`をクリックでこのソースの内容が見れます。
  - 外部にexportされていないメソッドも見たい場合、URL末尾に`?m=all`を追加してください。

## 開発
### ライブラリの追加

```go:xxx.go
import (
    "gorm.io/gorm"
)
```

- `$ go get`を実行し、ライブラリを追加する
- `$ go mod tidy`で依存関係を解決する

### ロギングの実装
- `logging/logger.go`をimportし、`Log`を使用してください
- 例:
```go
import (
  "github.com/backend-guchitter-app/logging"
  "github.com/bloom42/rz-go"
)
...
func hoge() {
  logging.Log.Info("hoge() started.")
  ...
  if err != nil {
    logging.Log.Error("Failed at FindAll()", rz.Err(err))
  }
}
```

### Migration
- 1.Migrationファイルの作成
```sh
make create-migration
```
- 2.作成されたMigrationにSQLを記載
  - ⚠Gormの規約で、モデル名を複数形にしたテーブル名にしなければいけない⚠
  - 例：モデル名:`User` => テーブル名は`users`
- 3.Up
```sh
make migrateup
```
  - ※新しく作成したMigrationファイルに対するMigrationについて、初回はエラーが出るので二度実施する(一度目でMigrate自体は実施されるだが、schema_migration の関係でエラーが出る。二度目はMigrateが実施されないが、エラーが出ない(dirty が解消される))

## 環境
- 設定ファイルは下記。`GUCHITTER_ENV`の値に応じて読み込まれる。
  - 本番:`.env.production`
  - 開発:`.env.development`


