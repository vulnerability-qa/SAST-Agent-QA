# CWE-611: XXE via lxml without disabling entities
from lxml import etree
def parse_config(xml_data):
    parser = etree.XMLParser(resolve_entities=False, no_network=True, dtd_validation=False, load_dtd=False)  # entities enabled by default
    return etree.fromstring(xml_data, parser)
