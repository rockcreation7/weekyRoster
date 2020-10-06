import Vue from "vue"
import Vuex from "vuex"

Vue.use(Vuex)
export const store = new Vuex.Store({
  state: {
    showSnackbar: false,
    position: "center",
    duration: 4000,
    isInfinity: false,
    error: "",
  },
  mutations: {
    errorMessage(state, err) {
      console.log(state, err, 'err')
      state.error = err
    },
  },
})

// https://dev.to/stephannv/how-to-create-a-global-snackbar-using-nuxt-vuetify-and-vuex-1bda
