<template>
  <div class="col-sm-12 col-lg-12">
    <div class="card card-body">
      <div class="row w-100">
        <div class="col-xs-12 col-md-1 d-flex align-items-center">
          <span class="align-baseline"><i class="fa fa-search mr-1" />검색</span>
        </div>
        <div class="col-xs-12 col-md-4">
          <input
            v-model="keyword"
            type="text"
            class="form-control"
            name="example-text-input"
            placeholder="검색어 입력">
        </div>
        <div class="col-xs-12 d-block d-md-none mt-1" />
        <div class="col-xs-12 col-md-1 d-flex align-items-center">
          <span><i class="fa fa-tag mr-1" />태그</span>
        </div>
        <div class="col-xs-12 col-md-4">
          <select ref="select" class="form-select form-control">
            <option v-for="(tag, idx) in myTags" :key="idx" :value="tag.id">{{ tag.name }}</option>
          </select>
        </div>
        <div class="col-xs-12 d-block d-md-none mt-1" />
        <div class="col-xs-12 col-md-2">
          <button
            class="btn btn-primary btn-block"
            @click="onSearch"
          >
            <i class="fa fa-search mr-1" />검색
          </button>
        </div>
      </div>
    </div>
    <div class="card card-body">
      <div v-if="total" class="card-title hr-text">검색결과: {{ formatNumber(total) }} 건</div>
    </div>
    <div class="row mt-2">
      <archive-item
        v-for="(archive, idx) in archives"
        :key="idx"
        :archive="archive"
        :tags="myTags"
      />
    </div>
    <div v-if="total > archives.length" class="row card card-body cursor-pointer" @click="onMore">
      <div class="hr-text mt-0 mb-0 text-primary">more<i class="fa fa-chevron-down" /></div>
    </div>
  </div>
</template>

<script>
import { get, toNumber, first } from 'lodash';
import ArchiveItem from 'components/ArchiveItem';
import $ from 'jquery';
import 'selectize';
import 'selectize/dist/css/selectize.css';

export default {
  name: 'Preview',
  components: {
    ArchiveItem,
  },
  props: {
    total: {
      type: Number,
      default: 0,
    },
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
      selectize: null,
      tagId: null,
      keyword: '',
    };
  },
  computed: {
    alreadyArchived() {
      return get(this.preview, 'id', 0) > 0;
    },
  },
  watch: {
    myTags() {
      this.myTags.forEach((tag) => {
        this.selectize.addOption({ value: `${tag.id}`, text: tag.name });
      });
    },
  },
  beforeDestroy() {
    if (this.selectize) {
      this.selectize.off('change');
    }
  },
  mounted() {
    const select = $(this.$refs.select).selectize({
      maxItems: 1,
    });
    this.selectize = select[0].selectize;

    this.selectize.on('change', () => {
      const value = first(this.selectize.items);
      this.tagId = toNumber(value);
    });
  },
  methods: {
    onSearch() {
      const { keyword, tagId } = this;
      this.$emit('search', { keyword, tagId });
    },
    onMore() {
      this.$emit('more');
    },
  },
};
</script>

<style scoped>
.cursor-pointer {
  cursor: pointer;
}
.hr-text {
  font-size: 0.75rem;
}
</style>
