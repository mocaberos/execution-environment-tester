# 実行環境確認＆テスト用dockerイメージ

[![Build Status](https://codebuild.ap-northeast-1.amazonaws.com/badges?uuid=eyJlbmNyeXB0ZWREYXRhIjoibUNrWmtLOWRTTFArWTc0eUZQcFZxWHJCZ0lTbko3eDhFbWFvUXZxbjEzdWJuRjc5akxqSGFwT1pCVzBCcUFHYjJFdERMYTFFL1BNRk5hS3NVdHcydWJRPSIsIml2UGFyYW1ldGVyU3BlYyI6ImpDb0xyUEpPaExDdUIrOU4iLCJtYXRlcmlhbFNldFNlcmlhbCI6MX0%3D&branch=master)](https://ap-northeast-1.console.aws.amazon.com/codesuite/codebuild/085041388644/projects/execution-environment-tester)

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
