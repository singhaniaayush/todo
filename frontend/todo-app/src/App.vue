<template>
  <div>Todo's</div>
  <todo-first/>
</template>

<script>
import TodoFirst from "./components/TodoFirst.vue"

export default {
  name: "App",
  components: {
    TodoFirst
  },
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}

.done {
  text-decoration: line-through;
}

.todo {
  cursor: pointer;
}
</style>
