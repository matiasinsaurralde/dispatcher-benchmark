import json

def process(serialized_object):
    object = json.loads(serialized_object)
    object['Message'] = "modified message"
    serialized_object = json.dumps(object)
    return serialized_object
