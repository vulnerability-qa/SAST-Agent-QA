// CWE-89: NoSQL Injection in MongoDB
const { MongoClient } = require('mongodb');
async function findUser(username, password) {
  const client = new MongoClient('mongodb://localhost');
  const db = client.db('app');
  // Attacker passes {$gt: ''} as password to bypass auth
  return db.collection('users').findOne({ username, password });
}
