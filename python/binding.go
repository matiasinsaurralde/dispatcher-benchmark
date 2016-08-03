package python

/*
#cgo pkg-config: python3

#include <Python.h>

#include <stdio.h>
#include <stdlib.h>

#include "binding.h"

static int Python_Init() {
  Py_Initialize();
  printf("Py_Initialize\n");
  return Py_IsInitialized();
}

static int Python_LoadDispatcher() {
  PyObject *module_name = PyUnicode_FromString( dispatcher_module_name );
  dispatcher_module = PyImport_Import( module_name );


  if( dispatcher_module == NULL ) {
  printf("dispatcher_module error\n");
    PyErr_Print();
    return -1;
  }

  dispatcher_dict = PyModule_GetDict(dispatcher_module);

  if( dispatcher_dict == NULL ) {
  printf("dispatcher_dict error\n");
    PyErr_Print();
    return -1;
  }

  dispatcher_dispatch = PyDict_GetItemString(dispatcher_dict, hook_name);

  if( dispatcher_dispatch == NULL ) {
    PyErr_Print();
    return -1;
  }

  if( !PyCallable_Check(dispatcher_dispatch) ) {
    return -1;
  }

  return 0;
}

static void Python_SetEnv(char* python_path) {
  printf("Setting PYTHONPATH to: %s\n", python_path );
  setenv("PYTHONPATH", python_path, 1 );
}

static void Python_Dispatch(char* object) {
  printf("dispatching: %s\n", object);
}
*/
import "C"

import(
  "errors"
)

func Init() error {
  var err error
  result := int(C.Python_Init())
  if result != 0 {
    err = errors.New("Couldn't initialize Python")
  }
  return err
}

func LoadDispatcher() error {
  var err error
  result := int(C.Python_LoadDispatcher())
  if result != 0 {
    err = errors.New("Couldn't load dispatcher")
  }
  return err
}

func SetPath(path string) {
  CPath := C.CString(path)
  C.Python_SetEnv((CPath))
}

func DispatchString(object []byte) {
  s := string(object)
  CObject := C.CString(s)
  C.Python_Dispatch(CObject)
}

func DispatchBytes(object []byte) {
}
