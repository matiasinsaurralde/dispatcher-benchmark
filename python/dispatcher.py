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
    print(serialized_object, "(Python)")
    object = msgpack.unpackb(serialized_object, use_list=True, encoding='utf-8')
    # object[b'name'] = b'theOtherObject'
    # object[b'message'] = b'someothermessage'
    # object[b'ts'] = 40
    object['nested_object']['string_field'] = "new string value"
    new_object = msgpack.packb(object, encoding='utf-8')
    print(new_object, "(Python, modified object)")
    return new_object
