package dispatcher

import(
  "testing"
  "time"
  "fmt"
)

func init() {
  fmt.Println("Benchmark setup")

  PythonSetPath("/Users/matias/dev/dispatcher/python")

  PythonInit()
  PythonLoadDispatcher()
}

func BenchmarkJsonMode(b *testing.B) {
  d := NewDispatcher(JsonMode)

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
