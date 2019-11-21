<template>
  <vs-row>
    <div :class="hasPrev ? 'paginate' : 'paginate disable'" @click="prev">
      <vs-icon icon-pack="fas" icon="fa-chevron-left" />
    </div>
    <div class="page-content">
      <vs-card v-for="(memo, idx) in memos" :key="`memo-${idx}`">
        <div class="content">
          <div>
            <div class="memo">{{ memo.content }}</div>
            <div class="writer">- {{ memo.writer.email }} 님 -</div>
            <div class="write-at"><i class="far fa-clock" />{{ formatDateTime(memo.createdAt) }}</div>
            <div class="meta">
              <div class="reaction" @click="reaction(memo, true)">
                <vs-icon icon-pack="far" icon="fa-thumbs-up" color="primary" />
                <span class="count">{{ formatNumber(memo.like) }}</span>
              </div>
              <div class="reaction" @click="reaction(memo, false)">
                <vs-icon icon-pack="far" icon="fa-thumbs-up" color="danger" class="fa-rotate-180" />
                <span class="count">{{ formatNumber(memo.unlike) }}</span>
              </div>
            </div>
          </div>
        </div>
      </vs-card>
    </div>
    <div :class="hasNext ? 'paginate' : 'paginate disable'" @click="next">
      <vs-icon icon-pack="fas" icon="fa-chevron-right" />
    </div>
  </vs-row>
</template>

<script>
import { slice, isNull } from 'lodash';

export default {
  name: 'MemoSlider',
  props: {
    paging: {
      type: Object,
      default: () => ({
        data: [],
        offset: 0,
        count: 1,
        lastOffset: null,
      }),
    },
  },
  computed: {
    memos() {
      const { data, count } = this.paging;
      return slice(data, 0, count);
    },
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
    reaction(memo, isLike) {
      this.$emit('reaction', { isLike, memoId: memo.id });
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

.content {
  display: flex;
  justify-content: center;
  text-align: center;
  .memo {
    font-weight: 400;
    font-size: 1.15em;
    padding: 10px;
  }
  .writer {
    font-size: 0.95em;
  }
  .write-at {
    display: none;
    line-height: 0.9em;
    font-size: 0.7em;
    > i {
      margin-right: 3px;
    }
  }
  .meta {
    margin-top: 10px;
    display: flex;
    align-items: center;
    justify-content: center;
    > i {
      margin-right: 5px;
    }
    > .reaction {
      margin-left: 10px;
      cursor: pointer;
      display: inline;
      > .count {
        margin-left: 5px;
        font-weight: 400;
      }
    }
  }
}
</style>
