<template>
  <nav id="ribbon" class="navbar fixed-bottom">
    <div v-if="loading">
      <span class="spinner-border spinner-border-sm"></span>
    </div>
    <div v-if="infos !== null">
      <span v-if="up" class="text-danger">API down</span>
      <span v-else class="text-success">Services up and running</span>
      -
      <span>{{infos.Version}}</span>
    </div>
    <div v-else>
      <span class="text-danger">API down</span>
    </div>
  </nav>
</template>

<script>
export default {
  name: "Ribbon",
  computed: {
    up() {
      if (this.health !== null) {
        for (let i = 0; i < this.health.length; i++) {
          let api = this.health[i];
          if (api.check === false) {
            return false
          }
        }
        return true;
      }

      return false;
    }
  },
  data() {
    return {
      loading: false,

      infos: null,
      health: null,
    }
  },
  mounted() {
    let self = this;
    self.loading = true;
    window.backend.Connector.Infos()
      .then(function(infos) {
        self.loading = false;
        self.infos = infos;
      })
      .catch(function(err) {
        self.loading = false;
        console.log(err)
      })
  }
}
</script>

<style scoped>
#ribbon {
  font-size: 14px;
  border-top: solid 1px white;
  padding-top: 1px;
  padding-bottom: 1px;
}

.error {
  color: red
}
</style>