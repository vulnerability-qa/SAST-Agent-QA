# CWE-611: XXE via lxml without disabling entities
from lxml import etree
def parse_config(xml_data):
    parser = etree.XMLParser()  # entities enabled by default
    return etree.fromstring(xml_data, parser)
