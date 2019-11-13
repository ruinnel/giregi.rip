<template>
  <home-section>
    <div class="spotlight">
      <div class="content">
        <header class="major">
          <h2>
            기록의 중요성
          </h2>
        </header>
        <p>
          나의 흑역사가 기록으로 남는다면?
          적어도 부끄러움을 아는 사람에게는 조금이라도 브레이크가 되지 않을까요?
        </p>
      </div>
      <span class="large-icon"><i class="fas fa-edit" /></span>
    </div>
  </home-section>
</template>

<script>
import HomeSection from 'views/home/module/HomeSection';

export default {
  name: 'ImportanceOfRecord',
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
