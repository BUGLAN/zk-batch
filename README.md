# zk-batch

[![Codacy Badge](https://api.codacy.com/project/badge/Grade/4bdf02c2fdab431ca2bd97f2b64515f0)](https://app.codacy.com/manual/BUGLAN/zk-batch?utm_source=github.com&utm_medium=referral&utm_content=BUGLAN/zk-batch&utm_campaign=Badge_Grade_Dashboard)
[![Go Report Card](https://goreportcard.com/badge/github.com/BUGLAN/zk-batch)](https://goreportcard.com/report/github.com/BUGLAN/zk-batch)


> zk-batch help easy to import or export zookeeper data

## install

```bash
go get github.com/BUGLAN/zk-batch
```

## import data

```bash
zk-batch -s locahost:2181 -u digest -a admin:admin  import -f a.txt
```

## export data

```bash
zk-batch  -s localhost:2181  -u digest -a admin:admin export -p / -f a.txt
```