# Sample Go Echo

## 実行方法

```sh
# ローカル環境
docker-compose up -d
docker-compose exec app bash
```

```sh
# 以降Dockerコンテナ内
# 即時実行
go run .

# ビルド
go build

# ビルド結果の実行
./api-server
```
