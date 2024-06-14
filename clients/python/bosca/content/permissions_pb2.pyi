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
