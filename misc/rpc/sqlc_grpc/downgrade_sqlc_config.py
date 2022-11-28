#!/usr/bin/env python

import yaml
import os
import sys


def validate():
  if len(sys.argv) != 3:
    print("downgrade_sqlc_config.py new_yaml old_yaml , your argv[{}]".format(sys.argv))
    exit(1)

  new_yaml_path = sys.argv[1]
  if not os.path.isfile(new_yaml_path):
    print("missing {}".format(new_yaml_path))
    exit(1)

  old_yaml_path = sys.argv[2]
  if os.path.exists(old_yaml_path) and not os.path.isfile(old_yaml_path):
    print("invalid path: {}".format(old_yaml_path))
    exit(1)

def gen_old_config():
  return {
    "version": "1",
    "packages": [],
  }

def downgrade(v2config):
  v1config = v2config.copy()
  if "gen" in v1config:
    gen_field = v1config.get("gen", {})
    go_field = gen_field.get("go")
    if go_field is None:
      return None
    v1config["name"] = go_field.get("package")
    v1config["path"] = go_field.get("out")

    del v1config["gen"]

  if v1config.get("name") is None:
    return None

  if v1config.get("path") is None:
    return None

  return v1config

if __name__ == "__main__":
  validate()
  new_yaml_path = sys.argv[1]
  old_yaml_path = sys.argv[2]
  with open(new_yaml_path, "r") as new_yaml_file, open(old_yaml_path, "w") as old_yaml_file:
    old_config = gen_old_config()
    new_yaml = yaml.safe_load(new_yaml_file)
    sql_fields = new_yaml.get("sql")
    for sql_field in sql_fields:
      old_sql_field = downgrade(sql_field)
      if old_sql_field is not None:
        old_config["packages"].append(old_sql_field)
    
    print(old_config)
    old_yaml_str = yaml.dump(old_config, allow_unicode=True)
    old_yaml_file.write(old_yaml_str)



