// CWE-798: Hardcoded AWS credentials
public class HardcodedCreds02 {
    private static final String AWS_KEY = "AKIAIOSFODNN7EXAMPLE";
    private static final String AWS_SECRET = "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY";
    public String getKey() { return AWS_KEY; }
}
