<template>
  <div id="connection" class="container">
    <h3 style="font-family: Cinzel, sans-serif">Connection</h3>
    <div class="form-group">
      <label for="username">Username</label>
      <input v-model="usernameInput" type="text" class="form-control" id="username" placeholder="Username...">
    </div>

    <div class="form-group">
      <label for="password">Password</label>
      <input v-model="passwordInput" type="password" class="form-control" id="password" placeholder="Password...">
    </div>
    <div class="form-group">
      <button @click="submit" type="button" class="btn btn-primary">Sign in</button>
    </div>
  </div>
</template>

<script>
import Backend from "@/backend";

export default {
  name: "Connection",
  data() {
    return {
      usernameInput: "",
      passwordInput: ""
    }
  },
  mounted() {
  },
  methods: {
    submit() {
      let self = this;
      window.backend.Logger.Loading("signing in...")
      Backend.connector().SignIn(this.usernameInput, this.passwordInput)
        .then(function() {
          window.backend.Logger.Info("successfully signed in")
          self.$emit('connect')
        })
        .catch(function(err) {
          window.backend.Logger.Error(err)
        })
    }
  }
}
</script>

<style scoped>
#connection {
  width: 20%;
  margin-top: 10%;

  padding: 30px;
  border: solid 5px white;
  border-radius: 30px;

  font-family: Arial, sans-serif;
}
</style>