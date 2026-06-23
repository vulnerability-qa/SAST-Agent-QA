// CWE-94: EL Injection via Spring Expression Language
import org.springframework.expression.*;
import org.springframework.expression.spel.standard.*;
public class ELInjection01 {
    ExpressionParser parser = new SpelExpressionParser();
    public Object eval(String expr) {
        return parser.parseExpression(expr).getValue(); // arbitrary SpEL execution
    }
}
