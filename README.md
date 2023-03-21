# 実行手順

1. docker環境の立ち上げ
   - `docker-compose up -d`
2. grpcurlのinstall
   - `brew install grpcurl`
3. ストーリー実行前のユーザー情報とミッション達成状況の表示
   - `./check-data.sh`
4. ストーリーの実行
   - `./grpcurl-exec.sh`
5. ストーリー実行後のユーザー情報とミッション達成状況の表示
   - `./check-data.sh`
