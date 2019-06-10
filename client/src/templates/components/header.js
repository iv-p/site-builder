const layout = {
  name: "Heading",
  html: `
    <h1>{{ content }}</h1>
    `,
  css: `
        h1 {
            font-size: 2rem;
        }
    `,
  js: [],
  props: {
    content: {
      name: "content",
      label: "Text",
      type: "text",
      default: "defalut header"
    }
  }
};

export default layout;
