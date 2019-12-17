<template>
  <router-view />
</template>

<script>
import ApiClient, { API } from 'api/client';

export default {
  name: 'App',
  components: {},
  mixins: [ApiClient],
  mounted() {
    const AuthApi = this.getApi(API.AUTH);
    AuthApi.check();
  },
};
</script>

<style lang="scss">
</style>
