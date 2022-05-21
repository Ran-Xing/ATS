<template>
  <h3>hello Profile</h3>
</template>

<script>
import { mapState } from 'vuex';

export default {
  computed: mapState({
    UserInfo: (state) => state.userModule.UserInfo,
  }),
};
</script>
<style>
</style>
