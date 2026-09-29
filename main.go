package main

import( 

	"fmt"
	"net/http"

)

func h8080(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}
func h8081(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}
func h8082(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func main() {

	serverMuxA := http.NewServeMux()
	serverMuxA.HandleFunc("/", h8080)


	serverMuxB := http.NewServeMux()
	serverMuxB.HandleFunc("/", h8081)


	serverMuxC := http.NewServeMux()
	serverMuxC.HandleFunc("/", h8082)


	go func() {
		fmt.Println("8080")
		http.ListenAndServe(":8080", serverMuxA)
	}()
	go func() {
		fmt.Println("8081")
		http.ListenAndServe(":8081", serverMuxB)
	}()
	fmt.Println("8082")
	http.ListenAndServe(":8082", serverMuxC)

}
