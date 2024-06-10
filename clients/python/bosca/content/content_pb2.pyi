from google.api import annotations_pb2 as _annotations_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from bosca.content import metadata_pb2 as _metadata_pb2
from bosca import empty_pb2 as _empty_pb2
from bosca import requests_pb2 as _requests_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class PermissionRelation(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    viewers: _ClassVar[PermissionRelation]
    discoverers: _ClassVar[PermissionRelation]
    editors: _ClassVar[PermissionRelation]
    managers: _ClassVar[PermissionRelation]
    serviceaccounts: _ClassVar[PermissionRelation]
    owners: _ClassVar[PermissionRelation]

class PermissionAction(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    view: _ClassVar[PermissionAction]
    list: _ClassVar[PermissionAction]
    edit: _ClassVar[PermissionAction]
    manage: _ClassVar[PermissionAction]
    service: _ClassVar[PermissionAction]
    delete: _ClassVar[PermissionAction]

class PermissionObjectType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    unknown_object_type: _ClassVar[PermissionObjectType]
    collection_type: _ClassVar[PermissionObjectType]
    metadata_type: _ClassVar[PermissionObjectType]
    system_resource_type: _ClassVar[PermissionObjectType]
    workflow_type: _ClassVar[PermissionObjectType]
    workflow_state_type: _ClassVar[PermissionObjectType]

class PermissionSubjectType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    unknown_subject_type: _ClassVar[PermissionSubjectType]
    user: _ClassVar[PermissionSubjectType]
    group: _ClassVar[PermissionSubjectType]
    service_account: _ClassVar[PermissionSubjectType]

class ItemType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    unknown: _ClassVar[ItemType]
    collection: _ClassVar[ItemType]
    metadata: _ClassVar[ItemType]
    metadata_variant: _ClassVar[ItemType]

class CollectionType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    standard: _ClassVar[CollectionType]
    folder: _ClassVar[CollectionType]
    root: _ClassVar[CollectionType]

class WorkflowStateType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    unknown_workflow_type: _ClassVar[WorkflowStateType]
    processing: _ClassVar[WorkflowStateType]
    draft: _ClassVar[WorkflowStateType]
    pending: _ClassVar[WorkflowStateType]
    approval: _ClassVar[WorkflowStateType]
    approved: _ClassVar[WorkflowStateType]
    published: _ClassVar[WorkflowStateType]
    failure: _ClassVar[WorkflowStateType]
viewers: PermissionRelation
discoverers: PermissionRelation
editors: PermissionRelation
managers: PermissionRelation
serviceaccounts: PermissionRelation
owners: PermissionRelation
view: PermissionAction
list: PermissionAction
edit: PermissionAction
manage: PermissionAction
service: PermissionAction
delete: PermissionAction
unknown_object_type: PermissionObjectType
collection_type: PermissionObjectType
metadata_type: PermissionObjectType
system_resource_type: PermissionObjectType
workflow_type: PermissionObjectType
workflow_state_type: PermissionObjectType
unknown_subject_type: PermissionSubjectType
user: PermissionSubjectType
group: PermissionSubjectType
service_account: PermissionSubjectType
unknown: ItemType
collection: ItemType
metadata: ItemType
metadata_variant: ItemType
standard: CollectionType
folder: CollectionType
root: CollectionType
unknown_workflow_type: WorkflowStateType
processing: WorkflowStateType
draft: WorkflowStateType
pending: WorkflowStateType
approval: WorkflowStateType
approved: WorkflowStateType
published: WorkflowStateType
failure: WorkflowStateType

class TransitionWorkflowRequest(_message.Message):
    __slots__ = ("metadata_id", "state_id", "status", "retry")
    METADATA_ID_FIELD_NUMBER: _ClassVar[int]
    STATE_ID_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    RETRY_FIELD_NUMBER: _ClassVar[int]
    metadata_id: str
    state_id: str
    status: str
    retry: bool
    def __init__(self, metadata_id: _Optional[str] = ..., state_id: _Optional[str] = ..., status: _Optional[str] = ..., retry: bool = ...) -> None: ...

class CompleteTransitionWorkflowRequest(_message.Message):
    __slots__ = ("metadata_id", "status", "success")
    METADATA_ID_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    SUCCESS_FIELD_NUMBER: _ClassVar[int]
    metadata_id: str
    status: str
    success: bool
    def __init__(self, metadata_id: _Optional[str] = ..., status: _Optional[str] = ..., success: bool = ...) -> None: ...

class AddMetadataRelationshipRequest(_message.Message):
    __slots__ = ("metadata_id1", "metadata_id2", "relationship")
    METADATA_ID1_FIELD_NUMBER: _ClassVar[int]
    METADATA_ID2_FIELD_NUMBER: _ClassVar[int]
    RELATIONSHIP_FIELD_NUMBER: _ClassVar[int]
    metadata_id1: str
    metadata_id2: str
    relationship: str
    def __init__(self, metadata_id1: _Optional[str] = ..., metadata_id2: _Optional[str] = ..., relationship: _Optional[str] = ...) -> None: ...

class AddSupplementaryRequest(_message.Message):
    __slots__ = ("id", "type", "name", "content_type", "content_length")
    ID_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    CONTENT_TYPE_FIELD_NUMBER: _ClassVar[int]
    CONTENT_LENGTH_FIELD_NUMBER: _ClassVar[int]
    id: str
    type: str
    name: str
    content_type: str
    content_length: int
    def __init__(self, id: _Optional[str] = ..., type: _Optional[str] = ..., name: _Optional[str] = ..., content_type: _Optional[str] = ..., content_length: _Optional[int] = ...) -> None: ...

class SupplementaryIdRequest(_message.Message):
    __slots__ = ("id", "type")
    ID_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    id: str
    type: str
    def __init__(self, id: _Optional[str] = ..., type: _Optional[str] = ...) -> None: ...

class PermissionCheckRequest(_message.Message):
    __slots__ = ("object", "object_type", "subject", "subject_type", "action")
    OBJECT_FIELD_NUMBER: _ClassVar[int]
    OBJECT_TYPE_FIELD_NUMBER: _ClassVar[int]
    SUBJECT_FIELD_NUMBER: _ClassVar[int]
    SUBJECT_TYPE_FIELD_NUMBER: _ClassVar[int]
    ACTION_FIELD_NUMBER: _ClassVar[int]
    object: str
    object_type: PermissionObjectType
    subject: str
    subject_type: PermissionSubjectType
    action: PermissionAction
    def __init__(self, object: _Optional[str] = ..., object_type: _Optional[_Union[PermissionObjectType, str]] = ..., subject: _Optional[str] = ..., subject_type: _Optional[_Union[PermissionSubjectType, str]] = ..., action: _Optional[_Union[PermissionAction, str]] = ...) -> None: ...

class PermissionCheckResponse(_message.Message):
    __slots__ = ("allowed",)
    ALLOWED_FIELD_NUMBER: _ClassVar[int]
    allowed: bool
    def __init__(self, allowed: bool = ...) -> None: ...

class Permissions(_message.Message):
    __slots__ = ("id", "permissions")
    ID_FIELD_NUMBER: _ClassVar[int]
    PERMISSIONS_FIELD_NUMBER: _ClassVar[int]
    id: str
    permissions: _containers.RepeatedCompositeFieldContainer[Permission]
    def __init__(self, id: _Optional[str] = ..., permissions: _Optional[_Iterable[_Union[Permission, _Mapping]]] = ...) -> None: ...

class Permission(_message.Message):
    __slots__ = ("id", "subject", "subject_type", "relation")
    ID_FIELD_NUMBER: _ClassVar[int]
    SUBJECT_FIELD_NUMBER: _ClassVar[int]
    SUBJECT_TYPE_FIELD_NUMBER: _ClassVar[int]
    RELATION_FIELD_NUMBER: _ClassVar[int]
    id: str
    subject: str
    subject_type: PermissionSubjectType
    relation: PermissionRelation
    def __init__(self, id: _Optional[str] = ..., subject: _Optional[str] = ..., subject_type: _Optional[_Union[PermissionSubjectType, str]] = ..., relation: _Optional[_Union[PermissionRelation, str]] = ...) -> None: ...

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

class AddCollectionRequest(_message.Message):
    __slots__ = ("parent", "collection")
    PARENT_FIELD_NUMBER: _ClassVar[int]
    COLLECTION_FIELD_NUMBER: _ClassVar[int]
    parent: str
    collection: Collection
    def __init__(self, parent: _Optional[str] = ..., collection: _Optional[_Union[Collection, _Mapping]] = ...) -> None: ...

class Workflows(_message.Message):
    __slots__ = ("workflows",)
    WORKFLOWS_FIELD_NUMBER: _ClassVar[int]
    workflows: _containers.RepeatedCompositeFieldContainer[Workflow]
    def __init__(self, workflows: _Optional[_Iterable[_Union[Workflow, _Mapping]]] = ...) -> None: ...

class Workflow(_message.Message):
    __slots__ = ("id", "name", "description", "queue", "configuration")
    class ConfigurationEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: str
        def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    QUEUE_FIELD_NUMBER: _ClassVar[int]
    CONFIGURATION_FIELD_NUMBER: _ClassVar[int]
    id: str
    name: str
    description: str
    queue: str
    configuration: _containers.ScalarMap[str, str]
    def __init__(self, id: _Optional[str] = ..., name: _Optional[str] = ..., description: _Optional[str] = ..., queue: _Optional[str] = ..., configuration: _Optional[_Mapping[str, str]] = ...) -> None: ...

class WorkflowStates(_message.Message):
    __slots__ = ("states",)
    STATES_FIELD_NUMBER: _ClassVar[int]
    states: _containers.RepeatedCompositeFieldContainer[WorkflowState]
    def __init__(self, states: _Optional[_Iterable[_Union[WorkflowState, _Mapping]]] = ...) -> None: ...

class WorkflowState(_message.Message):
    __slots__ = ("id", "name", "description", "type", "configuration", "workflow_id", "entry_workflow_id", "exit_workflow_id")
    class ConfigurationEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: str
        def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    CONFIGURATION_FIELD_NUMBER: _ClassVar[int]
    WORKFLOW_ID_FIELD_NUMBER: _ClassVar[int]
    ENTRY_WORKFLOW_ID_FIELD_NUMBER: _ClassVar[int]
    EXIT_WORKFLOW_ID_FIELD_NUMBER: _ClassVar[int]
    id: str
    name: str
    description: str
    type: WorkflowStateType
    configuration: _containers.ScalarMap[str, str]
    workflow_id: str
    entry_workflow_id: str
    exit_workflow_id: str
    def __init__(self, id: _Optional[str] = ..., name: _Optional[str] = ..., description: _Optional[str] = ..., type: _Optional[_Union[WorkflowStateType, str]] = ..., configuration: _Optional[_Mapping[str, str]] = ..., workflow_id: _Optional[str] = ..., entry_workflow_id: _Optional[str] = ..., exit_workflow_id: _Optional[str] = ...) -> None: ...

class WorkflowStateTransitions(_message.Message):
    __slots__ = ("transitions",)
    TRANSITIONS_FIELD_NUMBER: _ClassVar[int]
    transitions: _containers.RepeatedCompositeFieldContainer[WorkflowStateTransition]
    def __init__(self, transitions: _Optional[_Iterable[_Union[WorkflowStateTransition, _Mapping]]] = ...) -> None: ...

class WorkflowStateTransition(_message.Message):
    __slots__ = ("from_state_id", "to_state_id", "description")
    FROM_STATE_ID_FIELD_NUMBER: _ClassVar[int]
    TO_STATE_ID_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    from_state_id: str
    to_state_id: str
    description: str
    def __init__(self, from_state_id: _Optional[str] = ..., to_state_id: _Optional[str] = ..., description: _Optional[str] = ...) -> None: ...

class Traits(_message.Message):
    __slots__ = ("traits",)
    TRAITS_FIELD_NUMBER: _ClassVar[int]
    traits: _containers.RepeatedCompositeFieldContainer[Trait]
    def __init__(self, traits: _Optional[_Iterable[_Union[Trait, _Mapping]]] = ...) -> None: ...

class Trait(_message.Message):
    __slots__ = ("id", "name", "workflow_id")
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    WORKFLOW_ID_FIELD_NUMBER: _ClassVar[int]
    id: str
    name: str
    workflow_id: str
    def __init__(self, id: _Optional[str] = ..., name: _Optional[str] = ..., workflow_id: _Optional[str] = ...) -> None: ...

class Collection(_message.Message):
    __slots__ = ("id", "name", "type", "category_ids", "tags", "attributes", "created", "modified")
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
    TAGS_FIELD_NUMBER: _ClassVar[int]
    ATTRIBUTES_FIELD_NUMBER: _ClassVar[int]
    CREATED_FIELD_NUMBER: _ClassVar[int]
    MODIFIED_FIELD_NUMBER: _ClassVar[int]
    id: str
    name: str
    type: CollectionType
    category_ids: _containers.RepeatedScalarFieldContainer[str]
    tags: _containers.RepeatedScalarFieldContainer[str]
    attributes: _containers.ScalarMap[str, str]
    created: _timestamp_pb2.Timestamp
    modified: _timestamp_pb2.Timestamp
    def __init__(self, id: _Optional[str] = ..., name: _Optional[str] = ..., type: _Optional[_Union[CollectionType, str]] = ..., category_ids: _Optional[_Iterable[str]] = ..., tags: _Optional[_Iterable[str]] = ..., attributes: _Optional[_Mapping[str, str]] = ..., created: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., modified: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...

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

class AddMetadataRequest(_message.Message):
    __slots__ = ("collection", "metadata")
    COLLECTION_FIELD_NUMBER: _ClassVar[int]
    METADATA_FIELD_NUMBER: _ClassVar[int]
    collection: str
    metadata: _metadata_pb2.Metadata
    def __init__(self, collection: _Optional[str] = ..., metadata: _Optional[_Union[_metadata_pb2.Metadata, _Mapping]] = ...) -> None: ...

class Metadatas(_message.Message):
    __slots__ = ("metadata",)
    METADATA_FIELD_NUMBER: _ClassVar[int]
    metadata: _containers.RepeatedCompositeFieldContainer[_metadata_pb2.Metadata]
    def __init__(self, metadata: _Optional[_Iterable[_Union[_metadata_pb2.Metadata, _Mapping]]] = ...) -> None: ...

class AddMetadataTraitRequest(_message.Message):
    __slots__ = ("metadata_id", "trait_id")
    METADATA_ID_FIELD_NUMBER: _ClassVar[int]
    TRAIT_ID_FIELD_NUMBER: _ClassVar[int]
    metadata_id: str
    trait_id: str
    def __init__(self, metadata_id: _Optional[str] = ..., trait_id: _Optional[str] = ...) -> None: ...
