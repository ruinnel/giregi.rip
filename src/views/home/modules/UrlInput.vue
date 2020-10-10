<template>
  <validation-observer ref="validator">
    <div class="row row-deck row-cards">
      <div class="col-sm-12 col-lg-12">
        <div class="card">
          <div class="card-body">
            <label class="form-label" for="url-input">아카이브 할 URL을 입력해 주세요.</label>
            <validator rules="required|url" name="URL">
              <div class="input-group input-group-flat">
                <input
                  id="url-input"
                  v-model="url"
                  type="text"
                  class="form-control is-valid-lite"
                  @keyup.enter="preview"
                >
                <div v-if="url" class="input-group-append">
                  <span class="input-group-text">
                    <a class="link-secondary" title="Clear search" data-toggle="tooltip" @click="clearUrl">
                      <i class="fas fa-times pl-2 pr-2" />
                    </a>
                  </span>
                </div>
              </div>
            </validator>

            <div class="align-items-center mt-2">
              <button class="btn btn-outline-primary btn-block" :disabled="!isUrl(url)" @click="preview">
                <i class="far fa-eye mr-1" />
                미리보기
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </validation-observer>
</template>

<script>
import isURL from 'validator/lib/isURL';
export default {
  name: 'UrlInput',
  props: {
    value: {
      type: String,
      required: true,
    },
  },
  data() {
    return {
      url: this.value,
    };
  },
  watch: {
    url() {
      this.$emit('input', this.url);
    },
  },
  methods: {
    isUrl(url) {
      return isURL(url);
    },
    clearUrl() {
      this.url = '';
    },
    preview() {
      const valid = this.$refs.validator.validate();
      if (valid) {
        this.$emit('preview', this.url);
      }
    },
  },
};
</script>

<style scoped>

</style>
