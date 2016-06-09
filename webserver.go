package main
import(
  "net/http"
  "fmt"
)

func main() {
  http.HandleFunc("/", testHandler)
  http.ListenAndServe(":8080", nil)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("hello world"))
}

func serverReady() {
  fmt.Println("Server ready")
}
