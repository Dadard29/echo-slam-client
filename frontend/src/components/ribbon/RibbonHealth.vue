<template>
  <div v-if="health === null && loading === true">
    <span class="spinner-border spinner-border-sm"></span>
  </div>
  <div v-else-if="health !== null && loading === false">
    <span v-if="up" class="badge badge-success"
          data-toggle="tooltip" data-placement="top"
          title="Connection to services established">
      CONNECTED
    </span>
    <span v-else class="badge badge-danger"
          data-toggle="tooltip" data-placement="top"
          title="Some services are not available at this time, some features may not be working properly">
      API ERROR
    </span>
  </div>
  <div v-else-if="health === null  && loading === false">
    <span class="badge badge-danger"
          data-toggle="tooltip" data-placement="top"
          title="No internet connection can be established with the remote services">
      OFFLINE
    </span>
  </div>
</template>

<script>

import Backend from "@/backend";
import $ from "jquery";

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
      Backend.connector().HealthApis()
        .then(function(h) {
          self.loading = false;
          self.health = h.list;

          $(function () {
            $('[data-toggle="tooltip"]').tooltip()
          })
        })
        .catch(function(err) {
          self.loading = false;
          console.log(err)

          $(function () {
            $('[data-toggle="tooltip"]').tooltip()
          })
        })
    }
  }
}
</script>

<style scoped>
</style>