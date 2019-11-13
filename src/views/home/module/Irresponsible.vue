<template>
  <home-section>
    <div class="spotlight">
      <div class="content">
        <header class="major">
          <h2>
            누가 적었을까요?
          </h2>
        </header>
        <p>
          논두렁 시계 기사를 쓴 기자가 누구였는지 기억하시나요?
          조국 사태에 수없이 쏟아진 악의적인 보도를 한 기자가 누구인지 기억 하시나요?
          사람들은 보도한 언론사는 기억하지만 작성한 기자는 기억하지 못하는 경우가 많습니다.
        </p>
      </div>
      <span class="large-icon"><i class="fas fa-user-secret" /></span>
    </div>
  </home-section>
</template>

<script>
import HomeSection from 'views/home/module/HomeSection';

export default {
  name: 'Irresponsible',
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
