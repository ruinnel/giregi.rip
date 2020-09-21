<template>
  <div class="col-sm-12 col-lg-12">
    <div class="card">
      <div class="card-body">
        <h3 class="card-title">Preview</h3>
        <div>
          입력한 URL을 분석한 결과가 표시 됩니다.
        </div>
        <div v-if="data" class="align-items-center mt-2">
          <button class="btn btn-primary btn-block">
            <i class="fas fa-archive" />
            아카이브
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Preview',
  props: {
    data: {
      type: Object,
      default: () => null,
    },
  },
};
</script>

<style scoped>

</style>
