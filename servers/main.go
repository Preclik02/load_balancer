package main

import( 

	"fmt"
	"strconv"
	"flag"
	"log"
	"sync"
	"net"
	"os"

)

func make_server_listen(server string, wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()
		listener, err := net.Listen("tcp", ":"+server)
		if err != nil {
			log.Printf("[+] err - %s\n", err)	
			return
		}
		defer listener.Close()
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Printf("[+] err - %s\n", err)
				continue
			}
			go func(c net.Conn) {
				defer c.Close()
				c.Write([]byte("Hello "+server))
			}(conn)
		}
	}()
}
func main() {

	// -- initialization of variables -- //
	var wg sync.WaitGroup
	
	// -- making sure we wait for all goroutines to finish before closing the program -- //
	defer wg.Wait()

	// -- setting up the flag using format of ("name", default_value, "help value") -- //
	servers_to_run := flag.Int("servers", 0, "int")
	flag.Parse()
	
	// -- initilizing the variable to print into the file and writing into the file -- //
	data_to_file := []byte(fmt.Sprintf("%d", *servers_to_run))
	err := os.WriteFile("servers_count", data_to_file, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// -- checks the limit of servers we can start -- //
	if *servers_to_run > 999 {
		return
	}

	/////////////////////////////////////
	// -- RUN SERVERS IN A FOR LOOP -- //
	// -- THIS MAKES THEM "DYNAMIC" -- //
	for server_selected := 0; server_selected < *servers_to_run; server_selected++ {
		wg.Add(1)
		server_selected_int := 8000 + server_selected
		make_server_listen(strconv.Itoa(server_selected_int), &wg)
	}

}
