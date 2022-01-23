#!/usr/bin/env python3

import sys
import xml.etree.ElementTree as ET

from schablone_server_client import AuthenticatedClient
from schablone_server_client.models import *
from schablone_server_client.api.default import post_group_create, post_macro_create

client = None

# Import XML
def importXML(xml_file_name: str):
    tree = ET.parse(xml_file_name)
    root = tree.getroot()

    group_name = xml_file_name.split("/")[::-1][0].split(".")[0]
    response: int = post_group_create.sync(client=client, name=group_name, parent_group_id=1)
    print("Added Group", group_name, "with ID", response)

    # Groups
    for menu in root.iter("script"):
        content=" "
        try:
            content = texts.find("body").text
        except:
            pass
        name = menu.find("name").text
        macroId : int = post_macro_create.sync(
            client=client,
            name=name,
            content=content,
            initial_group=response,
        )
        print("Added Macro", name, "with ID", macroId)

# Start of script
# Syntax ./import_tempaltes.py ${BASE_URL_OF_SERVER} ${TOKEN} ${XML_FILES}
if __name__ == "__main__":
    base_url = sys.argv[1]
    token = sys.argv[2]
    client = AuthenticatedClient(base_url, token=token)
    client.headers = { "X-API-Key" : token}
    print(client.get_headers())
    for i in range(3, len(sys.argv)):
        xml_path = sys.argv[i]
        importXML(xml_path)
