from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class SignedUrlHeader(_message.Message):
    __slots__ = ("name", "value")
    NAME_FIELD_NUMBER: _ClassVar[int]
    VALUE_FIELD_NUMBER: _ClassVar[int]
    name: str
    value: str
    def __init__(self, name: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...

class SignedUrl(_message.Message):
    __slots__ = ("id", "url", "method", "headers", "attributes")
    class AttributesEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: str
        def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
    ID_FIELD_NUMBER: _ClassVar[int]
    URL_FIELD_NUMBER: _ClassVar[int]
    METHOD_FIELD_NUMBER: _ClassVar[int]
    HEADERS_FIELD_NUMBER: _ClassVar[int]
    ATTRIBUTES_FIELD_NUMBER: _ClassVar[int]
    id: str
    url: str
    method: str
    headers: _containers.RepeatedCompositeFieldContainer[SignedUrlHeader]
    attributes: _containers.ScalarMap[str, str]
    def __init__(self, id: _Optional[str] = ..., url: _Optional[str] = ..., method: _Optional[str] = ..., headers: _Optional[_Iterable[_Union[SignedUrlHeader, _Mapping]]] = ..., attributes: _Optional[_Mapping[str, str]] = ...) -> None: ...
