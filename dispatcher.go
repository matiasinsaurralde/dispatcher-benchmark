package dispatcher

import(
  "github.com/matiasinsaurralde/dispatcher-benchmark/python"

  "encoding/json"
  "unsafe"
  // "fmt"
)

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "dispatcher.h"

void* createNativeObject(char* name, char* message, int timestamp) {
  int len = sizeof(NativeObject) + strlen(name) + strlen(message) + sizeof(int) + 1;
  NativeObject* obj = (NativeObject*) malloc( len );

  obj->name = name;
  obj->message = message;
  obj->timestamp = timestamp;

  return obj;
};
*/
import "C"

type NativeObject C.struct_NativeObject

const(
	_ = iota
  JsonStringMode
  UJsonStringMode
	NativeMode
	MsgPackMode
	ProtobufMode
)

type Dispatcher struct {
  Mode int
}

/*go:generate msgp
type NestedObject struct {
  NestedStringField string `msg:"string_field"`
  NestedIntField int `msg:"int_field"`
}*/

//go:generate msgp
type Object struct {
  Name string `msg:"name"`
  Message string `msg:"message"`
  Timestamp int `msg:"ts"`
  // NestedObject NestedObject `msg:"nested_object"`
}

func NewDispatcher(mode int) Dispatcher {
  d := Dispatcher{
    Mode: mode,
  }
  // fmt.Println("NewDispatcher")
  return d
}

func PythonInit() error {
  return python.Init()
}

func PythonLoadDispatcher() error {
  return python.LoadDispatcher()
}

func DispatchJsonString(object []byte) string {
  return python.DispatchJsonString(object)
}

func DispatchUjsonString(object []byte) string {
  return python.DispatchJsonString(object)
}

func PythonSetPath(path string) {
  python.SetPath(path)
}

func (d *Dispatcher) Dispatch(o *Object) (interface{}, error) {
  var data []byte
  var err error
  var output interface{}

  switch d.Mode {
  case JsonStringMode:
    data, err = json.Marshal(o)
    // log.Print(string(data), len(data))
    output = DispatchJsonString(data)
  case UJsonStringMode:
    data, err = json.Marshal(o)
    output = python.DispatchUJsonString(data)
  case NativeMode:
    /*
    var name *C.char
    name = C.CString(o.Name)
    var message *C.char
    message = C.CString(o.Message)
    var timestamp C.int
    timestamp = C.int(o.Timestamp)

    var objectPtr unsafe.Pointer
    objectPtr = C.createNativeObject( name, message, timestamp )
    var newObject python.Object = python.DispatchNativeObject(objectPtr)
    o.Name = newObject.Name
    o.Message = newObject.Message
    o.Timestamp = newObject.Timestamp

    output = o*/
  case MsgPackMode:
    data, err = o.MarshalMsg(nil)
    // fmt.Println(string(data), "(Go)")
    dataStr := string(data)
    var CData *C.char
    CData = C.CString(dataStr)

    // *** EXPERIMENTAL ***
    // var someObject Object
    // someObject.UnmarshalMsg(data)
    // fmt.Println("someObject = ", someObject)
    //

    var outputBytes []byte
    outputBytes = python.DispatchMsgPackObject(unsafe.Pointer(CData))

    var newObject Object
    newObject.UnmarshalMsg(outputBytes)
    output = newObject

    // fmt.Println("o2 = ", newObject)

  }

  return output, err
}
