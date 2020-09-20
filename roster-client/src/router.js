import Vue from "vue"
import Router from "vue-router"

Vue.use(Router)

export default new Router({
  mode: "history", 
  routes: [
    {
      path: "/",
      alias: "/rosters",
      name: "rosters",
      component: () => import("./components/RostersList"),
    },
    {
      path: "/roster/:date",
      name: "roster-details",
      component: () => import("./components/Roster"),
    },
    {
      path: "/addroster",
      name: "addroster",
      component: () => import("./components/AddRoster"),
    },

    {
      path: "/products",
      name: "products",
      component: () => import("./components/ProductsList"),
    },

    {
      path: "/addProduct",
      name: "addProduct",
      component: () => import("./components/AddProduct"),
    },
    {
      path: "/product/:id",
      name: "product-details",
      component: () => import("./components/Product"),
    },
    {
      path: "/cart/",
      name: "product-cart",
      component: () => import("./components/Cart"),
    },
    /*{
        path: "/add",
        name: "add",
        component: () => import("./components/AddTutorial")
      }, */
  ],
})
