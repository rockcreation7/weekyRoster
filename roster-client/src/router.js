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
    /* {
        path: "/tutorials/:id",
        name: "tutorial-details",
        component: () => import("./components/Tutorial")
      }, */
    /*{
        path: "/add",
        name: "add",
        component: () => import("./components/AddTutorial")
      }, */
  ],
})
