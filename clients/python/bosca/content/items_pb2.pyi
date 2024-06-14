from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class ItemType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    unknown: _ClassVar[ItemType]
    collection: _ClassVar[ItemType]
    metadata: _ClassVar[ItemType]
    metadata_variant: _ClassVar[ItemType]
unknown: ItemType
collection: ItemType
metadata: ItemType
metadata_variant: ItemType

class Item(_message.Message):
    __slots__ = ("id", "name", "category_ids", "tags", "attributes", "created", "modified")
    class AttributesEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: str
        def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    CATEGORY_IDS_FIELD_NUMBER: _ClassVar[int]
    TAGS_FIELD_NUMBER: _ClassVar[int]
    ATTRIBUTES_FIELD_NUMBER: _ClassVar[int]
    CREATED_FIELD_NUMBER: _ClassVar[int]
    MODIFIED_FIELD_NUMBER: _ClassVar[int]
    id: str
    name: str
    category_ids: _containers.RepeatedScalarFieldContainer[str]
    tags: _containers.RepeatedScalarFieldContainer[str]
    attributes: _containers.ScalarMap[str, str]
    created: _timestamp_pb2.Timestamp
    modified: _timestamp_pb2.Timestamp
    def __init__(self, id: _Optional[str] = ..., name: _Optional[str] = ..., category_ids: _Optional[_Iterable[str]] = ..., tags: _Optional[_Iterable[str]] = ..., attributes: _Optional[_Mapping[str, str]] = ..., created: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., modified: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...
