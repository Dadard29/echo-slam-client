<template>
  <div class="navbar-custom">
    <router-link to="/">
      <img src="../assets/icons/main.png" alt="main">
    </router-link>
    <div class="nav flex-column nav-pills" id="v-pills-tab" role="tablist" aria-orientation="vertical">
      <router-link to="/" class="nav-link active" id="v-pills-home-tab"
                   data-toggle="pill" role="tab">
        Home
      </router-link>
      <router-link to="/about" class="nav-link" id="v-pills-profile-tab"
                   data-toggle="pill" role="tab">
        About
      </router-link>
    </div>
  </div>
</template>

<script>
export default {
  name: "Navbar"
}
</script>

<style scoped>
  .active {
    color: red
  }

  .navbar-custom {
    padding: 15px;
  }

</style>