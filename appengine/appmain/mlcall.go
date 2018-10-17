
package mlcall

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "test!")
}
