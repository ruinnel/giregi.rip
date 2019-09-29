<template>
  <nav id="nav">
    <ul>
      <li><a href="#recent"><i class="fas fa-rss" /><span class="nav-title">최근기사</span></a></li>
      <li><a href="#ranking"><i class="fas fa-medal"></i><span class="nav-title">랭킹</span></a></li>
      <li><a href="#grave"><i class="fas fa-ban"></i><span class="nav-title">묘지</span></a></li>
    </ul>
  </nav>
</template>

<script>
export default {
  name: 'HomeNav',
};
</script>

<style scoped>
.nav-title {
  margin-left: 5px;
}
</style>
