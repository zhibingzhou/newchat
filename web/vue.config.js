const path = require('path')

// compression-webpack-plugin，添加压缩文件类型
const CompressionWebpackPlugin = require('compression-webpack-plugin')
const productionGzipExtensions = ['js', 'css']
const isProd = process.env.NODE_ENV === 'production'

function resolve(dir) {
  return path.join(__dirname, dir)
}

const assetsCDN = {
  externals: {
    vue: 'Vue',
    'vue-router': 'VueRouter',
    vuex: 'Vuex',
    axios: 'axios',
  },
  css: [

  ],
  js: [
    '//cdn.jsdelivr.net/npm/vue@2.6.10/dist/vue.min.js',
    '//cdn.jsdelivr.net/npm/vue-router@3.4.9/dist/vue-router.min.js',
    '//cdn.jsdelivr.net/npm/vuex@3.5.1/dist/vuex.min.js',
    '//cdn.jsdelivr.net/npm/axios@0.19.0/dist/axios.min.js',
  ]
}

// vue.config.js
const vueConfig = {
  // 公共路径(必须有的)
  publicPath: "/",
  // 输出文件目录
  outputDir: "dist",
  // 静态资源存放的文件夹(相对于ouputDir)
  assetsDir: "static",
  runtimeCompiler: false,
  devServer: {
    port: 8000,
    proxy: {
      // 把key的路径代理到target位置
      // detail: https://cli.vuejs.org/config/#devserver-proxy
      [process.env.VUE_APP_BASE_API]: { //需要代理的路径   例如 '/api'
        target: `http://127.0.0.1:8888`, //代理到 目标路径
        changeOrigin: true,
        pathRewrite: { // 修改路径数据
          ['^' + process.env.VUE_APP_BASE_API]: '' // 举例 '^/api:""' 把路径中的/api字符串删除
        }
      }
    },
  },
  configureWebpack: {
    // webpack plugins
    plugins: [
      // 配置大文件压缩相关参数
      new CompressionWebpackPlugin({
        filename: '[path][base].gz',
        algorithm: 'gzip',
        test: new RegExp('\\.(' + productionGzipExtensions.join('|') + ')$'),
        threshold: 10240,
        minRatio: 0.8
      })
    ],

    // if prod, add externals
    externals: isProd ? assetsCDN.externals : {},
  },

  chainWebpack: (config) => {
    config.resolve.alias
      .set('@', resolve('src'))

    const svgRule = config.module.rule('svg')
    svgRule.uses.clear()
    svgRule
      .oneOf('inline')
      .resourceQuery(/inline/)
      .use('vue-svg-icon-loader')
      .loader('vue-svg-icon-loader')
      .end()
      .end()
      .oneOf('external')
      .use('file-loader')
      .loader('file-loader')
      .options({
        name: 'assets/[name].[hash:8].[ext]'
      })

    // if prod is on
    // assets require on cdn
    if (isProd) {
      config.plugin('html').tap(args => {
        args[0].cdn = assetsCDN
        return args
      })
    }
  },
  pluginOptions: {
    "style-resources-loader": {
      preProcessor: "less",
      patterns: [
        //全局加载 less 变量
        path.resolve(__dirname, "./src/assets/css/variable.less")
      ]
    }
  },
  css: {
    loaderOptions: {
      less: {
        lessOptions: {
          modifyVars: {},
          javascriptEnabled: true
        }
      }
    }
  },

  // 是否为生产环境构建生成 source map？
  productionSourceMap: false,
  // lintOnSave: undefined,
  lintOnSave: false,
  // babel-loader no-ignore node_modules/*
  transpileDependencies: []
}

module.exports = vueConfig