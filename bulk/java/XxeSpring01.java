// CWE-611: XXE via Spring's RestTemplate XML response
import org.springframework.web.client.RestTemplate;
public class XxeSpring01 {
    public String fetchXml(String url) {
        RestTemplate rt = new RestTemplate();
        return rt.getForObject(url, String.class); // XML body parsed without disabling entities
    }
}
