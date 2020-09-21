<template>
  <validation-observer ref="validator">
    <div class="container-xl">
      <!-- Page title -->
      <div class="page-header text-white">
        <div class="row align-items-center">
          <div class="col-auto">
            <h2 class="page-title">
              Giregi R.I.P
              <i class="fas fa-grip-lines-vertical" />
              <span class="small">기억하기 보다 기록하자.</span>
            </h2>
          </div>
        </div>
      </div>
      <div class="row row-deck row-cards">
        <div class="col-sm-12 col-lg-12">
          <div class="card">
            <div class="card-body">
              <label class="form-label" for="url-input">아카이브 할 URL을 입력해 주세요.</label>
              <validator rules="required|url" name="URL">
                <div class="input-group input-group-flat">
                  <input id="url-input" v-model="url" type="text" class="form-control is-valid-lite">
                  <div v-if="url" class="input-group-append">
                    <span class="input-group-text">
                      <a class="link-secondary" title="Clear search" data-toggle="tooltip" @click="clearUrl">
                        <i class="fas fa-times" />&nbsp;
                      </a>
                    </span>
                  </div>
                </div>
              </validator>

              <div class="align-items-center mt-2">
                <button class="btn btn-primary btn-block" :disabled="!isUrl(url)" @click="loadPreview">
                  <i class="far fa-eye" />
                  미리보기
                </button>
              </div>
            </div>
          </div>
        </div>
        <preview />
      </div>
    </div>
  </validation-observer>
</template>

<script>
import isURL from 'validator/lib/isURL';
import Preview from './modules/Preview';

export default {
  name: 'Home',
  components: {
    Preview,
  },
  data() {
    return {
      url: '',
    };
  },
  computed: {
    previewLoaded() {
      // TODO:
      return false;
    },
  },
  methods: {
    loadPreview() {
      // TODO:
    },
    isUrl(url) {
      return isURL(url);
    },
    clearUrl() {
      this.url = '';
    },
  },
};
</script>

<style lang="scss" scoped>
i {
  &.fas {
    margin-right: 5px;
  }
  &.far {
    margin-right: 5px;
  }
}
</style>
