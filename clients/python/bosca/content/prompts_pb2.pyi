from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Prompt(_message.Message):
    __slots__ = ("id", "name", "description", "prompt")
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    PROMPT_FIELD_NUMBER: _ClassVar[int]
    id: str
    name: str
    description: str
    prompt: str
    def __init__(self, id: _Optional[str] = ..., name: _Optional[str] = ..., description: _Optional[str] = ..., prompt: _Optional[str] = ...) -> None: ...

class Prompts(_message.Message):
    __slots__ = ("prompts",)
    PROMPTS_FIELD_NUMBER: _ClassVar[int]
    prompts: Prompt
    def __init__(self, prompts: _Optional[_Union[Prompt, _Mapping]] = ...) -> None: ...
