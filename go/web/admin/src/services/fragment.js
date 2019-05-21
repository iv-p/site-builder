const fragments = {
  1: {
    id: 1,
    name: 'test-fragment-name',
    html: '<h1 class="heading">{{ .Data.Heading }}</h1>',
    css: '.heading { font-size: 3rem }',
    config: {},
  },
  2: {
    id: 2,
    name: 'test-fragment-name-2',
    html: '<h1 class="heading">{{ .Data.Heading }}</h1>',
    css: '.heading { font-size: 3rem }',
    config: {},
  },
  3: {
    id: 3,
    name: 'test-fragment-name-3',
    html: '<h1 class="heading">{{ .Data.Heading }}</h1>',
    css: '.heading { font-size: 3rem }',
    config: { },
  },
};

export default {
  getFragments: async () => Promise.resolve(fragments),
  getFragment: async id => Promise.resolve(fragments[id]),
};
