
import json
import os
import sys


class Protobuf:
  def __init__(self, service_name):
    self.service_name = service_name
    self.apis = []

  def register(self, api_name, response_parametors, request_parametors):
    self.apis.append(
      {
        "name": api_name,
        "response_parametors": response_parametors,
        "request_parametors": request_parametors,
      }
    )

  def to_s_rpc_definitions(self):
    buf = ""
    for api in self.apis:
      buf += "\n\trpc {api_name}({api_name}Request) returns ({api_name}Response) {{}};".format(
        api_name=api["name"]
      )
    return buf

  def to_message(self, name, parametors):
    buf = "message {} {{".format(name)

    count = 0
    for p in parametors:
      count += 1
      buf += "\n\t{type} {name} = {count}".format(
        type=p["type"],
        name=p["name"],
        count=count
      )

    buf += "\n}\n"
    return buf


  def __str__(self):
    header = "service {service_name} {{\n{rpc_definitions}\n}}".format(
      service_name=self.service_name,
      rpc_definitions=self.to_s_rpc_definitions(),
    )

    body = "\n"
    for api in self.apis:
      name = api["name"]
      response_parametors = api["response_parametors"]
      request_parametors = api["request_parametors"]
      body += self.to_message("{}Response".format(name), response_parametors)
      body += self.to_message("{}Request".format(name), request_parametors)

    return header + body


def validate():
  if len(sys.argv) != 5:
    print("python3 generator.py type.json protobuf service_name package_name, your argv[{}]".format(sys.argv))
    exit(1)

  type_json_path = sys.argv[1]
  if not os.path.isfile(type_json_path):
    print("missing {}".format(type_json_path))
    exit(1)

  protobuf_path = sys.argv[2]
  if os.path.exists(protobuf_path) and not os.path.isfile(protobuf_path):
    print("invalid path: {}".format(protobuf_path))
    exit(1)

def to_protobuf_type(query_type):
  # FIXME: match
  if query_type == "serial":
    return "int32"
  elif query_type == "uuid":
    return "string"
  elif query_type == "varchar":
    return "string"
  elif query_type == "text":
    return "string"
  else:
    return "string"

def to_protobuf_api(query):
  api_name = query["name"]
  response_parametors = []
  request_parametors = []

  for column in query["columns"]:
    response_parametors.append(
      {
        "name": column["name"],
        "type": to_protobuf_type(column["type"]["name"]),
      }
    )
  for param in query["params"]:  # might need to sort?
    column = param["column"]
    request_parametors.append(
      {
        "name": column["name"],
        "type": to_protobuf_type(column["type"]["name"]),
      }
    )

  return api_name, \
    response_parametors, \
    request_parametors

if __name__ == "__main__":
  validate()
  type_json_path = sys.argv[1]
  protobuf_path = sys.argv[2]
  service_name = sys.argv[3]
  package_name = sys.argv[4]

  with open(type_json_path, "r") as type_json_file, open(protobuf_path, "w") as protobuf_file:
    protobuf = Protobuf(service_name)
    type_json = json.load(type_json_file)
    queries = type_json.get("queries", [])
    for q in queries:
      protobuf.register(*to_protobuf_api(q))
      
    protobuf_str = """
syntax = "proto3";
package {};

{}
""".format(package_name, str(protobuf))

    protobuf_file.write(protobuf_str)


