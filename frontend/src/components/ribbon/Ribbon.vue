<template>
  <nav id="ribbon" class="navbar fixed-bottom">
    <RibbonHealth/>
    <RibbonInfos/>
  </nav>
</template>

<script>
import RibbonInfos from "@/components/ribbon/RibbonInfos";
import RibbonHealth from "@/components/ribbon/RibbonHealth";
export default {
  name: "Ribbon",
  components: {RibbonHealth, RibbonInfos},
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