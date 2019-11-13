<template>
  <div :style="style" class="main-content floating-preview">
    <div class="reporter-info">
      <vs-list>
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
    </div>
    <div class="bottom-buttons">
      <button class="button small" @click="close">닫기</button>
    </div>
  </div>
</template>
<script>
import { get } from 'lodash';

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
  },
  data: () => ({}),
  computed: {
    style() {
      const { width } = this.rect;
      return {
        width: `${width}px`,
        display: this.active ? 'block' : 'none',
      };
    },
  },
  methods: {
    getAgency(reporter) {
      return get(reporter.agencies, '[0].name');
    },
    selectReporter(reporter) {
      this.$emit('select', reporter);
    },
    close() {
      this.$emit('close');
    },
  },
};
</script>
