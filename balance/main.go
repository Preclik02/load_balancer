package main

import (

	"flag"
	"fmt"
	"sync"
	"net"


)

func connect_user(user int, port string) {

	// -- CONNECTING A USER TO A SERVER (IN THIS CASE A PORT) -- //
	conn, err := net.Dial("tcp", "localhost:"+port)
	if err != nil {
		fmt.Printf("[+] An error happened when connecting user %d to port %s\n", user, port)
		return
	}

	defer conn.Close()

	fmt.Printf("[+] User %d connected successfully to port %s\n", user, port)

}


func main() {

	var wg sync.WaitGroup

	server_8080 := []int{}
	server_8081 := []int{}
	server_8082 := []int{}

	users_to_split := flag.Int("users", 0, "int")

	flag.Parse()

	// -- SELECTING SERVER FOR EACH USER -- // 
	selected_server := -1
	for i := 0; i < *users_to_split; i++ {
	
		selected_server += 1
		switch {
		case selected_server == 0:
			server_8080 = append(server_8080, i)
		case selected_server == 1:
			server_8081 = append(server_8081, i)
		case selected_server == 2:
			server_8082 = append(server_8082, i)
			selected_server = -1
		default:
			fmt.Println("[+] Error while selecting a server")
		}
	
	
	}
	fmt.Println(server_8080)
	fmt.Println(server_8081)
	fmt.Println(server_8082)

	
	// -- At the end of the program wait for goroutines to end -- //
	defer wg.Wait()
	
	// -- FOR loops for connection to the server -- //
	// -- 1ST FOR LOOP -- //
	for _, user := range server_8080 {
		wg.Add(1)
		go func (user int) {
			defer wg.Done()
			connect_user(user, "8080")
		}(user)
	}

	// -- 2ND FOR LOOP -- //
	for _, user := range server_8081 {
		wg.Add(1)
		go func(user int){
			defer wg.Done()
			connect_user(user, "8081")
		}(user)
	}

	// -- 3RD FOR LOOP -- // 
	for _, user := range server_8082 {
		wg.Add(1)
		go func(user int) {
			defer wg.Done()
			connect_user(user, "8082")
		}(user)
	}
	// -- ALL USERS "CONNECTED" -- //

}
