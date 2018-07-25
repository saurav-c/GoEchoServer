package main 
import "net" 
import "fmt" 
import "bufio" 


var clients = make(map[net.Conn]int)
var count int = 0

func serve(conn net.Conn){
	fmt.Println("Accept conn :", conn)
	clients[conn] = count
	count += 1
	for { 
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Print("conn error: ", err)
			return 
		}
		fmt.Print("Message received: ", string(message))
		go broadcast(message)
	}
}

func main() {   
	fmt.Println("Launching server...")
	ln, _ := net.Listen("tcp", "localhost:5000")
	for {
		conn, err := ln.Accept()
		if err != nil {
			return
		}
		go serve(conn)
	}
}

func broadcast(message string) {
	for conn, _ := range clients {
		go func(conn net.Conn, message string) {
			_, err := conn.Write([]byte(message + "\n"))
			if err != nil {
				return
			}
		}(conn, message)
	}
}