using System.Web;
using System.Web.UI;

namespace SastSamples
{
    public class PossibleXSS : Page
    {
        protected void Page_Load(object sender, System.EventArgs e)
        {
            var name = Request.QueryString["name"];
            // CWE-79: user input written to response without HTML encoding.
            Response.Write("<h1>Hello, " + name + "!</h1>");
        }
    }
}
