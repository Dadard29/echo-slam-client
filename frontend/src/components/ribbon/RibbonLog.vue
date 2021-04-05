<template>
<div>
  <span v-if="level === levelDebug" class="text-muted">
    {{msg}}
  </span>
  <span v-else-if="level === levelError" class="text-danger">
    <img src="../../assets/icons/icons8-cancel-48.png" alt="error" style="width: 1rem">
    {{msg}}
  </span>
  <span v-else-if="level === levelLoading" class="text-primary">
    <span class="spinner-border spinner-border-sm"></span>
    {{msg}}
  </span>
  <span v-else-if="level === levelInfo" class="text-success">
    <img src="../../assets/icons/icons8-ok-48.png" alt="info" style="width: 1rem">
    {{msg}}
  </span>
</div>
</template>

<script>
export default {
  name: "RibbonLog",
  data() {
    return {
      msg: "",
      level: "DEBUG",
      levelDebug: "DEBUG",
      levelError: "ERROR",
      levelInfo: "INFO",
      levelLoading: "LOADING"
    }
  },
  mounted() {
    this.getLastMessage();
    let self = this;

    // eslint-disable-next-line no-undef
    wails.Events.On("log",function() {
      self.getLastMessage()
    })
  },
  methods: {
    getLastMessage() {
      let self = this;
      window.backend.Logger.GetLastMsg()
        .then(function(logMsg) {
          self.msg = logMsg.msg;
          self.level = logMsg.level;
        })
    }
  }
}
</script>

<style scoped>

</style>