import json, ujson, msgpack

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

def dispatch_native_object(name, message, timestamp):
    message = "modified message"
    return name, message, str(timestamp)

def dispatch_msgpack_object(serialized_object):
    # print("serialized_object:", len(serialized_object))
    # print(serialized_object)
    object = msgpack.unpackb(serialized_object, use_list=False)
    # object[b'message'] = b'modified message'
    # object[b'ts'] = 100
    new_object = msgpack.packb(object)
    # print("new_object:", len(new_object))
    # print(new_object)
    return new_object
