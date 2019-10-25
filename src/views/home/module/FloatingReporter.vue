<template>
  <div :style="style" class="main-content floating-preview">
    <div class="reporter-info">
      <vs-list v-if="showList">
        <vs-list-item
          v-for="item in reporters"
          :key="item.id"
          :title="item.name"
          :subtitle="getAgency(item)"
          class="reporter-item"
          @click.native="selectReporter(item)"
        >
          <vs-avatar slot="avatar" />
        </vs-list-item>
      </vs-list>
      <div v-if="!showList">
        <vs-list-item
          :title="reporter.name"
          :subtitle="getAgency(reporter)"
        >
          <vs-avatar slot="avatar" />
        </vs-list-item>
        <vs-table stripe :data="memos" class="table">
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
      </div>
    </div>
    <div v-if="!showList">
      <vs-divider />
      <vs-collapse>
        <vs-collapse-item>
          <div slot="header" class="edit-title"><vs-icon icon="edit" />기자 메모</div>
          <vs-textarea v-model="memo" placeholder="기자 메모" />
        </vs-collapse-item>
      </vs-collapse>
    </div>
    <vs-divider />
    <div class="bottom-buttons">
      <button v-if="showSaveButton" class="button small primary" @click="save">등록</button>
      <button class="button small" @click="close">닫기</button>
    </div>
  </div>
</template>
<script>
import { size, isEmpty, get } from 'lodash';

export default {
  name: 'FloatingReport',
  props: {
    rect: {
      type: Object,
      default: () => {},
    },
    active: {
      type: Boolean,
      default: false,
    },
    reporters: {
      type: Array,
      required: true,
    },
    reporter: {
      type: Object,
      required: true,
    },
  },
  data: () => ({
    cords: {},
    memo: '',
  }),
  computed: {
    style() {
      const { width } = this.rect;
      return {
        width: `${width}px`,
        display: this.active ? 'block' : 'none',
      };
    },
    showList() {
      return isEmpty(this.reporter);
    },
    memos() {
      return get(this.reporter, 'memos');
    },
    showSaveButton() {
      return size(this.memo) > 0;
    },
  },
  methods: {
    getAgency(reporter) {
      return get(reporter.agencies, '[0].name');
    },
    selectReporter(reporter) {
      console.log('select - ', reporter);
      this.$emit('select', reporter);
    },
    save() {
      const data = { reporter: this.reporter, memo: this.memo };
      this.$emit('register-memo', data);
      this.memo = '';
    },
    close() {
      this.$emit('close');
    },
    like(memo) {
      this.$emit('reaction', { isLike: true, memoId: memo.id });
    },
    unlike(memo) {
      this.$emit('reaction', { isLike: false, memoId: memo.id });
    },
  },
};
</script>
