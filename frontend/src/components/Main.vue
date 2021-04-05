<template>
  <div id="main">
    <CredCheck v-if="connected === null"/>
    <div v-else-if="connected === false">
      <Connection @connect="connected = true"/>
    </div>
    <div v-else-if="connected === true">
      <Navbar @disconnect="connected = false"/>
      <RouteView/>
    </div>

    <Ribbon/>
  </div>
</template>

<script>
import Navbar from "@/components/navbar/Navbar";
import RouteView from "@/components/RouteView";
import Ribbon from "@/components/ribbon/Ribbon";
import Backend from "@/backend";
import CredCheck from "@/components/connection/CredCheck";
import Connection from "@/components/connection/Connection";
export default {
name: "Main",
  components: {Connection, CredCheck, Ribbon, RouteView, Navbar},
  data() {
    return {
      connected: null,
    }
  },
  mounted() {
    this.checkSignIn()
  },
  methods: {
    checkSignIn() {
      let self = this;

      self.connected = null;
      window.backend.Logger.Loading("signing in...")

      Backend.connector().SignInCheck()
        .then(function() {
          window.backend.Logger.Info("successfully signed in");
          self.connected = true;
        })
        .catch(function() {
          window.backend.Logger.Error("need credentials");
          self.connected = false;
        })
    }
  }
}
</script>

<style scoped>

</style>