const webpack = require('webpack');
const path = require('path');

module.exports = {
  entry: "./src/index.ts",
  mode: "development",
  output: {
    path: path.resolve(__dirname, 'dist'),
    filename: 'bundle.js'
  },
  devtool: 'inline-source-map',
  module: {
    rules: [
      {
        test: /\.ts$/,
        include: /src/,
        exclude: /node_modules/,
        loader: "ts-loader"
      }
    ]
  },
  resolve: {
    alias: {
      "cmd": path.resolve(__dirname, '../../cmd'),
      "pkg": path.resolve(__dirname, '../../pkg'),
    },
    extensions: [".ts", ".js"],
    modules: [path.resolve(__dirname, 'node_modules'), 'node_modules']
  },
  plugins: [
    new webpack.DefinePlugin({
      'USE_TLS': process.env.USE_TLS !== undefined
    })
  ]
};
