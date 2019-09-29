<template>
  <header id="header" class="alt">
    <input
      v-model="url"
      type="text"
      class="url-input"
      placeholder="https://"
      @keyup.enter="onEnter"
    />
    <h1 class="title">R . I . P</h1>
    <p class="sub-title">삼가 기레기의 명복을 빕니다.</p>
    <p class="photo-by">
      Photo by
      <a target="_blank" href="https://unsplash.com/@mattbotsford?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText">Matt Botsford</a>
      on <a target="_blank" href="https://unsplash.com/search/photos/tombstone?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText">Unsplash</a>
    </p>
  </header>
</template>

<script>
import { get } from 'lodash';
export default {
  name: 'HomeHeader',
  data() {
    return {
      url: 'http://news.chosun.com/site/data/html_dir/2019/08/27/2019082700128.html',
    };
  },
  computed: {},
  methods: {
    onEnter() {
      this.$vs.loading();
      this.$refs.iframe.onload = this.loadComplete;
    },
    loadComplete() {
      console.log('load complete');
      this.$refs.iframe.onload = () => {};
      this.$vs.loading.close();

      const contentWindow = this.$refs.iframe.contentWindow;
      const contentDocument = this.$refs.iframe.contentDocument;
      let html = get(contentWindow, 'body.innerHTML');
      if (!html) {
        html = get(contentDocument, 'body.innerHTML');
      }
      console.log(html);
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
