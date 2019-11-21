<template>
  <validation-observer ref="validator">
    <header id="header" class="alt">
      <div class="input-container">
        <validator rules="required" name="기사 URL or 기자이름">
          <input
            ref="input"
            v-model="input"
            type="text"
            class="url-input"
            placeholder="기사 URL or 기자이름"
            @keyup.enter="onEnter"
            @blur="onBlur"
          />
        </validator>
        <floating-news
          :rect="rect"
          :active="showPreview"
          :preview="preview"
          @register="createNews"
          @close="closePreview()"
          @reaction="reaction"
        />
        <floating-reporter
          :rect="rect"
          :active="showReporter"
          :reporters="reporters"
          @select="getReporter($event.id)"
          @close="closeReporter()"
        />
      </div>
      <h1 class="title">R . I . P</h1>
      <p class="sub-title">삼가 기레기의 명복을 빕니다.</p>
      <p class="photo-by">
        Photo by
        <a target="_blank" href="https://unsplash.com/@mattbotsford?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText">Matt Botsford</a>
        on <a target="_blank" href="https://unsplash.com/search/photos/tombstone?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText">Unsplash</a>
      </p>
    </header>

    <!-- popup -->
    <reporter-detail
      :show="showReporterDetail"
      :reporter="reporter"
      @close="clearData"
    />
  </validation-observer>
</template>

<script>
import { debounce, isEmpty, size, first, map } from 'lodash';
import NewsApi from 'api/news';
import ReporterApi from 'api/reporter';
import MemoApi from 'api/memo';
import CommentApi from 'api/comment';
import ReactionApi from 'api/reaction';
import FloatingNews from './FloatingNews';
import FloatingReporter from './FloatingReporter';
import ReporterDetail from 'views/reporter/Detail';

export default {
  name: 'HomeHeader',
  components: {
    FloatingNews,
    FloatingReporter,
    ReporterDetail,
  },
  data() {
    return {
      input: '', // 'http://news.chosun.com/site/data/html_dir/2019/08/27/2019082700128.html',
      rect: {},
      onResize: null,
      showPreview: false,
      showReporter: false,
      preview: {},
      reporters: [],
      reporter: {},
      onEnter: debounce(this.getPreview, 300),
    };
  },
  computed: {
    showReporterDetail() {
      return !isEmpty(this.reporter);
    },
  },
  mounted() {
    this.onResize = debounce(() => {
      this.refreshRect();
    }, 100);
    window.addEventListener('resize', this.onResize);
    this.refreshRect();
  },
  beforeDestroy() {
    window.removeEventListener('resize', this.onResize);
  },
  methods: {
    refreshRect() {
      const inputElem = this.$refs.input;
      const { x, y, width, height } = inputElem.getBoundingClientRect();
      this.rect = {
        left: x,
        top: y + height,
        width,
      };
    },
    clearData() {
      this.input = '';
      this.reporter = {};
      this.reporters = [];
      this.preview = {};
      this.showReporter = false;
      this.showPreview = false;
    },
    async getPreview() {
      const valid = await this.$refs.validator.validate();
      if (valid) {
        const regex = /^http(s*):\/\/.+$/i;
        const isUrl = this.input.match(regex);
        if (isUrl) {
          this.showPreview = true;
          this.showReporter = false;
          this.$vs.loading();
          this.preview = await NewsApi.preview(this.input)
            .catch((err) => {
              console.log('preview error', err.response);
              this.$vs.notify({
                color: 'warning',
                title: '지원되지 않는 url 입니다.',
                text: 'naver, daum 뉴스 url만 지원합니다.',
              });
            })
            .finally(() => this.$vs.loading.close());
        } else {
          this.showPreview = false;
          this.$vs.loading();
          this.reporters = await ReporterApi.search({ name: this.input })
            .then(async (reporters) => {
              if (size(reporters) === 1) {
                this.reporter = await ReporterApi.get(first(reporters).id);
                this.showReporter = false;
              } else if (size(reporters) > 1) {
                this.showReporter = true;
              }
              return reporters;
            })
            .catch(() => {
              const options = {
                color: 'warning',
                title: '요청이 실패하였습니다.',
                text: '데이터가 존재하지 않습니다.',
              };
              this.$vs.notify(options);
            })
            .finally(() => this.$vs.loading.close());

          if (size(this.reporter) === 0) {
            const options = {
              color: 'warning',
              title: '검색 실패',
              text: '데이터가 존재하지 않습니다.',
            };
            this.$vs.notify(options);
          }
        }
      }
    },
    async getReporter(id) {
      this.reporter = await ReporterApi.get(id);
    },
    onBlur() {
      if (isEmpty(this.input)) {
        this.showPreview = false;
        this.showReporter = false;
      }
    },
    async createNews({ memo, comment }) {
      const { registered, agency, reporter, parsed, news } = this.preview;
      this.$vs.loading();
      try {
        if (registered) {
          if (memo) await MemoApi.create({ reporterId: reporter.id, memo });
          if (comment) await CommentApi.create({ newsId: news.id, comment });
        } else {
          const data = {
            url: this.input,
            title: parsed.title,
            reportedAt: parsed.reportedAt,
            lastUpdatedAt: parsed.lastUpdatedAt,
            agencyId: agency.id,
            memo,
            comment,
          };
          if (reporter) {
            data.reporterId = reporter.id;
          } else {
            data.reporterName = parsed.reporter;
          }
          await NewsApi.create(data);
        }
        this.input = '';
        this.preview = {};
      } catch (e) {
        this.$vs.notify({
          color: 'warning',
          title: '등록실패',
          text: '다시 시도해 주세요.',
        });
      } finally {
        this.$vs.loading.close();
        this.showPreview = false;
        this.input = '';
        this.openCompleteDialog();
      }
    },
    openCompleteDialog() {
      this.$vs.dialog({
        type: 'confirm',
        color: 'danger',
        title: '등록 완료',
        text: '제보해 주셔서 감사합니다.',
        accept: () => this.clearData(),
      });
    },
    closePreview() {
      this.showPreview = false;
    },
    closeReporter() {
      this.showReporter = false;
    },
    async reaction({ isLike, memoId, commentId }) {
      const mode = memoId ? 'memo' : 'comment';
      if (memoId) {
        await ReactionApi.toggle({ mode, id: memoId, isLike });
        const memo = await MemoApi.get(memoId);
        const newMemos = map(this.reporter.memos, (m) => {
          if (m.id === memo.id) return memo;
          return m;
        });
        console.log(newMemos);
        this.$set(this.reporter, 'memos', newMemos);
      }
      if (commentId) {
        await ReactionApi.toggle({ mode, id: commentId, isLike });
        const comment = await CommentApi.get(commentId);
        const newComments = map(this.preview.comments, (c) => {
          if (c.id === comment.id) return comment;
          return c;
        });
        this.$set(this.preview, 'comments', newComments);
      }
    },
  },
};
</script>

<style>
#header h1.title {
  color: #1e252d !important;
  padding-top: 20px;
  font-weight: bold;
}
#header p.sub-title {
  color: #1e252d !important;
  font-weight: 400;
}
#header p.photo-by {
  font-size: 13px !important;
  color: #1e252d !important;
}
#header input.url-input {
  font-size: 22px;
  color: #1e252d !important;
  background-color: rgba(255, 255, 255, 0.75) !important;
}
#header input.url-input::placeholder {
  color: #717981 !important;
}
</style>
