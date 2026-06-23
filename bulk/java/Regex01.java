// CWE-1333: ReDoS via vulnerable regular expression
public class Regex01 {
    public boolean validate(String input) {
        return input.matches("^(\\w+\\s?)+$"); // catastrophic backtracking on long non-matching input
    }
}
