<template>
  <router-view />
</template>

<script>
import ApiClient, { API } from 'api/client';

export default {
  name: 'App',
  components: {},
  mixins: [ApiClient],
  async mounted() {
    const AuthApi = this.getApi(API.AUTH);
    const user = await AuthApi.check();
    await this.$store.dispatch('user/set', user);
  },
};
</script>

<style lang="scss">
</style>
