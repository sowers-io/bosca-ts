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
from bosca.content import collections_pb2 as bosca_dot_content_dot_collections__pb2
from bosca.content import metadata_pb2 as bosca_dot_content_dot_metadata__pb2
from bosca.content import model_pb2 as bosca_dot_content_dot_model__pb2
from bosca.content import workflows_pb2 as bosca_dot_content_dot_workflows__pb2
from bosca.content import storage_systems_pb2 as bosca_dot_content_dot_storage__systems__pb2
from bosca.content import traits_pb2 as bosca_dot_content_dot_traits__pb2
from bosca.content import prompts_pb2 as bosca_dot_content_dot_prompts__pb2
from bosca.content import permissions_pb2 as bosca_dot_content_dot_permissions__pb2
from bosca.content import url_pb2 as bosca_dot_content_dot_url__pb2
from bosca.content import sources_pb2 as bosca_dot_content_dot_sources__pb2
from bosca import empty_pb2 as bosca_dot_empty__pb2
from bosca import requests_pb2 as bosca_dot_requests__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1b\x62osca/content/content.proto\x12\rbosca.content\x1a\x1cgoogle/api/annotations.proto\x1a\x1f\x62osca/content/collections.proto\x1a\x1c\x62osca/content/metadata.proto\x1a\x19\x62osca/content/model.proto\x1a\x1d\x62osca/content/workflows.proto\x1a#bosca/content/storage_systems.proto\x1a\x1a\x62osca/content/traits.proto\x1a\x1b\x62osca/content/prompts.proto\x1a\x1f\x62osca/content/permissions.proto\x1a\x17\x62osca/content/url.proto\x1a\x1b\x62osca/content/sources.proto\x1a\x11\x62osca/empty.proto\x1a\x14\x62osca/requests.proto2\xa6+\n\x0e\x43ontentService\x12O\n\nGetSources\x12\x0c.bosca.Empty\x1a\x16.bosca.content.Sources\"\x1b\x82\xd3\xe4\x93\x02\x15\x12\x13/v1/content/sources\x12V\n\tGetSource\x12\x10.bosca.IdRequest\x1a\x15.bosca.content.Source\" \x82\xd3\xe4\x93\x02\x1a\x12\x18/v1/content/sources/{id}\x12U\n\x0cGetWorkflows\x12\x0c.bosca.Empty\x1a\x18.bosca.content.Workflows\"\x1d\x82\xd3\xe4\x93\x02\x17\x12\x15/v1/content/workflows\x12\\\n\x0bGetWorkflow\x12\x10.bosca.IdRequest\x1a\x17.bosca.content.Workflow\"\"\x82\xd3\xe4\x93\x02\x1c\x12\x1a/v1/content/workflows/{id}\x12L\n\tGetModels\x12\x0c.bosca.Empty\x1a\x15.bosca.content.Models\"\x1a\x82\xd3\xe4\x93\x02\x14\x12\x12/v1/content/models\x12S\n\x08GetModel\x12\x10.bosca.IdRequest\x1a\x14.bosca.content.Model\"\x1f\x82\xd3\xe4\x93\x02\x19\x12\x17/v1/content/models/{id}\x12O\n\nGetPrompts\x12\x0c.bosca.Empty\x1a\x16.bosca.content.Prompts\"\x1b\x82\xd3\xe4\x93\x02\x15\x12\x13/v1/content/prompts\x12V\n\tGetPrompt\x12\x10.bosca.IdRequest\x1a\x15.bosca.content.Prompt\" \x82\xd3\xe4\x93\x02\x1a\x12\x18/v1/content/prompts/{id}\x12\x64\n\x11GetStorageSystems\x12\x0c.bosca.Empty\x1a\x1d.bosca.content.StorageSystems\"\"\x82\xd3\xe4\x93\x02\x1c\x12\x1a/v1/content/storagesystems\x12k\n\x10GetStorageSystem\x12\x10.bosca.IdRequest\x1a\x1c.bosca.content.StorageSystem\"\'\x82\xd3\xe4\x93\x02!\x12\x1f/v1/content/storagesystems/{id}\x12~\n\x16GetStorageSystemModels\x12\x10.bosca.IdRequest\x1a\".bosca.content.StorageSystemModels\".\x82\xd3\xe4\x93\x02(\x12&/v1/content/storagesystems/{id}/models\x12k\n\x10GetWorkflowState\x12\x10.bosca.IdRequest\x1a\x1c.bosca.content.WorkflowState\"\'\x82\xd3\xe4\x93\x02!\x12\x1f/v1/content/workflowstates/{id}\x12\x64\n\x11GetWorkflowStates\x12\x0c.bosca.Empty\x1a\x1d.bosca.content.WorkflowStates\"\"\x82\xd3\xe4\x93\x02\x1c\x12\x1a/v1/content/workflowstates\x12L\n\tGetTraits\x12\x0c.bosca.Empty\x1a\x15.bosca.content.Traits\"\x1a\x82\xd3\xe4\x93\x02\x14\x12\x12/v1/content/traits\x12S\n\x08GetTrait\x12\x10.bosca.IdRequest\x1a\x14.bosca.content.Trait\"\x1f\x82\xd3\xe4\x93\x02\x19\x12\x17/v1/content/traits/{id}\x12\x93\x01\n\x1cGetWorkflowActivityInstances\x12\x10.bosca.IdRequest\x1a(.bosca.content.WorkflowActivityInstances\"7\x82\xd3\xe4\x93\x02\x31\x12//v1/content/workflows/{id}/activities/instances\x12\xd1\x01\n!GetWorkflowActivityStorageSystems\x12(.bosca.content.WorkflowActivityIdRequest\x1a-.bosca.content.WorkflowActivityStorageSystems\"S\x82\xd3\xe4\x93\x02M\x12K/v1/content/workflows/{workflow_id}/activities/{activity_id}/storagesystems\x12\xbc\x01\n\x1aGetWorkflowActivityPrompts\x12(.bosca.content.WorkflowActivityIdRequest\x1a&.bosca.content.WorkflowActivityPrompts\"L\x82\xd3\xe4\x93\x02\x46\x12\x44/v1/content/workflows/{workflow_id}/activities/{activity_id}/prompts\x12g\n\x16GetRootCollectionItems\x12\x0c.bosca.Empty\x1a\x1e.bosca.content.CollectionItems\"\x1f\x82\xd3\xe4\x93\x02\x19\x12\x17/v1/content/collections\x12r\n\x12GetCollectionItems\x12\x10.bosca.IdRequest\x1a\x1e.bosca.content.CollectionItems\"*\x82\xd3\xe4\x93\x02$\x12\"/v1/content/collections/{id}/items\x12k\n\rAddCollection\x12#.bosca.content.AddCollectionRequest\x1a\x11.bosca.IdResponse\"\"\x82\xd3\xe4\x93\x02\x1c\"\x17/v1/content/collections:\x01*\x12n\n\x0e\x41\x64\x64\x43ollections\x12$.bosca.content.AddCollectionsRequest\x1a\x12.bosca.IdResponses\"\"\x82\xd3\xe4\x93\x02\x1c\x1a\x17/v1/content/collections:\x01*\x12\x62\n\rGetCollection\x12\x10.bosca.IdRequest\x1a\x19.bosca.content.Collection\"$\x82\xd3\xe4\x93\x02\x1e\x12\x1c/v1/content/collections/{id}\x12X\n\x10\x44\x65leteCollection\x12\x10.bosca.IdRequest\x1a\x0c.bosca.Empty\"$\x82\xd3\xe4\x93\x02\x1e*\x1c/v1/content/collections/{id}\x12z\n\x18GetCollectionPermissions\x12\x10.bosca.IdRequest\x1a\x1a.bosca.content.Permissions\"0\x82\xd3\xe4\x93\x02*\x12(/v1/content/collections/{id}/permissions\x12w\n\x17\x41\x64\x64\x43ollectionPermission\x12\x19.bosca.content.Permission\x1a\x0c.bosca.Empty\"3\x82\xd3\xe4\x93\x02-\"(/v1/content/collections/{id}/permissions:\x01*\x12\x83\x01\n\x11\x41\x64\x64\x43ollectionItem\x12\'.bosca.content.AddCollectionItemRequest\x1a\x0c.bosca.Empty\"7\x82\xd3\xe4\x93\x02\x31\",/v1/content/collections/{collection_id}/item:\x01*\x12\x86\x01\n\x0f\x43heckPermission\x12%.bosca.content.PermissionCheckRequest\x1a&.bosca.content.PermissionCheckResponse\"$\x82\xd3\xe4\x93\x02\x1e\x12\x1c/v1/content/permission/check\x12[\n\x0bGetMetadata\x12\x10.bosca.IdRequest\x1a\x17.bosca.content.Metadata\"!\x82\xd3\xe4\x93\x02\x1b\x12\x19/v1/content/metadata/{id}\x12Y\n\x0cGetMetadatas\x12\x11.bosca.IdsRequest\x1a\x18.bosca.content.Metadatas\"\x1c\x82\xd3\xe4\x93\x02\x16\x12\x14/v1/content/metadata\x12\x64\n\x0b\x41\x64\x64Metadata\x12!.bosca.content.AddMetadataRequest\x1a\x11.bosca.IdResponse\"\x1f\x82\xd3\xe4\x93\x02\x19\"\x14/v1/content/metadata:\x01*\x12h\n\x0c\x41\x64\x64Metadatas\x12\".bosca.content.AddMetadatasRequest\x1a\x12.bosca.IdResponses\" \x82\xd3\xe4\x93\x02\x1a\"\x15/v1/content/metadatas:\x01*\x12\x91\x01\n\x10\x41\x64\x64MetadataTrait\x12&.bosca.content.AddMetadataTraitRequest\x1a\x17.bosca.content.Metadata\"<\x82\xd3\xe4\x93\x02\x36\"4/v1/content/metadata/{metadata_id}/traits/{trait_id}\x12S\n\x0e\x44\x65leteMetadata\x12\x10.bosca.IdRequest\x1a\x0c.bosca.Empty\"!\x82\xd3\xe4\x93\x02\x1b*\x19/v1/content/metadata/{id}\x12i\n\x14GetMetadataUploadUrl\x12\x10.bosca.IdRequest\x1a\x18.bosca.content.SignedUrl\"%\x82\xd3\xe4\x93\x02\x1f\"\x1d/v1/content/metadata/{id}/url\x12k\n\x16GetMetadataDownloadUrl\x12\x10.bosca.IdRequest\x1a\x18.bosca.content.SignedUrl\"%\x82\xd3\xe4\x93\x02\x1f\x12\x1d/v1/content/metadata/{id}/url\x12\x9f\x01\n\x18\x41\x64\x64MetadataSupplementary\x12&.bosca.content.AddSupplementaryRequest\x1a\x18.bosca.content.SignedUrl\"A\x82\xd3\xe4\x93\x02;\"9/v1/content/metadata/{id}/url/supplementary/{type}/upload\x12\xab\x01\n#GetMetadataSupplementaryDownloadUrl\x12%.bosca.content.SupplementaryIdRequest\x1a\x18.bosca.content.SignedUrl\"C\x82\xd3\xe4\x93\x02=\x12;/v1/content/metadata/{id}/url/supplementary/{type}/download\x12\x8e\x01\n\x1b\x44\x65leteMetadataSupplementary\x12%.bosca.content.SupplementaryIdRequest\x1a\x0c.bosca.Empty\":\x82\xd3\xe4\x93\x02\x34*2/v1/content/metadata/{id}/url/supplementary/{type}\x12\x61\n\x13SetMetadataUploaded\x12\x10.bosca.IdRequest\x1a\x0c.bosca.Empty\"*\x82\xd3\xe4\x93\x02$\"\"/v1/content/metadata/{id}/uploaded\x12u\n\x16GetMetadataPermissions\x12\x10.bosca.IdRequest\x1a\x1a.bosca.content.Permissions\"-\x82\xd3\xe4\x93\x02\'\x12%/v1/content/metadata/{id}/permissions\x12t\n\x16\x41\x64\x64MetadataPermissions\x12\x1a.bosca.content.Permissions\x1a\x0c.bosca.Empty\"0\x82\xd3\xe4\x93\x02*\"%/v1/content/metadata/{id}/permissions:\x01*\x12q\n\x15\x41\x64\x64MetadataPermission\x12\x19.bosca.content.Permission\x1a\x0c.bosca.Empty\"/\x82\xd3\xe4\x93\x02)\"$/v1/content/metadata/{id}/permission:\x01*\x12\x94\x01\n\x17\x42\x65ginTransitionWorkflow\x12(.bosca.content.TransitionWorkflowRequest\x1a\x0c.bosca.Empty\"A\x82\xd3\xe4\x93\x02;\"6/v1/content/metadata/{metadata_id}/workflow/{state_id}:\x01*\x12\xa8\x01\n\x1a\x43ompleteTransitionWorkflow\x12\x30.bosca.content.CompleteTransitionWorkflowRequest\x1a\x0c.bosca.Empty\"J\x82\xd3\xe4\x93\x02\x44\"?/v1/content/metadata/{metadata_id}/workflow/transition/complete:\x01*\x12\x98\x01\n\x17\x41\x64\x64MetadataRelationship\x12#.bosca.content.MetadataRelationship\x1a\x0c.bosca.Empty\"J\x82\xd3\xe4\x93\x02\x44\"?/v1/content/metadata/{metadata_id1}/relationship/{metadata_id2}:\x01*\x12\xb1\x01\n\x18GetMetadataRelationships\x12,.bosca.content.MetadataRelationshipIdRequest\x1a$.bosca.content.MetadataRelationships\"A\x82\xd3\xe4\x93\x02;\"6/v1/content/metadata/{id}/relationships/{relationship}:\x01*B%Z#bosca.io/api/protobuf/bosca/contentb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'bosca.content.content_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z#bosca.io/api/protobuf/bosca/content'
  _globals['_CONTENTSERVICE'].methods_by_name['GetSources']._loaded_options = None
  _globals['_CONTENTSERVICE'].methods_by_name['GetSources']._serialized_options = b'\202\323\344\223\002\025\022\023/v1/content/sources'
  _globals['_CONTENTSERVICE'].methods_by_name['GetSource']._loaded_options = None
  _globals['_CONTENTSERVICE'].methods_by_name['GetSource']._serialized_options = b'\202\323\344\223\002\032\022\030/v1/content/sources/{id}'
  _globals['_CONTENTSERVICE'].methods_by_name['GetWorkflows']._loaded_options = None
  _globals['_CONTENTSERVICE'].methods_by_name['GetWorkflows']._serialized_options = b'\202\323\344\223\002\027\022\025/v1/content/workflows'
  _globals['_CONTENTSERVICE'].methods_by_name['GetWorkflow']._loaded_options = None
  _globals['_CONTENTSERVICE'].methods_by_name['GetWorkflow']._serialized_options = b'\202\323\344\223\002\034\022\032/v1/content/workflows/{id}'
  _globals['_CONTENTSERVICE'].methods_by_name['GetModels']._loaded_options = None
  _globals['_CONTENTSERVICE'].methods_by_name['GetModels']._serialized_options = b'\202\323\344\223\002\024\022\022/v1/content/models'
  _globals['_CONTENTSERVICE'].methods_by_name['GetModel']._loaded_options = None
  _globals['_CONTENTSERVICE'].methods_by_name['GetModel']._serialized_options = b'\202\323\344\223\002\031\022\027/v1/content/models/{id}'
  _globals['_CONTENTSERVICE'].methods_by_name['GetPrompts']._loaded_options = None
  _globals['_CONTENTSERVICE'].methods_by_name['GetPrompts']._serialized_options = b'\202\323\344\223\002\025\022\023/v1/content/prompts'
  _globals['_CONTENTSERVICE'].methods_by_name['GetPrompt']._loaded_options = None
  _globals['_CONTENTSERVICE'].methods_by_name['GetPrompt']._serialized_options = b'\202\323\344\223\002\032\022\030/v1/content/prompts/{id}'
  _globals['_CONTENTSERVICE'].methods_by_name['GetStorageSystems']._loaded_options = None
  _globals['_CONTENTSERVICE'].methods_by_name['GetStorageSystems']._serialized_options = b'\202\323\344\223\002\034\022\032/v1/content/storagesystems'
  _globals['_CONTENTSERVICE'].methods_by_name['GetStorageSystem']._loaded_options = None
  _globals['_CONTENTSERVICE'].methods_by_name['GetStorageSystem']._serialized_options = b'\202\323\344\223\002!\022\037/v1/content/storagesystems/{id}'
  _globals['_CONTENTSERVICE'].methods_by_name['GetStorageSystemModels']._loaded_options = None
  _globals['_CONTENTSERVICE'].methods_by_name['GetStorageSystemModels']._serialized_options = b'\202\323\344\223\002(\022&/v1/content/storagesystems/{id}/models'
  _globals['_CONTENTSERVICE'].methods_by_name['GetWorkflowState']._loaded_options = None
  _globals['_CONTENTSERVICE'].methods_by_name['GetWorkflowState']._serialized_options = b'\202\323\344\223\002!\022\037/v1/content/workflowstates/{id}'
  _globals['_CONTENTSERVICE'].methods_by_name['GetWorkflowStates']._loaded_options = None
  _globals['_CONTENTSERVICE'].methods_by_name['GetWorkflowStates']._serialized_options = b'\202\323\344\223\002\034\022\032/v1/content/workflowstates'
  _globals['_CONTENTSERVICE'].methods_by_name['GetTraits']._loaded_options = None
  _globals['_CONTENTSERVICE'].methods_by_name['GetTraits']._serialized_options = b'\202\323\344\223\002\024\022\022/v1/content/traits'
  _globals['_CONTENTSERVICE'].methods_by_name['GetTrait']._loaded_options = None
  _globals['_CONTENTSERVICE'].methods_by_name['GetTrait']._serialized_options = b'\202\323\344\223\002\031\022\027/v1/content/traits/{id}'
  _globals['_CONTENTSERVICE'].methods_by_name['GetWorkflowActivityInstances']._loaded_options = None
  _globals['_CONTENTSERVICE'].methods_by_name['GetWorkflowActivityInstances']._serialized_options = b'\202\323\344\223\0021\022//v1/content/workflows/{id}/activities/instances'
  _globals['_CONTENTSERVICE'].methods_by_name['GetWorkflowActivityStorageSystems']._loaded_options = None
  _globals['_CONTENTSERVICE'].methods_by_name['GetWorkflowActivityStorageSystems']._serialized_options = b'\202\323\344\223\002M\022K/v1/content/workflows/{workflow_id}/activities/{activity_id}/storagesystems'
  _globals['_CONTENTSERVICE'].methods_by_name['GetWorkflowActivityPrompts']._loaded_options = None
  _globals['_CONTENTSERVICE'].methods_by_name['GetWorkflowActivityPrompts']._serialized_options = b'\202\323\344\223\002F\022D/v1/content/workflows/{workflow_id}/activities/{activity_id}/prompts'
  _globals['_CONTENTSERVICE'].methods_by_name['GetRootCollectionItems']._loaded_options = None
  _globals['_CONTENTSERVICE'].methods_by_name['GetRootCollectionItems']._serialized_options = b'\202\323\344\223\002\031\022\027/v1/content/collections'
  _globals['_CONTENTSERVICE'].methods_by_name['GetCollectionItems']._loaded_options = None
  _globals['_CONTENTSERVICE'].methods_by_name['GetCollectionItems']._serialized_options = b'\202\323\344\223\002$\022\"/v1/content/collections/{id}/items'
  _globals['_CONTENTSERVICE'].methods_by_name['AddCollection']._loaded_options = None
  _globals['_CONTENTSERVICE'].methods_by_name['AddCollection']._serialized_options = b'\202\323\344\223\002\034\"\027/v1/content/collections:\001*'
  _globals['_CONTENTSERVICE'].methods_by_name['AddCollections']._loaded_options = None
  _globals['_CONTENTSERVICE'].methods_by_name['AddCollections']._serialized_options = b'\202\323\344\223\002\034\032\027/v1/content/collections:\001*'
  _globals['_CONTENTSERVICE'].methods_by_name['GetCollection']._loaded_options = None
  _globals['_CONTENTSERVICE'].methods_by_name['GetCollection']._serialized_options = b'\202\323\344\223\002\036\022\034/v1/content/collections/{id}'
  _globals['_CONTENTSERVICE'].methods_by_name['DeleteCollection']._loaded_options = None
  _globals['_CONTENTSERVICE'].methods_by_name['DeleteCollection']._serialized_options = b'\202\323\344\223\002\036*\034/v1/content/collections/{id}'
  _globals['_CONTENTSERVICE'].methods_by_name['GetCollectionPermissions']._loaded_options = None
  _globals['_CONTENTSERVICE'].methods_by_name['GetCollectionPermissions']._serialized_options = b'\202\323\344\223\002*\022(/v1/content/collections/{id}/permissions'
  _globals['_CONTENTSERVICE'].methods_by_name['AddCollectionPermission']._loaded_options = None
  _globals['_CONTENTSERVICE'].methods_by_name['AddCollectionPermission']._serialized_options = b'\202\323\344\223\002-\"(/v1/content/collections/{id}/permissions:\001*'
  _globals['_CONTENTSERVICE'].methods_by_name['AddCollectionItem']._loaded_options = None
  _globals['_CONTENTSERVICE'].methods_by_name['AddCollectionItem']._serialized_options = b'\202\323\344\223\0021\",/v1/content/collections/{collection_id}/item:\001*'
  _globals['_CONTENTSERVICE'].methods_by_name['CheckPermission']._loaded_options = None
  _globals['_CONTENTSERVICE'].methods_by_name['CheckPermission']._serialized_options = b'\202\323\344\223\002\036\022\034/v1/content/permission/check'
  _globals['_CONTENTSERVICE'].methods_by_name['GetMetadata']._loaded_options = None
  _globals['_CONTENTSERVICE'].methods_by_name['GetMetadata']._serialized_options = b'\202\323\344\223\002\033\022\031/v1/content/metadata/{id}'
  _globals['_CONTENTSERVICE'].methods_by_name['GetMetadatas']._loaded_options = None
  _globals['_CONTENTSERVICE'].methods_by_name['GetMetadatas']._serialized_options = b'\202\323\344\223\002\026\022\024/v1/content/metadata'
  _globals['_CONTENTSERVICE'].methods_by_name['AddMetadata']._loaded_options = None
  _globals['_CONTENTSERVICE'].methods_by_name['AddMetadata']._serialized_options = b'\202\323\344\223\002\031\"\024/v1/content/metadata:\001*'
  _globals['_CONTENTSERVICE'].methods_by_name['AddMetadatas']._loaded_options = None
  _globals['_CONTENTSERVICE'].methods_by_name['AddMetadatas']._serialized_options = b'\202\323\344\223\002\032\"\025/v1/content/metadatas:\001*'
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
  _globals['_CONTENTSERVICE'].methods_by_name['GetMetadataRelationships']._loaded_options = None
  _globals['_CONTENTSERVICE'].methods_by_name['GetMetadataRelationships']._serialized_options = b'\202\323\344\223\002;\"6/v1/content/metadata/{id}/relationships/{relationship}:\001*'
  _globals['_CONTENTSERVICE']._serialized_start=420
  _globals['_CONTENTSERVICE']._serialized_end=5962
# @@protoc_insertion_point(module_scope)
