from bosca.content import model_pb2 as _model_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class StorageSystemType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    unknown_storage_system: _ClassVar[StorageSystemType]
    vector_storage_system: _ClassVar[StorageSystemType]
    search_storage_system: _ClassVar[StorageSystemType]
    metadata_storage_system: _ClassVar[StorageSystemType]
    supplementary_storage_system: _ClassVar[StorageSystemType]
unknown_storage_system: StorageSystemType
vector_storage_system: StorageSystemType
search_storage_system: StorageSystemType
metadata_storage_system: StorageSystemType
supplementary_storage_system: StorageSystemType

class StorageSystem(_message.Message):
    __slots__ = ("id", "type", "name", "description", "configuration")
    class ConfigurationEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: str
        def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
    ID_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    CONFIGURATION_FIELD_NUMBER: _ClassVar[int]
    id: str
    type: StorageSystemType
    name: str
    description: str
    configuration: _containers.ScalarMap[str, str]
    def __init__(self, id: _Optional[str] = ..., type: _Optional[_Union[StorageSystemType, str]] = ..., name: _Optional[str] = ..., description: _Optional[str] = ..., configuration: _Optional[_Mapping[str, str]] = ...) -> None: ...

class StorageSystems(_message.Message):
    __slots__ = ("systems",)
    SYSTEMS_FIELD_NUMBER: _ClassVar[int]
    systems: _containers.RepeatedCompositeFieldContainer[StorageSystem]
    def __init__(self, systems: _Optional[_Iterable[_Union[StorageSystem, _Mapping]]] = ...) -> None: ...

class StorageSystemModel(_message.Message):
    __slots__ = ("model", "type")
    MODEL_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    model: _model_pb2.Model
    type: str
    def __init__(self, model: _Optional[_Union[_model_pb2.Model, _Mapping]] = ..., type: _Optional[str] = ...) -> None: ...

class StorageSystemModels(_message.Message):
    __slots__ = ("models",)
    MODELS_FIELD_NUMBER: _ClassVar[int]
    models: _containers.RepeatedCompositeFieldContainer[StorageSystemModel]
    def __init__(self, models: _Optional[_Iterable[_Union[StorageSystemModel, _Mapping]]] = ...) -> None: ...
