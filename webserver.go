package main
import(
  "net/http"
  "fmt"
)

type httpHandler func(w http.ResponseWriter, r *http.Request)

func main() {
  initWebServer();
}

func testHandler(w http.ResponseWriter, r *http.Request) {
  if(r.URL.Path != "/") {
    w.Write([]byte("404 Page Not Found"))
  } else {
    w.Write([]byte(fmt.Sprintf("Hello %s", r.Host)))
  }
}

func initWebServer() {
  http.Handle("/", http.FileServer(http.Dir("./static/")))
  http.HandleFunc("*", testHandler)
  http.ListenAndServe(":8080", nil)
  fmt.Println("Server listening")
}
