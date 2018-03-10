<template>
  <div id="app">
    <header v-if="loggedIn">
      <div class="navbar navbar-dark bg-dark box-shadow">
        <div class="container-fluid">
          <a href="#" class="navbar-brand">
            <strong>App name</strong>
          </a>

          <ul class="navbar-nav ml-auto">
            <li class="nav-item active">
              <router-link :to="{ name: 'logout' }" class="nav-link">Sair</router-link>
            </li>
          </ul>
        </div>
      </div>
    </header>

    <router-view/>
  </div>
</template>

<script>
export default {
  name: 'App',
  computed: {
    loggedIn () {
      return this.$store.getters.loggedIn
    }
  }
}
</script>

<style type="scss">
@import 'bootstrap/dist/css/bootstrap.min.css';
</style>
