<template>
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
            <div class="input-group input-group-flat">
              <input id="url-input" v-model="url" type="text" class="form-control">
              <span class="input-group-text">
                <a class="link-secondary" title="Clear search" data-toggle="tooltip" @click="clearUrl">
                  <i class="fas fa-times" />
                </a>
              </span>
            </div>
            <div class="align-items-center mt-2">
              <button class="btn btn-primary btn-block" @click="loadPreview">
                <i class="far fa-eye" />
                미리보기
              </button>
            </div>
          </div>
        </div>
      </div>
      <div class="col-sm-12 col-lg-12">
        <div class="card">
          <div class="card-body">
            <h3 class="card-title">Preview</h3>
            <div>
              입력한 URL을 분석한 결과가 표시 됩니다.
            </div>
            <div v-if="previewLoaded" class="align-items-center mt-2">
              <button class="btn btn-primary btn-block">
                <i class="fas fa-archive" />
                아카이브
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Home',
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
    clearUrl() {
      // TODO:
      console.log('clear url.');
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
