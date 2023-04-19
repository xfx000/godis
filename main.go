package godis

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func ListenAndServe(address string) {
	//绑定监听地址
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(fmt.Sprintf("listen err:%v", err))
	}
	defer listener.Close()
	log.Println(fmt.Sprintf("bind:%s,start listening...", address))

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(fmt.Sprintf("accept err:%v", err))
		}
		go Handler(conn)
	}
}

func Handler(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		//ReadString 会一直阻塞遇到分隔符 '\n'
		msg, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				log.Println("connection close")
			} else {
				log.Println(err)
			}
			return

		}
		b := []byte(msg)
		//将接受到的消息发送给客户端
		conn.Write(b)
	}
}

func main() {
	ListenAndServe(":8000")

}
