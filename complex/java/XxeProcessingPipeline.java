// CWE-611: XXE via Chained XML Processing Pipeline
// Vulnerability: the first processor strips the DOCTYPE declaration (looks safe),
// but the second processor re-parses the intermediate string output with a
// permissive factory, allowing a parameter entity injected in element content
// to be expanded. Fixing only the first parser misses the re-parse.

import javax.xml.parsers.*;
import javax.xml.transform.*;
import javax.xml.transform.dom.DOMSource;
import javax.xml.transform.stream.StreamResult;
import javax.xml.transform.stream.StreamSource;
import org.w3c.dom.*;
import java.io.*;
import java.nio.charset.StandardCharsets;

public class XxeProcessingPipeline {

    // Stage 1: strips DOCTYPE — developer believes this neutralises XXE
    private static String stripDoctype(String xml) throws Exception {
        // Naive regex strip: doesn't handle parameter entities in content
        return xml.replaceAll("(?s)<!DOCTYPE[^>]*(?:>|\\[.*?\\]\\s*>)", "");
    }

    // Stage 2: applies an XSLT transform using a separate, permissive factory
    private static String applyTransform(String xml, String xslt) throws Exception {
        TransformerFactory tf = TransformerFactory.newInstance();
        // VULNERABLE: external entities NOT disabled on this factory
        Transformer transformer = tf.newTransformer(
            new StreamSource(new StringReader(xslt))
        );

        DocumentBuilderFactory dbf = DocumentBuilderFactory.newInstance();
        // VULNERABLE: dbf does not set FEATURE_SECURE_PROCESSING or disallow DOCTYPE
        DocumentBuilder db = dbf.newDocumentBuilder();
        Document doc = db.parse(new ByteArrayInputStream(xml.getBytes(StandardCharsets.UTF_8)));

        StringWriter out = new StringWriter();
        transformer.transform(new DOMSource(doc), new StreamResult(out));
        return out.toString();
    }

    // Entry point: receives raw XML from client upload
    public static String processDocument(String userXml, String xsltTemplate) throws Exception {
        // Stage 1 looks safe — DOCTYPE stripped
        String stage1 = stripDoctype(userXml);

        // Stage 2 re-parses with a vulnerable factory; attacker smuggles
        // the entity reference inside element text, not in DOCTYPE:
        // <data>&xxe;</data> where xxe was defined via parameter entity in XSLT
        return applyTransform(stage1, xsltTemplate); // VULNERABLE
    }

    public static void main(String[] args) throws Exception {
        String maliciousXml = "<root><data>&xxe;</data></root>";
        String maliciousXslt =
            "<?xml version='1.0'?>" +
            "<!DOCTYPE xsl:stylesheet [<!ENTITY xxe SYSTEM 'file:///etc/passwd'>]>" +
            "<xsl:stylesheet xmlns:xsl='http://www.w3.org/1999/XSL/Transform' version='1.0'>" +
            "<xsl:template match='/'><xsl:value-of select='/'/></xsl:template>" +
            "</xsl:stylesheet>";
        System.out.println(processDocument(maliciousXml, maliciousXslt));
    }
}
