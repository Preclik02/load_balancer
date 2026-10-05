package main

import( 

	"log"
	"sync"
	"net"

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


	// -- "SERVER" 8080 WITH TCP LISTENING -- //
	wg.Add(1)
	make_server_listen("8080", &wg)

	// -- "SERVER" 8081 WITH TCP LISTENING -- //
	wg.Add(1)
	make_server_listen("8081", &wg)

	// -- "SERVER" 8082 WITH TCP LISTENING -- //
	wg.Add(1)
	make_server_listen("8082", &wg)
}
