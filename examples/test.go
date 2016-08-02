package main

import(
  "github.com/matiasinsaurralde/dispatcher-benchmark"

  "time"
  "fmt"
)

func main() {

  fmt.Println("main")

  shinyDispatcher := dispatcher.NewDispatcher(dispatcher.JsonMode)

  object := dispatcher.Object{
    Name: "theObject",
    Message: "the message",
    Timestamp: time.Now().Unix(),
  }

  shinyDispatcher.Dispatch(&object)
}
