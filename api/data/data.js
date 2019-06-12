module.exports = {
  fragmentInstance: {
    "fragmentInstance1": {
      id: "fragmentInstance1",
      siteId: "site1",
      templateId: "fragmentTemplate1",
      data: {
        "UUID-for-width-param": "value"
      },
      slots: {
        "header": "UUID-for-header-slot"
      },
      params: {
        "width": {
          id: "UUID-for-width-param",
          dynamic: false
        },
        "height": {
          id: "UUID-for-height-param",
          dynamic: true
        }
      }
    }
  },
  fragmentTemplate: {
    "fragmentTemplate1": {
      id: "fragmentTemplate1",
      html: "html",
      css: {
        global: ".{{ class }}",
        scoped: ".this-is-scoped { color: blue; }",
        files: ["fontawesome"]
      },
      js: {
        global: "",
        scoped: "",
        files: [""]
      },
      slots: {
        name: "header"
      },
      params: [
        {
          name: "width",
          type: "int",
          default: "32"
        },
        {
          name: "height",
          type: "int",
          default: "32"
        }
      ]
    }
  },
  template: {
    id: "template1",
    siteId: "site1",
    rootFragmentId: "fragmentTemplate1",
    fragmentIds: ["fragmentTemplate1"],
    structure: {
      fragmentTemplate1: {
        "UUID-for-header-slot":  "fragmentTemplate2"
      }
    },
    type: "page" // "site"
  },
  page: {
    id: "page1",
    siteId: "site1",
    template: "template1",
    content: {
      "UUID-for-height-param": "64"
    },
    config: {
      // misc config, meta
    }
  },
  site: {
    id: "site1",
    ownerGroup: "group1",
    template: "layout1",
    content: {
      "UUID-for-height-param": "64"
    },
    config: {

    }
  }
}
