// vue.config.js
module.exports = {
    configureWebpack: {
        resolve: {
            alias: {
                // Use compiled pica files from /dist folder
                pica: 'pica/dist/pica.js',
            },
        }
    }
}