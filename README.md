# zk-batch

> zk-batch help easy to import or export zookeeper data



## import data

```bash
go run main.go -s locahost:2181 -u digest -a admin:admin  import -f a.txt
```

## export data

```bash
go run main.go -s locahost:2181  -u digest -a admin:admin export -p / -f a.txt
```