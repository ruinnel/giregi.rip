<template>
  <div id="wrapper">
    <header id="header">
      <h1>기자</h1>
      <p>기자 목록</p>
    </header>
    <div id="main">
      <section id="content" class="main reporter">
        <card v-for="(item, idx) in reporters" :key="idx" :reporter="item" />
      </section>
    </div>
    <site-footer />

    <reporter-detail :active="showDetail" :reporter="reporter" />
  </div>
</template>

<script>
import { isEmpty } from 'lodash';
import ReportApi from 'api/reporter';
import Card from './module/Card';
import ReporterDetail from './Detail';

export default {
  name: 'Reporter',
  components: {
    Card,
    ReporterDetail,
  },
  props: {
    reporterId: {
      type: String,
      default: () => '',
    },
  },
  data: function () {
    return {
      active: true,
      reporters: [],
      reporter: {},
    };
  },
  computed: {
    showDetail() {
      return !isEmpty(this.reporter);
    },
  },
  created() {
    this.loadData();
  },
  methods: {
    async loadData() {
      this.reporters = await ReportApi.search({});
      console.log('reporters - ', this.reporters);
    },
    onClose() {
      this.$router.push('/');
    },
  },
};
</script>

<style scoped>
</style>
