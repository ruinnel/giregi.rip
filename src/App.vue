<template>
  <router-view />
</template>

<script>
import AuthApi from 'api/auth';

export default {
  name: 'App',
  components: {},
  mounted() {
    AuthApi.check();
  },
};
</script>

<style lang="scss">
</style>
