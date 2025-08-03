# ATYLAB API

## 概要

汎用性のあるAPI群を格納するためのサービス

## テスト方法

sh test.sh

## ビルド

sh build.sh

.env は現時点では自動生成されません。
.env_baseから、.envを生成し、必要な内容に書き換えてください

## systemd への登録

.util内のatylab-api.service を /etc/systemd/system に設置

下記を実行し、起動してください

```
sudo systemctl daemon-reload
sudo systemctl enable atylab-api
sudo systemctl start atylab-api
sudo systemctl status atylab-api
```
