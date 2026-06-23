// CWE-362: Race Condition in balance check/deduct
let balance = 1000;
async function withdraw(amount) {
  if (balance >= amount) {
    await processPayment(amount);
    balance -= amount; // race allows double-spend
  }
}
async function processPayment() {}
