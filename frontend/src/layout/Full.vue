<template>
  <div class="flex-fill d-flex flex-column justify-content-center">
    <router-view />
  </div>
</template>

<script>
export default {
  name: 'Full',
  beforeCreate() {
    document.body.className = 'antialiased border-top-wide border-primary d-flex flex-column';
  },
  beforeDestroy() {
    document.body.className = 'antialiased';
  },
};
</script>

<style scoped>

</style>
