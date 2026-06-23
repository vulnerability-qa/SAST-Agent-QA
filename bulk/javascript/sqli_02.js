// CWE-89: SQL Injection via template literal
const { Pool } = require('pg');
const pool = new Pool();
async function getOrder(orderId) {
  const res = await pool.query(`SELECT * FROM orders WHERE id = ${orderId}`);
  return res.rows;
}
