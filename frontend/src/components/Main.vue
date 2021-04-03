<template>
  <div id="main" class="container-fluid">
    <div class="row">
      <div class="col-2">
        <Navbar/>
      </div>
      <div class="col-8">
        <RouteView/>
      </div>
    </div>
    <Ribbon/>
  </div>
</template>

<script>
import Navbar from "@/components/Navbar";
import RouteView from "@/components/RouteView";
import Ribbon from "@/components/Ribbon";
export default {
name: "Main",
  components: {Ribbon, RouteView, Navbar}
}
</script>

<style scoped>

</style>