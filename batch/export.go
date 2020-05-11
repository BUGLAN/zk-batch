package batch

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"log"
)

var zkData []ZkData

// Export the path and value to file or console
func Export(ctx *cli.Context) error {
	conn, err := conn(ctx.String("server"))
	if err != nil {
		log.Fatal("连接zk时发生错误: ", err)
	}
	defer conn.Close()

	// auth
	err = checkAuth(ctx, conn)
	if err != nil {
		log.Fatal("zk 验证错误 ", err)
	}

	err = readZkNode(conn, ctx.String("path"))
	if err != nil {
		log.Fatal("读取zk节点时出错 ", err)
	}

	filename := ctx.String("filename")
	if filename == "" {
		console(zkData)
	} else {
		err = writeFile(filename, zkData)
		if err != nil {
			log.Fatal("写入文件时失败: ", err)
		}
	}

	return nil
}

func readZkNode(conn *zk.Conn, path string) (err error) {
	err = readRoot(conn, path)
	if err != nil {
		return
	}
	err = readZkChild(conn, path)
	if err != nil {
		return
	}
	return nil
}

func readRoot(conn *zk.Conn, path string) (err error) {
	data, _, err := conn.Get(path)
	zkData = append(zkData, ZkData{Path: path, Data: string(data)})
	return
}

func readZkChild(conn *zk.Conn, path string) (err error) {
	lists, _, err := conn.Children(path)
	if err != nil && err != zk.ErrNoChildrenForEphemerals {
		return
	}

	if len(lists) == 0 || err == zk.ErrNoChildrenForEphemerals {
		return
	}

	for _, p := range lists {
		if path == "/" {
			pa := path + p
			data, err2 := get(conn, pa)
			if err2 != nil {
				return err2
			}
			zkData = append(zkData, ZkData{Path: pa, Data: data})
			_ = readZkChild(conn, pa)
		} else {
			pa := path + "/" + p
			data, err2 := get(conn, pa)
			if err2 != nil {
				return err2
			}
			zkData = append(zkData, ZkData{Path: pa, Data: data})
			_ = readZkChild(conn, pa)
		}
	}
	return err
}

func get(conn *zk.Conn, path string) (data string, err error) {
	value, _, err := conn.Get(path)
	return string(value), err
}

func console(zkData []ZkData) {
	for _, v := range zkData {
		fmt.Println(v.Path + sep + v.Data)
	}
}

func writeFile(filename string, zkData []ZkData) (err error) {
	data := make([]byte, 0)
	for _, v := range zkData {
		line := []byte(v.Path + sep + v.Data + "\n")
		data = append(data, line...)
	}
	err = ioutil.WriteFile(filename, data, 0600)
	return err
}
