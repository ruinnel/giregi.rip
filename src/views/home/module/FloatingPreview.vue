<template>
  <div :style="style">
    floating preview
  </div>
</template>
<script>
import { assign } from 'lodash';
const defaultStyle = {
  position: 'absolute',
  transitionDelay: '1s',
  backgroundColor: 'rgba(255, 255, 255, 0.75)',
  borderRadius: '8px',
  display: 'none',
};

export default {
  name: 'FloatingPreview',
  props: {
    rect: {
      type: Object,
      default: () => defaultStyle,
    },
    active: {
      type: Boolean,
      default: false,
    },
  },
  data: () => ({
    cords: {},
  }),
  computed: {
    style() {
      const { top, left, width } = this.rect;
      return assign({}, defaultStyle, {
        left: `${left}px`,
        top: `${top}px`,
        width: `${width}px`,
        display: this.active ? 'block' : 'none',
      });
    },
  },
  mounted() {
    // utils.insertBody(this.$refs.vstooltip);
  },
  methods: {
  },
};
</script>
