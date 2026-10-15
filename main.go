package main 

import(

	"strconv"
	"fmt"
	"log"
	"flag"
	"os/exec"

)

func main() {

	// -- DEFINING && PREPARING FLAGS -- //
	
	var servers_int int
	var users_int int

	flag.IntVar(&servers_int, "servers", 0, "servers int")
	flag.IntVar(&users_int, "users", 0, "users int")
	flag.Parse()

	servers := strconv.Itoa(servers_int)
	users := strconv.Itoa(users_int)

	// -- STARTING SERVERS -- //
	server_startup := exec.Command("go", "run", "servers/main.go", "-servers", servers)
	
	err := server_startup.Start()
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("[+] PID >> %d\n", server_startup.Process.Pid)

	// --  BALANCING USERS -- //
	balance_users := exec.Command("go", "run", "balance/main.go", "-users", users)

	output, err := balance_users.Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("[+] PID >> %d\n", balance_users.Process.Pid)

	output_string := string(output)
	fmt.Println(output_string)

	server_startup_pid := strconv.Itoa(server_startup.Process.Pid)
	defer exec.Command("pkill", server_startup_pid).Start()


}
