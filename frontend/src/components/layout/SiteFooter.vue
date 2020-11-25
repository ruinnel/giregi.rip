<template>
  <footer class="footer footer-transparent">
    <div class="container">
      <div class="row text-center align-items-center flex-row-reverse">
        <div class="col-lg-auto ml-lg-auto">
          <ul class="list-inline list-inline-dots mb-0">
            <li v-for="(menu, idx) in menus" :key="idx" class="list-inline-item">
              <a :href="menu.href" target="_blank" class="link-secondary">
                <i :class="menu.icon" />
                <span>{{ menu.name }}</span>
              </a>
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
  { name: '오류제보', icon: 'fas fa-bug', href: 'https://github.com/ruinnel/giregi.rip/issues' },
  { name: 'Source', icon: 'fab fa-github', href: 'https://github.com/ruinnel/giregi.rip' },
];

export default {
  name: 'SiteFooter',
  data() {
    return { menus };
  },
};
</script>

<style scoped>
.link-secondary > i {
  margin: 0 5px 0 0;
}
</style>
