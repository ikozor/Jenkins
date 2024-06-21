package main 

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	res := fmt.Sprintf(`
	<h1>Do some Calc</h1>
	<form hx-post="/add" hx-target="#result" hx-swap="afterbegin">
		<label for="num1">Number 1:</label>
		<input type="number" name="n1">
		<br>
		<label for="num2">Number 2:</label>
		<input type="number" name="n2">
		<br>
		<button type="submit">Add</button>
	</form>
	<div id="result"></div>
	<script src="https://unpkg.com/htmx.org@2.0.0"></script>
	`)

	io.WriteString(w, res)
}

func getAdd(w http.ResponseWriter, r *http.Request) {
	n1, err := strconv.Atoi(r.FormValue("n1"))
	if err != nil {
		io.WriteString(w, "<div>ERROR</div>")
		return
	}

	n2, err := strconv.Atoi(r.FormValue("n2"))
	if err != nil {
		io.WriteString(w, "<div>ERROR</div>")
		return
	}
	res := fmt.Sprintf("<div>%d + %d = %d</div>", n1, n2, add(n1, n2))
	io.WriteString(w, res)
}

func add(x, y int) int {
	return x + y
}

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/add", getAdd)
	log.Print(http.ListenAndServe(":3000", nil))
}
