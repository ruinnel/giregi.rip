<template>
  <div id="wrapper">
    <home-header />
    <home-nav />
    <div id="main">
      <trust-news-media class="main" />
      <irresponsible class="main" />
      <responsibility-of-reporter class="main" />
      <importance-of-record class="main" />
      <influence-of-memos class="main" />
      <importance-of-archive class="main" />
    </div>
    <site-footer />
  </div>
</template>

<script>
import { map } from 'lodash';
import HomeHeader from 'views/home/module/HomeHeader';
import HomeNav from 'views/home/module/HomeNav';
import TrustNewsMedia from './module/TrustNewsMedia';
import Irresponsible from './module/Irresponsible';
import ResponsibilityOfReporter from './module/ResponsibilityOfReporter';
import ImportanceOfRecord from './module/ImportanceOfRecord';
import InfluenceOfMemos from './module/InfluenceOfMemos';
import ImportanceOfArchive from './module/ImportanceOfArchive';

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
    TrustNewsMedia,
    Irresponsible,
    ResponsibilityOfReporter,
    ImportanceOfRecord,
    InfluenceOfMemos,
    ImportanceOfArchive,
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
