from imgurpython import ImgurClient

import json
import sys
import random

if __name__ == "__main__":
    if len(sys.argv) < 4:
        sys.exit(1)
    with open('ImgurGo/imgur.json') as json_file:
        data = json.load(json_file)
        client_id = data["id"]
        client_secret = data["secret"]
    client = ImgurClient(client_id, client_secret)

    items = client.gallery_search(sys.argv[3], sort=sys.argv[1], window=sys.argv[2])
    if len(items) > 0:
        index = random.randint(0, len(items) - 1)
        print(items[index].link)