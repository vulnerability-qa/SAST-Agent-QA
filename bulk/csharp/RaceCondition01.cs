// CWE-362: Race Condition on shared balance
public class RaceCondition01 {
    private int _balance = 1000;
    public void Withdraw(int amount) {
        if (_balance >= amount) { // check
            Thread.Sleep(1); // simulate work
            _balance -= amount; // deduct — not synchronized
        }
    }
}
