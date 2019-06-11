const { gql } = require('apollo-server');

module.exports = {
  typeDefs: gql`

    type SiteTemplate {
      id: ID!
      fragments: [FragmentInstance]!
      slotFragments: [FragmentInstanceSlotFragments]!
      contentSlot: FragmentInstanceSlot!
    }

    type PageTemplate {
      id: ID!
      fragments: [FragmentInstance]!
      slotFragments: [FragmentInstanceSlotFragments]!
    }

    type FragmentInstance {
      id: ID!
      siteId: ID!
      template: FragmentTemplate!
      data: FragmentInstanceData!
      slots: [FragmentInstanceSlot]!
    }

    type FragmentInstanceData {
      params: [FragmentInstanceParameter]!
    }

    type FragmentInstanceParameter {
      id: ID!
      templateParameter: FragmentTemplateParameter!
      key: String!
      value: String!
      name: String!
      dynamic: Boolean!
    }

    type FragmentInstanceSlot {
      id: ID!
      fragmentTemplateSlot: FragmentTemplateSlot!
    }

    type FragmentInstanceSlotFragments {
      fragmentInstanceSlot: FragmentInstanceSlot!
      fragment: [FragmentInstance]!
    }

    type FragmentTemplate {
      id: ID!
      type: FragmentTemplateType!
      html: String!
      css: String
      js: [JSLibrary]
      params: [FragmentTemplateParameter]!
      slot: [FragmentTemplateSlot]
    }

    enum FragmentTemplateType {
      MENU
      SITE_LAYOUT
      LAYOUT
      CONTENT
    }

    type JSLibrary {
      id: ID!
      name: String!
      url: String!
    }

    type FragmentTemplateSlot {
      id: ID!
      name: String!
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
  `,

  resolvers: {
    
  }
};