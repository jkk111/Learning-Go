package main

import(
  "fmt"
  "strconv"
  "crypto/rand"
  "encoding/base64"
  "os"
)

func main() {
  if len(os.Args) < 2 {
    fmt.Println("Must specify a length!")
    return
  }
  parsed, err := strconv.Atoi(os.Args[1])
  if err != nil {
    fmt.Println("Parameter must be a number!")
    return
  } else if parsed != 0 {
    fmt.Println("Generating random string of length:", parsed)
    rand := generateRandomBytes(parsed)
    fmt.Println("Random string: ", encode(rand))
  }
}

func encode(bytes []byte) string {
  // base64.URLEncoding taken from:
  // https://elithrar.github.io/article/generating-secure-random-numbers-crypto-rand/
  return base64.URLEncoding.EncodeToString(bytes)
}

func generateRandomBytes(length int) []byte {
  str := make([]byte, length)
  _, err := rand.Read(str)
  if err != nil {
    fmt.Println("Error occured while generating random string")
    tmp := make([]byte, 0)
    return tmp
  }
  return str
}
