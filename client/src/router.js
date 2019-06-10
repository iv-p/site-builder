import Vue from "vue";
import Router from "vue-router";

Vue.use(Router);

export default new Router({
  base: process.env.BASE_URL,
  routes: [
    {
      path: "/",
      name: "onboard",
      component: () => import("./views/Onboard.vue")
    },
    {
      path: "/editor",
      name: "editor",
      component: () => import("./views/Editor.vue")
    }
  ]
});
