const layout = {
  name: "Amp layout",
  html: `
    <!doctype html>
    <html ⚡>
        <head>
            <meta charset="utf-8">
            <title> Hello World</title>
            <script async src="https://cdn.ampproject.org/v0.js"></script>
            <link rel="canonical" href="<% canonical %>">
            <meta name="viewport" content="width=device-width,minimum-scale=1,initial-scale=1">
            <style amp-custom>
                h1 {
                color: red;
                }
            </style>
            <style amp-boilerplate>body{-webkit-animation:-amp-start 8s steps(1,end) 0s 1 normal both;-moz-animation:-amp-start 8s steps(1,end) 0s 1 normal both;-ms-animation:-amp-start 8s steps(1,end) 0s 1 normal both;animation:-amp-start 8s steps(1,end) 0s 1 normal both}@-webkit-keyframes -amp-start{from{visibility:hidden}to{visibility:visible}}@-moz-keyframes -amp-start{from{visibility:hidden}to{visibility:visible}}@-ms-keyframes -amp-start{from{visibility:hidden}to{visibility:visible}}@-o-keyframes -amp-start{from{visibility:hidden}to{visibility:visible}}@keyframes -amp-start{from{visibility:hidden}to{visibility:visible}}</style><noscript><style amp-boilerplate>body{-webkit-animation:none;-moz-animation:none;-ms-animation:none;animation:none}</style></noscript>
            <style amp-custom>
                {{ css }}
            </style>
        </head>
        <body>
            <h1>Hello World!</h1>
            <amp-img src="blob:http://localhost:8080/7d206a06-08d5-44b4-99dd-2c135c6a28c8" width="1080" height="610" layout="responsive"></amp-img>
            {{ header }}
            {{ content }}
            {{ footer }}
        </body>
    </html>
    `,
  css: `
        h1 {
            font-size: 2rem;
        }
    `,
  js: ["amp-image"],
  props: {}
};

export default layout;
