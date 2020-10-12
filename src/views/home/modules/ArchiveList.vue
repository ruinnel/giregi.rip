<template>
  <div class="col-sm-12 col-lg-12">
    <div class="card">
      <div class="card-body">
        <h3 class="card-title"><i class="fa fa-archive mr-1" />아카이브 목록</h3>
        <div class="card card-body">
          <div class="form-selectgroup form-selectgroup-pills">
            <label v-for="(tag, idx) in myTags" :key="idx" class="form-selectgroup-item">
              <input type="checkbox"
                     :name="`tag-${tag.id}`"
                     :value="tag.id"
                     class="form-selectgroup-input"
              >
              <span class="form-selectgroup-label"><i class="fa fa-hashtag mr-1" />{{ tag.name }}</span>
            </label>
            <label v-for="(tag, idx) in myTags" :key="idx" class="form-selectgroup-item">
              <input type="checkbox"
                     :name="`tag-${tag.id}`"
                     :value="tag.id"
                     class="form-selectgroup-input"
              >
              <span class="form-selectgroup-label"><i class="fa fa-hashtag mr-1" />{{ tag.name }}</span>
            </label>
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
  </div>
</template>

<script>
import { map, get } from 'lodash';
import ArchiveUtil from 'utils/archive';
import ArchiveItem from 'views/home/modules/ArchiveItem';

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
</style>
