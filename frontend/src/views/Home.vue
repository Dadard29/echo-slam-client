<template>
  <div>
    <h1>Home</h1>
  </div>
</template>

<script>
export default {
  name: "Home",
  data() {
    return {
      msg: ""
    }
  },
  methods: {
    getMessage() {
      let self = this;
      window.backend.basic()
        .then(function(msg) {
          self.msg = msg;
        })
    }
  }
}
</script>

<style scoped>

</style>