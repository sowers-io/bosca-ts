from google.api import annotations_pb2 as _annotations_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class QueryStorageRequest(_message.Message):
    __slots__ = ("storage_system", "query")
    STORAGE_SYSTEM_FIELD_NUMBER: _ClassVar[int]
    QUERY_FIELD_NUMBER: _ClassVar[int]
    storage_system: str
    query: str
    def __init__(self, storage_system: _Optional[str] = ..., query: _Optional[str] = ...) -> None: ...

class QueryPromptRequest(_message.Message):
    __slots__ = ("prompt_id", "model_id", "arguments")
    class ArgumentsEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: str
        def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
    PROMPT_ID_FIELD_NUMBER: _ClassVar[int]
    MODEL_ID_FIELD_NUMBER: _ClassVar[int]
    ARGUMENTS_FIELD_NUMBER: _ClassVar[int]
    prompt_id: str
    model_id: str
    arguments: _containers.ScalarMap[str, str]
    def __init__(self, prompt_id: _Optional[str] = ..., model_id: _Optional[str] = ..., arguments: _Optional[_Mapping[str, str]] = ...) -> None: ...

class QueryResponse(_message.Message):
    __slots__ = ("response",)
    RESPONSE_FIELD_NUMBER: _ClassVar[int]
    response: str
    def __init__(self, response: _Optional[str] = ...) -> None: ...
