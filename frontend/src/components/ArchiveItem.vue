<template>
  <div class="col-md-6 col-xl-4">
    <div class="card">
      <div class="card-header">
        <div class="row w-100">
          <div class="card-title col-12">
            <i :class="`${icon} memo-icon mr-1`" />
            <span class="memo-text">{{ archive.memo }}</span>
          </div>
          <div class="text-black-50 small col-12">
            <i class="far fa-clock mr-1" />
            <span>{{ formatDate(archive.createdAt) }}</span>
          </div>
        </div>
      </div>
      <div class="card-body">
        <dl class="row">
          <div>
            <label class="form-label text-black-50"><i class="fas fa-heading" /> 제목</label>
            <div class="form-control form-control-flush pl-1">{{ archive.title }}</div>
          </div>
          <div>
            <label class="form-label text-black-50"><i class="fas fa-globe-asia" /> URL</label>
            <div class="form-control form-control-flush pl-1">
              <a :href="webPageUrl" target="_blank">
                <i class="fa fa-bookmark mr-1" /><span>URL</span>
              </a>
              <a v-if="hasArchive" :href="archiveUrl" target="_blank" class="ml-3">
                <i class="fa fa-archive mr-1" /><span>아카이브</span>
              </a>
            </div>
          </div>
          <div v-if="getTags.length > 0">
            <label class="form-label text-black-50"><i class="fa fa-tag" /> Tags</label>
            <div class="form-control form-control-flush pl-1">
              <span v-for="(tag, idx) in getTags" :key="idx" class="badge bg-gray mr-1">
                # {{ tag.name }}
              </span>
            </div>
          </div>
        </dl>
        <a
          role="button"
          data-toggle="collapse"
          :data-target="`#detail-info-${archive.id}`"
        >
          <div class="hr-text hr-text-right mt-0 mb-0 text-primary">상세정보 보기<i class="fa fa-chevron-down" /></div>
        </a>
        <div :id="`detail-info-${archive.id}`" class="collapse">
          <dl class="row pt-3">
            <div v-for="(item, idx) in summary" :key="idx">
              <label class="form-label text-black-50">
                <i v-if="item.icon" :class="item.icon" /> {{ item.name }}
              </label>
              <div class="form-control form-control-flush pl-1">{{ item.value }}</div>
            </div>
          </dl>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { find, get, isEmpty, map, compact } from 'lodash';
import config from 'config';
import ArchiveUtil from 'utils/archive';
export default {
  name: 'ArchiveItem',
  props: {
    archive: {
      type: Object,
      required: true,
    },
    tags: {
      type: Array,
      required: true,
    },
  },
  computed: {
    icon() {
      return this.archive.public ? 'fas fa-lock-open' : 'fas fa-lock';
    },
    webPageUrl() {
      return get(this.archive, 'webPage.url', '');
    },
    hasArchive() {
      return !isEmpty(this.archive.waybackId);
    },
    archiveUrl() {
      const { waybackId } = this.archive;
      if (!isEmpty(waybackId)) {
        return `${config.archivePrefix}${waybackId}`;
      } else {
        return '';
      }
    },
    getTags() {
      const tags = map(this.archive.tagIds, (id) => find(this.tags, (tag) => (tag.id === id)));
      return compact(tags);
    },
    summary() {
      return ArchiveUtil.convert(this.archive, this);
    },
  },
};
</script>

<style lang="scss" scoped>
.form-label {
  font-size: 0.75rem;
  margin-bottom: 0;
}
.memo-icon {
  padding-top: 4px;
  vertical-align: top;
}
.memo-text {
  display: inline-block;
  width: 90%;
  overflow-x: auto;
}
.created-at {
  min-width: 150px;
}
</style>
