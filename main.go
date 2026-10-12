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

	var wg sync.WaitGroup
	
	defer wg.Wait()

	servers_to_run := flag.Int("servers", 0, "int")
	
	flag.Parse()
	
	data_to_file := []byte(fmt.Sprintf("%d", *servers_to_run))
	err := os.WriteFile("balance/servers_count", data_to_file, 0644)
	if err != nil {
		log.Fatal(err)
	}

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
