<template>
  <div id="wrapper">
    <home-header />
    <home-nav />
    <div id="main">
      <recent id="recent" class="main" @onMore="openDetail" />
      <ranking id="ranking" class="main special" />
      <grave id="grave" class="main special" />
    </div>
    <site-footer />
  </div>
</template>

<script>
import { map } from 'lodash';
import HomeHeader from 'views/home/module/HomeHeader';
import HomeNav from 'views/home/module/HomeNav';
import SiteFooter from 'layout/SiteFooter';
import Recent from './module/Recent';
import Ranking from './module/Ranking';
import Grave from './module/Grave';

const BODY_STYLE = {
  backgroundImage: 'url(/images/matt-botsford-wxuPH6QRvJc-unsplash.jpg)',
  backgroundPosition: 'top',
  backgroundSize: 'contain',
  backgroundRepeat: 'no-repeat',
};

export default {
  name: 'Home',
  components: {
    HomeHeader,
    HomeNav,
    SiteFooter,
    Recent,
    Ranking,
    Grave,
  },
  data() {
    return {
      bodyStyle: {},
    };
  },
  created() {
    const body = document.body;
    map(BODY_STYLE, (val, key) => {
      this.bodyStyle[key] = body.style[key];
      body.style[key] = val;
    });
  },
  destroyed() {
    const body = document.body;
    map(this.bodyStyle, (val, key) => {
      body.style[key] = val;
    });
  },
  methods: {
    openDetail(mode) {
      console.log('open-detail - ', mode);
    },
  },
};
</script>
