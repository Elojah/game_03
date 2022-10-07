const webpack = require('webpack');
const path = require('path');
const Dotenv = require('dotenv-webpack');

module.exports = {
	entry: "./src/app.tsx",
	mode: "development",
	output: {
		path: path.resolve(__dirname, 'dist'),
		filename: 'bundle.js'
	},
	devtool: 'inline-source-map',
	module: {
		rules: [
			{
				test: /\.tsx$/,
				include: /src/,
				exclude: /node_modules/,
				loader: "ts-loader",
			},
			{
				test: /\.ts$/,
				include: /src/,
				exclude: /node_modules/,
				loader: "ts-loader",
			},
			{
				test: /\.css$/i,
				use: ["style-loader", "css-loader"],
			},
			{
				test: /\.s[ac]ss$/i,
				use: [
					"style-loader",
					"css-loader",
					"sass-loader",
				],
				include: /src/,
			},
		]
	},
	resolve: {
		alias: {
			"cmd": path.resolve(__dirname, '../../cmd'),
			"pkg": path.resolve(__dirname, '../../pkg'),
		},
		extensions: [".ts", ".tsx", ".js", ".json", ".css", ".scss", ".sass"],
		modules: [path.resolve(__dirname, 'node_modules'), 'node_modules']
	},
	plugins: [
		new Dotenv({
			path: path.resolve(__dirname, '../../config/browser/.env'),
		}),
	],
};
