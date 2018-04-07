package main

import (
  "fmt"
  "http"
)

func main()  {
  fmt.Println("Starting web server on port 8080")
  http.HandleFunc("/", )
  http.ListenAndServe(":8080", )
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
  fmt.Println("User requested root")

}
