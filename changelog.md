# 変更履歴

[0.2.0]
* ToDo
  * Webからデータベース内のデータ編集機能
  * Webからデータベースのデータ削除機能
  * Webからデータベースへの追加機能

[0.1.2]
* 2022/09/26
  * データベースへの追加機能
  * 読み取りでrowsを開いたときに閉じていなかったバグを修正
* 2022/09/27
  * sqlite3読み取りプログラムで個別処理が必要ないVersionを用意したけ未使用
  * Webからデータベースへの追加機能
[0.1.1]
* 2022/09/26
  * Webからの読み取りでステータスコードの調整
  * Messageの戻り値にステータスコードを追加
  * エラーの戻り値があるときにエラー出力するように変更
[0.1.0]
* 2022/09/25
  * データベースでid指定で読み取り
  * データベースでキーワード読み取り
  * log関数を使用するとエラー出力になるので、message関数を作ってerr出力でないlog出力にする

[0.0.2]
* 2022/09/19
  * Upload機能を追加
  * Configの環境変数情報を一括読み込み
* 2022/09/21
  * SQLite3のデータベース読み込み機能
    * booknames テーブルのみ
* 2022/09/22
  * SQLite3のデータベース読み込み機能
    * filelists
    * copyfile
* 2022/09/23
  * SQLLite3のデータベーステーブル作成機能
* 2022/09/24
  * データベース内のテーブルを特定して一括読み取り

[0.0.1]
* 仮リリース