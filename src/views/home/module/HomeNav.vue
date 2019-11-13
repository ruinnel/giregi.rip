<template>
  <nav id="nav">
    <ul>
      <li>
        <router-link to="/reporter">
          <i class="fas fa-user-edit" />
          <!-- i class="fab fa-earlybirds" /-->
          <span class="nav-title">기자</span>
        </router-link>
      </li>
      <li>
        <router-link to="/news">
          <i class="fas fa-newspaper" />
          <span class="nav-title">기사</span>
        </router-link>
      </li>
      <li>
        <router-link to="/mypage">
          <i class="fas fa-user" />
          <span class="nav-title">마이페이지</span>
        </router-link>
      </li>
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
