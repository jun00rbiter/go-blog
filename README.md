# go-blog - ブログAPIアプリ

Go言語フレームワーク gin を使った RestFUL API アプリ。
Go言語開発環境構築のためのサンプルプログラム。

## API

ヘルスチェックAPI
~~~
# GET /api/health-check
+ Response 200 (application/json)
  + Body
    { "status": "OK" }
~~~
