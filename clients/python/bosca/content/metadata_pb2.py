# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: bosca/content/metadata.proto
# Protobuf Python Version: 5.26.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1c\x62osca/content/metadata.proto\x12\rbosca.content\x1a\x1fgoogle/protobuf/timestamp.proto\"\xff\x03\n\x08Metadata\x12\x12\n\ndefault_id\x18\x01 \x01(\t\x12\n\n\x02id\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\x12\x14\n\x0c\x63ontent_type\x18\x04 \x01(\t\x12\x13\n\x06source\x18\x05 \x01(\tH\x00\x88\x01\x01\x12\x14\n\x0clanguage_tag\x18\x06 \x01(\t\x12\x16\n\x0e\x63ontent_length\x18\x07 \x01(\x03\x12\x11\n\ttrait_ids\x18\x0b \x03(\t\x12\x14\n\x0c\x63\x61tegory_ids\x18\x0c \x03(\t\x12\x0c\n\x04tags\x18\r \x03(\t\x12;\n\nattributes\x18\x0e \x03(\x0b\x32\'.bosca.content.Metadata.AttributesEntry\x12+\n\x07\x63reated\x18\x14 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12,\n\x08modified\x18\x15 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x19\n\x11workflow_state_id\x18\x1f \x01(\t\x12&\n\x19workflow_state_pending_id\x18  \x01(\tH\x01\x88\x01\x01\x1a\x31\n\x0f\x41ttributesEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\x42\t\n\x07_sourceB\x1c\n\x1a_workflow_state_pending_idB\x1fZ\x1d\x62osca.io/api/protobuf/contentb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'bosca.content.metadata_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z\035bosca.io/api/protobuf/content'
  _globals['_METADATA_ATTRIBUTESENTRY']._loaded_options = None
  _globals['_METADATA_ATTRIBUTESENTRY']._serialized_options = b'8\001'
  _globals['_METADATA']._serialized_start=81
  _globals['_METADATA']._serialized_end=592
  _globals['_METADATA_ATTRIBUTESENTRY']._serialized_start=502
  _globals['_METADATA_ATTRIBUTESENTRY']._serialized_end=551
# @@protoc_insertion_point(module_scope)
