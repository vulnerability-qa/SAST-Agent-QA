// CWE-190: Integer Overflow in allocation size
public class IntegerOverflow01 {
    public byte[] allocate(int count, int itemSize) {
        int total = count * itemSize; // overflows if count * itemSize > Integer.MAX_VALUE
        return new byte[total];
    }
}
