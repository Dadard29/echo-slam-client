<template>
  <div id="profileMenu" v-if="avatarUrl !== null" class="dropdown">
    <img id="profileImage" :src="avatarUrl" alt="profileImage" class="profileImage" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
    <div class="dropdown-menu dropdown-menu-right" aria-labelledby="profileImage" style="max-width: 400px;">
      <div class="media">
          <img :src="avatarUrl" class="profileImageBig" alt="profileImage">
          <div class="media-body">
            <span style="font-family: Cinzel, sans-serif">{{ profile.username }}</span>
            <div>salut</div>
          </div>
      </div>
      <router-link class="dropdown-item" to="/settings">
          Settings
      </router-link>
      <div @click="disconnect" class="dropdown-item">
        Logout
      </div>
    </div>
  </div>
</template>

<script>
import Backend from "@/backend";

export default {
name: "ProfileMenu",
  data() {
    return {
      profile: null,
      avatarUrl: null
    }
  },
  mounted() {
    this.getProfile();
  },
  methods: {
    disconnect() {
      let self = this;
      Backend.connector().LogOut()
        .then(function() {
          window.backend.Logger.Info("successfully disconnected");
          self.$emit('disconnect');
        })
        .catch(function(err) {
          window.backend.Logger.Error(err)
        })

    },

    getProfile() {
      let self = this;
      Backend.connector().GetProfile()
          .then(function (profile) {
            self.profile = profile;
            return Backend.connector().GetAvatarUrl(profile.avatar_name);
          })
          .then(function (url) {
            self.avatarUrl = url;
          })
          .catch(function (err) {
            window.backend.Logger.Error(err)
          })
    }
  }
}
</script>

<style scoped>
#profileMenu {
  font-family: Arial, sans-serif;
}

.profileImage {
  width: 48px;
  border-radius: 24px;
  cursor: pointer;
}

.profileImageBig {
  width: 64px;
  border-radius: 32px;
  margin: 5px;
}
</style>