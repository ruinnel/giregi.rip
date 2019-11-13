<template>
  <vs-row>
    <div :class="hasPrev ? 'paginate' : 'paginate disable'" @click="prev">
      <vs-icon icon-pack="fas" icon="fa-chevron-left" />
    </div>
    <div class="page-content">
      <vs-table stripe :data="paging.data" class="table">
        <template slot="header">
          <h4></h4>
        </template>
        <template slot="thead">
          <vs-th>메모</vs-th>
          <vs-th class="reaction-cell" />
        </template>
        <template slot-scope="{data}">
          <vs-tr v-for="(row, idx) in data" :key="idx">
            <vs-td>
              <span class="memo-content">{{ row.content }}</span>
              <div class="memo-meta vs-row">
                <vs-col class="vs-sm-12"><vs-icon icon="person" />{{ row.writer.userEmail }}</vs-col>
                <vs-col class="vs-sm-12"><vs-icon icon="access_time" />{{ formatDateTime(row.createdAt) }}</vs-col>
              </div>
            </vs-td>
            <vs-td>
              <div class="reaction" @click="like(row)">
                <vs-icon icon-pack="far" icon="fa-thumbs-up" color="primary" />
                <span class="count">{{ formatNumber(row.like) }}</span>
              </div>
              <div class="reaction" @click="unlike(row)">
                <vs-icon icon-pack="far" icon="fa-thumbs-up" color="danger" class="fa-rotate-180" />
                <span class="count">{{ formatNumber(row.unlike) }}</span>
              </div>
            </vs-td>
          </vs-tr>
        </template>
      </vs-table>
    </div>
    <div :class="hasNext ? 'paginate' : 'paginate disable'" @click="next">
      <vs-icon icon-pack="fas" icon="fa-chevron-right" />
    </div>
  </vs-row>
</template>

<script>
import { isNull } from 'lodash';

export default {
  name: 'MemoList',
  props: {
    paging: {
      type: Object,
      default: () => ({
        data: [],
        offset: 0,
        count: 3,
        lastOffset: null,
      }),
    },
  },
  computed: {
    hasPrev() {
      const { offset } = this.paging;
      return offset > 0;
    },
    hasNext() {
      const { offset, count, lastOffset } = this.paging;
      return ((offset + count) < lastOffset) || isNull(lastOffset);
    },
  },
  methods: {
    prev() {
      const { offset, count } = this.paging;
      let error = null;
      let nextOffset = offset - count;
      if (!this.hasPrev) {
        nextOffset = 0;
        error = 'start';
      }
      this.$emit('paginate', { offset: nextOffset, count, error });
    },
    next() {
      const { offset, count } = this.paging;
      let error = null;
      let nextOffset = offset + count;
      if (!this.hasNext) {
        nextOffset = offset;
        error = 'end';
      }
      this.$emit('paginate', { offset: nextOffset, count, error });
    },
  },
};
</script>

<style scoped lang="scss">
.paginate {
  &.disable {
    cursor: not-allowed;
  }
  > i {
    font-size: 2.2em;
  }
  display: flex;
  align-items: center;
  justify-content: center;
  width: 50px;
  cursor: pointer;
}
.paginate:hover {
  background-color: rgba(222,222,222,0.25);
}

.page-content {
  width: calc(100% - 100px);
}
</style>
