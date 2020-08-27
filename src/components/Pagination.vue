<template>
  <div class="con-vs-pagination vs-pagination-primary">
    <nav class="vs-pagination--nav">
      <button :class="prevClasses" :disabled="!hasPrev" @click="prev">
        <i class="vs-icon notranslate icon-scale material-icons null">chevron_left</i>
      </button>
      <span class="page-number">{{ page }} / {{ totalPage }}page</span>
      <button :class="nextClasses" :disabled="!hasNext" @click="next">
        <i class="vs-icon notranslate icon-scale material-icons null">chevron_right</i>
      </button>
    </nav>
  </div>
</template>

<script>
import { ceil } from 'lodash';

export default {
  name: 'Pagination',
  props: {
    paging: {
      type: Object,
      default: () => ({
        total: 0,
        offset: 0,
        count: 5,
      }),
    },
  },
  computed: {
    totalPage() {
      const { total, count } = this.paging;
      return ceil(total / count);
    },
    page() {
      const { offset, count } = this.paging;
      return (offset / count) + 1;
    },
    hasPrev() {
      const { offset } = this.paging;
      return offset > 0;
    },
    hasNext() {
      const { total, count, page } = this.paging;
      const lastPage = ceil(total / count);
      return page < lastPage;
    },
    prevClasses() {
      const classes = 'vs-pagination--buttons btn-prev-pagination vs-pagination--button-prev round-button';
      if (this.hasPrev) {
        return classes;
      }
      return `${classes} disabled`;
    },
    nextClasses() {
      const classes = 'vs-pagination--buttons btn-next-pagination vs-pagination--button-next round-button';
      if (this.hasNext) {
        return classes;
      }
      return `${classes} disabled`;
    },
  },
  methods: {
    prev() {
      this.$emit('prev');
    },
    next() {
      this.$emit('next');
    },
  },
};
</script>

<style scoped>
.page-number {
  margin: 0px 10px 0px 10px;
  font-weight: bold;
}

.round-button {
  border-radius: 100px !important;
}
</style>
