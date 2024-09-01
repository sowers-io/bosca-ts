import { GraphQLResolveInfo, GraphQLScalarType, GraphQLScalarTypeConfig } from 'graphql';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
export type MakeEmpty<T extends { [key: string]: unknown }, K extends keyof T> = { [_ in K]?: never };
export type Incremental<T> = T | { [P in keyof T]?: P extends ' $fragmentName' | '__typename' ? T[P] : never };
export type Omit<T, K extends keyof T> = Pick<T, Exclude<keyof T, K>>;
export type RequireFields<T, K extends keyof T> = Omit<T, K> & { [P in K]-?: NonNullable<T[P]> };
/** All built-in and custom scalars, mapped to their actual values */
export interface Scalars {
  ID: { input: string; output: string; }
  String: { input: string; output: string; }
  Boolean: { input: boolean; output: boolean; }
  Int: { input: number; output: number; }
  Float: { input: number; output: number; }
  Date: { input: any; output: any; }
  JSON: { input: any; output: any; }
  JSONObject: { input: any; output: any; }
}

export interface Collection {
  __typename?: 'Collection';
  attributes: Array<Kv>;
  categoryIds: Array<Scalars['String']['output']>;
  created: Scalars['Date']['output'];
  id: Scalars['ID']['output'];
  items?: Maybe<Array<Maybe<CollectionItem>>>;
  labels: Array<Scalars['String']['output']>;
  modified: Scalars['Date']['output'];
  name: Scalars['String']['output'];
  permissions: Array<Permission>;
  traitIds: Array<Scalars['String']['output']>;
  type: CollectionType;
}


export interface CollectionItemsArgs {
  filter?: InputMaybe<CollectionItemFilter>;
}

export interface CollectionInput {
  name: Scalars['String']['input'];
}

export type CollectionItem = Collection | Metadata;

export interface CollectionItemFilter {
  created?: InputMaybe<Scalars['Date']['input']>;
}

export enum CollectionType {
  Folder = 'folder',
  Root = 'root',
  Standard = 'standard'
}

export interface Find {
  __typename?: 'Find';
  collections: Array<Collection>;
  metadata: Array<Metadata>;
}


export interface FindCollectionsArgs {
  query?: InputMaybe<FindQuery>;
}


export interface FindMetadataArgs {
  query?: InputMaybe<FindQuery>;
}

export interface FindInputAttribute {
  name: Scalars['String']['input'];
  value?: InputMaybe<Scalars['String']['input']>;
}

export interface FindQuery {
  attributes: Array<FindInputAttribute>;
}

export interface Kv {
  __typename?: 'KV';
  key: Scalars['String']['output'];
  value: Scalars['String']['output'];
}

export interface KvInput {
  key: Scalars['String']['input'];
  value: Scalars['String']['input'];
}

export interface Metadata {
  __typename?: 'Metadata';
  attributes: Array<Kv>;
  content?: Maybe<MetadataContent>;
  contentLength?: Maybe<Scalars['Int']['output']>;
  contentType: Scalars['String']['output'];
  created: Scalars['Date']['output'];
  downloadUrl: SignedUrl;
  id: Scalars['ID']['output'];
  labels: Array<Scalars['String']['output']>;
  languageTag: Scalars['String']['output'];
  modified: Scalars['Date']['output'];
  name: Scalars['String']['output'];
  parentId?: Maybe<Scalars['ID']['output']>;
  permissions: Array<Permission>;
  sourceId?: Maybe<Scalars['String']['output']>;
  sourceIdentifier?: Maybe<Scalars['String']['output']>;
  supplementaries: Array<Supplementary>;
  supplementary?: Maybe<Supplementary>;
  traitIds: Array<Scalars['String']['output']>;
  uploadUrl: SignedUrl;
  workflowState: MetadataWorkflowState;
}


export interface MetadataSupplementariesArgs {
  key?: InputMaybe<Array<Scalars['String']['input']>>;
}


export interface MetadataSupplementaryArgs {
  key: Scalars['String']['input'];
}

export interface MetadataContent {
  __typename?: 'MetadataContent';
  json?: Maybe<Scalars['JSONObject']['output']>;
  text?: Maybe<Scalars['String']['output']>;
}

export interface MetadataInput {
  contentLength?: InputMaybe<Scalars['Int']['input']>;
  contentType: Scalars['String']['input'];
  languageTag: Scalars['String']['input'];
  name: Scalars['String']['input'];
  traitIds?: InputMaybe<Array<Scalars['String']['input']>>;
}

export interface MetadataWorkflowState {
  __typename?: 'MetadataWorkflowState';
  deleteWorkflowId?: Maybe<Scalars['String']['output']>;
  id: Scalars['String']['output'];
  pendingId?: Maybe<Scalars['String']['output']>;
}

export interface Model {
  __typename?: 'Model';
  configuration: Array<Kv>;
  description?: Maybe<Scalars['String']['output']>;
  id: Scalars['ID']['output'];
  name?: Maybe<Scalars['String']['output']>;
  type?: Maybe<Scalars['String']['output']>;
}

export interface ModelInput {
  configuration: Array<KvInput>;
  description?: InputMaybe<Scalars['String']['input']>;
  name?: InputMaybe<Scalars['String']['input']>;
  type?: InputMaybe<Scalars['String']['input']>;
}

export interface Mutation {
  __typename?: 'Mutation';
  addCollection?: Maybe<Collection>;
  addCollectionPermissions?: Maybe<Collection>;
  addMetadata?: Maybe<Metadata>;
  addMetadataPermissions?: Maybe<Metadata>;
  addModel: Model;
  addPrompt: Prompt;
  addSource: Source;
  addStorageSystem: StorageSystem;
  addWorkflow: Workflow;
  addWorkflowActivity: WorkflowActivity;
  addWorkflowState: WorkflowState;
  addWorkflowStateTransition: WorkflowStateTransition;
  deleteCollection?: Maybe<Scalars['Boolean']['output']>;
  deleteCollectionPermissions?: Maybe<Collection>;
  deleteMetadata?: Maybe<Scalars['Boolean']['output']>;
  deleteMetadataPermissions?: Maybe<Metadata>;
  deleteModel?: Maybe<Scalars['Boolean']['output']>;
  deletePrompt: Scalars['Boolean']['output'];
  deleteSource: Scalars['Boolean']['output'];
  deleteStorageSystem?: Maybe<Scalars['Boolean']['output']>;
  deleteWorkflow: Scalars['Boolean']['output'];
  deleteWorkflowActivity: Scalars['Boolean']['output'];
  deleteWorkflowTransition?: Maybe<Scalars['Boolean']['output']>;
  editModel: Model;
  editPrompt: Prompt;
  editSource: Source;
  editStorageSystem: StorageSystem;
  editWorkflow: Workflow;
  editWorkflowActivity: WorkflowActivity;
  executeWorkflow: Scalars['String']['output'];
  login?: Maybe<Scalars['String']['output']>;
  setMetadataJSONContent?: Maybe<Metadata>;
  setMetadataReady?: Maybe<Metadata>;
  setMetadataTextContent?: Maybe<Metadata>;
  setPassword?: Maybe<Scalars['Boolean']['output']>;
  signup?: Maybe<Scalars['String']['output']>;
}


export interface MutationAddCollectionArgs {
  collection: CollectionInput;
  parent?: InputMaybe<Scalars['String']['input']>;
}


export interface MutationAddCollectionPermissionsArgs {
  id: Scalars['String']['input'];
  permissions: Array<PermissionInput>;
}


export interface MutationAddMetadataArgs {
  metadata: MetadataInput;
  parent?: InputMaybe<Scalars['String']['input']>;
}


export interface MutationAddMetadataPermissionsArgs {
  id: Scalars['String']['input'];
  permissions: Array<PermissionInput>;
}


export interface MutationAddModelArgs {
  model: ModelInput;
}


export interface MutationAddPromptArgs {
  prompt: PromptInput;
}


export interface MutationAddSourceArgs {
  source: SourceInput;
}


export interface MutationAddStorageSystemArgs {
  storageSystem: StorageSystemInput;
}


export interface MutationAddWorkflowArgs {
  workflow?: InputMaybe<WorkflowInput>;
}


export interface MutationAddWorkflowActivityArgs {
  activity?: InputMaybe<WorkflowActivityInput>;
}


export interface MutationAddWorkflowStateArgs {
  workflow?: InputMaybe<WorkflowStateInput>;
}


export interface MutationAddWorkflowStateTransitionArgs {
  workflow: WorkflowStateTransitionInput;
}


export interface MutationDeleteCollectionArgs {
  id: Scalars['String']['input'];
}


export interface MutationDeleteCollectionPermissionsArgs {
  id: Scalars['String']['input'];
  permissions: Array<PermissionInput>;
}


export interface MutationDeleteMetadataArgs {
  id: Scalars['String']['input'];
}


export interface MutationDeleteMetadataPermissionsArgs {
  id: Scalars['String']['input'];
  permissions: Array<PermissionInput>;
}


export interface MutationDeleteModelArgs {
  id: Scalars['ID']['input'];
}


export interface MutationDeletePromptArgs {
  id: Scalars['ID']['input'];
}


export interface MutationDeleteSourceArgs {
  id: Scalars['ID']['input'];
}


export interface MutationDeleteStorageSystemArgs {
  id: Scalars['ID']['input'];
}


export interface MutationDeleteWorkflowArgs {
  id: Scalars['String']['input'];
}


export interface MutationDeleteWorkflowActivityArgs {
  id: Scalars['ID']['input'];
}


export interface MutationDeleteWorkflowTransitionArgs {
  fromStateId: Scalars['String']['input'];
  toStateId: Scalars['String']['input'];
}


export interface MutationEditModelArgs {
  id: Scalars['ID']['input'];
  model?: InputMaybe<ModelInput>;
}


export interface MutationEditPromptArgs {
  id: Scalars['ID']['input'];
  prompt: PromptInput;
}


export interface MutationEditSourceArgs {
  id: Scalars['ID']['input'];
  source: SourceInput;
}


export interface MutationEditStorageSystemArgs {
  id: Scalars['ID']['input'];
  model?: InputMaybe<StorageSystemInput>;
}


export interface MutationEditWorkflowArgs {
  workflow?: InputMaybe<WorkflowInput>;
}


export interface MutationEditWorkflowActivityArgs {
  activity?: InputMaybe<WorkflowActivityInput>;
  id: Scalars['ID']['input'];
}


export interface MutationExecuteWorkflowArgs {
  request?: InputMaybe<WorkflowExecutionInput>;
}


export interface MutationLoginArgs {
  password: Scalars['String']['input'];
  username: Scalars['String']['input'];
}


export interface MutationSetMetadataJsonContentArgs {
  id: Scalars['String']['input'];
  json?: InputMaybe<Scalars['JSONObject']['input']>;
}


export interface MutationSetMetadataReadyArgs {
  id: Scalars['String']['input'];
}


export interface MutationSetMetadataTextContentArgs {
  id: Scalars['String']['input'];
  text?: InputMaybe<Scalars['String']['input']>;
}


export interface MutationSetPasswordArgs {
  password: Scalars['String']['input'];
}


export interface MutationSignupArgs {
  email: Scalars['String']['input'];
  firstName: Scalars['String']['input'];
  lastName: Scalars['String']['input'];
  password: Scalars['String']['input'];
}

export interface Permission {
  __typename?: 'Permission';
  relation: PermissionRelation;
  subject: Scalars['String']['output'];
  subjectType: PermissionSubjectType;
}

export interface PermissionInput {
  relation: PermissionRelation;
  subject: Scalars['String']['input'];
  subjectType: PermissionSubjectType;
}

export enum PermissionRelation {
  Discoverers = 'discoverers',
  Editors = 'editors',
  Managers = 'managers',
  Owners = 'owners',
  Serviceaccounts = 'serviceaccounts',
  Viewers = 'viewers'
}

export enum PermissionSubjectType {
  Group = 'group',
  Serviceaccount = 'serviceaccount',
  User = 'user'
}

export interface Prompt {
  __typename?: 'Prompt';
  description?: Maybe<Scalars['String']['output']>;
  id?: Maybe<Scalars['String']['output']>;
  inputType?: Maybe<Scalars['String']['output']>;
  name?: Maybe<Scalars['String']['output']>;
  outputType?: Maybe<Scalars['String']['output']>;
  systemPrompt?: Maybe<Scalars['String']['output']>;
  userPrompt?: Maybe<Scalars['String']['output']>;
}

export interface PromptInput {
  description?: InputMaybe<Scalars['String']['input']>;
  inputType?: InputMaybe<Scalars['String']['input']>;
  name?: InputMaybe<Scalars['String']['input']>;
  outputType?: InputMaybe<Scalars['String']['input']>;
  systemPrompt?: InputMaybe<Scalars['String']['input']>;
  userPrompt?: InputMaybe<Scalars['String']['input']>;
}

export interface Query {
  __typename?: 'Query';
  collection?: Maybe<Collection>;
  find?: Maybe<Find>;
  metadata?: Maybe<Metadata>;
  source?: Maybe<Source>;
  sources: Array<Source>;
  trait?: Maybe<Trait>;
  traits: Array<Trait>;
  workflows: Workflows;
}


export interface QueryCollectionArgs {
  id: Scalars['ID']['input'];
}


export interface QueryMetadataArgs {
  id: Scalars['ID']['input'];
}


export interface QuerySourceArgs {
  id: Scalars['ID']['input'];
}


export interface QueryTraitArgs {
  id: Scalars['ID']['input'];
}

export interface SignedUrl {
  __typename?: 'SignedUrl';
  headers?: Maybe<Array<SignedUrlHeader>>;
  id: Scalars['String']['output'];
  method: Scalars['String']['output'];
  url: Scalars['String']['output'];
}

export interface SignedUrlHeader {
  __typename?: 'SignedUrlHeader';
  name: Scalars['String']['output'];
  value: Scalars['String']['output'];
}

export interface Source {
  __typename?: 'Source';
  configuration: Array<Kv>;
  description: Scalars['String']['output'];
  id: Scalars['ID']['output'];
  name: Scalars['String']['output'];
}

export interface SourceInput {
  configuration: Array<KvInput>;
  description: Scalars['String']['input'];
  name: Scalars['String']['input'];
}

export interface StorageSystem {
  __typename?: 'StorageSystem';
  configuration: Array<Kv>;
  description?: Maybe<Scalars['String']['output']>;
  id: Scalars['ID']['output'];
  models: Array<StorageSystemModel>;
  name?: Maybe<Scalars['String']['output']>;
  type?: Maybe<Scalars['String']['output']>;
}

export interface StorageSystemInput {
  configuration: Array<KvInput>;
  description?: InputMaybe<Scalars['String']['input']>;
  models: Array<StorageSystemModelInput>;
  name?: InputMaybe<Scalars['String']['input']>;
  type?: InputMaybe<StorageSystemType>;
}

export interface StorageSystemModel {
  __typename?: 'StorageSystemModel';
  configuration: Array<Kv>;
  model: Model;
}

export interface StorageSystemModelInput {
  configuration: Array<KvInput>;
  id: Scalars['ID']['input'];
}

export enum StorageSystemType {
  Search = 'search',
  Vector = 'vector'
}

export interface Supplementary {
  __typename?: 'Supplementary';
  content?: Maybe<MetadataContent>;
  contentLength?: Maybe<Scalars['Int']['output']>;
  contentType: Scalars['String']['output'];
  created: Scalars['Date']['output'];
  downloadUrl: SignedUrl;
  key: Scalars['ID']['output'];
  metadataId: Scalars['ID']['output'];
  modified: Scalars['Date']['output'];
  name: Scalars['String']['output'];
  sourceId?: Maybe<Scalars['String']['output']>;
  sourceIdentifier?: Maybe<Scalars['String']['output']>;
  traitIds: Array<Scalars['String']['output']>;
  uploadUrl: SignedUrl;
}

export interface Trait {
  __typename?: 'Trait';
  id: Scalars['ID']['output'];
  name: Scalars['String']['output'];
  workflowIds: Array<Scalars['ID']['output']>;
}

export interface Workflow {
  __typename?: 'Workflow';
  activities: Array<WorkflowActivity>;
  configuration: Array<Kv>;
  description?: Maybe<Scalars['String']['output']>;
  id: Scalars['String']['output'];
  name: Scalars['String']['output'];
}

export interface WorkflowActivity {
  __typename?: 'WorkflowActivity';
  activityId: Scalars['String']['output'];
  childWorkflowId?: Maybe<Scalars['String']['output']>;
  configuration: Array<Kv>;
  executionGroup: Scalars['Int']['output'];
  id: Scalars['ID']['output'];
  inputs: Array<Kv>;
  models: Array<WorkflowActivityModel>;
  outputs: Array<Kv>;
  prompts: Array<WorkflowActivityPrompt>;
  queue: Scalars['String']['output'];
  storageSystems: Array<WorkflowActivityStorageSystem>;
}

export interface WorkflowActivityInput {
  activityId: Scalars['String']['input'];
  childWorkflowId?: InputMaybe<Scalars['String']['input']>;
  configuration: Array<KvInput>;
  executionGroup: Scalars['Int']['input'];
  inputs: Array<KvInput>;
  models?: InputMaybe<Array<WorkflowActivityModelInput>>;
  outputs: Array<KvInput>;
  prompts: Array<WorkflowActivityPromptInput>;
  queue: Scalars['String']['input'];
  storageSystems: Array<WorkflowActivityStorageSystemInput>;
}

export interface WorkflowActivityModel {
  __typename?: 'WorkflowActivityModel';
  configuration: Array<Kv>;
  model: Model;
}

export interface WorkflowActivityModelInput {
  configuration: Array<KvInput>;
  id: Scalars['ID']['input'];
}

export interface WorkflowActivityPrompt {
  __typename?: 'WorkflowActivityPrompt';
  configuration: Array<Kv>;
  prompt: Prompt;
}

export interface WorkflowActivityPromptInput {
  configuration: Array<KvInput>;
  id: Scalars['ID']['input'];
}

export interface WorkflowActivityStorageSystem {
  __typename?: 'WorkflowActivityStorageSystem';
  configuration: Array<Kv>;
  storageSystem: StorageSystem;
}

export interface WorkflowActivityStorageSystemInput {
  configuration: Array<KvInput>;
  id: Scalars['ID']['input'];
}

export interface WorkflowExecutionInput {
  context?: InputMaybe<Array<KvInput>>;
  metadataId?: InputMaybe<Scalars['String']['input']>;
  version?: InputMaybe<Scalars['Int']['input']>;
  workflowId: Scalars['String']['input'];
}

export interface WorkflowInput {
  configuration: Array<KvInput>;
  description: Scalars['String']['input'];
  id: Scalars['String']['input'];
  name: Scalars['String']['input'];
  queue: Scalars['String']['input'];
}

export interface WorkflowState {
  __typename?: 'WorkflowState';
  configuration: Array<Kv>;
  description: Scalars['String']['output'];
  entryWorkflowId?: Maybe<Scalars['String']['output']>;
  exitWorkflowId?: Maybe<Scalars['String']['output']>;
  id: Scalars['String']['output'];
  name: Scalars['String']['output'];
  queue: Scalars['String']['output'];
  type: WorkflowStateType;
  workflowId?: Maybe<Scalars['String']['output']>;
}

export interface WorkflowStateInput {
  configuration: Array<KvInput>;
  description: Scalars['String']['input'];
  entryWorkflowId?: InputMaybe<Scalars['String']['input']>;
  exitWorkflowId?: InputMaybe<Scalars['String']['input']>;
  id: Scalars['String']['input'];
  name: Scalars['String']['input'];
  queue: Scalars['String']['input'];
  type: WorkflowStateType;
  workflowId?: InputMaybe<Scalars['String']['input']>;
}

export interface WorkflowStateTransition {
  __typename?: 'WorkflowStateTransition';
  description: Scalars['String']['output'];
  fromStateId: Scalars['String']['output'];
  toStateId: Scalars['String']['output'];
}

export interface WorkflowStateTransitionInput {
  description: Scalars['String']['input'];
  fromStateId: Scalars['String']['input'];
  toStateId: Scalars['String']['input'];
}

export enum WorkflowStateType {
  Approval = 'approval',
  Approved = 'approved',
  Draft = 'draft',
  Failure = 'failure',
  Pending = 'pending',
  Processing = 'processing',
  Published = 'published'
}

export interface Workflows {
  __typename?: 'Workflows';
  models: Array<Model>;
  prompts: Array<Prompt>;
  states: Array<WorkflowState>;
  storageSystems: Array<StorageSystem>;
  transitions: Array<WorkflowStateTransition>;
  workflow: Workflow;
  workflows: Array<Workflow>;
}


export interface WorkflowsWorkflowArgs {
  id: Scalars['ID']['input'];
}

export type WithIndex<TObject> = TObject & Record<string, any>;
export type ResolversObject<TObject> = WithIndex<TObject>;

export type ResolverTypeWrapper<T> = Promise<T> | T;


export type ResolverWithResolve<TResult, TParent, TContext, TArgs> = {
  resolve: ResolverFn<TResult, TParent, TContext, TArgs>;
};
export type Resolver<TResult, TParent = {}, TContext = {}, TArgs = {}> = ResolverFn<TResult, TParent, TContext, TArgs> | ResolverWithResolve<TResult, TParent, TContext, TArgs>;

export type ResolverFn<TResult, TParent, TContext, TArgs> = (
  parent: TParent,
  args: TArgs,
  context: TContext,
  info: GraphQLResolveInfo
) => Promise<TResult> | TResult;

export type SubscriptionSubscribeFn<TResult, TParent, TContext, TArgs> = (
  parent: TParent,
  args: TArgs,
  context: TContext,
  info: GraphQLResolveInfo
) => AsyncIterable<TResult> | Promise<AsyncIterable<TResult>>;

export type SubscriptionResolveFn<TResult, TParent, TContext, TArgs> = (
  parent: TParent,
  args: TArgs,
  context: TContext,
  info: GraphQLResolveInfo
) => TResult | Promise<TResult>;

export interface SubscriptionSubscriberObject<TResult, TKey extends string, TParent, TContext, TArgs> {
  subscribe: SubscriptionSubscribeFn<{ [key in TKey]: TResult }, TParent, TContext, TArgs>;
  resolve?: SubscriptionResolveFn<TResult, { [key in TKey]: TResult }, TContext, TArgs>;
}

export interface SubscriptionResolverObject<TResult, TParent, TContext, TArgs> {
  subscribe: SubscriptionSubscribeFn<any, TParent, TContext, TArgs>;
  resolve: SubscriptionResolveFn<TResult, any, TContext, TArgs>;
}

export type SubscriptionObject<TResult, TKey extends string, TParent, TContext, TArgs> =
  | SubscriptionSubscriberObject<TResult, TKey, TParent, TContext, TArgs>
  | SubscriptionResolverObject<TResult, TParent, TContext, TArgs>;

export type SubscriptionResolver<TResult, TKey extends string, TParent = {}, TContext = {}, TArgs = {}> =
  | ((...args: any[]) => SubscriptionObject<TResult, TKey, TParent, TContext, TArgs>)
  | SubscriptionObject<TResult, TKey, TParent, TContext, TArgs>;

export type TypeResolveFn<TTypes, TParent = {}, TContext = {}> = (
  parent: TParent,
  context: TContext,
  info: GraphQLResolveInfo
) => Maybe<TTypes> | Promise<Maybe<TTypes>>;

export type IsTypeOfResolverFn<T = {}, TContext = {}> = (obj: T, context: TContext, info: GraphQLResolveInfo) => boolean | Promise<boolean>;

export type NextResolverFn<T> = () => Promise<T>;

export type DirectiveResolverFn<TResult = {}, TParent = {}, TContext = {}, TArgs = {}> = (
  next: NextResolverFn<TResult>,
  parent: TParent,
  args: TArgs,
  context: TContext,
  info: GraphQLResolveInfo
) => TResult | Promise<TResult>;

/** Mapping of union types */
export type ResolversUnionTypes<_RefType extends Record<string, unknown>> = ResolversObject<{
  CollectionItem: ( Omit<Collection, 'items'> & { items?: Maybe<Array<Maybe<_RefType['CollectionItem']>>> } ) | ( Metadata );
}>;


/** Mapping between all available schema types and the resolvers types */
export type ResolversTypes = ResolversObject<{
  Boolean: ResolverTypeWrapper<Scalars['Boolean']['output']>;
  Collection: ResolverTypeWrapper<Omit<Collection, 'items'> & { items?: Maybe<Array<Maybe<ResolversTypes['CollectionItem']>>> }>;
  CollectionInput: CollectionInput;
  CollectionItem: ResolverTypeWrapper<ResolversUnionTypes<ResolversTypes>['CollectionItem']>;
  CollectionItemFilter: CollectionItemFilter;
  CollectionType: CollectionType;
  Date: ResolverTypeWrapper<Scalars['Date']['output']>;
  Find: ResolverTypeWrapper<Omit<Find, 'collections'> & { collections: Array<ResolversTypes['Collection']> }>;
  FindInputAttribute: FindInputAttribute;
  FindQuery: FindQuery;
  ID: ResolverTypeWrapper<Scalars['ID']['output']>;
  Int: ResolverTypeWrapper<Scalars['Int']['output']>;
  JSON: ResolverTypeWrapper<Scalars['JSON']['output']>;
  JSONObject: ResolverTypeWrapper<Scalars['JSONObject']['output']>;
  KV: ResolverTypeWrapper<Kv>;
  KVInput: KvInput;
  Metadata: ResolverTypeWrapper<Metadata>;
  MetadataContent: ResolverTypeWrapper<MetadataContent>;
  MetadataInput: MetadataInput;
  MetadataWorkflowState: ResolverTypeWrapper<MetadataWorkflowState>;
  Model: ResolverTypeWrapper<Model>;
  ModelInput: ModelInput;
  Mutation: ResolverTypeWrapper<{}>;
  Permission: ResolverTypeWrapper<Permission>;
  PermissionInput: PermissionInput;
  PermissionRelation: PermissionRelation;
  PermissionSubjectType: PermissionSubjectType;
  Prompt: ResolverTypeWrapper<Prompt>;
  PromptInput: PromptInput;
  Query: ResolverTypeWrapper<{}>;
  SignedUrl: ResolverTypeWrapper<SignedUrl>;
  SignedUrlHeader: ResolverTypeWrapper<SignedUrlHeader>;
  Source: ResolverTypeWrapper<Source>;
  SourceInput: SourceInput;
  StorageSystem: ResolverTypeWrapper<StorageSystem>;
  StorageSystemInput: StorageSystemInput;
  StorageSystemModel: ResolverTypeWrapper<StorageSystemModel>;
  StorageSystemModelInput: StorageSystemModelInput;
  StorageSystemType: StorageSystemType;
  String: ResolverTypeWrapper<Scalars['String']['output']>;
  Supplementary: ResolverTypeWrapper<Supplementary>;
  Trait: ResolverTypeWrapper<Trait>;
  Workflow: ResolverTypeWrapper<Workflow>;
  WorkflowActivity: ResolverTypeWrapper<WorkflowActivity>;
  WorkflowActivityInput: WorkflowActivityInput;
  WorkflowActivityModel: ResolverTypeWrapper<WorkflowActivityModel>;
  WorkflowActivityModelInput: WorkflowActivityModelInput;
  WorkflowActivityPrompt: ResolverTypeWrapper<WorkflowActivityPrompt>;
  WorkflowActivityPromptInput: WorkflowActivityPromptInput;
  WorkflowActivityStorageSystem: ResolverTypeWrapper<WorkflowActivityStorageSystem>;
  WorkflowActivityStorageSystemInput: WorkflowActivityStorageSystemInput;
  WorkflowExecutionInput: WorkflowExecutionInput;
  WorkflowInput: WorkflowInput;
  WorkflowState: ResolverTypeWrapper<WorkflowState>;
  WorkflowStateInput: WorkflowStateInput;
  WorkflowStateTransition: ResolverTypeWrapper<WorkflowStateTransition>;
  WorkflowStateTransitionInput: WorkflowStateTransitionInput;
  WorkflowStateType: WorkflowStateType;
  Workflows: ResolverTypeWrapper<Workflows>;
}>;

/** Mapping between all available schema types and the resolvers parents */
export type ResolversParentTypes = ResolversObject<{
  Boolean: Scalars['Boolean']['output'];
  Collection: Omit<Collection, 'items'> & { items?: Maybe<Array<Maybe<ResolversParentTypes['CollectionItem']>>> };
  CollectionInput: CollectionInput;
  CollectionItem: ResolversUnionTypes<ResolversParentTypes>['CollectionItem'];
  CollectionItemFilter: CollectionItemFilter;
  Date: Scalars['Date']['output'];
  Find: Omit<Find, 'collections'> & { collections: Array<ResolversParentTypes['Collection']> };
  FindInputAttribute: FindInputAttribute;
  FindQuery: FindQuery;
  ID: Scalars['ID']['output'];
  Int: Scalars['Int']['output'];
  JSON: Scalars['JSON']['output'];
  JSONObject: Scalars['JSONObject']['output'];
  KV: Kv;
  KVInput: KvInput;
  Metadata: Metadata;
  MetadataContent: MetadataContent;
  MetadataInput: MetadataInput;
  MetadataWorkflowState: MetadataWorkflowState;
  Model: Model;
  ModelInput: ModelInput;
  Mutation: {};
  Permission: Permission;
  PermissionInput: PermissionInput;
  Prompt: Prompt;
  PromptInput: PromptInput;
  Query: {};
  SignedUrl: SignedUrl;
  SignedUrlHeader: SignedUrlHeader;
  Source: Source;
  SourceInput: SourceInput;
  StorageSystem: StorageSystem;
  StorageSystemInput: StorageSystemInput;
  StorageSystemModel: StorageSystemModel;
  StorageSystemModelInput: StorageSystemModelInput;
  String: Scalars['String']['output'];
  Supplementary: Supplementary;
  Trait: Trait;
  Workflow: Workflow;
  WorkflowActivity: WorkflowActivity;
  WorkflowActivityInput: WorkflowActivityInput;
  WorkflowActivityModel: WorkflowActivityModel;
  WorkflowActivityModelInput: WorkflowActivityModelInput;
  WorkflowActivityPrompt: WorkflowActivityPrompt;
  WorkflowActivityPromptInput: WorkflowActivityPromptInput;
  WorkflowActivityStorageSystem: WorkflowActivityStorageSystem;
  WorkflowActivityStorageSystemInput: WorkflowActivityStorageSystemInput;
  WorkflowExecutionInput: WorkflowExecutionInput;
  WorkflowInput: WorkflowInput;
  WorkflowState: WorkflowState;
  WorkflowStateInput: WorkflowStateInput;
  WorkflowStateTransition: WorkflowStateTransition;
  WorkflowStateTransitionInput: WorkflowStateTransitionInput;
  Workflows: Workflows;
}>;

export type CollectionResolvers<ContextType = any, ParentType extends ResolversParentTypes['Collection'] = ResolversParentTypes['Collection']> = ResolversObject<{
  attributes?: Resolver<Array<ResolversTypes['KV']>, ParentType, ContextType>;
  categoryIds?: Resolver<Array<ResolversTypes['String']>, ParentType, ContextType>;
  created?: Resolver<ResolversTypes['Date'], ParentType, ContextType>;
  id?: Resolver<ResolversTypes['ID'], ParentType, ContextType>;
  items?: Resolver<Maybe<Array<Maybe<ResolversTypes['CollectionItem']>>>, ParentType, ContextType, Partial<CollectionItemsArgs>>;
  labels?: Resolver<Array<ResolversTypes['String']>, ParentType, ContextType>;
  modified?: Resolver<ResolversTypes['Date'], ParentType, ContextType>;
  name?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  permissions?: Resolver<Array<ResolversTypes['Permission']>, ParentType, ContextType>;
  traitIds?: Resolver<Array<ResolversTypes['String']>, ParentType, ContextType>;
  type?: Resolver<ResolversTypes['CollectionType'], ParentType, ContextType>;
  __isTypeOf?: IsTypeOfResolverFn<ParentType, ContextType>;
}>;

export type CollectionItemResolvers<ContextType = any, ParentType extends ResolversParentTypes['CollectionItem'] = ResolversParentTypes['CollectionItem']> = ResolversObject<{
  __resolveType: TypeResolveFn<'Collection' | 'Metadata', ParentType, ContextType>;
}>;

export interface DateScalarConfig extends GraphQLScalarTypeConfig<ResolversTypes['Date'], any> {
  name: 'Date';
}

export type FindResolvers<ContextType = any, ParentType extends ResolversParentTypes['Find'] = ResolversParentTypes['Find']> = ResolversObject<{
  collections?: Resolver<Array<ResolversTypes['Collection']>, ParentType, ContextType, Partial<FindCollectionsArgs>>;
  metadata?: Resolver<Array<ResolversTypes['Metadata']>, ParentType, ContextType, Partial<FindMetadataArgs>>;
  __isTypeOf?: IsTypeOfResolverFn<ParentType, ContextType>;
}>;

export interface JsonScalarConfig extends GraphQLScalarTypeConfig<ResolversTypes['JSON'], any> {
  name: 'JSON';
}

export interface JsonObjectScalarConfig extends GraphQLScalarTypeConfig<ResolversTypes['JSONObject'], any> {
  name: 'JSONObject';
}

export type KvResolvers<ContextType = any, ParentType extends ResolversParentTypes['KV'] = ResolversParentTypes['KV']> = ResolversObject<{
  key?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  value?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  __isTypeOf?: IsTypeOfResolverFn<ParentType, ContextType>;
}>;

export type MetadataResolvers<ContextType = any, ParentType extends ResolversParentTypes['Metadata'] = ResolversParentTypes['Metadata']> = ResolversObject<{
  attributes?: Resolver<Array<ResolversTypes['KV']>, ParentType, ContextType>;
  content?: Resolver<Maybe<ResolversTypes['MetadataContent']>, ParentType, ContextType>;
  contentLength?: Resolver<Maybe<ResolversTypes['Int']>, ParentType, ContextType>;
  contentType?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  created?: Resolver<ResolversTypes['Date'], ParentType, ContextType>;
  downloadUrl?: Resolver<ResolversTypes['SignedUrl'], ParentType, ContextType>;
  id?: Resolver<ResolversTypes['ID'], ParentType, ContextType>;
  labels?: Resolver<Array<ResolversTypes['String']>, ParentType, ContextType>;
  languageTag?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  modified?: Resolver<ResolversTypes['Date'], ParentType, ContextType>;
  name?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  parentId?: Resolver<Maybe<ResolversTypes['ID']>, ParentType, ContextType>;
  permissions?: Resolver<Array<ResolversTypes['Permission']>, ParentType, ContextType>;
  sourceId?: Resolver<Maybe<ResolversTypes['String']>, ParentType, ContextType>;
  sourceIdentifier?: Resolver<Maybe<ResolversTypes['String']>, ParentType, ContextType>;
  supplementaries?: Resolver<Array<ResolversTypes['Supplementary']>, ParentType, ContextType, Partial<MetadataSupplementariesArgs>>;
  supplementary?: Resolver<Maybe<ResolversTypes['Supplementary']>, ParentType, ContextType, RequireFields<MetadataSupplementaryArgs, 'key'>>;
  traitIds?: Resolver<Array<ResolversTypes['String']>, ParentType, ContextType>;
  uploadUrl?: Resolver<ResolversTypes['SignedUrl'], ParentType, ContextType>;
  workflowState?: Resolver<ResolversTypes['MetadataWorkflowState'], ParentType, ContextType>;
  __isTypeOf?: IsTypeOfResolverFn<ParentType, ContextType>;
}>;

export type MetadataContentResolvers<ContextType = any, ParentType extends ResolversParentTypes['MetadataContent'] = ResolversParentTypes['MetadataContent']> = ResolversObject<{
  json?: Resolver<Maybe<ResolversTypes['JSONObject']>, ParentType, ContextType>;
  text?: Resolver<Maybe<ResolversTypes['String']>, ParentType, ContextType>;
  __isTypeOf?: IsTypeOfResolverFn<ParentType, ContextType>;
}>;

export type MetadataWorkflowStateResolvers<ContextType = any, ParentType extends ResolversParentTypes['MetadataWorkflowState'] = ResolversParentTypes['MetadataWorkflowState']> = ResolversObject<{
  deleteWorkflowId?: Resolver<Maybe<ResolversTypes['String']>, ParentType, ContextType>;
  id?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  pendingId?: Resolver<Maybe<ResolversTypes['String']>, ParentType, ContextType>;
  __isTypeOf?: IsTypeOfResolverFn<ParentType, ContextType>;
}>;

export type ModelResolvers<ContextType = any, ParentType extends ResolversParentTypes['Model'] = ResolversParentTypes['Model']> = ResolversObject<{
  configuration?: Resolver<Array<ResolversTypes['KV']>, ParentType, ContextType>;
  description?: Resolver<Maybe<ResolversTypes['String']>, ParentType, ContextType>;
  id?: Resolver<ResolversTypes['ID'], ParentType, ContextType>;
  name?: Resolver<Maybe<ResolversTypes['String']>, ParentType, ContextType>;
  type?: Resolver<Maybe<ResolversTypes['String']>, ParentType, ContextType>;
  __isTypeOf?: IsTypeOfResolverFn<ParentType, ContextType>;
}>;

export type MutationResolvers<ContextType = any, ParentType extends ResolversParentTypes['Mutation'] = ResolversParentTypes['Mutation']> = ResolversObject<{
  addCollection?: Resolver<Maybe<ResolversTypes['Collection']>, ParentType, ContextType, RequireFields<MutationAddCollectionArgs, 'collection'>>;
  addCollectionPermissions?: Resolver<Maybe<ResolversTypes['Collection']>, ParentType, ContextType, RequireFields<MutationAddCollectionPermissionsArgs, 'id' | 'permissions'>>;
  addMetadata?: Resolver<Maybe<ResolversTypes['Metadata']>, ParentType, ContextType, RequireFields<MutationAddMetadataArgs, 'metadata'>>;
  addMetadataPermissions?: Resolver<Maybe<ResolversTypes['Metadata']>, ParentType, ContextType, RequireFields<MutationAddMetadataPermissionsArgs, 'id' | 'permissions'>>;
  addModel?: Resolver<ResolversTypes['Model'], ParentType, ContextType, RequireFields<MutationAddModelArgs, 'model'>>;
  addPrompt?: Resolver<ResolversTypes['Prompt'], ParentType, ContextType, RequireFields<MutationAddPromptArgs, 'prompt'>>;
  addSource?: Resolver<ResolversTypes['Source'], ParentType, ContextType, RequireFields<MutationAddSourceArgs, 'source'>>;
  addStorageSystem?: Resolver<ResolversTypes['StorageSystem'], ParentType, ContextType, RequireFields<MutationAddStorageSystemArgs, 'storageSystem'>>;
  addWorkflow?: Resolver<ResolversTypes['Workflow'], ParentType, ContextType, Partial<MutationAddWorkflowArgs>>;
  addWorkflowActivity?: Resolver<ResolversTypes['WorkflowActivity'], ParentType, ContextType, Partial<MutationAddWorkflowActivityArgs>>;
  addWorkflowState?: Resolver<ResolversTypes['WorkflowState'], ParentType, ContextType, Partial<MutationAddWorkflowStateArgs>>;
  addWorkflowStateTransition?: Resolver<ResolversTypes['WorkflowStateTransition'], ParentType, ContextType, RequireFields<MutationAddWorkflowStateTransitionArgs, 'workflow'>>;
  deleteCollection?: Resolver<Maybe<ResolversTypes['Boolean']>, ParentType, ContextType, RequireFields<MutationDeleteCollectionArgs, 'id'>>;
  deleteCollectionPermissions?: Resolver<Maybe<ResolversTypes['Collection']>, ParentType, ContextType, RequireFields<MutationDeleteCollectionPermissionsArgs, 'id' | 'permissions'>>;
  deleteMetadata?: Resolver<Maybe<ResolversTypes['Boolean']>, ParentType, ContextType, RequireFields<MutationDeleteMetadataArgs, 'id'>>;
  deleteMetadataPermissions?: Resolver<Maybe<ResolversTypes['Metadata']>, ParentType, ContextType, RequireFields<MutationDeleteMetadataPermissionsArgs, 'id' | 'permissions'>>;
  deleteModel?: Resolver<Maybe<ResolversTypes['Boolean']>, ParentType, ContextType, RequireFields<MutationDeleteModelArgs, 'id'>>;
  deletePrompt?: Resolver<ResolversTypes['Boolean'], ParentType, ContextType, RequireFields<MutationDeletePromptArgs, 'id'>>;
  deleteSource?: Resolver<ResolversTypes['Boolean'], ParentType, ContextType, RequireFields<MutationDeleteSourceArgs, 'id'>>;
  deleteStorageSystem?: Resolver<Maybe<ResolversTypes['Boolean']>, ParentType, ContextType, RequireFields<MutationDeleteStorageSystemArgs, 'id'>>;
  deleteWorkflow?: Resolver<ResolversTypes['Boolean'], ParentType, ContextType, RequireFields<MutationDeleteWorkflowArgs, 'id'>>;
  deleteWorkflowActivity?: Resolver<ResolversTypes['Boolean'], ParentType, ContextType, RequireFields<MutationDeleteWorkflowActivityArgs, 'id'>>;
  deleteWorkflowTransition?: Resolver<Maybe<ResolversTypes['Boolean']>, ParentType, ContextType, RequireFields<MutationDeleteWorkflowTransitionArgs, 'fromStateId' | 'toStateId'>>;
  editModel?: Resolver<ResolversTypes['Model'], ParentType, ContextType, RequireFields<MutationEditModelArgs, 'id'>>;
  editPrompt?: Resolver<ResolversTypes['Prompt'], ParentType, ContextType, RequireFields<MutationEditPromptArgs, 'id' | 'prompt'>>;
  editSource?: Resolver<ResolversTypes['Source'], ParentType, ContextType, RequireFields<MutationEditSourceArgs, 'id' | 'source'>>;
  editStorageSystem?: Resolver<ResolversTypes['StorageSystem'], ParentType, ContextType, RequireFields<MutationEditStorageSystemArgs, 'id'>>;
  editWorkflow?: Resolver<ResolversTypes['Workflow'], ParentType, ContextType, Partial<MutationEditWorkflowArgs>>;
  editWorkflowActivity?: Resolver<ResolversTypes['WorkflowActivity'], ParentType, ContextType, RequireFields<MutationEditWorkflowActivityArgs, 'id'>>;
  executeWorkflow?: Resolver<ResolversTypes['String'], ParentType, ContextType, Partial<MutationExecuteWorkflowArgs>>;
  login?: Resolver<Maybe<ResolversTypes['String']>, ParentType, ContextType, RequireFields<MutationLoginArgs, 'password' | 'username'>>;
  setMetadataJSONContent?: Resolver<Maybe<ResolversTypes['Metadata']>, ParentType, ContextType, RequireFields<MutationSetMetadataJsonContentArgs, 'id'>>;
  setMetadataReady?: Resolver<Maybe<ResolversTypes['Metadata']>, ParentType, ContextType, RequireFields<MutationSetMetadataReadyArgs, 'id'>>;
  setMetadataTextContent?: Resolver<Maybe<ResolversTypes['Metadata']>, ParentType, ContextType, RequireFields<MutationSetMetadataTextContentArgs, 'id'>>;
  setPassword?: Resolver<Maybe<ResolversTypes['Boolean']>, ParentType, ContextType, RequireFields<MutationSetPasswordArgs, 'password'>>;
  signup?: Resolver<Maybe<ResolversTypes['String']>, ParentType, ContextType, RequireFields<MutationSignupArgs, 'email' | 'firstName' | 'lastName' | 'password'>>;
}>;

export type PermissionResolvers<ContextType = any, ParentType extends ResolversParentTypes['Permission'] = ResolversParentTypes['Permission']> = ResolversObject<{
  relation?: Resolver<ResolversTypes['PermissionRelation'], ParentType, ContextType>;
  subject?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  subjectType?: Resolver<ResolversTypes['PermissionSubjectType'], ParentType, ContextType>;
  __isTypeOf?: IsTypeOfResolverFn<ParentType, ContextType>;
}>;

export type PromptResolvers<ContextType = any, ParentType extends ResolversParentTypes['Prompt'] = ResolversParentTypes['Prompt']> = ResolversObject<{
  description?: Resolver<Maybe<ResolversTypes['String']>, ParentType, ContextType>;
  id?: Resolver<Maybe<ResolversTypes['String']>, ParentType, ContextType>;
  inputType?: Resolver<Maybe<ResolversTypes['String']>, ParentType, ContextType>;
  name?: Resolver<Maybe<ResolversTypes['String']>, ParentType, ContextType>;
  outputType?: Resolver<Maybe<ResolversTypes['String']>, ParentType, ContextType>;
  systemPrompt?: Resolver<Maybe<ResolversTypes['String']>, ParentType, ContextType>;
  userPrompt?: Resolver<Maybe<ResolversTypes['String']>, ParentType, ContextType>;
  __isTypeOf?: IsTypeOfResolverFn<ParentType, ContextType>;
}>;

export type QueryResolvers<ContextType = any, ParentType extends ResolversParentTypes['Query'] = ResolversParentTypes['Query']> = ResolversObject<{
  collection?: Resolver<Maybe<ResolversTypes['Collection']>, ParentType, ContextType, RequireFields<QueryCollectionArgs, 'id'>>;
  find?: Resolver<Maybe<ResolversTypes['Find']>, ParentType, ContextType>;
  metadata?: Resolver<Maybe<ResolversTypes['Metadata']>, ParentType, ContextType, RequireFields<QueryMetadataArgs, 'id'>>;
  source?: Resolver<Maybe<ResolversTypes['Source']>, ParentType, ContextType, RequireFields<QuerySourceArgs, 'id'>>;
  sources?: Resolver<Array<ResolversTypes['Source']>, ParentType, ContextType>;
  trait?: Resolver<Maybe<ResolversTypes['Trait']>, ParentType, ContextType, RequireFields<QueryTraitArgs, 'id'>>;
  traits?: Resolver<Array<ResolversTypes['Trait']>, ParentType, ContextType>;
  workflows?: Resolver<ResolversTypes['Workflows'], ParentType, ContextType>;
}>;

export type SignedUrlResolvers<ContextType = any, ParentType extends ResolversParentTypes['SignedUrl'] = ResolversParentTypes['SignedUrl']> = ResolversObject<{
  headers?: Resolver<Maybe<Array<ResolversTypes['SignedUrlHeader']>>, ParentType, ContextType>;
  id?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  method?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  url?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  __isTypeOf?: IsTypeOfResolverFn<ParentType, ContextType>;
}>;

export type SignedUrlHeaderResolvers<ContextType = any, ParentType extends ResolversParentTypes['SignedUrlHeader'] = ResolversParentTypes['SignedUrlHeader']> = ResolversObject<{
  name?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  value?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  __isTypeOf?: IsTypeOfResolverFn<ParentType, ContextType>;
}>;

export type SourceResolvers<ContextType = any, ParentType extends ResolversParentTypes['Source'] = ResolversParentTypes['Source']> = ResolversObject<{
  configuration?: Resolver<Array<ResolversTypes['KV']>, ParentType, ContextType>;
  description?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  id?: Resolver<ResolversTypes['ID'], ParentType, ContextType>;
  name?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  __isTypeOf?: IsTypeOfResolverFn<ParentType, ContextType>;
}>;

export type StorageSystemResolvers<ContextType = any, ParentType extends ResolversParentTypes['StorageSystem'] = ResolversParentTypes['StorageSystem']> = ResolversObject<{
  configuration?: Resolver<Array<ResolversTypes['KV']>, ParentType, ContextType>;
  description?: Resolver<Maybe<ResolversTypes['String']>, ParentType, ContextType>;
  id?: Resolver<ResolversTypes['ID'], ParentType, ContextType>;
  models?: Resolver<Array<ResolversTypes['StorageSystemModel']>, ParentType, ContextType>;
  name?: Resolver<Maybe<ResolversTypes['String']>, ParentType, ContextType>;
  type?: Resolver<Maybe<ResolversTypes['String']>, ParentType, ContextType>;
  __isTypeOf?: IsTypeOfResolverFn<ParentType, ContextType>;
}>;

export type StorageSystemModelResolvers<ContextType = any, ParentType extends ResolversParentTypes['StorageSystemModel'] = ResolversParentTypes['StorageSystemModel']> = ResolversObject<{
  configuration?: Resolver<Array<ResolversTypes['KV']>, ParentType, ContextType>;
  model?: Resolver<ResolversTypes['Model'], ParentType, ContextType>;
  __isTypeOf?: IsTypeOfResolverFn<ParentType, ContextType>;
}>;

export type SupplementaryResolvers<ContextType = any, ParentType extends ResolversParentTypes['Supplementary'] = ResolversParentTypes['Supplementary']> = ResolversObject<{
  content?: Resolver<Maybe<ResolversTypes['MetadataContent']>, ParentType, ContextType>;
  contentLength?: Resolver<Maybe<ResolversTypes['Int']>, ParentType, ContextType>;
  contentType?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  created?: Resolver<ResolversTypes['Date'], ParentType, ContextType>;
  downloadUrl?: Resolver<ResolversTypes['SignedUrl'], ParentType, ContextType>;
  key?: Resolver<ResolversTypes['ID'], ParentType, ContextType>;
  metadataId?: Resolver<ResolversTypes['ID'], ParentType, ContextType>;
  modified?: Resolver<ResolversTypes['Date'], ParentType, ContextType>;
  name?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  sourceId?: Resolver<Maybe<ResolversTypes['String']>, ParentType, ContextType>;
  sourceIdentifier?: Resolver<Maybe<ResolversTypes['String']>, ParentType, ContextType>;
  traitIds?: Resolver<Array<ResolversTypes['String']>, ParentType, ContextType>;
  uploadUrl?: Resolver<ResolversTypes['SignedUrl'], ParentType, ContextType>;
  __isTypeOf?: IsTypeOfResolverFn<ParentType, ContextType>;
}>;

export type TraitResolvers<ContextType = any, ParentType extends ResolversParentTypes['Trait'] = ResolversParentTypes['Trait']> = ResolversObject<{
  id?: Resolver<ResolversTypes['ID'], ParentType, ContextType>;
  name?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  workflowIds?: Resolver<Array<ResolversTypes['ID']>, ParentType, ContextType>;
  __isTypeOf?: IsTypeOfResolverFn<ParentType, ContextType>;
}>;

export type WorkflowResolvers<ContextType = any, ParentType extends ResolversParentTypes['Workflow'] = ResolversParentTypes['Workflow']> = ResolversObject<{
  activities?: Resolver<Array<ResolversTypes['WorkflowActivity']>, ParentType, ContextType>;
  configuration?: Resolver<Array<ResolversTypes['KV']>, ParentType, ContextType>;
  description?: Resolver<Maybe<ResolversTypes['String']>, ParentType, ContextType>;
  id?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  name?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  __isTypeOf?: IsTypeOfResolverFn<ParentType, ContextType>;
}>;

export type WorkflowActivityResolvers<ContextType = any, ParentType extends ResolversParentTypes['WorkflowActivity'] = ResolversParentTypes['WorkflowActivity']> = ResolversObject<{
  activityId?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  childWorkflowId?: Resolver<Maybe<ResolversTypes['String']>, ParentType, ContextType>;
  configuration?: Resolver<Array<ResolversTypes['KV']>, ParentType, ContextType>;
  executionGroup?: Resolver<ResolversTypes['Int'], ParentType, ContextType>;
  id?: Resolver<ResolversTypes['ID'], ParentType, ContextType>;
  inputs?: Resolver<Array<ResolversTypes['KV']>, ParentType, ContextType>;
  models?: Resolver<Array<ResolversTypes['WorkflowActivityModel']>, ParentType, ContextType>;
  outputs?: Resolver<Array<ResolversTypes['KV']>, ParentType, ContextType>;
  prompts?: Resolver<Array<ResolversTypes['WorkflowActivityPrompt']>, ParentType, ContextType>;
  queue?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  storageSystems?: Resolver<Array<ResolversTypes['WorkflowActivityStorageSystem']>, ParentType, ContextType>;
  __isTypeOf?: IsTypeOfResolverFn<ParentType, ContextType>;
}>;

export type WorkflowActivityModelResolvers<ContextType = any, ParentType extends ResolversParentTypes['WorkflowActivityModel'] = ResolversParentTypes['WorkflowActivityModel']> = ResolversObject<{
  configuration?: Resolver<Array<ResolversTypes['KV']>, ParentType, ContextType>;
  model?: Resolver<ResolversTypes['Model'], ParentType, ContextType>;
  __isTypeOf?: IsTypeOfResolverFn<ParentType, ContextType>;
}>;

export type WorkflowActivityPromptResolvers<ContextType = any, ParentType extends ResolversParentTypes['WorkflowActivityPrompt'] = ResolversParentTypes['WorkflowActivityPrompt']> = ResolversObject<{
  configuration?: Resolver<Array<ResolversTypes['KV']>, ParentType, ContextType>;
  prompt?: Resolver<ResolversTypes['Prompt'], ParentType, ContextType>;
  __isTypeOf?: IsTypeOfResolverFn<ParentType, ContextType>;
}>;

export type WorkflowActivityStorageSystemResolvers<ContextType = any, ParentType extends ResolversParentTypes['WorkflowActivityStorageSystem'] = ResolversParentTypes['WorkflowActivityStorageSystem']> = ResolversObject<{
  configuration?: Resolver<Array<ResolversTypes['KV']>, ParentType, ContextType>;
  storageSystem?: Resolver<ResolversTypes['StorageSystem'], ParentType, ContextType>;
  __isTypeOf?: IsTypeOfResolverFn<ParentType, ContextType>;
}>;

export type WorkflowStateResolvers<ContextType = any, ParentType extends ResolversParentTypes['WorkflowState'] = ResolversParentTypes['WorkflowState']> = ResolversObject<{
  configuration?: Resolver<Array<ResolversTypes['KV']>, ParentType, ContextType>;
  description?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  entryWorkflowId?: Resolver<Maybe<ResolversTypes['String']>, ParentType, ContextType>;
  exitWorkflowId?: Resolver<Maybe<ResolversTypes['String']>, ParentType, ContextType>;
  id?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  name?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  queue?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  type?: Resolver<ResolversTypes['WorkflowStateType'], ParentType, ContextType>;
  workflowId?: Resolver<Maybe<ResolversTypes['String']>, ParentType, ContextType>;
  __isTypeOf?: IsTypeOfResolverFn<ParentType, ContextType>;
}>;

export type WorkflowStateTransitionResolvers<ContextType = any, ParentType extends ResolversParentTypes['WorkflowStateTransition'] = ResolversParentTypes['WorkflowStateTransition']> = ResolversObject<{
  description?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  fromStateId?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  toStateId?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  __isTypeOf?: IsTypeOfResolverFn<ParentType, ContextType>;
}>;

export type WorkflowsResolvers<ContextType = any, ParentType extends ResolversParentTypes['Workflows'] = ResolversParentTypes['Workflows']> = ResolversObject<{
  models?: Resolver<Array<ResolversTypes['Model']>, ParentType, ContextType>;
  prompts?: Resolver<Array<ResolversTypes['Prompt']>, ParentType, ContextType>;
  states?: Resolver<Array<ResolversTypes['WorkflowState']>, ParentType, ContextType>;
  storageSystems?: Resolver<Array<ResolversTypes['StorageSystem']>, ParentType, ContextType>;
  transitions?: Resolver<Array<ResolversTypes['WorkflowStateTransition']>, ParentType, ContextType>;
  workflow?: Resolver<ResolversTypes['Workflow'], ParentType, ContextType, RequireFields<WorkflowsWorkflowArgs, 'id'>>;
  workflows?: Resolver<Array<ResolversTypes['Workflow']>, ParentType, ContextType>;
  __isTypeOf?: IsTypeOfResolverFn<ParentType, ContextType>;
}>;

export type Resolvers<ContextType = any> = ResolversObject<{
  Collection?: CollectionResolvers<ContextType>;
  CollectionItem?: CollectionItemResolvers<ContextType>;
  Date?: GraphQLScalarType;
  Find?: FindResolvers<ContextType>;
  JSON?: GraphQLScalarType;
  JSONObject?: GraphQLScalarType;
  KV?: KvResolvers<ContextType>;
  Metadata?: MetadataResolvers<ContextType>;
  MetadataContent?: MetadataContentResolvers<ContextType>;
  MetadataWorkflowState?: MetadataWorkflowStateResolvers<ContextType>;
  Model?: ModelResolvers<ContextType>;
  Mutation?: MutationResolvers<ContextType>;
  Permission?: PermissionResolvers<ContextType>;
  Prompt?: PromptResolvers<ContextType>;
  Query?: QueryResolvers<ContextType>;
  SignedUrl?: SignedUrlResolvers<ContextType>;
  SignedUrlHeader?: SignedUrlHeaderResolvers<ContextType>;
  Source?: SourceResolvers<ContextType>;
  StorageSystem?: StorageSystemResolvers<ContextType>;
  StorageSystemModel?: StorageSystemModelResolvers<ContextType>;
  Supplementary?: SupplementaryResolvers<ContextType>;
  Trait?: TraitResolvers<ContextType>;
  Workflow?: WorkflowResolvers<ContextType>;
  WorkflowActivity?: WorkflowActivityResolvers<ContextType>;
  WorkflowActivityModel?: WorkflowActivityModelResolvers<ContextType>;
  WorkflowActivityPrompt?: WorkflowActivityPromptResolvers<ContextType>;
  WorkflowActivityStorageSystem?: WorkflowActivityStorageSystemResolvers<ContextType>;
  WorkflowState?: WorkflowStateResolvers<ContextType>;
  WorkflowStateTransition?: WorkflowStateTransitionResolvers<ContextType>;
  Workflows?: WorkflowsResolvers<ContextType>;
}>;

