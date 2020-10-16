<template>
  <div class="container-xl">
    <div class="card">
      <div class="card-header">
        <h3 class="card-title"><i class="fas fa-archive mr-1" />내 아카이브 목록</h3>
      </div>
      <div class="card-body">
        <archive-list :archives="archives" :my-tags="tags" />
      </div>
    </div>
  </div>
</template>

<script>
import ArchiveList from './modules/ArchiveList';
import ApiClient, { API } from 'api/client';
import { concat } from 'lodash';

export default {
  name: 'Archives',
  components: {
    ArchiveList,
  },
  mixins: [ApiClient],
  data() {
    return {
      tags: [],
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
      this.tags = concat(this.tags, this.tags, this.tags, this.tags, this.tags);
    },
    async getArchives() {
      const UserApi = this.getApi(API.USER);
      const { data } = await UserApi.archives();
      this.archives = data;
    },
  },
};
</script>

<style lang="scss" scoped>

</style>
