import logging

from bosca.content.content_pb2 import SignedUrl

import requests


def download_file(signed_url: SignedUrl) -> str:
    logging.debug("downloading %s", signed_url.url)
    headers = {}
    for header in signed_url.headers:
        headers[header.name] = header.value
    response = requests.get(signed_url.url, headers=headers)
    logging.debug("downloaded file: %s -> %s", signed_url.url, response.text)
    return response.text
