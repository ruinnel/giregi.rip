<template>
  <div id="wrapper" class="padding-top">
    <navigator />
    <header id="header">
      <h1>기자</h1>
      <p>기자 목록</p>
    </header>
    <validator rules="required" name="기사 URL or 기자이름">
      <input
        ref="input"
        v-model="input"
        type="text"
        class="url-input"
        placeholder="기자이름"
        @keyup.enter="onEnter"
      />
    </validator>
    <div id="main">
      <section id="content" class="main reporter">
        <pagination class="pagination" :paging="paging" @prev="prev" @next="next" />
        <reporter-card
          v-for="(item, idx) in paging.data"
          :key="idx"
          :reporter="item"
          @click.native="selectReporter(item)"
          @memo="onMemo(item, $event)"
          @reaction="onReaction"
        />
        <div v-if="!paging.data">데이터가 없습니다.</div>
        <pagination class="pagination" :paging="paging" @prev="prev" @next="next" />
      </section>
    </div>

    <site-footer />

    <reporter-detail v-if="showDetail" :reporter="reporter" @close="closeDetail" />
  </div>
</template>

<script>
import { isEmpty, size, map, assign, find, ceil } from 'lodash';
import Navigator from 'components/Navigator';
import ReporterCard from 'components/ReporterCard';
import ReporterDetail from './Detail';
import ApiClient, { API } from 'api/client';

const REPORTER_COUNT = 5;
export default {
  name: 'Reporter',
  components: {
    Navigator,
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
      input: '',
      reporters: {},
      reporter: {},
      paging: {
        total: 0,
        data: [],
        page: 1,
        offset: 0,
        count: 5,
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
      const { total, count, page } = this.paging;
      const lastPage = ceil(total / count);
      return page < lastPage;
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
      let { total, data } = await ReporterApi.search({ name: this.input, offset, count });
      if (size(data) > 0) {
        const reporterIds = map(data, (reporter) => reporter.id);
        const myMemos = await MemoApi.my(reporterIds);
        data = map(data, (reporter) => assign(reporter, {
          myMemo: find(myMemos, (memo) => (memo.reporter.id === reporter.id)),
        }));
      }

      this.$vs.loading.close();
      return { total, data };
    },
    async paginate({ page = 1, offset = 0, count = REPORTER_COUNT, error, refresh = false }) {
      if (size(error) > 0) {
        const text = error === 'start' ? '첫 페이지 입니다.' : '마지막 페이지 입니다.';
        return this.$vs.notify({ color: 'warning', title: '페이지 이동 불가', text });
      }

      const paging = assign({}, this.paging, { page, offset, count });
      const data = this.reporters[page];
      if (!refresh && size(data) > 0) {
        paging.data = data;
      } else {
        const { total, data } = await this.getReporters({ offset, count });
        paging.total = total;
        paging.data = data;
        this.reporters[page] = data;
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
      const { page, offset, count } = this.paging;
      let error = null;
      let nextOffset = offset - count;
      let nextPage = page - 1;
      if (!this.hasPrev) {
        nextOffset = 0;
        nextPage = 1;
        error = 'start';
      }
      this.paginate({ page: nextPage, offset: nextOffset, count, error });
    },
    next() {
      const { page, offset, count } = this.paging;
      let error = null;
      let nextOffset = offset + count;
      let nextPage = page + 1;
      if (!this.hasNext) {
        nextOffset = offset;
        nextPage = page;
        error = 'end';
      }
      this.paginate({ page: nextPage, offset: nextOffset, count, error });
    },
    onEnter() {
      if (size(this.input) > 1) {
        this.reporters = {};
        this.paging = assign({}, this.paging, { data: [], offset: 0 });
        this.paginate({});
      } else {
        this.$vs.notify({ color: 'warning', title: '검색 실패', text: '2글자 이상 입력 해주세요.' });
      }
    },
    async onReaction({ isLike, reporterId }) {
      const ReactionApi = this.getApi(API.REACTION);
      if (reporterId) {
        await ReactionApi.toggle({ mode: 'reporter', id: reporterId, isLike });
        await this.paginate(assign({}, this.paging, { refresh: true }));
      }
    },
  },
};
</script>

<style scoped lang="scss">
.padding-top {
  padding-top: 50px;
}
.pagination {
  margin-bottom: 10px;
}
input.url-input {
  margin-bottom: 20px;
  font-size: 22px;
  color: #1e252d !important;
  background-color: rgba(255, 255, 255, 0.75) !important;
}
input.url-input::placeholder {
  color: #717981 !important;
}
</style>
