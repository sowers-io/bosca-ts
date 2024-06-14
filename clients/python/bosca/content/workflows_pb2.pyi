from bosca.content import metadata_pb2 as _metadata_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

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
unknown_workflow_type: WorkflowStateType
processing: WorkflowStateType
draft: WorkflowStateType
pending: WorkflowStateType
approval: WorkflowStateType
approved: WorkflowStateType
published: WorkflowStateType
failure: WorkflowStateType

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

class TraitWorkflow(_message.Message):
    __slots__ = ("trait_id", "workflow_id", "queue", "metadata")
    TRAIT_ID_FIELD_NUMBER: _ClassVar[int]
    WORKFLOW_ID_FIELD_NUMBER: _ClassVar[int]
    QUEUE_FIELD_NUMBER: _ClassVar[int]
    METADATA_FIELD_NUMBER: _ClassVar[int]
    trait_id: str
    workflow_id: str
    queue: str
    metadata: _metadata_pb2.Metadata
    def __init__(self, trait_id: _Optional[str] = ..., workflow_id: _Optional[str] = ..., queue: _Optional[str] = ..., metadata: _Optional[_Union[_metadata_pb2.Metadata, _Mapping]] = ...) -> None: ...
