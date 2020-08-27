<template>
  <vs-popup
    title="기자 정보"
    :active="true"
    class="popup-content"
    @close="onClose">
    <div>
      <reporter-card :reporter="reporterData" @reaction="reaction" @memo="onMemo">
        <memo-slider :paging="pagingMemo" @paginate="paginateMemo" @reaction="reaction" />
      </reporter-card>
      <news-list :paging="pagingNews" @paginate="paginateNews" @reaction="reaction" />
    </div>
  </vs-popup>
</template>

<script>
import { assign, concat, get, map, size, slice } from 'lodash';
import ReporterCard from 'components/ReporterCard';
import MemoSlider from './module/MemoSlider';
import NewsList from './module/NewsList';
import ApiClient, { API } from 'api/client';

const MEMO_COUNT = 1;
const NEWS_COUNT = 5;
export default {
  name: 'ReporterDetail',
  components: {
    NewsList,
    ReporterCard,
    MemoSlider,
  },
  mixins: [ApiClient],
  props: {
    reporter: {
      type: Object,
      default: () => ({}),
    },
  },
  data() {
    return {
      reporterData: this.reporter,
      memos: [],
      news: [],
      pagingNews: {
        data: [],
        offset: 0,
        count: NEWS_COUNT,
        lastOffset: null,
      },
      pagingMemo: {
        data: [],
        offset: 0,
        count: MEMO_COUNT,
        lastOffset: null,
      },
    };
  },
  created() {
    this.news = [];
    this.memos = [];
    this.paginateMemo({});
    this.paginateNews({});
    this.getMyMemo();
  },
  methods: {
    getAgency(reporter) {
      return get(reporter.agencies, '[0].name');
    },
    async getMemo({ offset = 0, count = MEMO_COUNT }) {
      const MemoApi = this.getApi(API.MEMO);
      const data = await MemoApi.search({ reporterId: this.reporterData.id, offset, count });
      return data;
    },
    async getNews({ offset = 0, count = NEWS_COUNT }) {
      const ReporterApi = this.getApi(API.REPORTER);
      const data = await ReporterApi.news(this.reporterData.id, { offset, count });
      return data;
    },
    async getMyMemo() {
      const MemoApi = this.getApi(API.MEMO);
      const data = await MemoApi.my([this.reporterData.id]);
      this.$set(this.reporterData, 'myMemo', data);
    },
    async reaction({ isLike, reporterId, memoId, newsId }) {
      const ReactionApi = this.getApi(API.REACTION);
      const ReporterApi = this.getApi(API.REPORTER);
      const MemoApi = this.getApi(API.MEMO);
      const NewsApi = this.getApi(API.NEWS);
      if (reporterId) {
        await ReactionApi.toggle({ mode: 'reporter', id: reporterId, isLike });
        this.reporterData = await ReporterApi.get(reporterId);
      }
      if (memoId) {
        await ReactionApi.toggle({ mode: 'memo', id: memoId, isLike });
        const memo = await MemoApi.get(memoId);
        this.memos = map(this.memos, (m) => {
          if (m.id === memo.id) return memo;
          return m;
        });
        this.pagingMemo.data = [memo];
      }
      if (newsId) {
        await ReactionApi.toggle({ mode: 'news', id: newsId, isLike });
        const news = await NewsApi.get(newsId);
        const newNewsList = map(this.news, (n) => {
          if (n.id === news.id) return news;
          return n;
        });
        this.news = newNewsList;
        const { offset, count } = this.pagingNews;
        this.pagingNews.data = slice(newNewsList, offset, offset + count);
      }
    },
    async onMemo({ id, memo }) {
      const MemoApi = this.getApi(API.MEMO);
      this.$vs.loading();
      if (size(id) > 0) {
        await MemoApi.update({ id, memo });
      } else {
        await MemoApi.create({ reporterId: this.reporterData.id, memo });
      }
      const newMemos = map(this.memos, (m) => {
        if (id === m.id) {
          return assign(m, { content: memo });
        } else {
          return m;
        }
      });
      this.$set(this.reporterData, 'myMemo', assign(this.reporterData.myMemo, { content: memo }));
      this.memos = newMemos;
      this.$vs.loading.close();
      this.$vs.notify({ color: 'success', title: '저장 완료', text: '메모가 저장 되었습니다.' });
    },
    async paginateMemo({ offset = 0, count = MEMO_COUNT, error }) {
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
      this.pagingMemo = paging;
    },
    async paginateNews({ offset = 0, count = NEWS_COUNT, error }) {
      if (size(error) > 0) {
        const text = error === 'start' ? '첫 페이지 입니다.' : '마지막 페이지 입니다.';
        return this.$vs.notify({ color: 'warning', title: '페이지 이동 불가', text });
      }

      const paging = assign({}, this.paging, { offset, count });
      if ((size(this.news) - 1) > offset) {
        paging.data = slice(this.news, offset, offset + count);
      } else {
        paging.data = await this.getNews({ offset, count });
        this.news = concat(this.news, paging.data);
        if (size(paging.data) < count) {
          paging.lastOffset = size(this.news) - 1;
        }
      }
      this.pagingNews = paging;
    },
    onClose() {
      this.$emit('close');
    },
  },
};
</script>

<style scoped>

</style>
