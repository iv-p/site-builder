const Site = require('./site');
const Page = require('./page');
const Fragment = require('./fragment');
const Common = require('./common');

module.exports = {
  typeDefs: [ Common.typeDefs, Site.typeDefs, Page.typeDefs, Fragment.typeDefs ],
  resolvers: {
    ...Common.resolvers,
    ...Site.resolvers,
    ...Page.resolvers,
    ...Fragment.resolvers
  }
}
