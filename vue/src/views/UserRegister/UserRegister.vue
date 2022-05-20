<template>
  <div class="Register">
    <b-row>
      <b-col>
        <b-card title="注册">
          <b-form-group
            id="input-group-1"
            label="Email address:"
            label-for="username"
            description="We'll never share your email with anyone else."
          >
            <b-form-input
              id="username"
              v-model="form.email"
              type="email"
              placeholder="Enter email"
              required
            ></b-form-input>
          </b-form-group>
          <b-form @submit.stop.prevent>
            <b-form-input
              type="password"
              id="text-password"
              v-model="form.password"
              aria-describedby="password-help-block"
              placeholder="Enter password"
            ></b-form-input>
            <b-form-text id="password-help-block">
              Your password must be 8-20 characters long, contain letters and
              numbers, and must not contain spaces, special characters, or
              emoji.
            </b-form-text>
          </b-form>
        </b-card>
      </b-col>
    </b-row>
  </div>
</template>

<script>
export default {
  data() {
    return {
      form: {
        email: '',
        password: '',
      },
    };
  },
};
</script>

<style>
</style>
