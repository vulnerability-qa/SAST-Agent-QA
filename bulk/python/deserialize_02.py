# CWE-502: Insecure deserialization via yaml.load
import yaml
def load_config(data):
    return yaml.load(data, Loader=yaml.Loader)  # should use SafeLoader
