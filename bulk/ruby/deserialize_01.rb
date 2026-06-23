# CWE-502: Insecure YAML deserialization
require 'yaml'
def load_config(data)
  YAML.load(data) # YAML.load executes Ruby objects; use YAML.safe_load
end
