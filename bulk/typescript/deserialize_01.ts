// CWE-502: Unsafe YAML deserialization
import * as yaml from 'js-yaml';
function loadConfig(yamlStr: string): unknown {
  return yaml.load(yamlStr); // DEFAULT_SCHEMA allows JS objects
}
