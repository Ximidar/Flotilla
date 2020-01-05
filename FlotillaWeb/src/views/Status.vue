<template>
  <div class="about">
    <h1>Status</h1>
    <p>Status go here</p>
  </div>
</template>

<script>
export default {
  name: 'FlotillaStatus',
  methods: {
    Hello () {
      console.log(Go)
      var go = new Go()
      console.log(go)
      var hellowasm = require("@/wasm_test/hellowasm.wasm")
      WebAssembly.instantiateStreaming(fetch(hellowasm), go.importObject).then((result) => {
        console.log("Result!!!")
        console.log(result)
        
        go.run(result.instance)
      })
      
      
    }
  },
  created(){
    this.Hello()
  }
}
</script>

<style>
.about{
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    color: White;
}
</style>