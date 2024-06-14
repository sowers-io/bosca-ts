from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Traits(_message.Message):
    __slots__ = ("traits",)
    TRAITS_FIELD_NUMBER: _ClassVar[int]
    traits: _containers.RepeatedCompositeFieldContainer[Trait]
    def __init__(self, traits: _Optional[_Iterable[_Union[Trait, _Mapping]]] = ...) -> None: ...

class Trait(_message.Message):
    __slots__ = ("id", "name", "description", "trait_workflow_ids")
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    TRAIT_WORKFLOW_IDS_FIELD_NUMBER: _ClassVar[int]
    id: str
    name: str
    description: str
    trait_workflow_ids: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, id: _Optional[str] = ..., name: _Optional[str] = ..., description: _Optional[str] = ..., trait_workflow_ids: _Optional[_Iterable[str]] = ...) -> None: ...

class TraitWorkflowStorageSystemRequest(_message.Message):
    __slots__ = ("trait_id", "workflow_id")
    TRAIT_ID_FIELD_NUMBER: _ClassVar[int]
    WORKFLOW_ID_FIELD_NUMBER: _ClassVar[int]
    trait_id: str
    workflow_id: str
    def __init__(self, trait_id: _Optional[str] = ..., workflow_id: _Optional[str] = ...) -> None: ...
