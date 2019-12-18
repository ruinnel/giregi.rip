<template>
  <vs-popup
    title="로그인"
    :active.sync="active"
    class="popup-content"
    @close="onClose"
  >
    <div id="login-buttons" ref="content" />
  </vs-popup>
</template>

<script>
import { isEmpty, get } from 'lodash';
import firebase from 'firebase/app';
import * as firebaseui from 'firebaseui';
import 'firebaseui/dist/firebaseui.css';
import ApiClient, { API } from 'api/client';

export default {
  name: 'Login',
  mixins: [ApiClient],
  data: function () {
    return {
      active: true,
    };
  },
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
          const AuthApi = this.getApi(API.AUTH);
          const { id, token } = await AuthApi.login(email, idToken);
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
        color: 'warning',
        title: '로그인이 필요합니다.',
        text: '로그인 페이지로 이동합니다.',
      };
      this.$vs.notify(options);
    },
    showSuccessMessage() {
      const options = {
        color: 'success',
        title: '로그인 완료',
        text: '로그인 되었습니다.',
      };
      this.$vs.notify(options);
    },
  },
};
</script>

<style scoped>
</style>
