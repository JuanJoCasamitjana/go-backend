<script>
import TodoList from './components/TodoList.vue';

export default {
  components: {
    TodoList
  },
  data() {
  return {
    LIST_URL:"/todos",
  }
  }
}

</script>

<template>
  <TodoList :API_URL="LIST_URL">  </TodoList>
</template>

<style scoped>
header {
  line-height: 1.5;
}

.logo {
  display: block;
  margin: 0 auto 2rem;
}

@media (min-width: 1024px) {
  header {
    display: flex;
    place-items: center;
    padding-right: calc(var(--section-gap) / 2);
  }

  .logo {
    margin: 0 2rem 0 0;
  }

  header .wrapper {
    display: flex;
    place-items: flex-start;
    flex-wrap: wrap;
  }
}
</style>
