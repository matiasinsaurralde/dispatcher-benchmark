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
    # print(serialized_object, "(Python)")
    object = msgpack.unpackb(serialized_object, use_list=False, encoding='utf-8')
    # object[b'name'] = b'theOtherObject'
    # object[b'message'] = b'someothermessage'
    # object[b'ts'] = 40
    # print("object", object)
    object['message'] = 'long stringlong stringlong stringlong stringlong stringlong stringlong stringlong string'
    # object['nested_object']['string_field'] = "long stringlong stringlong stringlong stringlong stringlong string"
    new_object = msgpack.packb(object, use_bin_type=True)
    # print(new_object, "(Python, modified object)")
    return new_object
