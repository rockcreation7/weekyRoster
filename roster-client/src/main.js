import Vue from "vue"
import App from "./App.vue"
import router from "./router"
import { MdTable, MdIcon, MdButton, MdList } from "vue-material/dist/components"
import "vue-material/dist/vue-material.min.css"
import "vue-material/dist/theme/default.css"
MdIcon

Vue.config.productionTip = false
Vue.use(MdTable)
Vue.use(MdIcon)
Vue.use(MdButton)
Vue.use(MdList) 
new Vue({
  router,
  render: (h) => h(App),
}).$mount("#app")
