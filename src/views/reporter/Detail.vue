<template>
  <vs-popup
    title="기자 정보"
    :active.sync="active"
    class="popup-content">
    <div>
      <card :reporter="reporter" @reaction="reaction" @memo="onMemo" />
      <memo-list :paging="paging" @paginate="paginate" />
    </div>
  </vs-popup>
</template>

<script>
import { get, size, assign, map, isEmpty, concat, slice } from 'lodash';
import MemoApi from 'api/memo';
import ReactionApi from 'api/reaction';
import Card from './module/Card';
import MemoList from './module/MemoList';

const COUNT = 3;
export default {
  name: 'Detail',
  components: {
    Card,
    MemoList,
  },
  props: {
    show: {
      type: Boolean,
      default: () => false,
    },
    reporter: {
      type: Object,
      default: () => '',
    },
  },
  data() {
    return {
      active: false,
      memo: '',
      memos: [],
      paging: {
        data: [],
        offset: 0,
        count: COUNT,
        lastOffset: null,
      },
    };
  },
  computed: {},
  watch: {
    show(val) {
      this.active = val;
    },
    active(val) {
      if (!val) {
        this.$emit('close', val);
      }
    },
    reporter(val) {
      if (!isEmpty(val)) {
        this.paginate({});
      }
    },
  },
  methods: {
    getAgency(reporter) {
      return get(reporter.agencies, '[0].name');
    },
    async getMemo({ offset = 0, count = COUNT }) {
      const data = await MemoApi.search({ reporterId: this.reporter.id, offset, count });
      return data;
    },
    async reaction({ isLike, reporterId }) {
      await ReactionApi.toggle({ mode: 'reporter', id: reporterId, isLike });
      this.$emit('refresh', this.reporter.id);
    },
    async onMemo({ id, memo }) {
      this.$vs.loading();
      if (size(id) > 0) {
        await MemoApi.update({ id, memo });
      } else {
        await MemoApi.create({ reporterId: this.reporter.id, memo });
      }
      const newMemos = map(this.memos, (m) => {
        if (id === m.id) {
          return assign(m, { content: memo });
        } else {
          return m;
        }
      });
      this.$set(this.reporter, 'myMemo', assign(this.reporter.myMemo, { content: memo }));
      this.memos = newMemos;
      this.$vs.loading.close();
      this.$vs.notify({ color: 'success', title: '저장 완료', text: '메모가 저장 되었습니다.' });
    },
    async paginate({ offset = 0, count = COUNT, error }) {
      if (size(error) > 0) {
        const text = error === 'start' ? '첫 페이지 입니다.' : '마지막 페이지 입니다.';
        return this.$vs.notify({ color: 'warning', title: '페이지 이동 불가', text });
      }

      const paging = assign({}, this.paging, { offset, count });
      if ((size(this.memos) - 1) > offset) {
        paging.data = slice(this.memos, offset, offset + count);
      } else {
        paging.data = await this.getMemo({ offset, count });
        this.memos = concat(this.memos, paging.data);
        if (size(paging.data) < count) {
          paging.lastOffset = size(this.memos) - 1;
        }
      }
      this.paging = paging;
    },
  },
};
</script>

<style scoped>

</style>
