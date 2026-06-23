// CWE-94: Dynamic code execution via Roslyn CSharpScript
using Microsoft.CodeAnalysis.CSharp.Scripting;
public class ELInjection01 {
    public async Task<object> Eval(string code) {
        return await CSharpScript.EvaluateAsync(code); // arbitrary code execution
    }
}
