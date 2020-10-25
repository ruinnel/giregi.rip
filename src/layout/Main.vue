<template>
  <div>
    <site-header />
    <div class="content">
      <router-view />
      <site-footer />
    </div>
  </div>
</template>

<script>
import SiteHeader from '../components/layout/SiteHeader';
import SiteFooter from '../components/layout/SiteFooter';
export default {
  name: 'Main',
  components: { SiteFooter, SiteHeader },
};
</script>

<style scoped>

</style>
