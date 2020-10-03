 

<template>
  <div class="page-container">
    <md-app md-mode="reveal">
      <md-app-toolbar class="md-primary">
        <md-button class="md-icon-button" @click="menuVisible = !menuVisible">
          <md-icon>menu</md-icon>
        </md-button>
        <router-link to="/rosters"
          ><span class="md-title">Store Tool</span></router-link
        >
      </md-app-toolbar>

      <md-app-drawer :md-active.sync="menuVisible">
        <md-toolbar class="md-transparent" md-elevation="0"
          >Menu</md-toolbar
        >

        <md-list>
          <md-list-item @click="menuVisible = !menuVisible">
            <md-icon>home</md-icon>
            <router-link to="/rosters"   class="nav-link"
              ><span class="md-list-item-text">Rosters</span></router-link
            >
          </md-list-item>

          <md-list-item @click="menuVisible = !menuVisible">
            <md-icon>add</md-icon>
            <router-link to="/addroster" class="nav-link"
              ><span class="md-list-item-text">+ Roster</span></router-link
            >
          </md-list-item>

          <md-list-item @click="menuVisible = !menuVisible">
            <md-icon>info</md-icon>
            <router-link to="/products" class="nav-link">
              <span class="md-list-item-text">Products</span>
            </router-link>
          </md-list-item>

          <md-list-item @click="menuVisible = !menuVisible">
            <md-icon>add</md-icon>
            <router-link to="/addproduct" class="nav-link">
              <span class="md-list-item-text">+ Product</span>
            </router-link>
          </md-list-item>

          <md-list-item @click="menuVisible = !menuVisible">
            <md-icon>info</md-icon>
            <router-link to="/cart" class="nav-link"
              ><span class="md-list-item-text">+Cart</span>
            </router-link>
          </md-list-item>
        </md-list>
      </md-app-drawer>

      <md-app-content>
        <router-view />
      </md-app-content>
    </md-app>
  </div>
</template> 

<script>
export default {
  name: "app",
  data: () => ({
    menuVisible: false,
  }),
};
</script>

<style>
.md-app {
/*   max-height: 400px;
  border: 1px solid rgba(#000, 0.12); */
  min-height: 100vh;
}

.md-drawer {
  width: 230px;
  max-width: calc(100vw - 125px);
}
</style>