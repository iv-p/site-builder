const { gql } = require('apollo-server-lambda');

module.exports = {
  typeDefs: gql`

    extend type Query {
      site(id: ID!): Site
    }

    type Site {
      id: ID!
      ownerGroup: ID!
      config: SiteConfig!
      template: SiteTemplate!
    }

    type SiteConfig {
      tracking: TrackingConfiguration!
      ads: AdsConfiguration!
      colors: SiteColors!
      meta: SiteMetaInformation!
    }

    type SiteMetaInformation {
      name: String!
    }

    type TrackingConfiguration {
      googleAnalyticsId: ID!
      facebookPixelId: ID!
    }

    type AdsConfiguration {
      adsense: AdsenseConfiguration
    }

    type AdsenseConfiguration {
      id: ID!
      client: String!
      slot: String!
      host: String
      test: String
      origin: String
      language: String
    }

    type SiteColors {
      primary: Color!
      secondary: Color!
      success: Color!
      danger: Color!
      warning: Color!
      info: Color!
    }
  `,

  resolvers: {}
};