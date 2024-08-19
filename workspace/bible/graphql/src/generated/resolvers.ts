import { GraphQLResolveInfo } from 'graphql';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
export type MakeEmpty<T extends { [key: string]: unknown }, K extends keyof T> = { [_ in K]?: never };
export type Incremental<T> = T | { [P in keyof T]?: P extends ' $fragmentName' | '__typename' ? T[P] : never };
export type RequireFields<T, K extends keyof T> = Omit<T, K> & { [P in K]-?: NonNullable<T[P]> };
/** All built-in and custom scalars, mapped to their actual values */
export interface Scalars {
  ID: { input: string; output: string; }
  String: { input: string; output: string; }
  Boolean: { input: boolean; output: boolean; }
  Int: { input: number; output: number; }
  Float: { input: number; output: number; }
}

export interface BibleMetadata {
  __typename?: 'BibleMetadata';
  abbreviation: Scalars['String']['output'];
  books: Array<BookMetadata>;
  id: Scalars['ID']['output'];
  language: Scalars['String']['output'];
  name: Scalars['String']['output'];
  systemId: Scalars['ID']['output'];
  version: Scalars['String']['output'];
}

export interface BookMetadata {
  __typename?: 'BookMetadata';
  chapters: Array<ChapterMetadata>;
  id: Scalars['ID']['output'];
  name: Scalars['String']['output'];
  usfm: Scalars['String']['output'];
}

export interface Chapter {
  __typename?: 'Chapter';
  html?: Maybe<Scalars['String']['output']>;
  id: Scalars['ID']['output'];
  name: Scalars['String']['output'];
  number?: Maybe<Scalars['String']['output']>;
  usfm: Scalars['String']['output'];
  usx?: Maybe<Scalars['String']['output']>;
}

export interface ChapterMetadata {
  __typename?: 'ChapterMetadata';
  id: Scalars['ID']['output'];
  name: Scalars['String']['output'];
  usfm: Scalars['String']['output'];
}

export interface Query {
  __typename?: 'Query';
  bibles: Array<BibleMetadata>;
  chapter?: Maybe<Chapter>;
  verse?: Maybe<Verse>;
  verses?: Maybe<Array<Maybe<Verse>>>;
}


export interface QueryChapterArgs {
  systemId: Scalars['ID']['input'];
  usfm: Scalars['String']['input'];
  version: Scalars['String']['input'];
}


export interface QueryVerseArgs {
  systemId: Scalars['ID']['input'];
  usfm: Scalars['String']['input'];
  version: Scalars['String']['input'];
}


export interface QueryVersesArgs {
  systemId: Scalars['ID']['input'];
  usfm: Array<Scalars['String']['input']>;
  version: Scalars['String']['input'];
}

export interface Verse {
  __typename?: 'Verse';
  html?: Maybe<Scalars['String']['output']>;
  id: Scalars['ID']['output'];
  name: Scalars['String']['output'];
  number?: Maybe<Scalars['String']['output']>;
  text?: Maybe<Scalars['String']['output']>;
  usfm: Scalars['String']['output'];
  usx?: Maybe<Scalars['String']['output']>;
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



/** Mapping between all available schema types and the resolvers types */
export type ResolversTypes = ResolversObject<{
  BibleMetadata: ResolverTypeWrapper<BibleMetadata>;
  BookMetadata: ResolverTypeWrapper<BookMetadata>;
  Boolean: ResolverTypeWrapper<Scalars['Boolean']['output']>;
  Chapter: ResolverTypeWrapper<Chapter>;
  ChapterMetadata: ResolverTypeWrapper<ChapterMetadata>;
  ID: ResolverTypeWrapper<Scalars['ID']['output']>;
  Query: ResolverTypeWrapper<{}>;
  String: ResolverTypeWrapper<Scalars['String']['output']>;
  Verse: ResolverTypeWrapper<Verse>;
}>;

/** Mapping between all available schema types and the resolvers parents */
export type ResolversParentTypes = ResolversObject<{
  BibleMetadata: BibleMetadata;
  BookMetadata: BookMetadata;
  Boolean: Scalars['Boolean']['output'];
  Chapter: Chapter;
  ChapterMetadata: ChapterMetadata;
  ID: Scalars['ID']['output'];
  Query: {};
  String: Scalars['String']['output'];
  Verse: Verse;
}>;

export type BibleMetadataResolvers<ContextType = any, ParentType extends ResolversParentTypes['BibleMetadata'] = ResolversParentTypes['BibleMetadata']> = ResolversObject<{
  abbreviation?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  books?: Resolver<Array<ResolversTypes['BookMetadata']>, ParentType, ContextType>;
  id?: Resolver<ResolversTypes['ID'], ParentType, ContextType>;
  language?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  name?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  systemId?: Resolver<ResolversTypes['ID'], ParentType, ContextType>;
  version?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  __isTypeOf?: IsTypeOfResolverFn<ParentType, ContextType>;
}>;

export type BookMetadataResolvers<ContextType = any, ParentType extends ResolversParentTypes['BookMetadata'] = ResolversParentTypes['BookMetadata']> = ResolversObject<{
  chapters?: Resolver<Array<ResolversTypes['ChapterMetadata']>, ParentType, ContextType>;
  id?: Resolver<ResolversTypes['ID'], ParentType, ContextType>;
  name?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  usfm?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  __isTypeOf?: IsTypeOfResolverFn<ParentType, ContextType>;
}>;

export type ChapterResolvers<ContextType = any, ParentType extends ResolversParentTypes['Chapter'] = ResolversParentTypes['Chapter']> = ResolversObject<{
  html?: Resolver<Maybe<ResolversTypes['String']>, ParentType, ContextType>;
  id?: Resolver<ResolversTypes['ID'], ParentType, ContextType>;
  name?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  number?: Resolver<Maybe<ResolversTypes['String']>, ParentType, ContextType>;
  usfm?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  usx?: Resolver<Maybe<ResolversTypes['String']>, ParentType, ContextType>;
  __isTypeOf?: IsTypeOfResolverFn<ParentType, ContextType>;
}>;

export type ChapterMetadataResolvers<ContextType = any, ParentType extends ResolversParentTypes['ChapterMetadata'] = ResolversParentTypes['ChapterMetadata']> = ResolversObject<{
  id?: Resolver<ResolversTypes['ID'], ParentType, ContextType>;
  name?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  usfm?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  __isTypeOf?: IsTypeOfResolverFn<ParentType, ContextType>;
}>;

export type QueryResolvers<ContextType = any, ParentType extends ResolversParentTypes['Query'] = ResolversParentTypes['Query']> = ResolversObject<{
  bibles?: Resolver<Array<ResolversTypes['BibleMetadata']>, ParentType, ContextType>;
  chapter?: Resolver<Maybe<ResolversTypes['Chapter']>, ParentType, ContextType, RequireFields<QueryChapterArgs, 'systemId' | 'usfm' | 'version'>>;
  verse?: Resolver<Maybe<ResolversTypes['Verse']>, ParentType, ContextType, RequireFields<QueryVerseArgs, 'systemId' | 'usfm' | 'version'>>;
  verses?: Resolver<Maybe<Array<Maybe<ResolversTypes['Verse']>>>, ParentType, ContextType, RequireFields<QueryVersesArgs, 'systemId' | 'usfm' | 'version'>>;
}>;

export type VerseResolvers<ContextType = any, ParentType extends ResolversParentTypes['Verse'] = ResolversParentTypes['Verse']> = ResolversObject<{
  html?: Resolver<Maybe<ResolversTypes['String']>, ParentType, ContextType>;
  id?: Resolver<ResolversTypes['ID'], ParentType, ContextType>;
  name?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  number?: Resolver<Maybe<ResolversTypes['String']>, ParentType, ContextType>;
  text?: Resolver<Maybe<ResolversTypes['String']>, ParentType, ContextType>;
  usfm?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  usx?: Resolver<Maybe<ResolversTypes['String']>, ParentType, ContextType>;
  __isTypeOf?: IsTypeOfResolverFn<ParentType, ContextType>;
}>;

export type Resolvers<ContextType = any> = ResolversObject<{
  BibleMetadata?: BibleMetadataResolvers<ContextType>;
  BookMetadata?: BookMetadataResolvers<ContextType>;
  Chapter?: ChapterResolvers<ContextType>;
  ChapterMetadata?: ChapterMetadataResolvers<ContextType>;
  Query?: QueryResolvers<ContextType>;
  Verse?: VerseResolvers<ContextType>;
}>;

