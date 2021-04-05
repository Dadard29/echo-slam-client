<template>
  <nav id="ribbon" class="navbar fixed-bottom">
    <RibbonHealth/>
    <RibbonLog/>
    <RibbonInfos/>
  </nav>
</template>

<script>
import RibbonInfos from "@/components/ribbon/RibbonInfos";
import RibbonHealth from "@/components/ribbon/RibbonHealth";
import RibbonLog from "@/components/ribbon/RibbonLog";
export default {
  name: "Ribbon",
  components: {RibbonLog, RibbonHealth, RibbonInfos},
}
</script>

<style scoped>
#ribbon {
  font-size: 14px;
  border-top: solid 1px white;
  padding-top: 1px;
  padding-bottom: 1px;

  font-family: monospace,sans-serif;
}
</style>