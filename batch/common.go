package batch

import (
	"github.com/samuel/go-zookeeper/zk"
	"github.com/urfave/cli/v2"
	"time"
)

const sep = "::"

var (
	flags int32 = 0
	acls        = zk.WorldACL(zk.PermAll)
)

type ZkData struct {
	Path string
	Data string
}

func conn(host string) (conn *zk.Conn, err error) {
	conn, _, err = zk.Connect([]string{host}, time.Second*3)
	return conn, err
}

func addAuth(conn *zk.Conn, digest string, auth string) (err error) {
	return conn.AddAuth(digest, []byte(auth))
}

func checkAuth(ctx *cli.Context, conn *zk.Conn) (err error) {
	return addAuth(conn, ctx.String("digest"), ctx.String("auth"))
}
