package dispatcher

import(
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

func (d *Dispatcher) Dispatch(o *Object) ([]byte, error) {
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
  return data, err
}
