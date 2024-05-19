from langchain.vectorstores import Qdrant
from langchain_community.embeddings import OllamaEmbeddings
from qdrant_client import QdrantClient

client = QdrantClient(url="http://localhost:6333")


def process():
    embeddings = OllamaEmbeddings(model="llama3:70b")
    doc_store = Qdrant.from_texts(
        texts, embeddings, url="<qdrant-url>", api_key="<qdrant-api-key>", collection_name="metadata"
    )