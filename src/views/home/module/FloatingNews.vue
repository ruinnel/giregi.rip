<template>
  <div :style="style" class="main-content floating-preview">
    <div v-if="showNewsInfo" class="news-info">
      <span class="title">
        <a :href="newsUrl" target="_blank">{{ newsInfo.title }}</a>
      </span>
      <div v-if="reported" text="등록된 기사" class="news-reported">
        <div>
          <div class="news-reported-at">입력: {{ formatDate(newsInfo.reportedAt) }}</div>
          <vs-icon icon="check_circle" color="success" />
          <span>등록됨({{ formatDate(newsInfo.createdAt) }})</span>
        </div>
      </div>
    </div>
    <reporter-card :reporter="reporter" @reaction="onReaction" @memo="onMemo" />
    <vs-divider />
    <div>
      <div class="edit-title"><vs-icon icon="edit" />기사 메모</div>
      <vs-input v-model="comment" placeholder="기사 메모" class="comment-input" maxlength="200" />
    </div>
    <vs-divider />
    <div class="bottom-buttons">
      <button class="button small primary" @click="save">{{ saveButtonText }}</button>
      <button class="button small" @click="close">닫기</button>
    </div>
  </div>
</template>
<script>
import { isEmpty, get, size, assign, isEqual } from 'lodash';
import ReporterCard from 'components/ReporterCard';
import ApiClient, { API } from 'api/client';

export default {
  name: 'FloatingNews',
  components: {
    ReporterCard,
  },
  mixins: [ApiClient],
  props: {
    rect: {
      type: Object,
      default: () => {},
    },
    active: {
      type: Boolean,
      default: false,
    },
    preview: {
      type: Object,
      required: true,
    },
  },
  data() {
    return {
      cords: {},
      newComment: '',
      prevPreview: this.preview,
    };
  },
  computed: {
    comment: {
      get() {
        if (size(this.newComment) > 0) {
          return this.newComment;
        } else {
          return get(this.preview, 'myComment.content');
        }
      },
      set(comment) {
        this.newComment = comment;
      },
    },
    style() {
      const { width } = this.rect;
      return {
        width: `${width}px`,
        display: this.active ? 'block' : 'none',
      };
    },
    reported() {
      return this.preview.reported;
    },
    showNewsInfo() {
      return !isEmpty(this.preview.parsed) || !isEmpty(this.preview.news);
    },
    saveButtonText() {
      const user = this.$store.state.user || {};
      const informer = get(this.preview, 'news.informer', {});
      if (this.preview.reported && user.id === informer.id) {
        return '저장';
      } else {
        return '제보';
      }
    },
    reporter() {
      const reporter = get(this.preview, 'reporter', {});
      if (!isEmpty(reporter)) {
        return reporter;
      } else {
        return {
          name: get(this.preview, 'parsed.reporter'),
          agency: { name: get(this.preview, 'parsed.agency') },
        };
      }
    },
    newsInfo() {
      if (this.preview.reported) {
        const { news } = this.preview;
        return news;
      } else {
        return this.preview.parsed;
      }
    },
    comments() {
      return get(this.preview, 'comments');
    },
    newsUrl() {
      const news = get(this.preview, 'news') || this.preview.parsed;
      const waybackId = get(news, 'archive.waybackId');
      if (waybackId) {
        return `http://web.archive.org${waybackId}`;
      } else {
        return news.url;
      }
    },
  },
  watch: {
    async preview(preview) {
      const MemoApi = this.getApi(API.MEMO);
      const CommentApi = this.getApi(API.COMMENT);
      if (!isEmpty(preview) && !isEqual(this.preview, this.prevPreview)) {
        const reporterId = get(preview, 'reporter.id');
        const newsId = get(preview, 'news.id');
        if (reporterId) {
          const memo = await MemoApi.my([reporterId]);
          if (size(memo) > 0) {
            this.$set(this.preview.reporter, 'myMemo', memo[0]);
          }
        }
        if (newsId) {
          const comment = await CommentApi.my(newsId);
          this.$set(this.preview, 'myComment', comment);
        }
      }

      this.prevPreview = preview;
    },
  },
  methods: {
    async save() {
      const { reported } = this.preview;
      if (size(this.newComment) > 0) {
        const CommentApi = this.getApi(API.COMMENT);
        const commentId = get(this.preview, 'myMemo.id');
        this.$vs.loading();
        if (size(commentId) > 0) {
          await CommentApi.update({ id: commentId, comment: this.newComment });
        } else {
          const newsId = this.preview.news.id;
          await CommentApi.create({ newsId, comment: this.newComment });
        }
        this.$vs.loading.close();
      }
      if (reported) {
        this.$emit('close', { clear: true });
      } else {
        this.$emit('report');
      }
    },
    async onReaction({ isLike, reporterId }) {
      const ReactionApi = this.getApi(API.REACTION);
      const ReporterApi = this.getApi(API.REPORTER);
      if (reporterId) {
        await ReactionApi.toggle({ mode: 'reporter', id: reporterId, isLike });
        const reporter = await ReporterApi.get(reporterId);
        this.$set(this.preview, 'reporter', reporter);
      }
    },
    async onMemo({ id, memo }) {
      const MemoApi = this.getApi(API.MEMO);
      this.$vs.loading();
      if (size(id) > 0) {
        await MemoApi.update({ id, memo });
      } else {
        await MemoApi.create({ reporterId: this.preview.reporter.id, memo });
      }
      this.$set(this.preview.reporter, 'myMemo', assign(this.preview.reporter.myMemo, { content: memo }));
      this.$vs.loading.close();
      this.$vs.notify({ color: 'success', title: '저장 완료', text: '메모가 저장 되었습니다.' });
    },
    close() {
      this.$emit('close', { clear: false });
    },
    like(comment) {
      this.$emit('reaction', { isLike: true, commentId: comment.id });
    },
    unlike(comment) {
      this.$emit('reaction', { isLike: false, commentId: comment.id });
    },
  },
};
</script>

<style scoped lang="scss">
.comment-input {
  width: 100% !important;
}
.vs-input--placeholder {
  top: 0px !important;
}
.news-info {
  align-items: center;
  margin-bottom: 10px;
  & .title {
    font-weight: bold;
  }
  & .info {
    font-size: 0.9em;
  }
  & div {
    display: flex;
    align-items: center;
  }
  & .vs-icon {
    margin-left: 5px;
    font-size: 1.1em;
  }
  & .news-reported-at {

  }
  & .news-reported {
    display: block;
    div {
      display: flex;
      font-size: 0.9em;
      align-items: center;
      & > span {
        font-size: 0.6em;
        font-weight: bold;
      }
    }
  }
}
</style>
