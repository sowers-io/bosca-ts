from bosca.content import metadata_pb2 as _metadata_pb2
from bosca.content import prompts_pb2 as _prompts_pb2
from bosca.content import storage_systems_pb2 as _storage_systems_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class WorkflowActivityParameterType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    unknown_activity_parameter_type: _ClassVar[WorkflowActivityParameterType]
    context: _ClassVar[WorkflowActivityParameterType]
    supplementary: _ClassVar[WorkflowActivityParameterType]
    supplementary_array: _ClassVar[WorkflowActivityParameterType]

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
unknown_activity_parameter_type: WorkflowActivityParameterType
context: WorkflowActivityParameterType
supplementary: WorkflowActivityParameterType
supplementary_array: WorkflowActivityParameterType
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

class WorkflowActivity(_message.Message):
    __slots__ = ("id", "name", "description", "child_workflow", "child_workflow_queue", "configuration", "inputs", "outputs")
    class ConfigurationEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: str
        def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
    class InputsEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: WorkflowActivityParameterType
        def __init__(self, key: _Optional[str] = ..., value: _Optional[_Union[WorkflowActivityParameterType, str]] = ...) -> None: ...
    class OutputsEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: WorkflowActivityParameterType
        def __init__(self, key: _Optional[str] = ..., value: _Optional[_Union[WorkflowActivityParameterType, str]] = ...) -> None: ...
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    CHILD_WORKFLOW_FIELD_NUMBER: _ClassVar[int]
    CHILD_WORKFLOW_QUEUE_FIELD_NUMBER: _ClassVar[int]
    CONFIGURATION_FIELD_NUMBER: _ClassVar[int]
    INPUTS_FIELD_NUMBER: _ClassVar[int]
    OUTPUTS_FIELD_NUMBER: _ClassVar[int]
    id: str
    name: str
    description: str
    child_workflow: bool
    child_workflow_queue: str
    configuration: _containers.ScalarMap[str, str]
    inputs: _containers.ScalarMap[str, WorkflowActivityParameterType]
    outputs: _containers.ScalarMap[str, WorkflowActivityParameterType]
    def __init__(self, id: _Optional[str] = ..., name: _Optional[str] = ..., description: _Optional[str] = ..., child_workflow: bool = ..., child_workflow_queue: _Optional[str] = ..., configuration: _Optional[_Mapping[str, str]] = ..., inputs: _Optional[_Mapping[str, WorkflowActivityParameterType]] = ..., outputs: _Optional[_Mapping[str, WorkflowActivityParameterType]] = ...) -> None: ...

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

class WorkflowActivityParameterValue(_message.Message):
    __slots__ = ("single_value", "array_value")
    SINGLE_VALUE_FIELD_NUMBER: _ClassVar[int]
    ARRAY_VALUE_FIELD_NUMBER: _ClassVar[int]
    single_value: str
    array_value: WorkflowActivityParameterValues
    def __init__(self, single_value: _Optional[str] = ..., array_value: _Optional[_Union[WorkflowActivityParameterValues, _Mapping]] = ...) -> None: ...

class WorkflowActivityParameterValues(_message.Message):
    __slots__ = ("values",)
    VALUES_FIELD_NUMBER: _ClassVar[int]
    values: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, values: _Optional[_Iterable[str]] = ...) -> None: ...

class WorkflowActivityInstance(_message.Message):
    __slots__ = ("id", "child_workflow", "child_workflow_queue", "execution_group", "configuration", "inputs", "outputs")
    class ConfigurationEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: str
        def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
    class InputsEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: WorkflowActivityParameterValue
        def __init__(self, key: _Optional[str] = ..., value: _Optional[_Union[WorkflowActivityParameterValue, _Mapping]] = ...) -> None: ...
    class OutputsEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: WorkflowActivityParameterValue
        def __init__(self, key: _Optional[str] = ..., value: _Optional[_Union[WorkflowActivityParameterValue, _Mapping]] = ...) -> None: ...
    ID_FIELD_NUMBER: _ClassVar[int]
    CHILD_WORKFLOW_FIELD_NUMBER: _ClassVar[int]
    CHILD_WORKFLOW_QUEUE_FIELD_NUMBER: _ClassVar[int]
    EXECUTION_GROUP_FIELD_NUMBER: _ClassVar[int]
    CONFIGURATION_FIELD_NUMBER: _ClassVar[int]
    INPUTS_FIELD_NUMBER: _ClassVar[int]
    OUTPUTS_FIELD_NUMBER: _ClassVar[int]
    id: str
    child_workflow: bool
    child_workflow_queue: str
    execution_group: int
    configuration: _containers.ScalarMap[str, str]
    inputs: _containers.MessageMap[str, WorkflowActivityParameterValue]
    outputs: _containers.MessageMap[str, WorkflowActivityParameterValue]
    def __init__(self, id: _Optional[str] = ..., child_workflow: bool = ..., child_workflow_queue: _Optional[str] = ..., execution_group: _Optional[int] = ..., configuration: _Optional[_Mapping[str, str]] = ..., inputs: _Optional[_Mapping[str, WorkflowActivityParameterValue]] = ..., outputs: _Optional[_Mapping[str, WorkflowActivityParameterValue]] = ...) -> None: ...

class WorkflowInstance(_message.Message):
    __slots__ = ("trait_id", "workflow", "activities", "queue", "metadata", "context")
    class ContextEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: str
        def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
    TRAIT_ID_FIELD_NUMBER: _ClassVar[int]
    WORKFLOW_FIELD_NUMBER: _ClassVar[int]
    ACTIVITIES_FIELD_NUMBER: _ClassVar[int]
    QUEUE_FIELD_NUMBER: _ClassVar[int]
    METADATA_FIELD_NUMBER: _ClassVar[int]
    CONTEXT_FIELD_NUMBER: _ClassVar[int]
    trait_id: str
    workflow: Workflow
    activities: _containers.RepeatedCompositeFieldContainer[WorkflowActivityInstance]
    queue: str
    metadata: _metadata_pb2.Metadata
    context: _containers.ScalarMap[str, str]
    def __init__(self, trait_id: _Optional[str] = ..., workflow: _Optional[_Union[Workflow, _Mapping]] = ..., activities: _Optional[_Iterable[_Union[WorkflowActivityInstance, _Mapping]]] = ..., queue: _Optional[str] = ..., metadata: _Optional[_Union[_metadata_pb2.Metadata, _Mapping]] = ..., context: _Optional[_Mapping[str, str]] = ...) -> None: ...

class WorkflowActivityExecutionContext(_message.Message):
    __slots__ = ("workflow_id", "trait_id", "activity", "metadata", "context", "inputs", "outputs")
    class ContextEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: WorkflowActivityParameterValue
        def __init__(self, key: _Optional[str] = ..., value: _Optional[_Union[WorkflowActivityParameterValue, _Mapping]] = ...) -> None: ...
    class InputsEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: WorkflowActivityParameterValue
        def __init__(self, key: _Optional[str] = ..., value: _Optional[_Union[WorkflowActivityParameterValue, _Mapping]] = ...) -> None: ...
    class OutputsEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: WorkflowActivityParameterValue
        def __init__(self, key: _Optional[str] = ..., value: _Optional[_Union[WorkflowActivityParameterValue, _Mapping]] = ...) -> None: ...
    WORKFLOW_ID_FIELD_NUMBER: _ClassVar[int]
    TRAIT_ID_FIELD_NUMBER: _ClassVar[int]
    ACTIVITY_FIELD_NUMBER: _ClassVar[int]
    METADATA_FIELD_NUMBER: _ClassVar[int]
    CONTEXT_FIELD_NUMBER: _ClassVar[int]
    INPUTS_FIELD_NUMBER: _ClassVar[int]
    OUTPUTS_FIELD_NUMBER: _ClassVar[int]
    workflow_id: str
    trait_id: str
    activity: WorkflowActivityInstance
    metadata: _metadata_pb2.Metadata
    context: _containers.MessageMap[str, WorkflowActivityParameterValue]
    inputs: _containers.MessageMap[str, WorkflowActivityParameterValue]
    outputs: _containers.MessageMap[str, WorkflowActivityParameterValue]
    def __init__(self, workflow_id: _Optional[str] = ..., trait_id: _Optional[str] = ..., activity: _Optional[_Union[WorkflowActivityInstance, _Mapping]] = ..., metadata: _Optional[_Union[_metadata_pb2.Metadata, _Mapping]] = ..., context: _Optional[_Mapping[str, WorkflowActivityParameterValue]] = ..., inputs: _Optional[_Mapping[str, WorkflowActivityParameterValue]] = ..., outputs: _Optional[_Mapping[str, WorkflowActivityParameterValue]] = ...) -> None: ...

class WorkflowActivityStorageSystem(_message.Message):
    __slots__ = ("storage_system", "configuration")
    class ConfigurationEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: str
        def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
    STORAGE_SYSTEM_FIELD_NUMBER: _ClassVar[int]
    CONFIGURATION_FIELD_NUMBER: _ClassVar[int]
    storage_system: _storage_systems_pb2.StorageSystem
    configuration: _containers.ScalarMap[str, str]
    def __init__(self, storage_system: _Optional[_Union[_storage_systems_pb2.StorageSystem, _Mapping]] = ..., configuration: _Optional[_Mapping[str, str]] = ...) -> None: ...

class WorkflowActivityStorageSystems(_message.Message):
    __slots__ = ("systems",)
    SYSTEMS_FIELD_NUMBER: _ClassVar[int]
    systems: _containers.RepeatedCompositeFieldContainer[WorkflowActivityStorageSystem]
    def __init__(self, systems: _Optional[_Iterable[_Union[WorkflowActivityStorageSystem, _Mapping]]] = ...) -> None: ...

class WorkflowActivityPrompt(_message.Message):
    __slots__ = ("prompt", "configuration")
    class ConfigurationEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: str
        def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
    PROMPT_FIELD_NUMBER: _ClassVar[int]
    CONFIGURATION_FIELD_NUMBER: _ClassVar[int]
    prompt: _prompts_pb2.Prompt
    configuration: _containers.ScalarMap[str, str]
    def __init__(self, prompt: _Optional[_Union[_prompts_pb2.Prompt, _Mapping]] = ..., configuration: _Optional[_Mapping[str, str]] = ...) -> None: ...

class WorkflowActivityPrompts(_message.Message):
    __slots__ = ("prompts",)
    PROMPTS_FIELD_NUMBER: _ClassVar[int]
    prompts: WorkflowActivityPrompt
    def __init__(self, prompts: _Optional[_Union[WorkflowActivityPrompt, _Mapping]] = ...) -> None: ...
