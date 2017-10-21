package main

import (
    "log"
    "net/http"
)

func main() {
	defer db.Close()

    router := NewRouter()

    log.Fatal(http.ListenAndServe(":8080", router))
}
