from bosca.content import metadata_pb2 as _metadata_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf import struct_pb2 as _struct_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class CollectionType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    standard: _ClassVar[CollectionType]
    folder: _ClassVar[CollectionType]
    root: _ClassVar[CollectionType]
standard: CollectionType
folder: CollectionType
root: CollectionType

class AddCollectionRequest(_message.Message):
    __slots__ = ("parent", "collection")
    PARENT_FIELD_NUMBER: _ClassVar[int]
    COLLECTION_FIELD_NUMBER: _ClassVar[int]
    parent: str
    collection: Collection
    def __init__(self, parent: _Optional[str] = ..., collection: _Optional[_Union[Collection, _Mapping]] = ...) -> None: ...

class AddCollectionsRequest(_message.Message):
    __slots__ = ("collections",)
    COLLECTIONS_FIELD_NUMBER: _ClassVar[int]
    collections: _containers.RepeatedCompositeFieldContainer[AddCollectionRequest]
    def __init__(self, collections: _Optional[_Iterable[_Union[AddCollectionRequest, _Mapping]]] = ...) -> None: ...

class AddCollectionItemRequest(_message.Message):
    __slots__ = ("collection_id", "child_collection_id", "child_metadata_id")
    COLLECTION_ID_FIELD_NUMBER: _ClassVar[int]
    CHILD_COLLECTION_ID_FIELD_NUMBER: _ClassVar[int]
    CHILD_METADATA_ID_FIELD_NUMBER: _ClassVar[int]
    collection_id: str
    child_collection_id: str
    child_metadata_id: str
    def __init__(self, collection_id: _Optional[str] = ..., child_collection_id: _Optional[str] = ..., child_metadata_id: _Optional[str] = ...) -> None: ...

class Collection(_message.Message):
    __slots__ = ("id", "name", "type", "category_ids", "labels", "attributes", "created", "modified", "metadata")
    class AttributesEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: str
        def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    CATEGORY_IDS_FIELD_NUMBER: _ClassVar[int]
    LABELS_FIELD_NUMBER: _ClassVar[int]
    ATTRIBUTES_FIELD_NUMBER: _ClassVar[int]
    CREATED_FIELD_NUMBER: _ClassVar[int]
    MODIFIED_FIELD_NUMBER: _ClassVar[int]
    METADATA_FIELD_NUMBER: _ClassVar[int]
    id: str
    name: str
    type: CollectionType
    category_ids: _containers.RepeatedScalarFieldContainer[str]
    labels: _containers.RepeatedScalarFieldContainer[str]
    attributes: _containers.ScalarMap[str, str]
    created: _timestamp_pb2.Timestamp
    modified: _timestamp_pb2.Timestamp
    metadata: _struct_pb2.Struct
    def __init__(self, id: _Optional[str] = ..., name: _Optional[str] = ..., type: _Optional[_Union[CollectionType, str]] = ..., category_ids: _Optional[_Iterable[str]] = ..., labels: _Optional[_Iterable[str]] = ..., attributes: _Optional[_Mapping[str, str]] = ..., created: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., modified: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., metadata: _Optional[_Union[_struct_pb2.Struct, _Mapping]] = ...) -> None: ...

class CollectionItems(_message.Message):
    __slots__ = ("items",)
    ITEMS_FIELD_NUMBER: _ClassVar[int]
    items: _containers.RepeatedCompositeFieldContainer[CollectionItem]
    def __init__(self, items: _Optional[_Iterable[_Union[CollectionItem, _Mapping]]] = ...) -> None: ...

class CollectionItem(_message.Message):
    __slots__ = ("metadata", "collection")
    METADATA_FIELD_NUMBER: _ClassVar[int]
    COLLECTION_FIELD_NUMBER: _ClassVar[int]
    metadata: _metadata_pb2.Metadata
    collection: Collection
    def __init__(self, metadata: _Optional[_Union[_metadata_pb2.Metadata, _Mapping]] = ..., collection: _Optional[_Union[Collection, _Mapping]] = ...) -> None: ...
