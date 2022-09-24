# Go言語による図書サーバ

# 概要について

# 機能について

* 機能APIについて

|url|Method|説明|備考|
|--|--|--|--|
|/health|*|healthチェック機能||
|/v1/upload|POST|ファイルのアップロード機能||
|/v1/read/[テーブル]/|LIST|データベース内のテーブルデータすべて読み取り||
|/v1/read/[テーブル]/[id]/GET|データベース内のIDを指定して読み取る||

# 設定可能な環境変数

|名前|説明|初期値|備考|
|--|--|--|--|
|PROTOCOL|プロトコル名|tcp||
|WEB_HOST|ホスト名|空白||
|WEB_PORT|解放ポート|8080||
|DB_NAME|SQLのデータベースタイプ|mysql|sqlite3も可能|
|DB_HOST|SQLのIPアドレス|127.0.0.1||
|DB_PORT|SQLの接続ポート|3306||
|DB_USER|SQLの接続ユーザ||
|DB_PASS|SQLの接続ユーザパスワード||
|DB_FILE|SQLite3の接続ファイルパス|test.db|