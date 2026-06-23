// CWE-78: Command Injection via child_process in TypeScript
import { execSync } from 'child_process';
export function resizeImage(filename: string, size: string): string {
  return execSync(`convert ${filename} -resize ${size} output.png`).toString();
}
