module.exports = {
    chainWebpack: (config) => {
      const svgRule = config.module.rule('svg');
  
      svgRule.uses.clear();
  
      svgRule
        .use('vue-svg-loader')
        .loader('vue-svg-loader');
        
      
      const wasmRule = config.module.rule('wasm')
      wasmRule.uses.clear()
      wasmRule
        .type('javascript/auto')
        .test(/\.wasm$/)
        .use('file-loader')
        .loader('file-loader')
      
      
      const rxlite = config.module.rule('./rx.lite');
      rxlite.uses.clear();
      rxlite
        .test(/rx\.lite\.aggregates\.js/)
        .use('imports-loader?define=>false')
        .loader('imports-loader?define=>false');
    },
    
  };
  