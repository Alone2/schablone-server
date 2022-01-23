#!/usr/bin/env python3

import sys
import xml.etree.ElementTree as ET

from schablone_server_client import AuthenticatedClient
from schablone_server_client.api.default import post_group_create, post_template_create

client = None

# Import XML
def importXML(xml_file_name: str):
    tree = ET.parse(xml_file_name)
    root = tree.getroot()
    # Groups
    for menu in root.iter("menu"):
        group_title = menu.find("title").text

        response: int = post_group_create.sync(client=client, name=group_title, parent_group_id=1)
        print("Added Group", group_title, "with ID", response)
        # Templates in these groups
        for texts in menu.iter("text"):
            template_content = " "
            template_name = " "
            template_subject = " "
            try:
                template_content = texts.find("body").text
                template_name = texts.find("name").text
                template_subject = texts.find("subject").text
            except:
                pass
            # send to server
            template_id: int = post_template_create.sync(
                client=client,
                name=template_name,
                subject=template_subject,
                content=template_content,
                initial_group=response,
            )
            print("Added Template", template_name, "with ID", template_id)


# Start of script
# Syntax ./import_templates.py ${BASE_URL_OF_SERVER} ${TOKEN} ${XML_FILES}
if __name__ == "__main__":
    base_url = sys.argv[1]
    token = sys.argv[2]
    client = AuthenticatedClient(base_url, token=token)
    client.headers = {"X-API-Key": token}
    print(client.get_headers())
    for i in range(3, len(sys.argv)):
        xml_path = sys.argv[i]
        importXML(xml_path)
