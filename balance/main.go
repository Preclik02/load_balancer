package main

import (

	"flag"
	"fmt"
	"sync"
	"net"
	"time"

)
///////////////////////////
// -- FUNC IS HEALTHY -- //
///////////////////////////
func is_healthy(healthy_servers []string, server string) int {

	var check_if_healthy int

	for _, i := range healthy_servers {
		if i == server {
			check_if_healthy += 1
		}
	}

	if check_if_healthy == 1 {
		return 1
	} else {
		return 0
	}

}

/////////////////////////////////
// -- FUNCTION CHECK SERVER -- //
/////////////////////////////////
func check_server(server string, timeout time.Duration) int {
	conn, err := net.DialTimeout("tcp", "localhost:"+server, timeout)
	if err != nil {
		return 0
	}

	defer conn.Close()

	return 1
}

////////////////////////////////
// -- FUNCION CONNECT USER -- //
////////////////////////////////
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

	servers := []string{"8080", "8081", "8082"}
	healthy_servers := []string{ }
	var healthy_servers_int int

	users_to_split := flag.Int("users", 0, "int")
	

	flag.Parse()

	// -- CHECK SERVERS FUNCTION -- //
	for _, i := range servers {
		
		if check_server(i, 500 * time.Millisecond) == 0 {
			fmt.Printf("[+] Server %s is down\n", i)
		} else if check_server(i, 500 * time.Millisecond) == 1 {
			healthy_servers = append(healthy_servers, i) 
			healthy_servers_int += 1
			fmt.Printf("[+] Server %s is up\n", i)
		}
	}
	if healthy_servers_int <= 0 { return }
	// ^^ CHECKS SERVERS FUNCTION ^^ //


	// -- SELECTING SERVER FOR EACH USER -- // 
	selected_server := -1
	for i := 0; i < *users_to_split; i++ {
	
		selected_server += 1
		switch {
		case selected_server == 0:
			if is_healthy(healthy_servers, "8080") == 1 {
				server_8080 = append(server_8080, i)
			} else {
				i -= 1
			}
		case selected_server == 1:
			if is_healthy(healthy_servers, "8081") == 1 {
				server_8081 = append(server_8081, i)
			} else {
				i -= 1
			}
		case selected_server == 2:
			if is_healthy(healthy_servers, "8082") == 1 {
				server_8082 = append(server_8082, i)
			} else {
				i -= 1
			}
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
