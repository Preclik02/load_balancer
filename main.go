package main

import( 

	"log"
	"sync"
	"net"

)


func main() {

	var wg sync.WaitGroup
	
	defer wg.Wait()


	// -- "SERVER" 8080 WITH TCP LISTENING -- //
	wg.Add(1)
	go func() {
		defer wg.Done()
		listener, err := net.Listen("tcp", ":8080")
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
				c.Write([]byte("Hello 8080"))
			}(conn)
		}
	}()

	// -- "SERVER" 8081 WITH TCP LISTENING -- //
	wg.Add(1)
	go func() {
		defer wg.Done()
		listener, err := net.Listen("tcp", ":8081")
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
				c.Write([]byte("Hello 8081"))
			}(conn)

		}
	}()


	// -- "SERVER" 8082 WITH TCP LISTENING -- //
	wg.Add(1)
	go func() {
		defer wg.Done()
		listener, err := net.Listen("tcp", ":8082")
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
				c.Write([]byte("Hello 8082"))
			}(conn)

		}
	}()
}
