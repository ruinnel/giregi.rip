<template>
  <div class="col-sm-12 col-lg-12">
    <div class="card">
      <div class="card-body">
        <h3 class="card-title">Preview <i v-if="alreadyArchived" class="fa fa-check text-success" /></h3>
        <div v-if="summary.length > 0">
          <dl class="row">
            <template v-if="alreadyArchived">
              <dt class="col-3">상태:</dt>
              <dd class="col-9">
                <span class="strong"> 이미 아카이브됨</span>
                <i class="fa fa-check text-success" />
              </dd>
            </template>
            <template v-for="(item, idx) in summary">
              <dt :key="`dt-${idx}`" class="col-3">{{ item.name }}:</dt>
              <dd :key="`dd-${idx}`" class="col-9">{{ item.value }}</dd>
            </template>
          </dl>
        </div>
        <div v-else>미리보기 결과가 표시됩니다.</div>

        <div v-if="showArchive" class="align-items-center mt-2">
          <button class="btn btn-primary btn-block" @click="onArchive">
            <i class="fas fa-archive" />
            아카이브
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { map, isEmpty, get } from 'lodash';
import ArchiveUtil from 'utils/archive';

export default {
  name: 'Preview',
  props: {
    preview: {
      type: Object,
      default: () => null,
    },
  },
  computed: {
    showArchive() {
      return !isEmpty(this.preview);
    },
    alreadyArchived() {
      return get(this.preview, 'id', 0) > 0;
    },
    summary() {
      return map(this.preview.summary, ({ name, value }) => {
        let converted = value;
        if (name === 'createdAt' || name === 'updatedAt') {
          converted = this.formatDateTime(value);
        }
        return { name: ArchiveUtil.labels[name], value: converted };
      });
    },
  },
  methods: {
    onArchive() {
      this.$emit('archive');
    },
  },
};
</script>

<style scoped>

</style>
