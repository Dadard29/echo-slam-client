<template>
  <nav id="navbar" class="navbar navbar-expand-lg">
    <router-link to="/" class="navbar-brand">
      <img class="brand" src="../../assets/icons/main.png" alt="main">
    </router-link>

    <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
      <span class="navbar-toggler-icon"></span>
    </button>

    <div class="collapse navbar-collapse" id="navbarSupportedContent">
      <ul class="navbar-nav mr-auto">
        <li class="nav-item">
          <router-link class="nav-link link" to="/" exact-active-class="active">Home</router-link>
        </li>
        <li class="nav-item">
          <router-link class="nav-link link" to="/about" exact-active-class="active">About</router-link>
        </li>
      </ul>
    </div>

    <ProfileMenu @disconnect="$emit('disconnect')"/>
  </nav>
</template>

<script>
import ProfileMenu from "@/components/navbar/ProfileMenu";
export default {
  name: "Navbar",
  components: {ProfileMenu}
}
</script>

<style scoped>
#navbar {
  border-bottom: solid 1px white;
  background: black;
}

.link {
  font-size: 24px;
  color: grey;
}

.active {
  color: white;
}

.brand {
  border: solid 1px white;
  border-radius: 5px;
}
</style>