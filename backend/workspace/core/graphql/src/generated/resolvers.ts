import { GraphQLResolveInfo, GraphQLScalarType, GraphQLScalarTypeConfig } from 'graphql';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
export type MakeEmpty<T extends { [key: string]: unknown }, K extends keyof T> = { [_ in K]?: never };
export type Incremental<T> = T | { [P in keyof T]?: P extends ' $fragmentName' | '__typename' ? T[P] : never };
export type RequireFields<T, K extends keyof T> = Omit<T, K> & { [P in K]-?: NonNullable<T[P]> };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: { input: string; output: string; }
  String: { input: string; output: string; }
  Boolean: { input: boolean; output: boolean; }
  Int: { input: number; output: number; }
  Float: { input: number; output: number; }
  Date: { input: any; output: any; }
};

export type Metadata = {
  __typename?: 'Metadata';
  attributes: Array<MetadataAttribute>;
  contentLength?: Maybe<Scalars['Int']['output']>;
  contentType: Scalars['String']['output'];
  created: Scalars['Date']['output'];
  downloadUrl: SignedUrl;
  id: Scalars['ID']['output'];
  labels: Array<Scalars['String']['output']>;
  languageTag: Scalars['String']['output'];
  modified: Scalars['Date']['output'];
  name: Scalars['String']['output'];
  parentId?: Maybe<Scalars['String']['output']>;
  sourceId?: Maybe<Scalars['String']['output']>;
  sourceIdentifier?: Maybe<Scalars['String']['output']>;
  uploadUrl: SignedUrl;
  workflowState: MetadataWorkflowState;
};

export type MetadataAttribute = {
  __typename?: 'MetadataAttribute';
  name: Scalars['String']['output'];
  value: Scalars['String']['output'];
};

export type MetadataInput = {
  contentLength?: InputMaybe<Scalars['Int']['input']>;
  contentType: Scalars['String']['input'];
  languageTag: Scalars['String']['input'];
  name: Scalars['String']['input'];
};

export type MetadataWorkflowState = {
  __typename?: 'MetadataWorkflowState';
  deleteWorkflowId?: Maybe<Scalars['String']['output']>;
  id: Scalars['String']['output'];
  pendingId?: Maybe<Scalars['String']['output']>;
};

export type Mutation = {
  __typename?: 'Mutation';
  addMetadata?: Maybe<Metadata>;
};


export type MutationAddMetadataArgs = {
  metadata: MetadataInput;
};

export type Query = {
  __typename?: 'Query';
  metadata?: Maybe<Metadata>;
  sources: Array<Source>;
};


export type QueryMetadataArgs = {
  id: Scalars['ID']['input'];
};

export type SignedUrl = {
  __typename?: 'SignedUrl';
  headers: Array<SignedUrlHeader>;
  id: Scalars['String']['output'];
  method: Scalars['String']['output'];
  url: Scalars['String']['output'];
};

export type SignedUrlHeader = {
  __typename?: 'SignedUrlHeader';
  name: Scalars['String']['output'];
  value: Scalars['String']['output'];
};

export type Source = {
  __typename?: 'Source';
  id: Scalars['ID']['output'];
  name: Scalars['String']['output'];
};

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



/** Mapping between all available schema types and the resolvers types */
export type ResolversTypes = ResolversObject<{
  Boolean: ResolverTypeWrapper<Scalars['Boolean']['output']>;
  Date: ResolverTypeWrapper<Scalars['Date']['output']>;
  ID: ResolverTypeWrapper<Scalars['ID']['output']>;
  Int: ResolverTypeWrapper<Scalars['Int']['output']>;
  Metadata: ResolverTypeWrapper<Metadata>;
  MetadataAttribute: ResolverTypeWrapper<MetadataAttribute>;
  MetadataInput: MetadataInput;
  MetadataWorkflowState: ResolverTypeWrapper<MetadataWorkflowState>;
  Mutation: ResolverTypeWrapper<{}>;
  Query: ResolverTypeWrapper<{}>;
  SignedUrl: ResolverTypeWrapper<SignedUrl>;
  SignedUrlHeader: ResolverTypeWrapper<SignedUrlHeader>;
  Source: ResolverTypeWrapper<Source>;
  String: ResolverTypeWrapper<Scalars['String']['output']>;
}>;

/** Mapping between all available schema types and the resolvers parents */
export type ResolversParentTypes = ResolversObject<{
  Boolean: Scalars['Boolean']['output'];
  Date: Scalars['Date']['output'];
  ID: Scalars['ID']['output'];
  Int: Scalars['Int']['output'];
  Metadata: Metadata;
  MetadataAttribute: MetadataAttribute;
  MetadataInput: MetadataInput;
  MetadataWorkflowState: MetadataWorkflowState;
  Mutation: {};
  Query: {};
  SignedUrl: SignedUrl;
  SignedUrlHeader: SignedUrlHeader;
  Source: Source;
  String: Scalars['String']['output'];
}>;

export interface DateScalarConfig extends GraphQLScalarTypeConfig<ResolversTypes['Date'], any> {
  name: 'Date';
}

export type MetadataResolvers<ContextType = any, ParentType extends ResolversParentTypes['Metadata'] = ResolversParentTypes['Metadata']> = ResolversObject<{
  attributes?: Resolver<Array<ResolversTypes['MetadataAttribute']>, ParentType, ContextType>;
  contentLength?: Resolver<Maybe<ResolversTypes['Int']>, ParentType, ContextType>;
  contentType?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  created?: Resolver<ResolversTypes['Date'], ParentType, ContextType>;
  downloadUrl?: Resolver<ResolversTypes['SignedUrl'], ParentType, ContextType>;
  id?: Resolver<ResolversTypes['ID'], ParentType, ContextType>;
  labels?: Resolver<Array<ResolversTypes['String']>, ParentType, ContextType>;
  languageTag?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  modified?: Resolver<ResolversTypes['Date'], ParentType, ContextType>;
  name?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  parentId?: Resolver<Maybe<ResolversTypes['String']>, ParentType, ContextType>;
  sourceId?: Resolver<Maybe<ResolversTypes['String']>, ParentType, ContextType>;
  sourceIdentifier?: Resolver<Maybe<ResolversTypes['String']>, ParentType, ContextType>;
  uploadUrl?: Resolver<ResolversTypes['SignedUrl'], ParentType, ContextType>;
  workflowState?: Resolver<ResolversTypes['MetadataWorkflowState'], ParentType, ContextType>;
  __isTypeOf?: IsTypeOfResolverFn<ParentType, ContextType>;
}>;

export type MetadataAttributeResolvers<ContextType = any, ParentType extends ResolversParentTypes['MetadataAttribute'] = ResolversParentTypes['MetadataAttribute']> = ResolversObject<{
  name?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  value?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
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
}>;

export type QueryResolvers<ContextType = any, ParentType extends ResolversParentTypes['Query'] = ResolversParentTypes['Query']> = ResolversObject<{
  metadata?: Resolver<Maybe<ResolversTypes['Metadata']>, ParentType, ContextType, RequireFields<QueryMetadataArgs, 'id'>>;
  sources?: Resolver<Array<ResolversTypes['Source']>, ParentType, ContextType>;
}>;

export type SignedUrlResolvers<ContextType = any, ParentType extends ResolversParentTypes['SignedUrl'] = ResolversParentTypes['SignedUrl']> = ResolversObject<{
  headers?: Resolver<Array<ResolversTypes['SignedUrlHeader']>, ParentType, ContextType>;
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

export type Resolvers<ContextType = any> = ResolversObject<{
  Date?: GraphQLScalarType;
  Metadata?: MetadataResolvers<ContextType>;
  MetadataAttribute?: MetadataAttributeResolvers<ContextType>;
  MetadataWorkflowState?: MetadataWorkflowStateResolvers<ContextType>;
  Mutation?: MutationResolvers<ContextType>;
  Query?: QueryResolvers<ContextType>;
  SignedUrl?: SignedUrlResolvers<ContextType>;
  SignedUrlHeader?: SignedUrlHeaderResolvers<ContextType>;
  Source?: SourceResolvers<ContextType>;
}>;

