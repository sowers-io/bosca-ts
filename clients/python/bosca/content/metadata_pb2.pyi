from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Metadata(_message.Message):
    __slots__ = ("default_id", "id", "name", "content_type", "source", "language_tag", "content_length", "trait_ids", "category_ids", "tags", "attributes", "created", "modified", "workflow_state_id", "workflow_state_pending_id")
    class AttributesEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: str
        def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
    DEFAULT_ID_FIELD_NUMBER: _ClassVar[int]
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    CONTENT_TYPE_FIELD_NUMBER: _ClassVar[int]
    SOURCE_FIELD_NUMBER: _ClassVar[int]
    LANGUAGE_TAG_FIELD_NUMBER: _ClassVar[int]
    CONTENT_LENGTH_FIELD_NUMBER: _ClassVar[int]
    TRAIT_IDS_FIELD_NUMBER: _ClassVar[int]
    CATEGORY_IDS_FIELD_NUMBER: _ClassVar[int]
    TAGS_FIELD_NUMBER: _ClassVar[int]
    ATTRIBUTES_FIELD_NUMBER: _ClassVar[int]
    CREATED_FIELD_NUMBER: _ClassVar[int]
    MODIFIED_FIELD_NUMBER: _ClassVar[int]
    WORKFLOW_STATE_ID_FIELD_NUMBER: _ClassVar[int]
    WORKFLOW_STATE_PENDING_ID_FIELD_NUMBER: _ClassVar[int]
    default_id: str
    id: str
    name: str
    content_type: str
    source: str
    language_tag: str
    content_length: int
    trait_ids: _containers.RepeatedScalarFieldContainer[str]
    category_ids: _containers.RepeatedScalarFieldContainer[str]
    tags: _containers.RepeatedScalarFieldContainer[str]
    attributes: _containers.ScalarMap[str, str]
    created: _timestamp_pb2.Timestamp
    modified: _timestamp_pb2.Timestamp
    workflow_state_id: str
    workflow_state_pending_id: str
    def __init__(self, default_id: _Optional[str] = ..., id: _Optional[str] = ..., name: _Optional[str] = ..., content_type: _Optional[str] = ..., source: _Optional[str] = ..., language_tag: _Optional[str] = ..., content_length: _Optional[int] = ..., trait_ids: _Optional[_Iterable[str]] = ..., category_ids: _Optional[_Iterable[str]] = ..., tags: _Optional[_Iterable[str]] = ..., attributes: _Optional[_Mapping[str, str]] = ..., created: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., modified: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., workflow_state_id: _Optional[str] = ..., workflow_state_pending_id: _Optional[str] = ...) -> None: ...
