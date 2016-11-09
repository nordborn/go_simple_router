// USAGE
// import (
//	sr "go_simple_router"
//	"log"
// 	"fmt"
// )
//
// Embed BasicHandler to implement all stub methods of CustomHttpHandler interface
// type Index struct{ *sr.BasicHandler }
//
// func (*Index) Get(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintln(w, "This is Index Get")
//}
//
// func (*Index) Post(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintln(w, "This is Index Post")
// }
// ...
// func main() {
// 	http.HandleFunc("/", sr.R{&Index{}}.RouteByMethod)
//	log.Println("Listen port 8000")
//	log.Fatal(http.ListenAndServe(":8000", nil))
// }

package go_simple_router

import (
	"net/http"
	"log"
)

// Basic routing struct with
// RouteByMethod() - main method to handle request methods and route them
// Must be initialized with specified handler struct with Get(), Post(), Put(), Patch(), Delete()
type R struct {
	Handler CustomHttpHandler // implementation of specific handler
}

func (rtr R) RouteByMethod(w http.ResponseWriter, r *http.Request) {
	log.Println("Handler Handle Func")

	if r.Method == "POST" {
		rtr.Handler.Post(w, r)
		return
	}

	if r.Method == "PUT" {
		rtr.Handler.Post(w, r)
		return
	}

	if r.Method == "PATCH" {
		rtr.Handler.Patch(w, r)
		return
	}

	if r.Method == "Delete" {
		rtr.Handler.Delete(w, r)
		return
	}

	rtr.Handler.Get(w, r)
}

// Interface with list of necessary methods
type CustomHttpHandler interface {
	Get (w http.ResponseWriter, r *http.Request)
	Post (w http.ResponseWriter, r *http.Request)
	Put (w http.ResponseWriter, r *http.Request)
	Patch (w http.ResponseWriter, r *http.Request)
	Delete (w http.ResponseWriter, r *http.Request)
}

// Struct to implement stub methods
// Embed it when you want to implement CustomHttpHandler interface
type BasicHandlers struct{}

func (*BasicHandlers) Get(w http.ResponseWriter, r *http.Request) {
	log.Println("Get Stub")
}

func (*BasicHandlers) Post(w http.ResponseWriter, r *http.Request) {
	log.Println("Post Stub")
}

func (*BasicHandlers) Put(w http.ResponseWriter, r *http.Request) {
	log.Println("Put Stub")
}

func (*BasicHandlers) Patch(w http.ResponseWriter, r *http.Request) {
	log.Println("Patch Stub")
}

func (*BasicHandlers) Delete (w http.ResponseWriter, r *http.Request) {
	log.Println("Delete Stub")
}

