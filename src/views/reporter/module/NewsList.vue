<template>
  <section>
    <div class="news-header">
      <h2 class="title">뉴스목록</h2>
      <pagination
        class="pagination"
        :paging="paging"
        @prev="prev"
        @next="next"
      />
    </div>
    <div class="table-wrapper">
      <table>
        <thead>
          <tr>
            <th>언론사</th>
            <th>제목</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(news, idx) in data" :key="`news-${idx}`">
            <td>{{ news.agency.name }}</td>
            <td>
              <div><a target="_blank" :href="newsUrl(news)">{{ news.title }}</a></div>
              <div class="meta">
                <i class="far fa-clock" />{{ formatDateTime(news.createdAt) }}
                <div class="reaction" @click="reaction(news, true)">
                  <vs-icon icon-pack="far" icon="fa-thumbs-up" color="primary" />
                  <span class="count">{{ formatNumber(news.like) }}</span>
                </div>
                <div class="reaction" @click="reaction(news, false)">
                  <vs-icon icon-pack="far" icon="fa-thumbs-up" color="danger" class="fa-rotate-180" />
                  <span class="count">{{ formatNumber(news.unlike) }}</span>
                </div>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
      <pagination
        class="pagination"
        :paging="paging"
        @prev="prev"
        @next="next"
      />
    </div>
  </section>
</template>

<script>
import { slice, get, isNil } from 'lodash';

export default {
  name: 'NewsList',
  props: {
    paging: {
      type: Object,
      default: () => ({
        data: [],
        offset: 0,
        count: 5,
        lastOffset: null,
      }),
    },
  },
  data() {
    return {
      page: 1,
    };
  },
  computed: {
    data() {
      const { data, count } = this.paging;
      return slice(data, 0, count);
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
    newsUrl(news) {
      const waybackId = get(news, 'archive.waybackId');
      if (waybackId) {
        return `http://web.archive.org${waybackId}`;
      } else {
        return news.url;
      }
    },
    reaction(news, isLike) {
      this.$emit('reaction', { isLike, newsId: news.id });
    },
  },
};
</script>

<style scoped lang="scss">
.news-header {
  .title {
    display: inline;
  }
  .pagination {
    display: inline;
  }
}
.meta {
  margin-top: 10px;
  display: flex;
  align-items: center;
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
</style>
