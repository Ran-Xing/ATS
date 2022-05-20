<template>
  <div class="Register">
    <b-row class="mt-5">
      <b-col
        md="8"
        offset-md="2"
        lg="6"
        offset-lg="3"
      >
        <b-card title="Register">
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
                @click="register"
                block
                variant="outline-primary"
              >Register</b-button>
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
    // 限制输入字符数量
    ValidityState(name) {
      const { $dirty, $error } = this.$v.form[name];
      return $dirty ? !$error : null;
    },
    register() {
      // 验证数据
      this.$v.form.$touch();
      if (this.$v.form.$anyError) {
        console.log(this.$v.form.$anyError);
      }
      // 请求
      const api = 'http://localhost:8081/api/auth/register';
      this.$axios
        .post(api, { ...this.form })
        .then((res) => {
          if (res.data.code === 422) {
            this.$bvToast.toast(res.data.data, {
              title: '请求失败',
              variant: 'danger',
              appendToast: true,
              autoHideDelay: 3000,
            });
            console.log(res.data.data);
          }
          if (res.data.code === 200) {
            this.$bvToast.toast(res.data.data.message, {
              title: '注册成功!',
              variant: 'success',
              appendToast: true,
              autoHideDelay: 3000,
            });
            console.log(res.data.data.message);
          }
          // 保存token

          // 跳转主页
        })
        .catch((err) => {
          this.$bvToast.toast(err.respone.data.data, {
            title: '请求失败',
            variant: 'danger',
            appendToast: true,
            autoHideDelay: 3000,
          });
          console.log(err.respone.data.data);
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
