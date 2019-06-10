// add to handler.js
const AWS = require('aws-sdk');
const dynamoDb = new AWS.DynamoDB.DocumentClient();

const { ApolloServer, gql } = require('apollo-server-lambda');

// Construct a schema, using GraphQL schema language
const typeDefs = gql`
  type Query {
    site(id: ID!): Site
  }

  type Site {
    id: ID!
    ownerGroup: ID!
    thumbnail: ID!
    config: SiteConfig!
    rootFragmentId: ID!
  }

  type SiteConfig {
    tracking: TrackingConfiguration!
    ads: AdsConfiguration!
    colors: SiteColors!
    meta: MetaInformation!
    opengraph: OpenGraphInformation!
  }

  type MetaInformation {
    name: String!
    title: String!
    description: String!
    image: ID!
  }

  type TrackingConfiguration {
    googleAnalyticsId: ID!
    facebookPixelId: ID!
  }

  type AdsConfiguration {
    id: ID!
    type: AdsProviderType!
  }

  enum AdsProviderType {
    Adsense
  }

  type AdsenseConfiguration implements AdsConfiguration {
    id: ID!
    type: AdsProviderType!
    client: String!
    slot: String!
    host: String?
    test: String?
    origin: String?
    language: String?
  }

  type SiteColors {
    primary: Color!
    secondary: Color!
    success: Color!
    danger: Color!
    warning: Color!
    info: Color!
  }

  type Color {
    value: String!
    alpha: Integer!
  }

  type Page {
    id: ID!
    siteId: ID!
    config: Config!
    slots: [ID]! 
  }

  type FragmentInstance {
    id: ID!
    siteId: ID!
    templateId: ID!
    data: FragmentInstanceData!
    slots: [FragmentInstanceSlot]!
  }

  type FragmentInstanceData {
    id: ID!
    FragmentTemplateParameterId: ID!
    key: String!
    value: String!
    name: String!
    dynamic: Boolean!
  }

  type FragmentInstanceSlot {
    id: ID!
    fragmentTempateSlotId: ID!
    fragments: [FragmentInstance]!
  }

  type FragmentTemplate {
    id: ID!
    type: FragmentTemplateType!
    html: String!
    css: String?
    js: [JSLibrary]
    params: [FragmentTemplateParameter]!
    slot: [FragmentTemplateSlot]
  }

  enum FragmentTemplateType {
    MENU
    LAYOUT
    CONTENT
    PAGE_CONTENT
  }

  type JSLibrary {
    id: ID!
    name: String!
    url: String!
  }
  
  type FragmentTemplateSlot {
    id: ID!
    name: String!
    multiple: Boolean!
  }

  type FragmentTemplateParameter {
    id: ID!
    name: String!
    type: FragmentTemplateParameterType! 
    multiple: Boolean!
  }

  enum FragmentTemplateParameterType {
    IMAGE
    TEXT
    NUMBER
    LINK
  }

  type Image {
    id: ID!
    siteId: ID!
    description: String?
    url: String!
    width: Integer!
    height: Integer!
    original: Boolean!
    createdFromId: ID?
  }

  type Mutation {

  }
`;

// Provide resolver functions for your schema fields
const resolvers = {
  Query: {
    hello: () => 'Hello world!',
  },
};

const server = new ApolloServer({
  typeDefs,
  resolvers,
  context: ({ event, context }) => ({
    headers: event.headers,
    functionName: context.functionName,
    event,
    context,
  }),
});

exports.graphqlHandler = server.createHandler({
  cors: {
    origin: true,
    credentials: true
  },
});