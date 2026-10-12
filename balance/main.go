package main
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
	
	healthy_servers := []string{ }
	var healthy_servers_int int

	users_to_split := flag.Int("users", 0, "int")
	

	flag.Parse()
	

	// -- CHECK SERVERS FUNCTION -- //
	server_selected := 0 
	for server_selected := 0; server_selected < servers_count; server_selected ++ {
		
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
	server_selected = 0 
	for i := 0; i < *users_to_split; i++ {
	
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













	
	// -- At the end of the program wait for goroutines to end -- //
	defer wg.Wait()
	
	// -- FOR loops for connection to the server -- //
	// -- 1ST FOR LOOP -- //
	for server_selected := 0; server_selected < servers_count; server_selected ++ {
		for _, user := range servers[server_selected] {
				wg.Add(1)
				go func (user int) {
					defer wg.Done()
					server_selected_int := server_selected + 8000
					connect_user(user, strconv.Itoa(server_selected_int))
				}(user)
		}
	}
	// -- ALL USERS "CONNECTED" -- //
	
}

