package utils

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
)

var (
	users             map[net.Conn]string
	chatSync          sync.Mutex
	saveMessageMutex  sync.Mutex
	messages          []string
)

func HandleConn(conn net.Conn) {
	defer func() {
		chatSync.Lock()
		name := users[conn]
		delete(users, conn)
		chatSync.Unlock()
		Announce(fmt.Sprintf("%s has left the chat...", name), conn)
		conn.Close()
	}()

	conn.Write([]byte("Welcome to TCP-Chat!\n" +
		"         _nnnn_\n" +
		"        dGGGGMMb\n" +
		"       @p~qp~~qMb\n" +
		"       M|@||@) M|\n" +
		"       @,----.JM|\n" +
		"      JS^\\__/  qKL\n" +
		"     dZP        qKRb\n" +
		"    dZP          qKKb\n" +
		"   fZP            SMMb\n" +
		"   HZM            MMMM\n" +
		"   FqM            MMMM\n" +
		" __| \".        |\\dS\"qML\n" +
		" |    .       | ' \\Zq\n" +
		"_)      \\.___.,|     .'\n" +
		"\\____   )MMMMMP|   .'\n" +
		"     -'       --'\n" +
		"[ENTER YOUR NAME]: "))

	name, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Println("Error reading name:", err)
		return
	}
	name = strings.TrimSpace(name)
	if name == "" {
		conn.Write([]byte("Invalid name. Connection closing.\n"))
		return
	}

	chatSync.Lock()
	users[conn] = name
	chatSync.Unlock()

	Announce(fmt.Sprintf("%s has joined our chat...\n", name), conn)
	LoadMessages(conn, messages)
	
	saveMessageMutex.Lock()
	messages = append(messages, fmt.Sprintf("%s has joined our chat...\n", name))
	saveMessageMutex.Unlock()
	

	go handleMessages(conn, name)

	select {} // Keep the connection alive until it's closed
}
