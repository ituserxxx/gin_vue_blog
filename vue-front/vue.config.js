module.exports = {
    devServer: {
      proxy: {
        '/blog': {
          target: 'http://0.0.0.0:6008',
          ws: true,
          changeOrigin: true
        },
      
      }
    }
  }