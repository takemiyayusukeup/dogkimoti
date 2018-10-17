package mlcall

import (
	"fmt"
	"net/http"
	"google.golang.org/appengine"
)

type Values struct {
	Atai1 int
	Atai2 int
	ans   int // カプセル化されている
}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "test!")
}
