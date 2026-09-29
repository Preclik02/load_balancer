package main

import (

	"fmt"

)

func main() {

	var users_to_split int
	//var servers int

	server_8080 := []int{}
	server_8081 := []int{}
	server_8082 := []int{}

	fmt.Printf("[-] How much users >> ")
	fmt.Scan(&users_to_split)


	// -- SELECTING SERVER FOR EACH USER -- // 
	selected_server := -1
	for i := 0; i < users_to_split; i++ {
	
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


	// -- FOR LOOPS TO CONNECT ALL USERS TO SERVERS || PORTS -- //

	


}
