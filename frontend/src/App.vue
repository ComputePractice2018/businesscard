<template>
  <div id="app" class="app">
    <b-container>
      <b-navbar toggleable="md" type="light" variant="light">
      <b-navbar-nav>
        <b-nav-item><router-link to="/">Главная</router-link></b-nav-item>
        <b-nav-item><router-link to="/about">О программе</router-link></b-nav-item>
      </b-navbar-nav>
      </b-navbar>
      <router-view/>
    </b-container>
  </div>
</template>
<style>
.app {
  background-color: #e2f7fd;
}
.navbar.navbar-light {
  background-color: #e2f7fd !important;
}
</style>
