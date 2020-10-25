<template>
  <div class="container-tight py-6">
    <div class="text-center mb-4">
      <img src="/image/logo/logo-black-horizontal.png" height="36" alt="logo">
    </div>
    <form class="card card-md" action="." method="get">
      <div class="card-body">
        <h2 class="mb-5 text-center">Login</h2>
        <div class="form-footer">
          <div id="login-buttons" ref="content" />
        </div>
      </div>
    </form>
  </div>
</template>

<script>
import firebase from 'firebase/app';
import * as firebaseui from 'firebaseui';
import 'firebaseui/dist/firebaseui.css';
import { get, isEmpty } from 'lodash';
import ApiClient, { API } from 'api/client';

export default {
  name: 'Login',
  mixins: [ApiClient],
  created() {
    firebase.auth().onAuthStateChanged(this.onAuthStateChanged, this.onAuthError);
  },
  mounted() {
    const ui = new firebaseui.auth.AuthUI(firebase.auth());
    const config = {
      signInSuccessUrl: window.location.href,
      signInOptions: [
        firebase.auth.GoogleAuthProvider.PROVIDER_ID,
      ],
    };
    ui.start('#login-buttons', config);
  },
  methods: {
    async onAuthStateChanged(user) {
      if (user) {
        // User is signed in.
        const { email } = user;
        const idToken = await user.getIdToken();

        let loginResult = false;
        if (isEmpty(idToken) > 0) {
          loginResult = false;
          this.showFailMessage();
        } else {
          loginResult = true;
          const UserApi = this.getApi(API.USER);
          const { id, token } = await UserApi.login(email, idToken);
          loginResult = !!(id && token);
          if (loginResult) {
            this.showSuccessMessage();
          } else {
            this.showFailMessage();
          }
        }

        if (loginResult) {
          const returnUrl = get(this.$router.currentRoute, 'query.returnUrl', '/');
          await this.$router.push(returnUrl);
        }
      } else {
        // User is signed out.
        this.showFailMessage();
        await this.$router.push('/login');
      }
    },
    onAuthError(err) {
      console.log('authFail - ', err);
      this.showFailMessage();
      this.$router.push('/');
    },
    onClose() {
      this.$router.push('/');
    },
    showFailMessage() {
      const options = {
        title: '로그인이 필요합니다.',
        msg: '로그인 페이지로 이동합니다.',
      };
      this.$toastr.w(options);
    },
    showSuccessMessage() {
      const options = {
        title: '로그인 완료',
        msg: '로그인 되었습니다.',
      };
      this.$toastr.s(options);
    },
  },
};
</script>

<style scoped>

</style>
