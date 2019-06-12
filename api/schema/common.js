const { gql } = require('apollo-server-lambda');

module.exports = {
  typeDefs: gql`

    type Query {
      _empty: String
    }

    type Color {
      value: String!
      alpha: Int!
    }

    type Image {
      id: ID!
      siteId: ID!
      description: String
      url: String!
      width: Int!
      height: Int!
      original: Boolean!
      thumbnail: Image!
      createdFromId: ID
    }
  `,

  resolvers: {}
};
