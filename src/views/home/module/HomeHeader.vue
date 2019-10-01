<template>
  <validation-observer ref="validator">
    <header id="header" class="alt">
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
      <h1 class="title">R . I . P</h1>
      <p class="sub-title">삼가 기레기의 명복을 빕니다.</p>
      <p class="photo-by">
        Photo by
        <a target="_blank" href="https://unsplash.com/@mattbotsford?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText">Matt Botsford</a>
        on <a target="_blank" href="https://unsplash.com/search/photos/tombstone?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText">Unsplash</a>
      </p>
      <floating-preview
        :rect="rect"
        :active="showPreview"
      />
    </header>
  </validation-observer>
</template>

<script>
import { debounce, startsWith } from 'lodash';
import FloatingPreview from './FloatingPreview';

export default {
  name: 'HomeHeader',
  components: {
    FloatingPreview,
  },
  data() {
    return {
      input: '', // 'http://news.chosun.com/site/data/html_dir/2019/08/27/2019082700128.html',
      rect: {},
      onResize: null,
      showPreview: false,
    };
  },
  computed: {},
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
      console.log(this.rect);
    },
    async onEnter() {
      if (startsWith(this.input, 'https')) {
        // TODO
        console.log(this.input);
      }
      const valid = await this.$refs.validator.validate();
      if (valid) {
        // this.$vs.loading();
        this.showPreview = true;
      }
    },
    onBlur() {
      this.showPreview = false;
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
