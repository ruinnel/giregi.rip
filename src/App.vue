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
    if (this.$route.name !== 'login') {
      const UserApi = this.getApi(API.USER);
      const user = await UserApi.my();
      await this.$store.dispatch('user/set', user);
    }
  },
};
</script>
<style lang="scss">
</style>
