package main

import "net/http"

func main() {
	app := NewApplication()
	app.infoLog.Println("starting sintapi server at :3000")
	if err := http.ListenAndServe(":3000", app.routes()); err != nil {
		app.errorLog.Fatal(err)
	}
}
