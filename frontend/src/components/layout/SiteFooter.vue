<template>
  <footer class="footer footer-transparent">
    <div class="container">
      <div class="row text-center align-items-center flex-row-reverse">
        <div class="col-lg-auto ml-lg-auto">
          <ul class="list-inline list-inline-dots mb-0">
            <li v-for="(menu, idx) in menus" :key="idx" class="list-inline-item">
              <a :href="menu.href" class="link-secondary">{{ menu.name }}</a>
            </li>
          </ul>
        </div>
        <div class="col-12 col-lg-auto mt-3 mt-lg-0">
          Copyright © 2020
          <a href="." class="link-secondary">ruinnel</a>.
          All rights reserved.
        </div>
      </div>
    </div>
  </footer>
</template>

<script>
const menus = [
  { name: 'Home', href: '/' },
  { name: 'Mypage', href: '/mypage' },
];

export default {
  name: 'SiteFooter',
  data() {
    return { menus };
  },
};
</script>

<style scoped>

</style>
