package main

import(
  "github.com/matiasinsaurralde/dispatcher-benchmark"

  // "time"
  "fmt"
)

func main() {

  fmt.Println("main")

  var err error

  dispatcher.PythonSetPath("/Users/matias/dev/dispatcher/python")

  err = dispatcher.PythonInit()
  err = dispatcher.PythonLoadDispatcher()

  if err != nil {
    panic(err)
  }

  shinyDispatcher := dispatcher.NewDispatcher(dispatcher.MsgPackMode)

  nestedObject := dispatcher.NestedObject{
    NestedStringField: "nested",
    NestedIntField: 22,
  }

  object := dispatcher.Object{
    Name: "theObject",
    Message: "the message",
    Timestamp: 20,
    NestedObject: nestedObject,
  }

  fmt.Println(object, "(Go, original object)")

  fmt.Print( shinyDispatcher.Dispatch(&object) )
  fmt.Println(" (Go, deserialized object)")
  /*
  var output interface{}
  output, err = shinyDispatcher.Dispatch(&object)
  var outputString string
  outputString = output.(string)

  fmt.Println("output = ", outputString)
  */
}
