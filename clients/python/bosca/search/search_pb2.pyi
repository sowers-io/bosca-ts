from google.api import annotations_pb2 as _annotations_pb2
from bosca.content import metadata_pb2 as _metadata_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class SearchRequest(_message.Message):
    __slots__ = ("query", "offset", "limit")
    QUERY_FIELD_NUMBER: _ClassVar[int]
    OFFSET_FIELD_NUMBER: _ClassVar[int]
    LIMIT_FIELD_NUMBER: _ClassVar[int]
    query: str
    offset: int
    limit: int
    def __init__(self, query: _Optional[str] = ..., offset: _Optional[int] = ..., limit: _Optional[int] = ...) -> None: ...

class SearchResponse(_message.Message):
    __slots__ = ("metadata",)
    METADATA_FIELD_NUMBER: _ClassVar[int]
    metadata: _containers.RepeatedCompositeFieldContainer[_metadata_pb2.Metadata]
    def __init__(self, metadata: _Optional[_Iterable[_Union[_metadata_pb2.Metadata, _Mapping]]] = ...) -> None: ...
