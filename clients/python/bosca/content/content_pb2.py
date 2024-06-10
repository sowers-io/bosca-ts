# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: bosca/content/content.proto
# Protobuf Python Version: 5.26.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from bosca.content import metadata_pb2 as bosca_dot_content_dot_metadata__pb2
from bosca import empty_pb2 as bosca_dot_empty__pb2
from bosca import requests_pb2 as bosca_dot_requests__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1b\x62osca/content/content.proto\x12\rbosca.content\x1a\x1cgoogle/api/annotations.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1c\x62osca/content/metadata.proto\x1a\x11\x62osca/empty.proto\x1a\x14\x62osca/requests.proto\"a\n\x19TransitionWorkflowRequest\x12\x13\n\x0bmetadata_id\x18\x01 \x01(\t\x12\x10\n\x08state_id\x18\x02 \x01(\t\x12\x0e\n\x06status\x18\x03 \x01(\t\x12\r\n\x05retry\x18\x04 \x01(\x08\"Y\n!CompleteTransitionWorkflowRequest\x12\x13\n\x0bmetadata_id\x18\x01 \x01(\t\x12\x0e\n\x06status\x18\x02 \x01(\t\x12\x0f\n\x07success\x18\x03 \x01(\x08\"b\n\x1e\x41\x64\x64MetadataRelationshipRequest\x12\x14\n\x0cmetadata_id1\x18\x01 \x01(\t\x12\x14\n\x0cmetadata_id2\x18\x02 \x01(\t\x12\x14\n\x0crelationship\x18\x03 \x01(\t\"o\n\x17\x41\x64\x64SupplementaryRequest\x12\n\n\x02id\x18\x01 \x01(\t\x12\x0c\n\x04type\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\x12\x14\n\x0c\x63ontent_type\x18\x04 \x01(\t\x12\x16\n\x0e\x63ontent_length\x18\x05 \x01(\x03\"2\n\x16SupplementaryIdRequest\x12\n\n\x02id\x18\x01 \x01(\t\x12\x0c\n\x04type\x18\x02 \x01(\t\"\xe0\x01\n\x16PermissionCheckRequest\x12\x0e\n\x06object\x18\x01 \x01(\t\x12\x38\n\x0bobject_type\x18\x02 \x01(\x0e\x32#.bosca.content.PermissionObjectType\x12\x0f\n\x07subject\x18\x03 \x01(\t\x12:\n\x0csubject_type\x18\x04 \x01(\x0e\x32$.bosca.content.PermissionSubjectType\x12/\n\x06\x61\x63tion\x18\x05 \x01(\x0e\x32\x1f.bosca.content.PermissionAction\"*\n\x17PermissionCheckResponse\x12\x0f\n\x07\x61llowed\x18\x01 \x01(\x08\"I\n\x0bPermissions\x12\n\n\x02id\x18\x01 \x01(\t\x12.\n\x0bpermissions\x18\x02 \x03(\x0b\x32\x19.bosca.content.Permission\"\x9a\x01\n\nPermission\x12\n\n\x02id\x18\x01 \x01(\t\x12\x0f\n\x07subject\x18\x02 \x01(\t\x12:\n\x0csubject_type\x18\x03 \x01(\x0e\x32$.bosca.content.PermissionSubjectType\x12\x33\n\x08relation\x18\x05 \x01(\x0e\x32!.bosca.content.PermissionRelation\"\x8b\x02\n\x04Item\x12\n\n\x02id\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\x12\x14\n\x0c\x63\x61tegory_ids\x18\x0c \x03(\t\x12\x0c\n\x04tags\x18\r \x03(\t\x12\x37\n\nattributes\x18\x0e \x03(\x0b\x32#.bosca.content.Item.AttributesEntry\x12+\n\x07\x63reated\x18\x14 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12,\n\x08modified\x18\x15 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x1a\x31\n\x0f\x41ttributesEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"U\n\x14\x41\x64\x64\x43ollectionRequest\x12\x0e\n\x06parent\x18\x01 \x01(\t\x12-\n\ncollection\x18\x02 \x01(\x0b\x32\x19.bosca.content.Collection\"7\n\tWorkflows\x12*\n\tworkflows\x18\x01 \x03(\x0b\x32\x17.bosca.content.Workflow\"\xc1\x01\n\x08Workflow\x12\n\n\x02id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x03 \x01(\t\x12\r\n\x05queue\x18\x04 \x01(\t\x12\x41\n\rconfiguration\x18\x05 \x03(\x0b\x32*.bosca.content.Workflow.ConfigurationEntry\x1a\x34\n\x12\x43onfigurationEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\">\n\x0eWorkflowStates\x12,\n\x06states\x18\x01 \x03(\x0b\x32\x1c.bosca.content.WorkflowState\"\x80\x03\n\rWorkflowState\x12\n\n\x02id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x03 \x01(\t\x12.\n\x04type\x18\x04 \x01(\x0e\x32 .bosca.content.WorkflowStateType\x12\x46\n\rconfiguration\x18\x05 \x03(\x0b\x32/.bosca.content.WorkflowState.ConfigurationEntry\x12\x18\n\x0bworkflow_id\x18\x06 \x01(\tH\x00\x88\x01\x01\x12\x1e\n\x11\x65ntry_workflow_id\x18\x07 \x01(\tH\x01\x88\x01\x01\x12\x1d\n\x10\x65xit_workflow_id\x18\x08 \x01(\tH\x02\x88\x01\x01\x1a\x34\n\x12\x43onfigurationEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\x42\x0e\n\x0c_workflow_idB\x14\n\x12_entry_workflow_idB\x13\n\x11_exit_workflow_id\"W\n\x18WorkflowStateTransitions\x12;\n\x0btransitions\x18\x01 \x03(\x0b\x32&.bosca.content.WorkflowStateTransition\"Z\n\x17WorkflowStateTransition\x12\x15\n\rfrom_state_id\x18\x01 \x01(\t\x12\x13\n\x0bto_state_id\x18\x02 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x03 \x01(\t\".\n\x06Traits\x12$\n\x06traits\x18\x01 \x03(\x0b\x32\x14.bosca.content.Trait\"6\n\x05Trait\x12\n\n\x02id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x13\n\x0bworkflow_id\x18\x03 \x01(\t\"\xc4\x02\n\nCollection\x12\n\n\x02id\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\x12+\n\x04type\x18\x05 \x01(\x0e\x32\x1d.bosca.content.CollectionType\x12\x14\n\x0c\x63\x61tegory_ids\x18\x0c \x03(\t\x12\x0c\n\x04tags\x18\r \x03(\t\x12=\n\nattributes\x18\x0e \x03(\x0b\x32).bosca.content.Collection.AttributesEntry\x12+\n\x07\x63reated\x18\x14 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12,\n\x08modified\x18\x15 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x1a\x31\n\x0f\x41ttributesEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"?\n\x0f\x43ollectionItems\x12,\n\x05items\x18\x01 \x03(\x0b\x32\x1d.bosca.content.CollectionItem\"v\n\x0e\x43ollectionItem\x12+\n\x08metadata\x18\x01 \x01(\x0b\x32\x17.bosca.content.MetadataH\x00\x12/\n\ncollection\x18\x03 \x01(\x0b\x32\x19.bosca.content.CollectionH\x00\x42\x06\n\x04Item\".\n\x0fSignedUrlHeader\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t\"\xd6\x01\n\tSignedUrl\x12\n\n\x02id\x18\x01 \x01(\t\x12\x0b\n\x03url\x18\x02 \x01(\t\x12\x0e\n\x06method\x18\x03 \x01(\t\x12/\n\x07headers\x18\x04 \x03(\x0b\x32\x1e.bosca.content.SignedUrlHeader\x12<\n\nattributes\x18\x05 \x03(\x0b\x32(.bosca.content.SignedUrl.AttributesEntry\x1a\x31\n\x0f\x41ttributesEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"S\n\x12\x41\x64\x64MetadataRequest\x12\x12\n\ncollection\x18\x01 \x01(\t\x12)\n\x08metadata\x18\x02 \x01(\x0b\x32\x17.bosca.content.Metadata\"6\n\tMetadatas\x12)\n\x08metadata\x18\x01 \x03(\x0b\x32\x17.bosca.content.Metadata\"@\n\x17\x41\x64\x64MetadataTraitRequest\x12\x13\n\x0bmetadata_id\x18\x01 \x01(\t\x12\x10\n\x08trait_id\x18\x02 \x01(\t*n\n\x12PermissionRelation\x12\x0b\n\x07viewers\x10\x00\x12\x0f\n\x0b\x64iscoverers\x10\x01\x12\x0b\n\x07\x65\x64itors\x10\x02\x12\x0c\n\x08managers\x10\x03\x12\x13\n\x0fserviceaccounts\x10\x04\x12\n\n\x06owners\x10\x05*U\n\x10PermissionAction\x12\x08\n\x04view\x10\x00\x12\x08\n\x04list\x10\x02\x12\x08\n\x04\x65\x64it\x10\x03\x12\n\n\x06manage\x10\x04\x12\x0b\n\x07service\x10\x05\x12\n\n\x06\x64\x65lete\x10\x06*\x9d\x01\n\x14PermissionObjectType\x12\x17\n\x13unknown_object_type\x10\x00\x12\x13\n\x0f\x63ollection_type\x10\x01\x12\x11\n\rmetadata_type\x10\x02\x12\x18\n\x14system_resource_type\x10\x03\x12\x11\n\rworkflow_type\x10\x04\x12\x17\n\x13workflow_state_type\x10\x05*[\n\x15PermissionSubjectType\x12\x18\n\x14unknown_subject_type\x10\x00\x12\x08\n\x04user\x10\x01\x12\t\n\x05group\x10\x02\x12\x13\n\x0fservice_account\x10\x03*K\n\x08ItemType\x12\x0b\n\x07unknown\x10\x00\x12\x0e\n\ncollection\x10\x01\x12\x0c\n\x08metadata\x10\x02\x12\x14\n\x10metadata_variant\x10\x03*4\n\x0e\x43ollectionType\x12\x0c\n\x08standard\x10\x00\x12\n\n\x06\x66older\x10\x01\x12\x08\n\x04root\x10\x02*\x8e\x01\n\x11WorkflowStateType\x12\x19\n\x15unknown_workflow_type\x10\x00\x12\x0e\n\nprocessing\x10\x01\x12\t\n\x05\x64raft\x10\x02\x12\x0b\n\x07pending\x10\x03\x12\x0c\n\x08\x61pproval\x10\x04\x12\x0c\n\x08\x61pproved\x10\x05\x12\r\n\tpublished\x10\x06\x12\x0b\n\x07\x66\x61ilure\x10\x07\x32\xb7\x1c\n\x0e\x43ontentService\x12`\n\x0fGetWorkflowById\x12\x10.bosca.IdRequest\x1a\x17.bosca.content.Workflow\"\"\x82\xd3\xe4\x93\x02\x1c\x12\x1a/v1/content/workflows/{id}\x12U\n\x0cGetWorkflows\x12\x0c.bosca.Empty\x1a\x18.bosca.content.Workflows\"\x1d\x82\xd3\xe4\x93\x02\x17\x12\x15/v1/content/workflows\x12o\n\x14GetWorkflowStateById\x12\x10.bosca.IdRequest\x1a\x1c.bosca.content.WorkflowState\"\'\x82\xd3\xe4\x93\x02!\x12\x1f/v1/content/workflowstates/{id}\x12\x64\n\x11GetWorkflowStates\x12\x0c.bosca.Empty\x1a\x1d.bosca.content.WorkflowStates\"\"\x82\xd3\xe4\x93\x02\x1c\x12\x1a/v1/content/workflowstates\x12W\n\x0cGetTraitById\x12\x10.bosca.IdRequest\x1a\x14.bosca.content.Trait\"\x1f\x82\xd3\xe4\x93\x02\x19\x12\x17/v1/content/traits/{id}\x12L\n\tGetTraits\x12\x0c.bosca.Empty\x1a\x15.bosca.content.Traits\"\x1a\x82\xd3\xe4\x93\x02\x14\x12\x12/v1/content/traits\x12g\n\x16GetRootCollectionItems\x12\x0c.bosca.Empty\x1a\x1e.bosca.content.CollectionItems\"\x1f\x82\xd3\xe4\x93\x02\x19\x12\x17/v1/content/collections\x12r\n\x12GetCollectionItems\x12\x10.bosca.IdRequest\x1a\x1e.bosca.content.CollectionItems\"*\x82\xd3\xe4\x93\x02$\x12\"/v1/content/collections/{id}/items\x12k\n\rAddCollection\x12#.bosca.content.AddCollectionRequest\x1a\x11.bosca.IdResponse\"\"\x82\xd3\xe4\x93\x02\x1c\"\x17/v1/content/collections:\x01*\x12\x62\n\rGetCollection\x12\x10.bosca.IdRequest\x1a\x19.bosca.content.Collection\"$\x82\xd3\xe4\x93\x02\x1e\x12\x1c/v1/content/collections/{id}\x12X\n\x10\x44\x65leteCollection\x12\x10.bosca.IdRequest\x1a\x0c.bosca.Empty\"$\x82\xd3\xe4\x93\x02\x1e*\x1c/v1/content/collections/{id}\x12z\n\x18GetCollectionPermissions\x12\x10.bosca.IdRequest\x1a\x1a.bosca.content.Permissions\"0\x82\xd3\xe4\x93\x02*\x12(/v1/content/collections/{id}/permissions\x12w\n\x17\x41\x64\x64\x43ollectionPermission\x12\x19.bosca.content.Permission\x1a\x0c.bosca.Empty\"3\x82\xd3\xe4\x93\x02-\"(/v1/content/collections/{id}/permissions:\x01*\x12\x86\x01\n\x0f\x43heckPermission\x12%.bosca.content.PermissionCheckRequest\x1a&.bosca.content.PermissionCheckResponse\"$\x82\xd3\xe4\x93\x02\x1e\x12\x1c/v1/content/permission/check\x12[\n\x0bGetMetadata\x12\x10.bosca.IdRequest\x1a\x17.bosca.content.Metadata\"!\x82\xd3\xe4\x93\x02\x1b\x12\x19/v1/content/metadata/{id}\x12Y\n\x0cGetMetadatas\x12\x11.bosca.IdsRequest\x1a\x18.bosca.content.Metadatas\"\x1c\x82\xd3\xe4\x93\x02\x16\x12\x14/v1/content/metadata\x12\x64\n\x0b\x41\x64\x64Metadata\x12!.bosca.content.AddMetadataRequest\x1a\x11.bosca.IdResponse\"\x1f\x82\xd3\xe4\x93\x02\x19\"\x14/v1/content/metadata:\x01*\x12\x91\x01\n\x10\x41\x64\x64MetadataTrait\x12&.bosca.content.AddMetadataTraitRequest\x1a\x17.bosca.content.Metadata\"<\x82\xd3\xe4\x93\x02\x36\"4/v1/content/metadata/{metadata_id}/traits/{trait_id}\x12S\n\x0e\x44\x65leteMetadata\x12\x10.bosca.IdRequest\x1a\x0c.bosca.Empty\"!\x82\xd3\xe4\x93\x02\x1b*\x19/v1/content/metadata/{id}\x12i\n\x14GetMetadataUploadUrl\x12\x10.bosca.IdRequest\x1a\x18.bosca.content.SignedUrl\"%\x82\xd3\xe4\x93\x02\x1f\"\x1d/v1/content/metadata/{id}/url\x12k\n\x16GetMetadataDownloadUrl\x12\x10.bosca.IdRequest\x1a\x18.bosca.content.SignedUrl\"%\x82\xd3\xe4\x93\x02\x1f\x12\x1d/v1/content/metadata/{id}/url\x12\x9f\x01\n\x18\x41\x64\x64MetadataSupplementary\x12&.bosca.content.AddSupplementaryRequest\x1a\x18.bosca.content.SignedUrl\"A\x82\xd3\xe4\x93\x02;\"9/v1/content/metadata/{id}/url/supplementary/{type}/upload\x12\xab\x01\n#GetMetadataSupplementaryDownloadUrl\x12%.bosca.content.SupplementaryIdRequest\x1a\x18.bosca.content.SignedUrl\"C\x82\xd3\xe4\x93\x02=\x12;/v1/content/metadata/{id}/url/supplementary/{type}/download\x12\x8e\x01\n\x1b\x44\x65leteMetadataSupplementary\x12%.bosca.content.SupplementaryIdRequest\x1a\x0c.bosca.Empty\":\x82\xd3\xe4\x93\x02\x34*2/v1/content/metadata/{id}/url/supplementary/{type}\x12\x61\n\x13SetMetadataUploaded\x12\x10.bosca.IdRequest\x1a\x0c.bosca.Empty\"*\x82\xd3\xe4\x93\x02$\"\"/v1/content/metadata/{id}/uploaded\x12u\n\x16GetMetadataPermissions\x12\x10.bosca.IdRequest\x1a\x1a.bosca.content.Permissions\"-\x82\xd3\xe4\x93\x02\'\x12%/v1/content/metadata/{id}/permissions\x12t\n\x16\x41\x64\x64MetadataPermissions\x12\x1a.bosca.content.Permissions\x1a\x0c.bosca.Empty\"0\x82\xd3\xe4\x93\x02*\"%/v1/content/metadata/{id}/permissions:\x01*\x12q\n\x15\x41\x64\x64MetadataPermission\x12\x19.bosca.content.Permission\x1a\x0c.bosca.Empty\"/\x82\xd3\xe4\x93\x02)\"$/v1/content/metadata/{id}/permission:\x01*\x12\x94\x01\n\x17\x42\x65ginTransitionWorkflow\x12(.bosca.content.TransitionWorkflowRequest\x1a\x0c.bosca.Empty\"A\x82\xd3\xe4\x93\x02;\"6/v1/content/metadata/{metadata_id}/workflow/{state_id}:\x01*\x12\xa8\x01\n\x1a\x43ompleteTransitionWorkflow\x12\x30.bosca.content.CompleteTransitionWorkflowRequest\x1a\x0c.bosca.Empty\"J\x82\xd3\xe4\x93\x02\x44\"?/v1/content/metadata/{metadata_id}/workflow/transition/complete:\x01*\x12\xa2\x01\n\x17\x41\x64\x64MetadataRelationship\x12-.bosca.content.AddMetadataRelationshipRequest\x1a\x0c.bosca.Empty\"J\x82\xd3\xe4\x93\x02\x44\"?/v1/content/metadata/{metadata_id1}/relationship/{metadata_id2}:\x01*B\x1fZ\x1d\x62osca.io/api/protobuf/contentb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'bosca.content.content_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z\035bosca.io/api/protobuf/content'
  _globals['_ITEM_ATTRIBUTESENTRY']._loaded_options = None
  _globals['_ITEM_ATTRIBUTESENTRY']._serialized_options = b'8\001'
  _globals['_WORKFLOW_CONFIGURATIONENTRY']._loaded_options = None
  _globals['_WORKFLOW_CONFIGURATIONENTRY']._serialized_options = b'8\001'
  _globals['_WORKFLOWSTATE_CONFIGURATIONENTRY']._loaded_options = None
  _globals['_WORKFLOWSTATE_CONFIGURATIONENTRY']._serialized_options = b'8\001'
  _globals['_COLLECTION_ATTRIBUTESENTRY']._loaded_options = None
  _globals['_COLLECTION_ATTRIBUTESENTRY']._serialized_options = b'8\001'
  _globals['_SIGNEDURL_ATTRIBUTESENTRY']._loaded_options = None
  _globals['_SIGNEDURL_ATTRIBUTESENTRY']._serialized_options = b'8\001'
  _globals['_CONTENTSERVICE'].methods_by_name['GetWorkflowById']._loaded_options = None
  _globals['_CONTENTSERVICE'].methods_by_name['GetWorkflowById']._serialized_options = b'\202\323\344\223\002\034\022\032/v1/content/workflows/{id}'
  _globals['_CONTENTSERVICE'].methods_by_name['GetWorkflows']._loaded_options = None
  _globals['_CONTENTSERVICE'].methods_by_name['GetWorkflows']._serialized_options = b'\202\323\344\223\002\027\022\025/v1/content/workflows'
  _globals['_CONTENTSERVICE'].methods_by_name['GetWorkflowStateById']._loaded_options = None
  _globals['_CONTENTSERVICE'].methods_by_name['GetWorkflowStateById']._serialized_options = b'\202\323\344\223\002!\022\037/v1/content/workflowstates/{id}'
  _globals['_CONTENTSERVICE'].methods_by_name['GetWorkflowStates']._loaded_options = None
  _globals['_CONTENTSERVICE'].methods_by_name['GetWorkflowStates']._serialized_options = b'\202\323\344\223\002\034\022\032/v1/content/workflowstates'
  _globals['_CONTENTSERVICE'].methods_by_name['GetTraitById']._loaded_options = None
  _globals['_CONTENTSERVICE'].methods_by_name['GetTraitById']._serialized_options = b'\202\323\344\223\002\031\022\027/v1/content/traits/{id}'
  _globals['_CONTENTSERVICE'].methods_by_name['GetTraits']._loaded_options = None
  _globals['_CONTENTSERVICE'].methods_by_name['GetTraits']._serialized_options = b'\202\323\344\223\002\024\022\022/v1/content/traits'
  _globals['_CONTENTSERVICE'].methods_by_name['GetRootCollectionItems']._loaded_options = None
  _globals['_CONTENTSERVICE'].methods_by_name['GetRootCollectionItems']._serialized_options = b'\202\323\344\223\002\031\022\027/v1/content/collections'
  _globals['_CONTENTSERVICE'].methods_by_name['GetCollectionItems']._loaded_options = None
  _globals['_CONTENTSERVICE'].methods_by_name['GetCollectionItems']._serialized_options = b'\202\323\344\223\002$\022\"/v1/content/collections/{id}/items'
  _globals['_CONTENTSERVICE'].methods_by_name['AddCollection']._loaded_options = None
  _globals['_CONTENTSERVICE'].methods_by_name['AddCollection']._serialized_options = b'\202\323\344\223\002\034\"\027/v1/content/collections:\001*'
  _globals['_CONTENTSERVICE'].methods_by_name['GetCollection']._loaded_options = None
  _globals['_CONTENTSERVICE'].methods_by_name['GetCollection']._serialized_options = b'\202\323\344\223\002\036\022\034/v1/content/collections/{id}'
  _globals['_CONTENTSERVICE'].methods_by_name['DeleteCollection']._loaded_options = None
  _globals['_CONTENTSERVICE'].methods_by_name['DeleteCollection']._serialized_options = b'\202\323\344\223\002\036*\034/v1/content/collections/{id}'
  _globals['_CONTENTSERVICE'].methods_by_name['GetCollectionPermissions']._loaded_options = None
  _globals['_CONTENTSERVICE'].methods_by_name['GetCollectionPermissions']._serialized_options = b'\202\323\344\223\002*\022(/v1/content/collections/{id}/permissions'
  _globals['_CONTENTSERVICE'].methods_by_name['AddCollectionPermission']._loaded_options = None
  _globals['_CONTENTSERVICE'].methods_by_name['AddCollectionPermission']._serialized_options = b'\202\323\344\223\002-\"(/v1/content/collections/{id}/permissions:\001*'
  _globals['_CONTENTSERVICE'].methods_by_name['CheckPermission']._loaded_options = None
  _globals['_CONTENTSERVICE'].methods_by_name['CheckPermission']._serialized_options = b'\202\323\344\223\002\036\022\034/v1/content/permission/check'
  _globals['_CONTENTSERVICE'].methods_by_name['GetMetadata']._loaded_options = None
  _globals['_CONTENTSERVICE'].methods_by_name['GetMetadata']._serialized_options = b'\202\323\344\223\002\033\022\031/v1/content/metadata/{id}'
  _globals['_CONTENTSERVICE'].methods_by_name['GetMetadatas']._loaded_options = None
  _globals['_CONTENTSERVICE'].methods_by_name['GetMetadatas']._serialized_options = b'\202\323\344\223\002\026\022\024/v1/content/metadata'
  _globals['_CONTENTSERVICE'].methods_by_name['AddMetadata']._loaded_options = None
  _globals['_CONTENTSERVICE'].methods_by_name['AddMetadata']._serialized_options = b'\202\323\344\223\002\031\"\024/v1/content/metadata:\001*'
  _globals['_CONTENTSERVICE'].methods_by_name['AddMetadataTrait']._loaded_options = None
  _globals['_CONTENTSERVICE'].methods_by_name['AddMetadataTrait']._serialized_options = b'\202\323\344\223\0026\"4/v1/content/metadata/{metadata_id}/traits/{trait_id}'
  _globals['_CONTENTSERVICE'].methods_by_name['DeleteMetadata']._loaded_options = None
  _globals['_CONTENTSERVICE'].methods_by_name['DeleteMetadata']._serialized_options = b'\202\323\344\223\002\033*\031/v1/content/metadata/{id}'
  _globals['_CONTENTSERVICE'].methods_by_name['GetMetadataUploadUrl']._loaded_options = None
  _globals['_CONTENTSERVICE'].methods_by_name['GetMetadataUploadUrl']._serialized_options = b'\202\323\344\223\002\037\"\035/v1/content/metadata/{id}/url'
  _globals['_CONTENTSERVICE'].methods_by_name['GetMetadataDownloadUrl']._loaded_options = None
  _globals['_CONTENTSERVICE'].methods_by_name['GetMetadataDownloadUrl']._serialized_options = b'\202\323\344\223\002\037\022\035/v1/content/metadata/{id}/url'
  _globals['_CONTENTSERVICE'].methods_by_name['AddMetadataSupplementary']._loaded_options = None
  _globals['_CONTENTSERVICE'].methods_by_name['AddMetadataSupplementary']._serialized_options = b'\202\323\344\223\002;\"9/v1/content/metadata/{id}/url/supplementary/{type}/upload'
  _globals['_CONTENTSERVICE'].methods_by_name['GetMetadataSupplementaryDownloadUrl']._loaded_options = None
  _globals['_CONTENTSERVICE'].methods_by_name['GetMetadataSupplementaryDownloadUrl']._serialized_options = b'\202\323\344\223\002=\022;/v1/content/metadata/{id}/url/supplementary/{type}/download'
  _globals['_CONTENTSERVICE'].methods_by_name['DeleteMetadataSupplementary']._loaded_options = None
  _globals['_CONTENTSERVICE'].methods_by_name['DeleteMetadataSupplementary']._serialized_options = b'\202\323\344\223\0024*2/v1/content/metadata/{id}/url/supplementary/{type}'
  _globals['_CONTENTSERVICE'].methods_by_name['SetMetadataUploaded']._loaded_options = None
  _globals['_CONTENTSERVICE'].methods_by_name['SetMetadataUploaded']._serialized_options = b'\202\323\344\223\002$\"\"/v1/content/metadata/{id}/uploaded'
  _globals['_CONTENTSERVICE'].methods_by_name['GetMetadataPermissions']._loaded_options = None
  _globals['_CONTENTSERVICE'].methods_by_name['GetMetadataPermissions']._serialized_options = b'\202\323\344\223\002\'\022%/v1/content/metadata/{id}/permissions'
  _globals['_CONTENTSERVICE'].methods_by_name['AddMetadataPermissions']._loaded_options = None
  _globals['_CONTENTSERVICE'].methods_by_name['AddMetadataPermissions']._serialized_options = b'\202\323\344\223\002*\"%/v1/content/metadata/{id}/permissions:\001*'
  _globals['_CONTENTSERVICE'].methods_by_name['AddMetadataPermission']._loaded_options = None
  _globals['_CONTENTSERVICE'].methods_by_name['AddMetadataPermission']._serialized_options = b'\202\323\344\223\002)\"$/v1/content/metadata/{id}/permission:\001*'
  _globals['_CONTENTSERVICE'].methods_by_name['BeginTransitionWorkflow']._loaded_options = None
  _globals['_CONTENTSERVICE'].methods_by_name['BeginTransitionWorkflow']._serialized_options = b'\202\323\344\223\002;\"6/v1/content/metadata/{metadata_id}/workflow/{state_id}:\001*'
  _globals['_CONTENTSERVICE'].methods_by_name['CompleteTransitionWorkflow']._loaded_options = None
  _globals['_CONTENTSERVICE'].methods_by_name['CompleteTransitionWorkflow']._serialized_options = b'\202\323\344\223\002D\"?/v1/content/metadata/{metadata_id}/workflow/transition/complete:\001*'
  _globals['_CONTENTSERVICE'].methods_by_name['AddMetadataRelationship']._loaded_options = None
  _globals['_CONTENTSERVICE'].methods_by_name['AddMetadataRelationship']._serialized_options = b'\202\323\344\223\002D\"?/v1/content/metadata/{metadata_id1}/relationship/{metadata_id2}:\001*'
  _globals['_PERMISSIONRELATION']._serialized_start=3468
  _globals['_PERMISSIONRELATION']._serialized_end=3578
  _globals['_PERMISSIONACTION']._serialized_start=3580
  _globals['_PERMISSIONACTION']._serialized_end=3665
  _globals['_PERMISSIONOBJECTTYPE']._serialized_start=3668
  _globals['_PERMISSIONOBJECTTYPE']._serialized_end=3825
  _globals['_PERMISSIONSUBJECTTYPE']._serialized_start=3827
  _globals['_PERMISSIONSUBJECTTYPE']._serialized_end=3918
  _globals['_ITEMTYPE']._serialized_start=3920
  _globals['_ITEMTYPE']._serialized_end=3995
  _globals['_COLLECTIONTYPE']._serialized_start=3997
  _globals['_COLLECTIONTYPE']._serialized_end=4049
  _globals['_WORKFLOWSTATETYPE']._serialized_start=4052
  _globals['_WORKFLOWSTATETYPE']._serialized_end=4194
  _globals['_TRANSITIONWORKFLOWREQUEST']._serialized_start=180
  _globals['_TRANSITIONWORKFLOWREQUEST']._serialized_end=277
  _globals['_COMPLETETRANSITIONWORKFLOWREQUEST']._serialized_start=279
  _globals['_COMPLETETRANSITIONWORKFLOWREQUEST']._serialized_end=368
  _globals['_ADDMETADATARELATIONSHIPREQUEST']._serialized_start=370
  _globals['_ADDMETADATARELATIONSHIPREQUEST']._serialized_end=468
  _globals['_ADDSUPPLEMENTARYREQUEST']._serialized_start=470
  _globals['_ADDSUPPLEMENTARYREQUEST']._serialized_end=581
  _globals['_SUPPLEMENTARYIDREQUEST']._serialized_start=583
  _globals['_SUPPLEMENTARYIDREQUEST']._serialized_end=633
  _globals['_PERMISSIONCHECKREQUEST']._serialized_start=636
  _globals['_PERMISSIONCHECKREQUEST']._serialized_end=860
  _globals['_PERMISSIONCHECKRESPONSE']._serialized_start=862
  _globals['_PERMISSIONCHECKRESPONSE']._serialized_end=904
  _globals['_PERMISSIONS']._serialized_start=906
  _globals['_PERMISSIONS']._serialized_end=979
  _globals['_PERMISSION']._serialized_start=982
  _globals['_PERMISSION']._serialized_end=1136
  _globals['_ITEM']._serialized_start=1139
  _globals['_ITEM']._serialized_end=1406
  _globals['_ITEM_ATTRIBUTESENTRY']._serialized_start=1357
  _globals['_ITEM_ATTRIBUTESENTRY']._serialized_end=1406
  _globals['_ADDCOLLECTIONREQUEST']._serialized_start=1408
  _globals['_ADDCOLLECTIONREQUEST']._serialized_end=1493
  _globals['_WORKFLOWS']._serialized_start=1495
  _globals['_WORKFLOWS']._serialized_end=1550
  _globals['_WORKFLOW']._serialized_start=1553
  _globals['_WORKFLOW']._serialized_end=1746
  _globals['_WORKFLOW_CONFIGURATIONENTRY']._serialized_start=1694
  _globals['_WORKFLOW_CONFIGURATIONENTRY']._serialized_end=1746
  _globals['_WORKFLOWSTATES']._serialized_start=1748
  _globals['_WORKFLOWSTATES']._serialized_end=1810
  _globals['_WORKFLOWSTATE']._serialized_start=1813
  _globals['_WORKFLOWSTATE']._serialized_end=2197
  _globals['_WORKFLOWSTATE_CONFIGURATIONENTRY']._serialized_start=1694
  _globals['_WORKFLOWSTATE_CONFIGURATIONENTRY']._serialized_end=1746
  _globals['_WORKFLOWSTATETRANSITIONS']._serialized_start=2199
  _globals['_WORKFLOWSTATETRANSITIONS']._serialized_end=2286
  _globals['_WORKFLOWSTATETRANSITION']._serialized_start=2288
  _globals['_WORKFLOWSTATETRANSITION']._serialized_end=2378
  _globals['_TRAITS']._serialized_start=2380
  _globals['_TRAITS']._serialized_end=2426
  _globals['_TRAIT']._serialized_start=2428
  _globals['_TRAIT']._serialized_end=2482
  _globals['_COLLECTION']._serialized_start=2485
  _globals['_COLLECTION']._serialized_end=2809
  _globals['_COLLECTION_ATTRIBUTESENTRY']._serialized_start=1357
  _globals['_COLLECTION_ATTRIBUTESENTRY']._serialized_end=1406
  _globals['_COLLECTIONITEMS']._serialized_start=2811
  _globals['_COLLECTIONITEMS']._serialized_end=2874
  _globals['_COLLECTIONITEM']._serialized_start=2876
  _globals['_COLLECTIONITEM']._serialized_end=2994
  _globals['_SIGNEDURLHEADER']._serialized_start=2996
  _globals['_SIGNEDURLHEADER']._serialized_end=3042
  _globals['_SIGNEDURL']._serialized_start=3045
  _globals['_SIGNEDURL']._serialized_end=3259
  _globals['_SIGNEDURL_ATTRIBUTESENTRY']._serialized_start=1357
  _globals['_SIGNEDURL_ATTRIBUTESENTRY']._serialized_end=1406
  _globals['_ADDMETADATAREQUEST']._serialized_start=3261
  _globals['_ADDMETADATAREQUEST']._serialized_end=3344
  _globals['_METADATAS']._serialized_start=3346
  _globals['_METADATAS']._serialized_end=3400
  _globals['_ADDMETADATATRAITREQUEST']._serialized_start=3402
  _globals['_ADDMETADATATRAITREQUEST']._serialized_end=3466
  _globals['_CONTENTSERVICE']._serialized_start=4197
  _globals['_CONTENTSERVICE']._serialized_end=7836
# @@protoc_insertion_point(module_scope)
