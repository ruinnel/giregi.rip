<template>
  <home-section>
    <div class="spotlight">
      <div class="content">
        <header class="major">
          <h2>Ipsum sed adipiscing</h2>
        </header>
        <p>
          Sed lorem ipsum dolor sit amet nullam consequat feugiat consequat magna
          adipiscing magna etiam amet veroeros. Lorem ipsum dolor tempus sit cursus.
          Tempus nisl et nullam lorem ipsum dolor sit amet aliquam.
        </p>
        <ul class="actions">
          <li><button class="button" @click="onMore">Learn More</button></li>
        </ul>
      </div>
      <span class="image"><img src="images/pic01.jpg" alt="" /></span>
    </div>
  </home-section>
</template>

<script>
import HomeSection from 'views/home/module/HomeSection';

export default {
  name: 'Recent',
  components: {
    HomeSection,
  },
  methods: {
    onMore() {
      this.$emit('onMore', 'recent');
    },
  },
};
</script>

<style scoped>

</style>
