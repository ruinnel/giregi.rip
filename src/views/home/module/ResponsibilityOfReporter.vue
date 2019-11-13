<template>
  <home-section>
    <div class="spotlight">
      <div class="content">
        <header class="major">
          <h2>
            기자의 책임
          </h2>
        </header>
        <p>
          펜의 힘은 강합니다. 기사 하나로 누군가는 인생이 망가집니다.
          하지만 기사를 잘못 써서 기자가 사과했다는 말은 들어 본적이 없습니다.
          그리고 사람들 기억에서 사라집니다.
        </p>
        <p>
          말도 안되는 기사를 올려놓고 회사 뒤로 숨고 잊혀지고 또 다시 같은 일을 반복 합니다.
          개개인의 기자는 본인의 행동에 책임을 져야하는 성인이고 한사람의 사회 구성원입니다.
        </p>
      </div>
      <span class="large-icon"><i class="fas fa-marker" /></span>
    </div>
  </home-section>
</template>

<script>
import HomeSection from 'views/home/module/HomeSection';

export default {
  name: 'ResponsibilityOfReporter',
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
