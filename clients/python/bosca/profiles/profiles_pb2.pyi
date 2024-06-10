from google.api import annotations_pb2 as _annotations_pb2
from google.protobuf import any_pb2 as _any_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from bosca import empty_pb2 as _empty_pb2
from bosca import requests_pb2 as _requests_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class ProfileVisibility(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    system: _ClassVar[ProfileVisibility]
    user: _ClassVar[ProfileVisibility]
    friends: _ClassVar[ProfileVisibility]
    friends_of_friends: _ClassVar[ProfileVisibility]
    public: _ClassVar[ProfileVisibility]
system: ProfileVisibility
user: ProfileVisibility
friends: ProfileVisibility
friends_of_friends: ProfileVisibility
public: ProfileVisibility

class ProfileConfiguration(_message.Message):
    __slots__ = ("avatar_template_url",)
    AVATAR_TEMPLATE_URL_FIELD_NUMBER: _ClassVar[int]
    avatar_template_url: str
    def __init__(self, avatar_template_url: _Optional[str] = ...) -> None: ...

class Profile(_message.Message):
    __slots__ = ("id", "principal", "name", "attributes", "visibility", "created")
    ID_FIELD_NUMBER: _ClassVar[int]
    PRINCIPAL_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    ATTRIBUTES_FIELD_NUMBER: _ClassVar[int]
    VISIBILITY_FIELD_NUMBER: _ClassVar[int]
    CREATED_FIELD_NUMBER: _ClassVar[int]
    id: str
    principal: str
    name: str
    attributes: _containers.RepeatedCompositeFieldContainer[ProfileAttribute]
    visibility: ProfileVisibility
    created: _timestamp_pb2.Timestamp
    def __init__(self, id: _Optional[str] = ..., principal: _Optional[str] = ..., name: _Optional[str] = ..., attributes: _Optional[_Iterable[_Union[ProfileAttribute, _Mapping]]] = ..., visibility: _Optional[_Union[ProfileVisibility, str]] = ..., created: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...

class ProfileAttributeType(_message.Message):
    __slots__ = ("id", "name", "description")
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    id: str
    name: str
    description: str
    def __init__(self, id: _Optional[str] = ..., name: _Optional[str] = ..., description: _Optional[str] = ...) -> None: ...

class ProfileAttributeTypes(_message.Message):
    __slots__ = ("types",)
    TYPES_FIELD_NUMBER: _ClassVar[int]
    types: _containers.RepeatedCompositeFieldContainer[ProfileAttributeType]
    def __init__(self, types: _Optional[_Iterable[_Union[ProfileAttributeType, _Mapping]]] = ...) -> None: ...

class ProfileAttribute(_message.Message):
    __slots__ = ("id", "type_id", "visibility", "value", "confidence", "priority", "source", "created", "expiration")
    ID_FIELD_NUMBER: _ClassVar[int]
    TYPE_ID_FIELD_NUMBER: _ClassVar[int]
    VISIBILITY_FIELD_NUMBER: _ClassVar[int]
    VALUE_FIELD_NUMBER: _ClassVar[int]
    CONFIDENCE_FIELD_NUMBER: _ClassVar[int]
    PRIORITY_FIELD_NUMBER: _ClassVar[int]
    SOURCE_FIELD_NUMBER: _ClassVar[int]
    CREATED_FIELD_NUMBER: _ClassVar[int]
    EXPIRATION_FIELD_NUMBER: _ClassVar[int]
    id: str
    type_id: str
    visibility: ProfileVisibility
    value: _any_pb2.Any
    confidence: float
    priority: float
    source: str
    created: _timestamp_pb2.Timestamp
    expiration: _timestamp_pb2.Timestamp
    def __init__(self, id: _Optional[str] = ..., type_id: _Optional[str] = ..., visibility: _Optional[_Union[ProfileVisibility, str]] = ..., value: _Optional[_Union[_any_pb2.Any, _Mapping]] = ..., confidence: _Optional[float] = ..., priority: _Optional[float] = ..., source: _Optional[str] = ..., created: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., expiration: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...
