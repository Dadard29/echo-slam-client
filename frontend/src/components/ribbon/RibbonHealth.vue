<template>
  <div v-if="health === null && loading === true">
    <span class="spinner-border spinner-border-sm"></span>
  </div>
  <div v-else-if="health !== null && loading === false">
    <span v-if="up" class="badge badge-success">CONNECTED</span>
    <span v-else class="badge badge-danger">API ERROR</span>
  </div>
  <div v-else-if="health === null  && loading === false">
    <span class="badge badge-danger">OFFLINE</span>
  </div>
</template>

<script>
export default {
  name: "RibbonHealth",
  computed: {
    up() {
      if (this.health !== null) {
        for (let i = 0; i < this.health.length; i++) {
          let apiHealth = this.health[i];
          if (apiHealth.check === false) {
            return false;
          }
        }

        return true;

      } else {
        return false;
      }
    }
  },
  data() {
    return {
      health: null,
      loading: false,
    }
  },
  mounted() {
    this.getHealth();
  },
  methods: {
    getHealth() {
      let self = this;
      self.loading = true;
      self.health = null;
      window.backend.Connector.HealthApis()
        .then(function(h) {
          self.loading = false;
          self.health = h.list
        })
    }
  }
}
</script>

<style scoped>

</style>