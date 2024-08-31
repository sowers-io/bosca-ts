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

export interface Attribute {
  __typename?: 'Attribute';
  name: Scalars['String']['output'];
  value: Scalars['String']['output'];
}

export interface Collection {
  __typename?: 'Collection';
  attributes: Array<Attribute>;
  categoryIds: Array<Scalars['String']['output']>;
  created: Scalars['Date']['output'];
  id: Scalars['ID']['output'];
  items?: Maybe<Array<Maybe<CollectionItem>>>;
  labels: Array<Scalars['String']['output']>;
  modified: Scalars['Date']['output'];
  name: Scalars['String']['output'];
  traitIds: Array<Scalars['String']['output']>;
  type: CollectionType;
}


export interface CollectionItemsArgs {
  filter?: InputMaybe<CollectionItemFilter>;
}

export interface CollectionInput {
  contentLength?: InputMaybe<Scalars['Int']['input']>;
  contentType: Scalars['String']['input'];
  languageTag: Scalars['String']['input'];
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

export interface Metadata {
  __typename?: 'Metadata';
  attributes: Array<Attribute>;
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
  sourceId?: Maybe<Scalars['String']['output']>;
  sourceIdentifier?: Maybe<Scalars['String']['output']>;
  supplementaries: Array<Supplementary>;
  supplementary?: Maybe<Supplementary>;
  traitIds: Array<Scalars['String']['output']>;
  uploadUrl: SignedUrl;
  workflowState: MetadataWorkflowState;
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

export interface Mutation {
  __typename?: 'Mutation';
  addMetadata?: Maybe<Metadata>;
  login?: Maybe<Scalars['String']['output']>;
  setMetadataJSONContent?: Maybe<Metadata>;
  setMetadataReady?: Maybe<Metadata>;
  setMetadataTextContent?: Maybe<Metadata>;
  setPassword?: Maybe<Scalars['Boolean']['output']>;
}


export interface MutationAddMetadataArgs {
  metadata: MetadataInput;
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

export interface Query {
  __typename?: 'Query';
  collection?: Maybe<Collection>;
  find?: Maybe<Find>;
  metadata?: Maybe<Metadata>;
  source?: Maybe<Source>;
  sources: Array<Source>;
  trait?: Maybe<Trait>;
  traits: Array<Trait>;
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
  id: Scalars['ID']['output'];
  name: Scalars['String']['output'];
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
  Attribute: ResolverTypeWrapper<Attribute>;
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
  Metadata: ResolverTypeWrapper<Metadata>;
  MetadataContent: ResolverTypeWrapper<MetadataContent>;
  MetadataInput: MetadataInput;
  MetadataWorkflowState: ResolverTypeWrapper<MetadataWorkflowState>;
  Mutation: ResolverTypeWrapper<{}>;
  Query: ResolverTypeWrapper<{}>;
  SignedUrl: ResolverTypeWrapper<SignedUrl>;
  SignedUrlHeader: ResolverTypeWrapper<SignedUrlHeader>;
  Source: ResolverTypeWrapper<Source>;
  String: ResolverTypeWrapper<Scalars['String']['output']>;
  Supplementary: ResolverTypeWrapper<Supplementary>;
  Trait: ResolverTypeWrapper<Trait>;
}>;

/** Mapping between all available schema types and the resolvers parents */
export type ResolversParentTypes = ResolversObject<{
  Attribute: Attribute;
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
  Metadata: Metadata;
  MetadataContent: MetadataContent;
  MetadataInput: MetadataInput;
  MetadataWorkflowState: MetadataWorkflowState;
  Mutation: {};
  Query: {};
  SignedUrl: SignedUrl;
  SignedUrlHeader: SignedUrlHeader;
  Source: Source;
  String: Scalars['String']['output'];
  Supplementary: Supplementary;
  Trait: Trait;
}>;

export type AttributeResolvers<ContextType = any, ParentType extends ResolversParentTypes['Attribute'] = ResolversParentTypes['Attribute']> = ResolversObject<{
  name?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  value?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  __isTypeOf?: IsTypeOfResolverFn<ParentType, ContextType>;
}>;

export type CollectionResolvers<ContextType = any, ParentType extends ResolversParentTypes['Collection'] = ResolversParentTypes['Collection']> = ResolversObject<{
  attributes?: Resolver<Array<ResolversTypes['Attribute']>, ParentType, ContextType>;
  categoryIds?: Resolver<Array<ResolversTypes['String']>, ParentType, ContextType>;
  created?: Resolver<ResolversTypes['Date'], ParentType, ContextType>;
  id?: Resolver<ResolversTypes['ID'], ParentType, ContextType>;
  items?: Resolver<Maybe<Array<Maybe<ResolversTypes['CollectionItem']>>>, ParentType, ContextType, Partial<CollectionItemsArgs>>;
  labels?: Resolver<Array<ResolversTypes['String']>, ParentType, ContextType>;
  modified?: Resolver<ResolversTypes['Date'], ParentType, ContextType>;
  name?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
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

export type MetadataResolvers<ContextType = any, ParentType extends ResolversParentTypes['Metadata'] = ResolversParentTypes['Metadata']> = ResolversObject<{
  attributes?: Resolver<Array<ResolversTypes['Attribute']>, ParentType, ContextType>;
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
  sourceId?: Resolver<Maybe<ResolversTypes['String']>, ParentType, ContextType>;
  sourceIdentifier?: Resolver<Maybe<ResolversTypes['String']>, ParentType, ContextType>;
  supplementaries?: Resolver<Array<ResolversTypes['Supplementary']>, ParentType, ContextType>;
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

export type MutationResolvers<ContextType = any, ParentType extends ResolversParentTypes['Mutation'] = ResolversParentTypes['Mutation']> = ResolversObject<{
  addMetadata?: Resolver<Maybe<ResolversTypes['Metadata']>, ParentType, ContextType, RequireFields<MutationAddMetadataArgs, 'metadata'>>;
  login?: Resolver<Maybe<ResolversTypes['String']>, ParentType, ContextType, RequireFields<MutationLoginArgs, 'password' | 'username'>>;
  setMetadataJSONContent?: Resolver<Maybe<ResolversTypes['Metadata']>, ParentType, ContextType, RequireFields<MutationSetMetadataJsonContentArgs, 'id'>>;
  setMetadataReady?: Resolver<Maybe<ResolversTypes['Metadata']>, ParentType, ContextType, RequireFields<MutationSetMetadataReadyArgs, 'id'>>;
  setMetadataTextContent?: Resolver<Maybe<ResolversTypes['Metadata']>, ParentType, ContextType, RequireFields<MutationSetMetadataTextContentArgs, 'id'>>;
  setPassword?: Resolver<Maybe<ResolversTypes['Boolean']>, ParentType, ContextType, RequireFields<MutationSetPasswordArgs, 'password'>>;
}>;

export type QueryResolvers<ContextType = any, ParentType extends ResolversParentTypes['Query'] = ResolversParentTypes['Query']> = ResolversObject<{
  collection?: Resolver<Maybe<ResolversTypes['Collection']>, ParentType, ContextType, RequireFields<QueryCollectionArgs, 'id'>>;
  find?: Resolver<Maybe<ResolversTypes['Find']>, ParentType, ContextType>;
  metadata?: Resolver<Maybe<ResolversTypes['Metadata']>, ParentType, ContextType, RequireFields<QueryMetadataArgs, 'id'>>;
  source?: Resolver<Maybe<ResolversTypes['Source']>, ParentType, ContextType, RequireFields<QuerySourceArgs, 'id'>>;
  sources?: Resolver<Array<ResolversTypes['Source']>, ParentType, ContextType>;
  trait?: Resolver<Maybe<ResolversTypes['Trait']>, ParentType, ContextType, RequireFields<QueryTraitArgs, 'id'>>;
  traits?: Resolver<Array<ResolversTypes['Trait']>, ParentType, ContextType>;
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
  id?: Resolver<ResolversTypes['ID'], ParentType, ContextType>;
  name?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
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

export type Resolvers<ContextType = any> = ResolversObject<{
  Attribute?: AttributeResolvers<ContextType>;
  Collection?: CollectionResolvers<ContextType>;
  CollectionItem?: CollectionItemResolvers<ContextType>;
  Date?: GraphQLScalarType;
  Find?: FindResolvers<ContextType>;
  JSON?: GraphQLScalarType;
  JSONObject?: GraphQLScalarType;
  Metadata?: MetadataResolvers<ContextType>;
  MetadataContent?: MetadataContentResolvers<ContextType>;
  MetadataWorkflowState?: MetadataWorkflowStateResolvers<ContextType>;
  Mutation?: MutationResolvers<ContextType>;
  Query?: QueryResolvers<ContextType>;
  SignedUrl?: SignedUrlResolvers<ContextType>;
  SignedUrlHeader?: SignedUrlHeaderResolvers<ContextType>;
  Source?: SourceResolvers<ContextType>;
  Supplementary?: SupplementaryResolvers<ContextType>;
  Trait?: TraitResolvers<ContextType>;
}>;

