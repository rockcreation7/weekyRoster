<template>
  <div id="app">
    <nav class="navbar navbar-expand">
      <a href="#" class="navbar-brand">Roster</a>
      <div class="navbar-nav mr-auto">
        <li class="nav-item">
          <a href="/rosters" class="nav-link">Rosters</a>
        </li>
        <li class="nav-item">
          <a href="/addroster" class="nav-link">Add Roster</a>
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
