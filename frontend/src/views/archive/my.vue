<template>
  <div class="container-xl">
    <div class="card">
      <div class="card-header">
        <h3 class="card-title">
          <span><i class="fas fa-archive mr-1" />내 아카이브 목록</span>
        </h3>
      </div>
      <div class="card-body">
        <archive-list
          :total="total"
          :archives="archives"
          :my-tags="tags"
          @search="onSearch"
          @more="onMore"
        />
      </div>
    </div>
  </div>
</template>

<script>
import { concat } from 'lodash';
import ArchiveList from './modules/ArchiveList';
import ApiClient, { API } from 'api/client';

export default {
  name: 'Archives',
  components: {
    ArchiveList,
  },
  mixins: [ApiClient],
  data() {
    return {
      tags: [],
      total: 0,
      search: {},
      archives: [],
    };
  },
  created() {
    this.getTags();
    this.getArchives();
  },
  methods: {
    async getTags() {
      const UserApi = this.getApi(API.USER);
      this.tags = await UserApi.tags();
    },
    async getArchives() {
      const loader = this.$loading.show();
      try {
        const UserApi = this.getApi(API.USER);
        const { total, nextCursor, data } = await UserApi.archives(this.search);
        this.total = total;
        this.search.nextCursor = nextCursor;
        this.archives = concat(this.archives, data);
      } catch (e) {
        console.warn('load archives fail.', e);
      } finally {
        loader.hide();
      }
    },
    async onSearch({ keyword, tagId }) {
      this.search = { keyword, tagId };
      this.total = 0;
      this.archives = [];
      await this.getArchives();
    },
    async onMore() {
      await this.getArchives();
    },
  },
};
</script>

<style lang="scss" scoped>

</style>
