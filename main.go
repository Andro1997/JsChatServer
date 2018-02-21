package main

import (
	"fmt" 	// for work with potocks
	"net"	// for work with sockets
	"bufio"	// for work with potocks
	"container/list" // for work with saved sockets
)



var clients *list.List //  save of clients

func handleClient(socket net.Conn) {
	for {
		buffer, err := bufio.NewReader(socket).ReadString('\n')
		if err != nil {
			fmt.Println("User is out!")
			socket.Close()
			return
		}
		for i:= clients.Front(); i != nil; i = i.Next() {
			fmt.Fprint(i.Value.(net.Conn), buffer)
			//bufio.NewWriter(i.Value.(net.Conn)).WriteString(buffer)
		}
	}
}

func main() {
	fmt.Println("server started")
	clients = list.New()

	server,err := net.Listen("tcp",":8080") //get connect with clients
	if err != nil {
		fmt.Println("Error: %s", err.Error()) // ОБРаботка ошибки)
		return
	}

	//CONNECTION
	for {
		client, err := server.Accept()

		if err!= nil {
			fmt.Println("Error: %s", err.Error())
			return
		}
		fmt.Println("Connect new user! :%s",client.RemoteAddr())
		clients.PushBack(client)

		go handleClient(client)//start in potock обработчик клиента ГОРУТИН
	}

}
