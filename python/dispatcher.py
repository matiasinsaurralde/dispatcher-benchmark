import json, ujson

def dispatch_json_string(serialized_object):
    object = json.loads(serialized_object)
    object['Message'] = "modified message"
    serialized_object = json.dumps(object)
    return serialized_object

def dispatch_ujson_string(serialized_object):
    object = ujson.loads(serialized_object)
    object['Message'] = "modified message"
    serialized_object = ujson.dumps(object)
    return serialized_object

def dispatch_native_object( name, message, timestamp ):
    message = "modified message"
    return name, message, str(timestamp)
