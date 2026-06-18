import java.util.Scanner;

public class CommandInjection {
    public static void main(String[] args) throws Exception {
        Scanner sc = new Scanner(System.in);
        System.out.print("Enter command: ");
        String cmd = sc.nextLine();

        // CWE-78: user input passed to Runtime.exec via shell without validation.
        Runtime.getRuntime().exec(new String[]{"/bin/sh", "-c", cmd});
    }
}
