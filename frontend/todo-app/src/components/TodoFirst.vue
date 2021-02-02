<template>
  <form @submit.prevent="addNewTodo">
    <label> Add Todo </label>
    <input v-model="newTodo" name="createTodo" />
    <button>Add</button>
  </form>

  <ul>
    <li v-for="(todo, index) in todos" :key="todo.id" class="todo">
      <h3 :class="{ done: todo.done }" @click="toggleDone(todo)">
        {{ todo.content }}
      </h3>
      <button @click="removeTodo(index)">Remove</button>
    </li>
  </ul>
</template>

<script>
import {ref} from 'vue'
export default {
  name: "TodoFirst",

  setup() {
    const newTodo = ref("");
    const todos = ref([]);

    function addNewTodo() {
      todos.value.push({
        id: Date.now(),
        done: false,
        content: newTodo.value,
      });
      newTodo.value = "";
    }

    function toggleDone(todo) {
      todo.done = !todo.done;
    }

    function removeTodo(index) {
      todos.value.splice(index, 1);
    }

    return {
      todos,
      addNewTodo,
      newTodo,
      toggleDone,
      removeTodo,
    };
  },
};
</script>

<style scoped>
.done {
  text-decoration: line-through;
}

.todo {
  cursor: pointer;
}
</style>