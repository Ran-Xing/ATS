<template>
  <div class="Login">
    <b-row class="mt-5">
      <b-col
        md="8"
        offset-md="2"
        lg="6"
        offset-lg="3"
      >
        <b-card title="Login">
          <b-form>
            <b-form-group label="Email">
              <b-form-input
                id="username"
                v-model="$v.form.email.$model"
                type="email"
                placeholder="Enter email"
                required
              ></b-form-input>
              <b-form-invalid-feedback :state="ValidityState('email')">
                Email 0 - 50
              </b-form-invalid-feedback>
              <b-form-valid-feedback :state="ValidityState('email')">
                Looks Good.
              </b-form-valid-feedback>
            </b-form-group>
            <b-form-group label="Password">
              <b-form-input
                id="text-password"
                v-model="$v.form.password.$model"
                type="password"
                placeholder="Enter password"
                required
              ></b-form-input>
              <b-form-invalid-feedback :state="ValidityState('password')">
                Password 6 - 30
              </b-form-invalid-feedback>
              <b-form-valid-feedback :state="ValidityState('password')">
                Looks Good.
              </b-form-valid-feedback>
            </b-form-group>
            <b-form-group>
              <b-button
                @click="login"
                block
                variant="outline-primary"
              >Login</b-button>
            </b-form-group>
          </b-form>
        </b-card>
      </b-col>
    </b-row>
  </div>
</template>

<script>
import {
  required, minLength, maxLength, email,
} from 'vuelidate/lib/validators';
import { mapActions } from 'vuex';

export default {
  data() {
    return {
      form: {
        email: '',
        password: '',
      },
    };
  },
  validations: {
    form: {
      email: {
        required,
        email,
        minLength: minLength(3),
        maxLength: maxLength(50),
      },
      password: {
        required,
        minLength: minLength(7),
        maxLength: maxLength(30),
      },
    },
  },
  methods: {
    ...mapActions('userModule', { userLogin: 'login' }),
    // 限制输入字符数量
    ValidityState(name) {
      const { $dirty, $error } = this.$v.form[name];
      return $dirty ? !$error : null;
    },
    login() {
      // 验证数据
      this.$v.form.$touch();
      if (this.$v.form.$anyError) {
        console.log(this.$v.form.$anyError);
        return;
      }
      // 请求
      this.userLogin(this.form)
        .then((respone) => {
          // 验证状态码
          if (respone.data.code === 422) {
            this.$bvToast.toast(respone.data.data, {
              title: '请求失败',
              variant: 'danger',
              appendToast: true,
              autoHideDelay: 3000,
            });
            console.log(respone.data.data);
          }
          if (respone.data.code === 200) {
            this.$bvToast.toast(respone.data.data.message, {
              title: '登录成功!',
              variant: 'success',
              appendToast: true,
              autoHideDelay: 3000,
            });
            console.log(respone.data.data.message);
          }
        })
        .then(() => {
          // 跳转主页
          this.$router.replace({ name: 'home' });
        })
        .catch((err) => {
          this.$bvToast.toast(err.response, {
            title: '请求失败',
            variant: 'danger',
            appendToast: true,
            autoHideDelay: 3000,
          });
          console.log(err.response);
        });
    },
  },
};
</script>

<style>
.toast:not(.show) {
  display: block;
}
</style>
