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

import os

from llama_index.core import VectorStoreIndex, Settings, ServiceContext
from llama_index.embeddings.ollama import OllamaEmbedding
from llama_index.llms.ollama import Ollama
from llama_index.vector_stores.qdrant import QdrantVectorStore
from qdrant_client import qdrant_client

from bosca.content.storage_systems_pb2 import StorageSystem
from bosca.content.model_pb2 import Model


class ChatContext(object):

    def __init__(self, model: Model, storage_system: StorageSystem, service_context: ServiceContext | None = None):
        client_connection_parts = os.environ["BOSCA_QDRANT_API_ADDRESS"].split(":")

        self.llm = Ollama(
            model=model.name,
            context_window=model.configuration["contextWindow"],
            base_url=os.environ["BOSCA_OLLAMA_API_ADDRESS"],
            request_timeout=120
        )
        self.embeddings = OllamaEmbedding(
            model_name=model.name,
            base_url=os.environ["BOSCA_OLLAMA_API_ADDRESS"],
            embed_batch_size=100
        )
        self.qdrant_vector_client = qdrant_client.QdrantClient(
            host=client_connection_parts[0],
            grpc_port=int(client_connection_parts[1])
        )

        Settings.chunk_size = 1024
        Settings.chunk_overlap = 20
        Settings.llm = self.llm
        Settings.embed_model = self.embeddings
        Settings.callback_manager = self.llm.callback_manager

        self.service_context = service_context
        self.vector_store = QdrantVectorStore(client=self.qdrant_vector_client,
                                              collection_name=storage_system.configuration["indexName"],
                                              parallel=2)

        self.vector_store_index = VectorStoreIndex.from_vector_store(vector_store=self.vector_store,
                                                                     embed_model=self.embeddings,
                                                                     service_context=self.service_context)
