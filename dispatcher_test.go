package dispatcher

import(
  "testing"
  "time"
  "fmt"

  "encoding/json"
)

const iterations int = 50000

func init() {
  fmt.Println("Benchmark setup")

  PythonSetPath("/Users/matias/dev/dispatcher/python")

  PythonInit()
  PythonLoadDispatcher()
}

func BenchmarkJsonStringMode(b *testing.B) {
  d := NewDispatcher(JsonStringMode)

  for i := 0; i < iterations; i++ {
    object := Object{
      Name: "theObject",
      Message: "the message",
      Timestamp: time.Now().Unix(),
    }
    // time.Sleep(1*time.Millisecond)
    var output interface{}
    output, _ = d.Dispatch(&object)
    outputString := output.(string)

    var deserializedObject Object
    json.Unmarshal( []byte(outputString), &deserializedObject)
  }

}

func BenchmarkUjsonStringMode(b *testing.B) {
  d := NewDispatcher(UJsonStringMode)

  for i := 0; i < iterations; i++ {
    object := Object{
      Name: "theObject",
      Message: "the message",
      Timestamp: time.Now().Unix(),
    }
    // time.Sleep(1*time.Millisecond)
    var output interface{}
    output, _ = d.Dispatch(&object)
    outputString := output.(string)

    var deserializedObject Object
    json.Unmarshal( []byte(outputString), &deserializedObject)
  }

}

func BenchmarkNativeMode(b *testing.B) {
  d := NewDispatcher(NativeMode)

  for i := 0; i < iterations; i++ {
    object := Object{
      Name: "theObject",
      Message: "the message",
      Timestamp: time.Now().Unix(),
    }
    // time.Sleep(1*time.Millisecond)
    d.Dispatch(&object)
  }

}
/*
func BenchmarkMsgpackMode(b *testing.B) {
  d := NewDispatcher(MsgPackMode)
  for i := 0; i < 200; i++ {
    object := Object{
      Name: "theObject",
      Message: "the message",
      Timestamp: time.Now().Unix(),
    }
    time.Sleep(1*time.Millisecond)
    d.Dispatch(&object)
  }
}
*/
