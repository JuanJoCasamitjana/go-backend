<template>
    <tr>
      <td>{{ todo.title }}</td>
      <td>{{ todo.description }}</td>
      <td>
        <input type="checkbox" v-model="todo.completed" />
      </td>
    </tr>
  </template>
  
  <script>
  export default {
    props: {
      todo: Object,
    },
  };
  </script>
  