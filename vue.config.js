const path = require('path');

module.exports = {
    publicPath: "/",
    outputDir: 'web/dist',
    pages: {
        index: {
            entry: "web/src/main.js",
        }
    },
    configureWebpack: {
        performance: {
            maxEntrypointSize: 500000000,
            maxAssetSize: 3000000
        }
    }
}