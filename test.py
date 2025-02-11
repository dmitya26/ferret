"""
The following outlines very basically a prototype of a spec compliant json parser and marshaller/demarshaller. It works to essentially divide the file into two kinds of kv pairs, those that deepen the json structure, and those that don't, it does this while keeping track of which non-deepening kv pairs belong to which deepening kv pairs; say for example: 'hello':'hello' belongs to 'heyo'.
"""
import re 
import functools 

test = "'hello':{{}, 'heyuh':'heyo'}"

splitted = re.split(r":\{|,\s*", test)
#print("\n".join(splitted))


def append_marshal(d: dict, deep_keys, value) -> dict:
    functools.reduce(lambda d, key: d.setdefault(key, {}), deep_keys[:-1], d)[deep_keys[-1]] = value

depth_queue=[]
json={}

for c in splitted:
    if ":" not in c:
        depth_queue.append(c.replace("{", ""))
        #print("Inside ", " -> ".join(depth_queue))
    else:
        tmp = re.sub(r"}|'|\"", "", c).split(":")

        keys = depth_queue
        keys.append(tmp[0])

        append_marshal(json, keys, tmp[1])
    for j, d in enumerate(c):
        if d == "}":
            depth_queue.pop(-1)

for key in json:
    print(key, "\n")
    print(json[key]["heyuh"], "\n")
