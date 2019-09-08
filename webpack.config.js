const path = require("path");
const HtmlWebPackPlugin = require("html-webpack-plugin");

module.exports = {
    mode: 'development',
    module: {
        rules: [
            {
                test: /\.(js|jsx)$/,
                exclude: /node_modules/,
                use: {
                    loader: "babel-loader"
                }
            },
            {
                test: /\.html$/,
                use: [
                    {
                        loader: "html-loader"
                    }
                ]
            }
        ]
    },
    output: {
        // the output bundle
        filename: 'build.js',
        // saves the files into the dist/static folder
        path: path.resolve(__dirname, 'views', 'components'),
        // set static as src="static/main.js as relative path
        publicPath: 'components/'
    },
    // plugins: [
    //     new HtmlWebPackPlugin({
    //         template: path.resolve(__dirname, 'views', 'test.html'),
    //         filename: path.resolve(__dirname, 'views', 'index1.html')
    //     })
    // ]

};