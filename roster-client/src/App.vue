<template>
  <div id="app">
    <router-link to="/rosters">Roster</router-link>
    <nav class="navbar navbar-expand">
      <div class="navbar-nav">
        <li class="nav-item">
          <router-link to="/rosters" class="nav-link">Home</router-link>
        </li>
        <li class="nav-item">
          <router-link to="/addroster" class="nav-link">+Roster</router-link>
        </li>
        <li class="nav-item">
          <router-link to="/products" class="nav-link">Products</router-link>
        </li>
        <li class="nav-item">
          <router-link to="/addproduct" class="nav-link">+Product</router-link>
        </li>
        <li class="nav-item">
          <router-link to="/cart" class="nav-link">Cart</router-link>
        </li>
      </div>
    </nav>

    <div class="container mt-3">
      <router-view />
    </div>
  </div>
</template>

<script>
export default {
  name: "app",
};
</script>
