from imgurpython import ImgurClient

import json
import sys

if __name__ == "__main__":
    if len(sys.argv) < 2:
        sys.exit(1)
    with open('imgur.json') as json_file:
        data = json.load(json_file)
        client_id = data["id"]
        client_secret = data["secret"]
    client = ImgurClient(client_id, client_secret)

    items = client.gallery_search(sys.argv[1])
    if len(items) > 0:
        print(items[0].link)