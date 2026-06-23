// CWE-89: SQL Injection in TypeScript/TypeORM raw query
import { getConnection } from 'typeorm';
async function findUser(username: string) {
  const conn = getConnection();
  return conn.query(`SELECT * FROM users WHERE username = '${username}'`);
}
