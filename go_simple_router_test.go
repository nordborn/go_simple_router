package go_simple_router

import (
	"fmt"
	"log"
	"net/http"
	"testing"

)

// Embed BasicHandler to implement stub methods of of CustomHttpHandler interface
type Index struct{ *BasicHandlers }

func (*Index) Get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is Index Get")
}

func (*Index) Post(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is Index Post")
}

type User struct{ *BasicHandlers }

func (*User) Get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is User Get")
}

func (u *User) Post(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Call User Get from Post")
	u.Get(w, r)
}

func main() {
	http.HandleFunc("/", R{&Index{}}.RouteByMethod)
	http.HandleFunc("/user/", R{&User{}}.RouteByMethod)

	log.Println("Listen port 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func Test(t *testing.T) {
	main()
}