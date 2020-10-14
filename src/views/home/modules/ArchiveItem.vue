<template>
  <div class="col-md-6 col-xl-4">
    <div class="card">
      <div class="card-header flex-row justify-content-between">
        <h6 class="card-title"><i class="far fa-sticky-note mr-1" />{{ archive.memo }}</h6>
        <span class="text-black-50 small"><i class="far fa-clock mr-1" />{{ formatDateTime(archive.createdAt) }}</span>
      </div>
      <div class="card-body">
        <dl class="row">
          <div>
            <label class="form-label text-black-50"><i class="fas fa-globe-asia" /> URL</label>
            <div class="form-control form-control-flush pl-1">
              <a :href="webPageUrl(archive)" target="_blank">
                {{ webPageUrl(archive) }}
              </a>
              <a v-if="hasArchive(archive)" :href="archiveUrl(archive)" target="_blank" class="ml-2">
                <i class="fa fa-archive mr-1" />Archive
              </a>
            </div>
          </div>
          <div v-if="getTags(archive).length > 0">
            <label class="form-label text-black-50"><i class="fa fa-tag" /> Tags</label>
            <div class="form-control form-control-flush pl-1">
              <span v-for="(tag, idx) in getTags(archive)" :key="idx" class="badge bg-gray mr-1">
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
          <div class="hr-text hr-text-right mt-0 mb-0 text-primary">more<i class="fa fa-chevron-down" /></div>
        </a>
        <div :id="`detail-info-${archive.id}`" class="collapse">
          <dl class="row pt-3">
            <div v-for="(item, idx) in summary(archive)" :key="idx">
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
  methods: {
    webPageUrl(archive) {
      return get(archive, 'webPage.url', '');
    },
    hasArchive(archive) {
      return !isEmpty(archive.waybackId);
    },
    archiveUrl(archive) {
      const { waybackId } = archive;
      if (!isEmpty(waybackId)) {
        return `${config.archivePrefix}${waybackId}`;
      } else {
        return '';
      }
    },
    getTags(archive) {
      const tags = map(archive.tagIds, (id) => find(this.tags, (tag) => (tag.id === id)));
      return compact(tags);
    },
    summary(archive) {
      return ArchiveUtil.convert(archive, this);
    },
  },
};
</script>

<style lang="scss" scoped>
.form-label {
  font-size: 0.75rem;
  margin-bottom: 0;
}
</style>
