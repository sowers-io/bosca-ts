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
from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1c\x62osca/content/metadata.proto\x12\rbosca.content\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1cgoogle/protobuf/struct.proto\"\xe6\x04\n\x08Metadata\x12\x12\n\ndefault_id\x18\x01 \x01(\t\x12\n\n\x02id\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\x12\x14\n\x0c\x63ontent_type\x18\x04 \x01(\t\x12\x16\n\tsource_id\x18\x05 \x01(\tH\x00\x88\x01\x01\x12\x1e\n\x11source_identifier\x18\x06 \x01(\tH\x01\x88\x01\x01\x12\x14\n\x0clanguage_tag\x18\x07 \x01(\t\x12\x16\n\x0e\x63ontent_length\x18\x08 \x01(\x03\x12\x11\n\ttrait_ids\x18\x0b \x03(\t\x12\x14\n\x0c\x63\x61tegory_ids\x18\x0c \x03(\t\x12\x0c\n\x04tags\x18\r \x03(\t\x12;\n\nattributes\x18\x0e \x03(\x0b\x32\'.bosca.content.Metadata.AttributesEntry\x12+\n\x07\x63reated\x18\x14 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12,\n\x08modified\x18\x15 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x19\n\x11workflow_state_id\x18\x1f \x01(\t\x12&\n\x19workflow_state_pending_id\x18  \x01(\tH\x02\x88\x01\x01\x12)\n\x08metadata\x18! \x01(\x0b\x32\x17.google.protobuf.Struct\x1a\x31\n\x0f\x41ttributesEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\x42\x0c\n\n_source_idB\x14\n\x12_source_identifierB\x1c\n\x1a_workflow_state_pending_id\"S\n\x12\x41\x64\x64MetadataRequest\x12\x12\n\ncollection\x18\x01 \x01(\t\x12)\n\x08metadata\x18\x02 \x01(\x0b\x32\x17.bosca.content.Metadata\"6\n\tMetadatas\x12)\n\x08metadata\x18\x01 \x03(\x0b\x32\x17.bosca.content.Metadata\"b\n\x1e\x41\x64\x64MetadataRelationshipRequest\x12\x14\n\x0cmetadata_id1\x18\x01 \x01(\t\x12\x14\n\x0cmetadata_id2\x18\x02 \x01(\t\x12\x14\n\x0crelationship\x18\x03 \x01(\t\"@\n\x17\x41\x64\x64MetadataTraitRequest\x12\x13\n\x0bmetadata_id\x18\x01 \x01(\t\x12\x10\n\x08trait_id\x18\x02 \x01(\t\"o\n\x17\x41\x64\x64SupplementaryRequest\x12\n\n\x02id\x18\x01 \x01(\t\x12\x0c\n\x04type\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\x12\x14\n\x0c\x63ontent_type\x18\x04 \x01(\t\x12\x16\n\x0e\x63ontent_length\x18\x05 \x01(\x03\"2\n\x16SupplementaryIdRequest\x12\n\n\x02id\x18\x01 \x01(\t\x12\x0c\n\x04type\x18\x02 \x01(\tB%Z#bosca.io/api/protobuf/bosca/contentb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'bosca.content.metadata_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z#bosca.io/api/protobuf/bosca/content'
  _globals['_METADATA_ATTRIBUTESENTRY']._loaded_options = None
  _globals['_METADATA_ATTRIBUTESENTRY']._serialized_options = b'8\001'
  _globals['_METADATA']._serialized_start=111
  _globals['_METADATA']._serialized_end=725
  _globals['_METADATA_ATTRIBUTESENTRY']._serialized_start=610
  _globals['_METADATA_ATTRIBUTESENTRY']._serialized_end=659
  _globals['_ADDMETADATAREQUEST']._serialized_start=727
  _globals['_ADDMETADATAREQUEST']._serialized_end=810
  _globals['_METADATAS']._serialized_start=812
  _globals['_METADATAS']._serialized_end=866
  _globals['_ADDMETADATARELATIONSHIPREQUEST']._serialized_start=868
  _globals['_ADDMETADATARELATIONSHIPREQUEST']._serialized_end=966
  _globals['_ADDMETADATATRAITREQUEST']._serialized_start=968
  _globals['_ADDMETADATATRAITREQUEST']._serialized_end=1032
  _globals['_ADDSUPPLEMENTARYREQUEST']._serialized_start=1034
  _globals['_ADDSUPPLEMENTARYREQUEST']._serialized_end=1145
  _globals['_SUPPLEMENTARYIDREQUEST']._serialized_start=1147
  _globals['_SUPPLEMENTARYIDREQUEST']._serialized_end=1197
# @@protoc_insertion_point(module_scope)
