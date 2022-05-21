<template>
  <div>
    <b-navbar
      toggleable="lg"
      type="dark"
      variant="info"
    >
      <b-container>
        <b-navbar-brand @click="$router.go(0)">简历管理系统</b-navbar-brand>
        <!-- <b-navbar-brand @click="$router.push({ name: 'home' })">简历管理系统</b-navbar-brand> -->
        <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>
        <b-collapse
          id="nav-collapse"
          is-nav
        >
          <b-navbar-nav class="ml-auto">
            <b-nav-item-dropdown v-if="UserInfo">
              <template #button-content>
                <em>{{ UserInfo.username }}</em>
              </template>
              <b-dropdown-item @click="$router.push({name: 'profile'})">Profile</b-dropdown-item>
              <b-dropdown-item @click="logout">Sign Out</b-dropdown-item>
            </b-nav-item-dropdown>
            <b-nav-item-dropdown v-if="!UserInfo">
              <template #button-content>
                <em>Login</em>
              </template>
              <b-dropdown-item
                v-if="$route.name != 'login'"
                @click="$router.replace({ name: 'login' })"
              >Login</b-dropdown-item>
              <b-dropdown-item
                v-if="$route.name != 'register'"
                @click="$router.replace({ name: 'register' })"
              >Register</b-dropdown-item>
            </b-nav-item-dropdown>
          </b-navbar-nav>
        </b-collapse>
      </b-container>
    </b-navbar>
  </div>
</template>

<script>
import { mapState, mapActions } from 'vuex';

export default {
  computed: mapState({
    UserInfo: (state) => state.userModule.UserInfo,
  }),
  methods: mapActions('userModule', ['logout']),
};
</script>

<style>
fieldset {
  margin-top: 1.7vh !important;
}

button.btn.btn-outline-primary.btn-block {
  width: 100%;
}
</style>
