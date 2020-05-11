package batch

import (
	"bufio"
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"strings"
)

// Import zookeeper data from file
func Import(ctx *cli.Context) error {
	conn, err := conn(ctx.String("server"))
	if err != nil {
		log.Fatal("连接zk时发生错误: ", err)
	}
	defer conn.Close()

	err = checkAuth(ctx, conn)
	if err != nil {
		log.Fatal("zk 验证错误")
	}

	list, err := parseFile(ctx.String("filename"))
	if err != nil {
		log.Fatal("解析文件失败: ", err)
	}

	for _, v := range list {
		path, err := createOrSet(conn, v.Path, v.Data)
		if err != nil {
			log.Fatal("创建path失败: ", err)
		} else {
			fmt.Println(path, sep, v.Data)
		}
	}

	return nil
}

func createOrSet(conn *zk.Conn, path string, data string) (string, error) {
	exist, state, err := conn.Exists(path)
	if exist && state != nil {
		_, err = conn.Set(path, []byte(data), state.Version)
	} else {
		_, err = conn.Create(path, []byte(data), flags, acls)
	}
	return path, err
}

func parseFile(filename string) (data []ZkData, err error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("打开文件失败: ", err)
	}

	list := make([]ZkData, 0)
	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		l := strings.Split(scanner.Text(), sep)
		list = append(list, ZkData{l[0], l[1]})
	}
	return list, nil
}
