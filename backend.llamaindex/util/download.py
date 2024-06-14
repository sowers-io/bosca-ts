#
# Copyright 2024 Sowers, LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

import logging

import requests

from bosca.content.url_pb2 import SignedUrl


def download_file(signed_url: SignedUrl) -> str:
    logging.debug("downloading %s", signed_url.url)
    headers = {}
    for header in signed_url.headers:
        headers[header.name] = header.value
    response = requests.get(signed_url.url, headers=headers)
    logging.debug("downloaded file: %s -> %s", signed_url.url, response.text)
    if not response.ok:
        raise Exception("failed to download file")
    return response.text
