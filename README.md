# mission-reward

## 動かすのに必要な環境

Go 1.20
Docker 24.0.2

## 動作確認手順
### 初期セットアップ方法

```
make setup
```

### 検証方法
#### サーバーの起動
1. サーバーの起動をします
```
docker-compose up -d
make migrate-up
make run-server
```

#### クライアントの起動
1. 別セッションを、立ち上げてクライアントの起動をします

```
make run-client
```
