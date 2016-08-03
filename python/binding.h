#include <Python.h>

#ifndef DISPATCHER
#define DISPATCHER

static int Python_Init();
static int Python_LoadDispatcher();
static void Python_SetEnv(char*);
static void Python_Dispatch(char*);

static char* dispatcher_module_name = "dispatcher";
static char* hook_name = "process";

static PyObject* dispatcher;
static PyObject* dispatcher_module;
static PyObject* dispatcher_dict;

static PyObject* dispatcher_dispatch;

#endif
