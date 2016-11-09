# go_simple_router

Extremely basic routing for Go (Golang) to handle requets methods
 - GET
 - POST
 - PUT
 - PATCH
 - DELETE

in http.HandleFunc

No overhead and custom routing implementations

**Just use empty structs with such methods**

### Example
```
package main

import (
	"fmt"
	"log"
	"net/http"
	sr "github.com/nordorn/go_simple_router"
)

// Embed BasicHandlers to implement stub methods
type Index struct{ *sr.BasicHandlers }

func (*Index) Get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is Index Get")
}

func (*Index) Post(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is Index Post")
}

type User struct{ *sr.BasicHandlers }

func (*User) Get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is User Get")
}

func (u *User) Post(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Call User Get from Post")
	u.Get(w, r)
}

func main() {
	http.HandleFunc("/", sr.R{&Index{}}.RouteByMethod)
	http.HandleFunc("/user/", sr.R{&User{}}.RouteByMethod)

	log.Println("Listen port 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
```
