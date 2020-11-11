import { isEmpty, isFunction, isNil } from 'lodash';
import Vue from 'vue';
import Dialog from './index.vue';
import DomUtil from 'utils/dom';

const DialogConstructor = Vue.extend(Dialog);

export default (Vue, options = {}) => {
  Vue.mixin({
    data() {
      return {
        dialogInstance: null,
      };
    },
    beforeCreate() {
      this.$dialog = {};
      this.$dialog.open = (props = {}) => {
        this.dialogInstance = new DialogConstructor();
        this.dialogInstance.vm = this.dialogInstance.$mount();
        DomUtil.insertBody(this.dialogInstance.vm.$el, props.parent);
        this.dialogInstance.$data.title = props.title;
        this.dialogInstance.$data.message = props.message;
        if (!isNil(props.showConfirm)) {
          this.dialogInstance.$data.showConfirm = props.showConfirm;
        }
        if (!isNil(props.showCancel)) {
          this.dialogInstance.$data.showCancel = props.showCancel;
        }
        if (!isEmpty(props.confirmText)) {
          this.dialogInstance.$data.confirmText = props.confirmText;
        }
        if (!isEmpty(props.cancelText)) {
          this.dialogInstance.$data.cancelText = props.cancelText;
        }
        if (isFunction(props.onConfirm)) {
          this.dialogInstance.$data.onConfirm = props.onConfirm;
        }
        if (isFunction(props.onCancel)) {
          this.dialogInstance.$data.onCancel = props.onCancel;
        }
        Vue.nextTick(() => {
          this.dialogInstance.$data.active = true;
          // this.dialogInstance.$data.parameters = props.parameters;
        });
      };
      this.$dialog.closeDialog = () => {
        this.dialogInstance.$data.active = false;
        DomUtil.removeBody(this.dialogInstance.vm.$el);
      };
    },
  });
};
