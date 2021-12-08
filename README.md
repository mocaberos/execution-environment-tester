# 実行環境確認＆テスト用dockerイメージ

AWS ECS でネットワークの接続確認や環境変数確認など、インフラ調整用のdockerイメージ

## ルーティング一覧
- `/`
  - 説明
- `/ip`
  - 接続して来たクライアントのIP表示
- `/env/:key`
  - 任意の環境変数を取得する
- `/usage`
  - システム使用状況を取得する

## 使い方
```text
# Public
FROM mocaberos/execution-environment-tester:latest

# Private
FROM 085041388644.dkr.ecr.ap-northeast-1.amazonaws.com/execution-environment-tester:latest
```
