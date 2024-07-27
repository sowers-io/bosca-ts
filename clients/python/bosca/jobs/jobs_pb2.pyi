from google.api import annotations_pb2 as _annotations_pb2
from bosca import empty_pb2 as _empty_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class PollRequest(_message.Message):
    __slots__ = ("queue", "timeout")
    QUEUE_FIELD_NUMBER: _ClassVar[int]
    TIMEOUT_FIELD_NUMBER: _ClassVar[int]
    queue: str
    timeout: int
    def __init__(self, queue: _Optional[str] = ..., timeout: _Optional[int] = ...) -> None: ...

class QueueRequest(_message.Message):
    __slots__ = ("queue", "json")
    QUEUE_FIELD_NUMBER: _ClassVar[int]
    JSON_FIELD_NUMBER: _ClassVar[int]
    queue: str
    json: bytes
    def __init__(self, queue: _Optional[str] = ..., json: _Optional[bytes] = ...) -> None: ...

class FinishRequest(_message.Message):
    __slots__ = ("queue", "id", "success")
    QUEUE_FIELD_NUMBER: _ClassVar[int]
    ID_FIELD_NUMBER: _ClassVar[int]
    SUCCESS_FIELD_NUMBER: _ClassVar[int]
    queue: str
    id: str
    success: bool
    def __init__(self, queue: _Optional[str] = ..., id: _Optional[str] = ..., success: bool = ...) -> None: ...

class Job(_message.Message):
    __slots__ = ("id", "json")
    ID_FIELD_NUMBER: _ClassVar[int]
    JSON_FIELD_NUMBER: _ClassVar[int]
    id: str
    json: bytes
    def __init__(self, id: _Optional[str] = ..., json: _Optional[bytes] = ...) -> None: ...

class QueueResponse(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: str
    def __init__(self, id: _Optional[str] = ...) -> None: ...
