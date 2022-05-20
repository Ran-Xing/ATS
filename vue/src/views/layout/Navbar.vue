<template>
  <div>
    <b-navbar
      toggleable="lg"
      type="dark"
      variant="info"
    >
      <b-container>
        <b-navbar-brand href="#">简历管理系统</b-navbar-brand>
        <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>
        <b-collapse
          id="nav-collapse"
          is-nav
        >
          <b-navbar-nav class="ml-auto">
            <b-nav-item @click="$router.replace({ name: 'login' })">登录</b-nav-item>
            <b-nav-item @click="$router.replace({ name: 'register' })">注册</b-nav-item>
          </b-navbar-nav>
        </b-collapse>
      </b-container>
    </b-navbar>
  </div>
</template>

<script>
export default {};
</script>

<style>
fieldset {
  margin-top: 1.7vh !important;
}

button.btn.btn-outline-primary.btn-block {
  width: 100%;
}
</style>
