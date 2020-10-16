<template>
  <div class="col-sm-12 col-lg-12">
    <div class="">
      <h6 class="card-title">
        <span><i class="fa fa-tag mr-1" />Tags</span>
        <a role="button" data-toggle="collapse" data-target="#tag-list" class="align-content-end">
          <span class="ml-1"><i class="fa fa-chevron-down" /></span>
        </a>
      </h6>
      <div id="tag-list" class="card card-sm collapse">
        <div class="card-body">
          <div class="form-selectgroup form-selectgroup-pills">
            <label v-for="(tag, idx) in myTags" :key="idx" class="form-selectgroup-item">
              <input
                type="checkbox"
                :name="`tag-${tag.id}`"
                :value="tag.id"
                class="form-selectgroup-input"
              >
              <span class="form-selectgroup-label"><i class="fa fa-hashtag mr-1" />{{ tag.name }}</span>
            </label>
          </div>
        </div>
      </div>
      <div class="row mt-2">
        <archive-item
          v-for="(archive, idx) in archives"
          :key="idx"
          :archive="archive"
          :tags="myTags"
        />
      </div>
    </div>
  </div>
</template>

<script>
import { get } from 'lodash';
import ArchiveItem from 'components/ArchiveItem';

export default {
  name: 'Preview',
  components: {
    ArchiveItem,
  },
  props: {
    archives: {
      type: Array,
      default: () => [],
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
    };
  },
  computed: {
    alreadyArchived() {
      return get(this.preview, 'id', 0) > 0;
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
.form-selectgroup-label {
  font-size: 0.5rem;
}
h6.card-title {
  font-size: 0.8rem;
}
</style>
