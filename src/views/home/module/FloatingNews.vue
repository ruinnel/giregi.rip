<template>
  <div :style="style" class="main-content floating-preview">
    <div v-if="showNewsInfo" class="news-info">
      <span class="title">{{ newsInfo.title }}</span>
      <div>
        <span class="info">{{ newsInfo.agency }} / {{ newsInfo.reporter }}</span>
        <span v-if="newsInfo.coreporter" class="info">, {{ newsInfo.coreporter }}</span>
        <div v-if="registered" text="등록된 기사" class="news-registered">
          <vs-icon icon="check_circle" color="success" />
          <span>등록됨({{ formatDate(newsInfo.createdAt) }})</span>
        </div>
      </div>
    </div>
    <vs-table stripe :data="comments" class="table">
      <template slot="header">
        <h4></h4>
      </template>
      <template slot="thead">
        <vs-th>댓글</vs-th>
        <vs-th class="reaction-cell" />
      </template>
      <template slot-scope="{data}">
        <vs-tr v-for="(row, idx) in data" :key="idx">
          <vs-td>
            <span class="memo-content">{{ row.content }}</span>
            <div class="memo-meta vs-row">
              <vs-col class="vs-sm-12"><vs-icon icon="person" />{{ row.userEmail }}</vs-col>
              <vs-col class="vs-sm-12"><vs-icon icon="access_time" />{{ formatDateTime(row.createdAt) }}</vs-col>
            </div>
          </vs-td>
          <vs-td>
            <div class="reaction" @click="like(row)">
              <vs-icon icon="thumb_up" color="primary" />
              <span class="count">{{ formatNumber(row.like) }}</span>
            </div>
            <div class="reaction" @click="unlike(row)">
              <vs-icon icon="thumb_down" color="primary" />
              <span class="count">{{ formatNumber(row.unlike) }}</span>
            </div>
          </vs-td>
        </vs-tr>
      </template>
    </vs-table>
    <vs-divider />
    <vs-collapse>
      <vs-collapse-item>
        <div slot="header" class="edit-title"><vs-icon icon="edit" />기자 메모</div>
        <vs-textarea v-model="memo" placeholder="기자 메모" />
      </vs-collapse-item>
      <vs-collapse-item>
        <div slot="header" class="edit-title"><vs-icon icon="edit" />기사 댓글</div>
        <vs-textarea v-model="comment" placeholder="기사 댓글" />
      </vs-collapse-item>
    </vs-collapse>
    <vs-divider />
    <div class="bottom-buttons">
      <button v-if="showSaveButton" class="button small primary" @click="save">등록</button>
      <button class="button small" @click="close">닫기</button>
    </div>
  </div>
</template>
<script>
import { isEmpty, first, get, size } from 'lodash';

export default {
  name: 'FloatingNews',
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
  data: () => ({
    cords: {},
    memo: '',
    comment: '',
  }),
  computed: {
    style() {
      const { width } = this.rect;
      return {
        width: `${width}px`,
        display: this.active ? 'block' : 'none',
      };
    },
    registered() {
      return this.preview.registered;
    },
    showNewsInfo() {
      return !isEmpty(this.preview.parsed) || !isEmpty(this.preview.news);
    },
    newsInfo() {
      if (this.preview.registered) {
        const { news, agency, reporter } = this.preview;
        return {
          title: news.title,
          agency: first(agency.names),
          reporter: reporter.name,
          createdAt: news.createdAt,
        };
      } else {
        return this.preview.parsed;
      }
    },
    comments() {
      return get(this.preview, 'comments');
    },
    showSaveButton() {
      if (this.preview.registered) {
        return size(this.memo) > 0 || size(this.comment) > 0;
      } else {
        return true;
      }
    },
  },
  methods: {
    save() {
      const data = { memo: this.memo, comment: this.comment };
      this.$emit('register', data);
      this.memo = '';
      this.comment = '';
    },
    close() {
      this.$emit('close');
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
