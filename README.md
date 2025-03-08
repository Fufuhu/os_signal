# OSシグナル処理サンプル

OSシグナルを受け取って処理をするためのサンプルです

## 実行方法

```bash
go build
./os_signal
```

## シグナルの送り方

```bash
$ ps -ef | grep os_signal # os_signalのPIDを確認
$ kill ${os_signalのPID}
```

`kill -9`はコード上では対応しているように見えますが、プロセスが強制的に