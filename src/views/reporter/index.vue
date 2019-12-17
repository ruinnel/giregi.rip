<template>
  <div id="wrapper">
    <header id="header">
      <h1>기자</h1>
      <p>기자 목록</p>
    </header>
    <div id="main">
      <section id="content" class="main reporter">
        <pagination class="pagination" :paging="paging" @prev="prev" @next="next" />
        <reporter-card
          v-for="(item, idx) in paging.data"
          :key="idx"
          :reporter="item"
          @click.native="selectReporter(item)"
          @memo="onMemo(item, $event)"
        />
        <pagination class="pagination" :paging="paging" @prev="prev" @next="next" />
      </section>
    </div>

    <site-footer />

    <reporter-detail :show="showDetail" :reporter="reporter" @close="closeDetail" />
  </div>
</template>

<script>
import { isEmpty, size, map, assign, find, isNil, slice, concat } from 'lodash';
import ReporterCard from 'components/ReporterCard';
import ReporterDetail from './Detail';
import ApiClient, { API } from 'api/client';

const REPORTER_COUNT = 5;
export default {
  name: 'Reporter',
  components: {
    ReporterCard,
    ReporterDetail,
  },
  mixins: [ApiClient],
  props: {
    reporterId: {
      type: String,
      default: () => '',
    },
  },
  data() {
    return {
      active: true,
      reporters: [],
      reporter: {},
      paging: {
        data: [],
        offset: 0,
        count: 5,
        lastOffset: null,
      },
    };
  },
  computed: {
    showDetail() {
      return !isEmpty(this.reporter);
    },
    hasPrev() {
      const { offset } = this.paging;
      return offset > 0;
    },
    hasNext() {
      const { offset, count, lastOffset } = this.paging;
      return ((offset + count) < lastOffset) || isNil(lastOffset);
    },
  },
  mounted() {
    this.paginate({});
  },
  methods: {
    async getReporters({ offset = 0, count = REPORTER_COUNT }) {
      this.$vs.loading();
      const MemoApi = this.getApi(API.MEMO);
      const ReporterApi = this.getApi(API.REPORTER);
      return ReporterApi.search({ offset, count })
        .then((reporters) => {
          const reporterIds = map(reporters, (reporter) => reporter.id);
          return Promise.all([reporters, MemoApi.my(reporterIds)]);
        })
        .then(([reporters, myMemos]) => {
          return map(reporters, (reporter) => assign(reporter, {
            myMemo: find(myMemos, (memo) => (memo.reporter.id === reporter.id)),
          }));
        })
        .catch((err) => {
          console.log('get reporter error', err);
          this.$vs.notify({
            color: 'warning',
            title: '로딩 실패',
            text: '기자 정보를 가져오는데 실패 하였습니다.',
          });
        })
        .finally(() => this.$vs.loading.close());
    },
    async paginate({ offset = 0, count = REPORTER_COUNT, error }) {
      if (size(error) > 0) {
        const text = error === 'start' ? '첫 페이지 입니다.' : '마지막 페이지 입니다.';
        return this.$vs.notify({ color: 'warning', title: '페이지 이동 불가', text });
      }

      const paging = assign({}, this.paging, { offset, count });
      if ((size(this.reporters) - 1) > offset) {
        paging.data = slice(this.reporters, offset, offset + count);
      } else {
        paging.data = await this.getReporters({ offset, count });
        this.reporters = concat(this.reporters, paging.data);
        if (size(paging.data) < count) {
          paging.lastOffset = size(this.reporters) - 1;
        }
      }
      this.paging = paging;
    },
    closeDetail() {
      this.reporter = {};
    },
    selectReporter(reporter) {
      this.reporter = reporter;
    },
    async onMemo(reporter, { id, memo }) {
      if (size(memo) === 0) return;
      const MemoApi = this.getApi('memo');
      let request;
      if (size(id) > 0) {
        request = MemoApi.update({ id, memo });
      } else {
        request = MemoApi.create({ reporterId: reporter.id, memo });
      }
      this.$vs.loading();
      request
        .then(() => {
          this.reporters = map(this.reporters, (r) => {
            if (r.id === reporter.id) {
              r.myMemo = { id, content: memo };
            }
            return assign({}, r);
          });
        })
        .catch(() => {
          this.$vs.notify({
            color: 'warning',
            title: '메모 저장 실패',
            text: '메모 저장 요청이 실패 하였습니다.',
          });
        })
        .finally(() => this.$vs.loading.close());
    },
    prev() {
      const { offset, count } = this.paging;
      let error = null;
      let nextOffset = offset - count;
      if (!this.hasPrev) {
        nextOffset = 0;
        error = 'start';
      }
      this.paginate({ offset: nextOffset, count, error });
    },
    next() {
      const { offset, count } = this.paging;
      let error = null;
      let nextOffset = offset + count;
      if (!this.hasNext) {
        nextOffset = offset;
        error = 'end';
      }
      this.paginate({ offset: nextOffset, count, error });
    },
  },
};
</script>

<style scoped>
.pagination {
  margin-bottom: 10px;
}
</style>
