package main

import (
	// importing this packages for formattingg, logging and handling http request and responses
	"fmt"
	"log"
	"net/http"
)


// This function handles HTTP POST requests submitted from a form. It parses the form data and prints the values of "name" and "address" fields.
func formhandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}
// This function handles HTTP GET requests to the "/hello" path. It simply prints "hello" as a response.
func hellohandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello"{
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET"{
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello")

}

// This is the entry point of the program. It sets up HTTP handlers for serving static files, handling form submissions, and responding to the "/hello" endpoint.
func main(){
	fileServer := http.FileServer(http.Dir("./static")) // fileserver = creates a file server  handler that serves files from the "./static" directory.
	http.Handle("/", fileServer) // this will give us index.html file
	http.HandleFunc("/form", formhandler) //this will give us forms.html file
	http.HandleFunc("/hello", hellohandler) // this will print out just hello on the screen

	fmt.Printf("starting server at port 8080 \n")
	//http.ListenAndServe(":8080", nil): Starts an HTTP server on port 8080. It listens for incoming requests and dispatches them to the appropriate handler based on the URL path.
	if error := http.ListenAndServe(":8080", nil); error != nil { 
		log.Fatal(error) //Errors are handled using log.Fatal if there is any issue starting the server.
	}

}