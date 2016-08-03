package dispatcher

import(
  "github.com/matiasinsaurralde/dispatcher-benchmark/python"

  "encoding/json"
  // "log"
)

const(
	_ = iota
  JsonMode
	NativeMode
	MsgPackMode
	ProtobufMode
)


type Dispatcher struct {
  Mode int
}

//go:generate msgp
type Object struct {
  Name string `msg:"name"`
  Message string `msg:"message"`
  Timestamp int64 `msg:"ts"`
}

func NewDispatcher(mode int) Dispatcher {
  d := Dispatcher{
    Mode: mode,
  }
  // log.Println("New dispatcher, using mode:", d.Mode)
  return d
}

func PythonInit() error {
  return python.Init()
}

func PythonLoadDispatcher() error {
  return python.LoadDispatcher()
}

func PythonDispatchString(object []byte) string {
  return python.DispatchString(object)
}

func PythonSetPath(path string) {
  python.SetPath(path)
}

func (d *Dispatcher) Dispatch(o *Object) (interface{}, error) {
  var data []byte
  var err error
  switch d.Mode {
  case JsonMode:
    data, err = json.Marshal(o)
    // log.Print(string(data), len(data))
  case MsgPackMode:
    data, err = o.MarshalMsg(nil)
    // log.Print(string(data), len(data))
  }

  var output interface{}
  output = python.DispatchString(data)

  return output, err
}
