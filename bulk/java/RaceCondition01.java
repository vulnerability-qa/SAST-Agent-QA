// CWE-362: Race Condition in account balance check
public class RaceCondition01 {
    private int balance = 500;
    public void withdraw(int amount) {
        if (balance >= amount) { // check — not synchronized
            try { Thread.sleep(10); } catch (InterruptedException e) {}
            balance -= amount; // deduct — race allows overdraft
        }
    }
}
