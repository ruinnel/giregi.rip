<template>
  <div class="col-sm-12 col-lg-12 mt-2">
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
            <dt class="col-3">제목:</dt>
            <dd class="col-9">
              <validator rules="required" name="제목">
                <input
                  v-model="title"
                  type="text"
                  maxlength="200"
                  class="form-control"
                  placeholder="제목"
                />
              </validator>
            </dd>
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
              <tag-input ref="tagInput" :my-tags="myTags" :tags="tags" @change="onTagChanged" />
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
import { get, find, isEmpty } from 'lodash';
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
      title: this.getTitle(),
      memo: '',
      tags: this.getTags(),
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
  watch: {
    preview() {
      this.title = this.getTitle();
      this.memo = '';
      this.tags = this.getTags();
      this.isPublic = 'false';
      this.$nextTick(() => {
        if (this.$refs.tagInput) {
          this.$refs.tagInput.refreshItems();
        }
      });
    },
  },
  methods: {
    onArchive() {
      this.$emit('archive', { title: this.title, memo: this.memo, tags: this.tags });
    },
    onTagChanged(tags) {
      this.tags = tags;
    },
    getTitle() {
      const summaries = get(this.preview, 'summary', []);
      const title = find(summaries, (summary) => summary.name === 'title');
      return get(title, 'value', '');
    },
    getTags() {
      const summaries = get(this.preview, 'summary', []);
      const writer = find(summaries, (summary) => summary.name === 'writer');
      const agency = find(summaries, (summary) => summary.name === 'agency');
      const cowriter = find(summaries, (summary) => summary.name === 'cowriter');

      const tags = [];
      if (!isEmpty(writer)) tags.push(this.convertToTag(this.myTags, writer));
      if (!isEmpty(agency)) tags.push(this.convertToTag(this.myTags, agency));
      if (!isEmpty(cowriter)) tags.push(this.convertToTag(this.myTags, cowriter));
      return tags;
    },
    convertToTag(myTags, summary) {
      const exists = find(myTags, (tag) => tag.name === summary.value);
      if (isEmpty(exists)) {
        return { id: -1, name: summary.value };
      } else {
        return exists;
      }
    },
  },
};
</script>

<style scoped>

</style>
