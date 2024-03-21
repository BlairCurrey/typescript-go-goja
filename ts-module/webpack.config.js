// webpack.config.js
const path = require("path");

module.exports = {
  mode: "production",
  entry: "./module.ts",
  output: {
    filename: "bundle.js",
    path: path.resolve(__dirname, "dist"),
    //  Important: Call from goja like MyModule.myfunction(...)
    library: "MyModule",
    libraryTarget: "umd",
    // Resolves this error in go: panic: ReferenceError: self is not defined at bundle.js:1:228(2)
    // https://github.com/gregberge/loadable-components/issues/276
    globalObject: `typeof self !== 'undefined' ? self : this`,
  },
  resolve: {
    extensions: [".ts", ".js"],
  },
  module: {
    rules: [
      {
        test: /\.ts$/,
        use: "ts-loader",
        exclude: /node_modules/,
      },
    ],
  },
  // Dont include external libraries here, they will be excluded from bundle
  // externals: {
  //   lodash: "_",
  // },
};
