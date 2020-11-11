<template>
  <div ref="dialog" class="modal modal-blur" role="dialog" :style="styles">
    <div class="modal-dialog modal-dialog-centered" role="document">
      <div class="modal-content">
        <div class="modal-header">
          <h5 class="modal-title">{{ title }}</h5>
          <button
            type="button"
            class="close"
            data-dismiss="modal"
            aria-label="Close"
            @click="close(false)"
          >
            <i class="fas fa-times" />
          </button>
        </div>
        <div class="modal-body">
          {{ message }}
        </div>
        <div class="modal-footer">
          <button
            v-if="showCancel"
            type="button"
            class="btn btn-danger"
            data-dismiss="modal"
            @click="close(false)">
            {{ cancelText }}
          </button>
          <button
            v-if="showConfirm"
            type="button"
            class="btn btn-primary"
            data-dismiss="modal"
            @click="close(true)">
            {{ confirmText }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import DomUtil from 'utils/dom';

export default {
  name: 'Dialog',
  props: {},
  data() {
    return {
      active: false,
      title: '',
      message: '',
      showCancel: false,
      showConfirm: true,
      confirmText: '확인',
      cancelText: '취소',
      onConfirm: () => {},
      onCancel: () => {},
    };
  },
  computed: {
    styles() {
      return {
        display: this.active ? 'block' : 'none',
      };
    },
  },
  watch: {},
  methods: {
    close(isConfirm) {
      if (isConfirm) {
        this.onConfirm();
      } else {
        this.onCancel();
      }
      DomUtil.removeBody(this.vm.$el);
    },
  },
};
</script>

<style scoped>

</style>
