# mission-reward

## 動かすのに必要な環境

Go 1.20
Docker 24.0.2

## 動作確認手順

### 初期セットアップ方法

```
make tools
docker compose up -d db
make migrate-up
```

### 検証方法

#### サーバーの起動

1. サーバーの起動をします

```
docker-compose up -d
```

#### クライアントの起動

1. 別セッションを、立ち上げてクライアントの起動をします

```
make run-client
```

## 残課題

- モンスターを倒すミッションの実装
- レベルアップミッションの実装
- アイテム取得ミッションの実装
- 冪等になるように upsert を使うようにする
- テストの追加
- チャンネルを適切に閉じる graceful shutdown
- 適切なエラーハンドリング
