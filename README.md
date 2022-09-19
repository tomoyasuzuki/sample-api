# sample-api
sample repository of rest api.

※ 免責事項
あくまで素振り用のリポジトリなのでローカルDBの機密情報などが含まれています。このリポジトリのコードを流用する場合でも、パブリッククラウドの認証情報などは絶対にアップロードしないようにしてください。

# Build 
```
docker build -t sample-api:latest -f app/docker/Dockerfile .
```

# Run
```
docker run -it -p 80:80 sample-api:latest
```

# License
MIT
