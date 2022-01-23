#!/usr/bin/env python3

import sys
import xml.etree.ElementTree as ET

from schablone_server_client import AuthenticatedClient
from schablone_server_client.models import *

client = None

# Import XML
def importXML(xml_file_name: str):
    tree = ET.parse(xml_file_name)
    root = tree.getroot()
    # Groups
    for menu in root.iter("script"):
        texts.find("body").text
        texts.find("name").text
        # TODO: send to server


# Start of script
# Syntax ./import_tempaltes.py ${BASE_URL_OF_SERVER} ${TOKEN} ${XML_FILES}
if __name__ == "__main__":
    base_url = sys.argv[1]
    token = sys.argv[2]
    client = AuthenticatedClient(base_url, token)
    for i in range(3, len(sys.argv)):
        xml_path = sys.argv[i]
        importXML(xml_path)
