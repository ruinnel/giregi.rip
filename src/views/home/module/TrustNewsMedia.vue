<template>
  <home-section>
    <div class="spotlight">
      <div class="content">
        <header class="major">
          <h2>
            <a href="https://www.statista.com/statistics/308468/importance-brand-journalist-creating-trust-news" target="_blank">
              언론 신뢰도 꼴지(2019년)
            </a>
          </h2>
        </header>
        <p>
          노무현 대통령 논두렁 시계사건, 세월호 관련 오보, 조국 법무부 장관 관련 엄청난 양의 아니면 그만 식의 기사들.
        </p>
      </div>
      <span class="large-icon"><i class="fas fa-thumbs-down" /></span>
    </div>
  </home-section>
</template>

<script>
import HomeSection from 'views/home/module/HomeSection';

export default {
  name: 'TrustNewsMedia',
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
