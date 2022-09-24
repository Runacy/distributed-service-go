- 起動

```
cd cmd/server/
go run ./main.go
```

```
# 適当なところでcurlしてみる(例)
# 書き込み
curl -X POST localhost:8880 -d '{"record":{"value":"vvfdss2311fdsfs8"}}'

# 読み込み
curl -X GET localhost:8880 -d '{"offset":0}'

```