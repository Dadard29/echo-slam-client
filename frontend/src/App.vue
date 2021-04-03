<template>
  <div id=app>
    <Main/>
  </div>
</template>

<script>
import Main from "@/components/Main";
export default {
  components: {Main}
}
</script>

<style>
@import '../src/assets/css/bootstrap.min.css';

@font-face {
  font-family: Cinzel;
  font-weight: bold;
  src: url("../src/assets/css/Cinzel-Regular.ttf");
}

body {
  font-family: Cinzel, sans-serif;
}
</style>