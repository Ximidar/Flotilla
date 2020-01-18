import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import Go from "@/wasm_test/go_wasm_exec"

Vue.config.productionTip = false

new Vue({
  router,
  store,
  Go,
  render: h => h(App)
}).$mount('#app')
