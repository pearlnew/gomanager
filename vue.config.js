const path = require('path');

module.exports = {
    publicPath: "/",
    outputDir: 'www/dist',
    pages: {
        index: {
            entry: "www/src/main.js",
        }
    },
    configureWebpack: {
        performance: {
            maxEntrypointSize: 500000000,
            maxAssetSize: 3000000
        }
    }
}