import json

from scriptlab import Scriptlab

sl = Scriptlab()

response = {
    "response": True,
    "data": "data"
}

sl.response(json.dumps(response))