const { gql } = require('apollo-server');

module.exports = {
  typeDefs: gql`

    type Page {
      id: ID!
      siteId: ID!
      config: PageConfig!
      template: PageTemplate!
    }

    type PageConfig {
      meta: PageMetaInformation!
    }

    type PageMetaInformation {
      name: String!
      title: String!
      description: String!
      image: Image!
    }
  `,

  resolvers: {}
};
