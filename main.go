// -- package main since this is really the main file -- //
package main 

import(

	"time"
	"strconv"
	"fmt"
	"log"
	"flag"
	"os/exec"

)

// -- starting off with a main function since there is no need for other functions -- //
func main() {

	// -- DEFINING && PREPARING FLAGS -- //
	var servers_int int
	var users_int int

	flag.IntVar(&servers_int, "servers", 0, "servers int")
	flag.IntVar(&users_int, "users", 0, "users int")
	flag.Parse()
	
	// -- converting Itoa (Int to ASCII) the flags -- //
	servers := strconv.Itoa(servers_int)
	users := strconv.Itoa(users_int)

	// -- STARTING SERVERS -- //
	server_startup := exec.Command("go", "run", "servers/main.go", "-servers", servers)
	
	// -- starting the servers -- //
	err := server_startup.Start()
	if err != nil {
		log.Fatal(err)
	}
	
	// -- geting servers startup process PID -- //
	fmt.Printf("[+] PID >> %d\n", server_startup.Process.Pid)

	// --  BALANCING USERS -- //
	// -- waiting for the servers to setup correctly -- //
	time.Sleep(1 * time.Second)

	balance_users := exec.Command("go", "run", "balance/main.go", "-users", users)

	// -- starting the balance prgoram with .Output() for catching output -- //
	output, err := balance_users.Output()
	if err != nil {
		log.Fatal(err)
	}

	// -- getting the balance process PID -- //
	fmt.Printf("[+] PID >> %d\n", balance_users.Process.Pid)

	// -- converting output into string and printing it -- //
	output_string := string(output)
	fmt.Println(output_string)

	// -- setting up and pkilling the servers process at the end -- //
	server_startup_pid := strconv.Itoa(server_startup.Process.Pid)
	defer exec.Command("pkill", server_startup_pid).Start()

}
