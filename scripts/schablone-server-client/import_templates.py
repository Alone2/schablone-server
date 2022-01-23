#!/usr/bin/env python3

import sys
import xml.etree.ElementTree as ET

import schablone_server_client.api.default as schablone
from schablone_server_client import AuthenticatedClient

client = None

# Import XML
async def importXML(xml_file_name: str):
    tree = ET.parse(xml_file_name)
    root = tree.getroot()
    # Groups
    for menu in root.iter("menu"):
        group_title = menu.find("title").text

        response: int = await schablone.post_group_create.asyncio(client=client, name=group_title)
        await schablone.post_group_add_parent_group.asyncio_detailed(
            client=client, parent_group_id=1, group_id=response
        )

        # Templates in these groups
        for texts in menu.iter("text"):
            template_content = ""
            template_name = ""
            template_name = ""
            try:
                template_content = texts.find("body").text
                template_name = texts.find("name").text
                template_subject = texts.find("subject").text
            except:
                pass
            # send to server
            schablone.post_template_create.asyncio_detailed(
                client=client,
                name=template_name,
                subject=template_subject,
                content=template_content,
                initial_group=response,
            )
            # TODO: Send to correct group


# Start of script
# Syntax ./import_templates.py ${BASE_URL_OF_SERVER} ${TOKEN} ${XML_FILES}
if __name__ == "__main__":
    base_url = sys.argv[1]
    token = sys.argv[2]
    client = AuthenticatedClient(base_url, token)
    for i in range(3, len(sys.argv)):
        xml_path = sys.argv[i]
        importXML(xml_path)
