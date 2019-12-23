<template>
  <nav id="nav" ref="nav" :class="fixNav ? 'alt' : 'absolute-nav'">
    <ul>
      <li>
        <router-link to="/">
          <i class="fas fa-home" />
          <!-- i class="fab fa-earlybirds" /-->
          <span class="nav-title">홈</span>
        </router-link>
      </li>
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
  name: 'Navigator',
  data() {
    return {
      fixNav: false,
    };
  },
  mounted() {
    window.addEventListener('scroll', this.onScroll);
  },
  beforeDestroy() {
    window.removeEventListener('scroll', this.onScroll);
  },
  methods: {
    onScroll(event) {
      const main = document.getElementById('main');
      const { top } = main.getBoundingClientRect();
      if (top < 0) {
        this.fixNav = true;
        this.$refs.nav.classList.add('alt');
      } else {
        this.fixNav = false;
        this.$refs.nav.classList.remove('alt');
      }
      console.log(`top - ${top}, event - `, event);
      console.log(`y - ${window.scrollY}`);
    },
  },
};
</script>

<style scoped>
.absolute-nav {
  position: absolute !important;
  max-width: 100% !important;
  top: 0;
  left: 0;
}
</style>
