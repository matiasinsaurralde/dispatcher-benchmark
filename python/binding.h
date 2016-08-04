#include <Python.h>

#ifndef DISPATCHER
#define DISPATCHER

static int Python_Init();
static int Python_LoadDispatcher();
static void Python_SetEnv(char*);

static char* Python_DispatchJsonString(char*);
static char* Python_DispatchUJsonString(char*);

static char* dispatcher_module_name = "dispatcher";
static char* hook_name = "process";

static char* dispatch_json_string = "dispatch_json_string";
static PyObject* dispatch_json_string_hook;

static char* dispatch_ujson_string = "dispatch_ujson_string";
static PyObject* dispatch_ujson_string_hook;

static char* dispatch_native_object = "dispatch_native_object";
static PyObject* dispatch_native_object_hook;

static PyObject* dispatcher;
static PyObject* dispatcher_module;
static PyObject* dispatcher_dict;

static PyObject* dispatcher_dispatch;

#endif
