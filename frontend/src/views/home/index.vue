<template>
  <div class="container-xl">
    <!-- Page title -->
    <div class="page-header text-white">
      <div class="row align-items-center">
        <div class="col-auto">
          <h2 class="page-title">
            Giregi R.I.P
            <i class="fas fa-archive pl-3" />
            <span class="small">기억하기 보다 기록하자.</span>
          </h2>
        </div>
      </div>
    </div>
    <url-input v-model="url" @preview="loadPreview" />
    <preview
      v-if="showPreview"
      :preview="preview"
      :my-tags="tags"
      @archive="onArchive"
    />
    <archive-list :archives="archives" :my-tags="tags" />
  </div>
</template>

<script>
import { isEmpty } from 'lodash';
import ApiClient, { API } from 'api/client';
import Preview from './modules/Preview';
import UrlInput from 'views/home/modules/UrlInput';
import ArchiveList from 'views/home/modules/ArchiveList';

export default {
  name: 'Home',
  components: {
    UrlInput,
    Preview,
    ArchiveList,
  },
  mixins: [ApiClient],
  data() {
    return {
      url: '',
      preview: {},
      tags: [],
      archives: [],
    };
  },
  computed: {
    showPreview() {
      return !isEmpty(this.preview);
    },
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
      const UserApi = this.getApi(API.USER);
      const { data } = await UserApi.archives();
      this.archives = data;
    },
    async loadPreview() {
      const loader = this.$loading.show();
      const ArchiveApi = this.getApi(API.ARCHIVE);
      try {
        const archived = await ArchiveApi.getByUrl(this.url);
        if (isEmpty(archived)) {
          this.preview = await ArchiveApi.preview(this.url);
        } else {
          this.preview = archived;
        }
      } catch (e) {
        console.warn('load preview fail.', e);
      } finally {
        loader.hide();
      }
    },
    clear() {
      this.url = '';
      this.preview = {};
    },
    async onArchive({ memo, tags }) {
      const ArchiveApi = this.getApi(API.ARCHIVE);
      const loader = this.$loading.show();
      try {
        await ArchiveApi.archive(this.url, memo, tags);
        this.$dialog.open({
          title: '아카이브 요청 완료',
          message: '아카이브 요청이 완료 되었습니다.\n30초에서 몇분정도 소요 됩니다.',
          onConfirm: () => this.clear(),
        });
      } catch (e) {
        console.warn('archive fail.', e);
      } finally {
        loader.hide();
      }
    },
  },
};
</script>

<style lang="scss" scoped>
i {
  &.fas {
    margin-right: 5px;
  }
  &.far {
    margin-right: 5px;
  }
}
</style>
