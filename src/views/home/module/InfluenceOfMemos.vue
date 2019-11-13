<template>
  <home-section>
    <div class="spotlight">
      <div class="content">
        <header class="major">
          <h2>
            메모의 영향력
          </h2>
        </header>
        <p>
          전체 기사를 기억할 순 없습니다.
          별명 혹은 상징적인 한두마디 말이(ex: 왜요? 기자) 훨씬 기억에 남습니다.
        </p>
        <p>
          잘 운영되는 커뮤니티들을 보면 회원들이 메모 기능을 적극 활용하고 있습니다.
          한사람의 커뮤니티에서의 행동을 자신만의 꼬리표를 붙여서 분류하는 거죠.
        </p>
        <p>
          기자한테 메모(꼬리표)를 붙이고 공유해 봅시다.
        </p>
      </div>
      <span class="large-icon"><i class="fas fa-user-tag" /></span>
    </div>
  </home-section>
</template>

<script>
import HomeSection from 'views/home/module/HomeSection';

export default {
  name: 'InfluenceOfMemos',
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
