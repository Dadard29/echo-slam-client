<template>
  <div>
    <div v-if="infos === null && loading === true">
      <span class="spinner-border spinner-border-sm"></span>
    </div>
    <div v-else-if="infos !== null && loading === false">
      <span class="badge badge-primary">{{infos.Version}}</span>
    </div>
    <div v-else-if="infos === null && loading === false">
      <span class="badge badge-danger">OFFLINE</span>
    </div>
  </div>
</template>

<script>
export default {
  name: "RibbonInfos",
  data() {
    return {
      loading: false,
      infos: null,
    }
  },
  mounted() {
    this.getInfos();
  },
  methods: {
    getInfos() {
      let self = this;
      self.loading = true;
      self.infos = null;

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
}
</script>

<style scoped>

</style>