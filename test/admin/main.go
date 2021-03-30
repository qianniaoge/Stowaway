/*
 * @Author: ph4ntom
 * @Date: 2021-03-11 16:10:51
 * @LastEditors: ph4ntom
 * @LastEditTime: 2021-03-26 17:39:03
 */
package main

import (
	"Stowaway/utils"
	"fmt"
	"io"
	"net"

	"github.com/fwhezfwhez/errorx"
	"github.com/xtaci/kcp-go"
)

type socks struct {
	aaa map[int]int
}

func main() {
	var b int
	fmt.Printf("addr is %p,b is %d\n", &b, b)
	b, err := utils.Str2Int("1")
	fmt.Printf("addr is %p,b is %d\n", &b, b)
	mapa := make(map[int]*socks)
	_, ok := mapa[1]
	print("ok is", ok)
	fmt.Println([]byte("6111111111111111111"))
	fmt.Println("kcp listens on 10000")
	lis, err := kcp.ListenWithOptions(":10000", nil, 10, 3)
	if err != nil {
		panic(err)
	}
	for {
		conn, e := lis.AcceptKCP()
		if e != nil {
			panic(e)
		}
		go func(conn net.Conn) {
			var buffer = make([]byte, 1024, 1024)
			for {
				n, e := conn.Read(buffer)
				if e != nil {
					if e == io.EOF {
						break
					}
					fmt.Println(errorx.Wrap(e))
					break
				}

				fmt.Println("receive from client:", buffer[:n])
			}
		}(conn)
	}
}