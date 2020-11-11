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
        <hr />
        <div class="row pl-1 pr-1">
          <dl class="row">
            <dt class="col-3">메모:</dt>
            <dd class="col-9">
              <validator rules="required" name="메모">
                <input
                  v-model="memo"
                  type="text"
                  maxlength="200"
                  class="form-control"
                  placeholder="메모"
                />
              </validator>
            </dd>
            <dt class="col-3">공개여부:</dt>
            <dd class="col-9">
              <label class="form-check form-check-inline">
                <input v-model="isPublic" class="form-check-input" type="radio" value="true" disabled>
                <span class="form-check-label">공개</span>
              </label>
              <label class="form-check form-check-inline" value="false">
                <input v-model="isPublic" class="form-check-input" type="radio" value="false">
                <span class="form-check-label">비공개</span>
              </label>
            </dd>
            <dt class="col-3">태그:</dt>
            <dd class="col-9">
              <tag-input :tags="myTags" @change="onTagChanged" />
            </dd>
          </dl>
        </div>
        <div class="align-items-center mt-2">
          <button class="btn btn-primary btn-block" :disabled="memo.length === 0" @click="onArchive">
            <i class="fas fa-archive mr-1" />
            아카이브
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { get } from 'lodash';
import TagInput from 'components/TagInput';
import ArchiveUtil from 'utils/archive';

export default {
  name: 'Preview',
  components: {
    TagInput,
  },
  props: {
    preview: {
      type: Object,
      default: () => null,
    },
    myTags: {
      type: Array,
      default: () => [],
    },
  },
  data() {
    return {
      memo: '',
      tags: [],
      isPublic: 'false',
    };
  },
  computed: {
    alreadyArchived() {
      return get(this.preview, 'id', 0) > 0;
    },
    summary() {
      return ArchiveUtil.convert(this.preview, this);
    },
  },
  methods: {
    onArchive() {
      this.$emit('archive', { memo: this.memo, tags: this.tags });
    },
    onTagChanged(tags) {
      this.tags = tags;
    },
  },
};
</script>

<style scoped>

</style>
