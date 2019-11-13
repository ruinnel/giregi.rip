<template>
  <home-section>
    <div class="spotlight">
      <div class="content">
        <header class="major">
          <h2>
            보존의 중요성
          </h2>
        </header>
        <p>
          포털이나 신문사의 페이지는 삭제되거난 내용을 변경 할 수 있습니다.
          조국 사태 이후 모 포탈에서의 기사 검색 결과 건수가 줄어 들고 있다는 의혹이 있었습니다.
        </p>
        <p>
          포털이나 신문사들이 기사를 내리더라도 그 당시 기사를 볼수 있게 보관(archive) 합니다.
        </p>
      </div>
      <span class="large-icon"><i class="fas fa-archive" /></span>
    </div>
  </home-section>
</template>

<script>
import HomeSection from 'views/home/module/HomeSection';

export default {
  name: 'ImportanceOfArchive',
  components: {
    HomeSection,
  },
  methods: {
    onMore() {
      this.$emit('onMore', 'about');
    },
  },
};
</script>

<style scoped>

</style>
