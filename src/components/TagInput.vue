<template>
  <select ref="select" name="tags-advanced" class="form-select" multiple>
    <option v-for="(tag, idx) in tags" :key="idx" :value="tag.id">{{ tag.name }}</option>
  </select>
</template>

<script>
import { isEmpty, includes, chain, toNumber } from 'lodash';
import $ from 'jquery';
import 'selectize';
import 'selectize/dist/css/selectize.css';

export default {
  name: 'TagInput',
  props: {
    tags: {
      type: Array,
      default: () => [],
    },
  },
  data() {
    return {
      selectize: null,
      input: null,
      offset: -1,
      text: '',
    };
  },
  beforeDestroy() {
    if (this.selectize) {
      this.selectize.off('change');
      this.selectize.off('type');
      this.selectize.off('item_remove');
    }
    if (this.input) {
      this.input.unbind('keyup');
    }
  },
  mounted() {
    const select = $(this.$refs.select).selectize({
      plugins: ['remove_button'],
    });
    this.selectize = select[0].selectize;
    this.selectize.on('change', () => {
      const values = this.selectize.items;
      const items = chain(this.selectize.options)
        .filter((opt) => includes(values, opt.value))
        .map((opt) => ({ id: toNumber(opt.value), name: opt.text }))
        .value();
      this.$emit('change', items);
    });
    this.selectize.on('type', (text) => {
      this.text = text;
    });
    this.selectize.on('item_remove', (value) => {
      if (value < 0) {
        this.selectize.removeOption(value);
        this.selectize.refreshOptions();
        this.selectize.removeItem(value);
        this.selectize.refreshItems();
      }
    });

    this.input = $($('.selectize-input > input')[0]);
    this.input.keyup((key) => {
      if (key.keyCode === 13) {
        if (!isEmpty(this.text)) {
          const offset = this.offset;
          this.selectize.addOption({ value: `${offset}`, text: this.text });
          this.selectize.refreshOptions();
          this.selectize.addItem(offset);
          this.selectize.refreshItems();
          this.text = '';
          this.input.val('');
          this.offset -= 1;
        }
      }
    });
  },
};
</script>

<style scoped>

</style>
