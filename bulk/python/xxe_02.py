# CWE-611: XXE via xml.etree.ElementTree (vulnerable to billion laughs)
import defusedxml.ElementTree as ET
def parse_feed(xml_string):
    return ET.fromstring(xml_string)
