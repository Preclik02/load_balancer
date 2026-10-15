// -- Using main package since we did not made more packages or files for now -- //
package main

// -- importing a couple of usefull libs because our program needs them -- //
import ( 
	
	"strconv" 
	"strings"
	"os"
	"log"
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

	// -- initilizing the variable -- // 
	var check_if_healthy int

	// -- for loop to check if the server is in slice "healthy_servers" than returning a value based of the status -- //
	for _, i := range healthy_servers {
		if i == server {
			check_if_healthy += 1
		}
	}

	// -- checks if the function did its task successfully or not than returnin value depending on the status -- //
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

	// -- tries to check connection while some specified time (timeout) with TCP -- //
	conn, err := net.DialTimeout("tcp", "localhost:"+server, timeout)
	if err != nil {
		return 0
	}

	// -- making sure we close connection when the function ends -- //
	defer conn.Close()

	// -- returns 1 as a sign of success 
	return 1
}

////////////////////////////////
// -- FUNCION CONNECT USER -- //
////////////////////////////////
func connect_user(user int, port string) {

	// -- CONNECTING A USER TO A SERVER (IN THIS CASE A PORT) USING TCP -- //
	conn, err := net.Dial("tcp", "localhost:"+port)
	if err != nil {
		// -- debug print wen the user doesn't connect sucesfully -- // 
		fmt.Printf("[+] An error happened when connecting user %d to port %s\n", user, port)
		return
	}

	// -- making sure that we close the connection when the function ends -- //
	defer conn.Close()

	// -- debug print when user sucessfully connects -- //
	fmt.Printf("[+] User %d connected successfully to port %s\n", user, port)
}


func main() {
	
	// -- reading temp file using os.ReadFile -- //
	servers_count_temp, err := os.ReadFile("servers_count")
	if err != nil {
		log.Fatal(err)
	}

	// -- CONVERTING THE os.ReadFile INTO INT -- //
	servers_count_string := strings.TrimSpace(string(servers_count_temp))
	servers_count, err := strconv.Atoi(servers_count_string)
	if err != nil {
		log.Fatal(err)
	}

	// -- INITILIZING 2D SLICE -- //
	var servers [][]int
	for i := 0; i < servers_count; i++ {
		row := []int{}
		servers = append(servers, row)	
	}
	
	// -- initialization of variables -- //
	healthy_servers := []string{ }
	var healthy_servers_int int
	var server_selected int
	var wg sync.WaitGroup
	// -- setting up flags while keeping this structure ("name", default_value, "help text") -- //
	users_to_split := flag.Int("users", 0, "int")
	flag.Parse()

	// -- CHECK SERVERS FUNCTION -- //
	// -- initialization of variables -- //
	server_selected = 0 
	for server_selected := 0; server_selected < servers_count; server_selected ++ {
		
		// -- converting int to ASCII (Itoa) && checking health && appending -- //
		server_selected_int := server_selected + 8000
		if check_server(strconv.Itoa(server_selected_int), 500 * time.Millisecond) == 0 {
			fmt.Printf("[+] Server %s is down\n", strconv.Itoa(server_selected_int))
		} else if check_server(strconv.Itoa(server_selected_int), 500 * time.Millisecond) == 1 {
			healthy_servers = append(healthy_servers, strconv.Itoa(server_selected_int)) 
			healthy_servers_int += 1
			fmt.Printf("[+] Server %s is up\n", strconv.Itoa(server_selected_int))
		}
	}
	if healthy_servers_int <= 0 { return }
	// ^^ CHECKS SERVERS FUNCTION ^^ //



	//////////////////////////////////////////
	// -- SELECTING SERVER FOR EACH USER -- // 
	// -- initialization of variables -- //
	server_selected = 0 
	for i := 0; i < *users_to_split; i++ {
	
		// -- converting int to ASCII (Itoa) && checking && appending if yes -- //
		server_selected_int := server_selected + 8000
		if is_healthy(healthy_servers, strconv.Itoa(server_selected_int)) == 1 {
			servers[server_selected] = append(servers[server_selected], i)
		} else {
			i -= 1
		}

		server_selected += 1
		
		if server_selected >= servers_count {
			server_selected = 0
		}
	}
	for i := 0; i < servers_count; i++ {
		fmt.Println(servers[i])
	}

	///////////////////////////////////////
	// -- CONNECTING SERVERS WITH TCP -- //
	// -- At the end of the program wait for goroutines to end -- //
	defer wg.Wait()
	
	// -- nested FOR loops for connection to the servers -- //
	for server_selected := 0; server_selected < servers_count; server_selected ++ {
		for _, user := range servers[server_selected] {
				// -- telling the wg that one goroutine is about to start -- //
				wg.Add(1)
				go func (user int) {
					// -- making sure we set this goroutine as finished when it ends its task -- //
					defer wg.Done()
					// -- converting the int to ASCII (Itoa) and connecting with function connect_user -- // 
					server_selected_int := server_selected + 8000
					connect_user(user, strconv.Itoa(server_selected_int))
				}(user)
		}
	}
	// -- ALL USERS "CONNECTED" -- //
	
}

