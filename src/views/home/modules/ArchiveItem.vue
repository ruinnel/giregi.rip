<template>
  <div class="col-md-6 col-xl-4">
    <div class="card">
      <div class="card-body">
        <dl class="row">
          <dt class="col-3">메모:</dt>
          <dd class="col-9">{{ archive.memo }}</dd>
          <dt class="col-3">URL:</dt>
          <dd class="col-9">
            <a :href="makeUrl(archive)" target="_blank">{{ makeUrl(archive) }}</a>
          </dd>
          <dt class="col-3">Tags:</dt>
          <dd class="col-9">
            <span v-for="(tagId, idx) in archive.tagIds" :key="idx" class="badge bg-gray mr-1">
              # {{ getTagName(tagId) }}
            </span>
          </dd>
        </dl>
        <a
          role="button"
          data-toggle="collapse"
          aria-expanded="false"
          :data-target="`#detail-info-${archive.id}`"
          :aria-controls="`detail-info-${archive.id}`"
        >
          <div class="hr-text hr-text-right mt-0 mb-0">more<i class="fa fa-chevron-down" /></div>
        </a>
        <div :id="`detail-info-${archive.id}`" class="collapse">
          <dl class="row pt-3">
            <template v-for="(item, idx) in archive.summary">
              <dt :key="`dt-${idx}`" class="col-3">{{ item.name }}:</dt>
              <dd :key="`dd-${idx}`" class="col-9">{{ item.value }}</dd>
            </template>
          </dl>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { find } from 'lodash';
import config from 'config';
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
    makeUrl(archive) {
      return `${config.archivePrefix}${archive.waybackId}`;
    },
    getTagName(tagId) {
      const tag = find(this.tags, (tag) => tag.id === tagId) || {};
      return tag.name;
    },
  },
};
</script>

<style lang="scss" scoped>
</style>
