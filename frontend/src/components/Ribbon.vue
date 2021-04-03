<template>
  <nav id="ribbon" class="navbar fixed-bottom">
    <div v-if="loading">
      <span class="spinner-border spinner-border-sm"></span>
    </div>
    <div v-else>
      <a class="navbar-brand" href="#">Fixed bottom</a>
    </div>
  </nav>
</template>

<script>
export default {
  name: "Ribbon",
  data() {
    return {
      loading: false
    }
  }
}
</script>

<style scoped>
#ribbon {
  font-size: 8px;
  border-top: solid 1px white;
  padding-top: 1px;
  padding-bottom: 1px;
}
</style>