<template>
    <div>
      <button @click="fetchTodos">Actualizar Lista</button>
      <table>
        <thead>
          <tr>
            <th>Titulo</th>
            <th>Descripción</th>
            <th>Completado</th>
          </tr>
        </thead>
        <tbody>
          <todo-item v-for="todo in todos" :key="todo.id" :todo="todo" />
        </tbody>
      </table>
    </div>
  </template>
  
  <script>
  import TodoItem from './TodoItem.vue';
  
  export default {
    components: {
      TodoItem,
    },
    data() {
      return {
        todos: [],
      };
    },
    methods: {
      async fetchTodos() {
        try {
          const response = await fetch('/todos');
          if (response.ok) {
            this.todos = await response.json();
          } else {
            console.error('Error al cargar la lista de tareas:', response.statusText);
          }
        } catch (error) {
          console.error('Error al cargar la lista de tareas:', error);
        }
      },
    },
    setup() {
      const { fetchTodos, todos } = this;
  
      watchEffect(() => {
        fetchTodos(); // Cargar la lista de tareas al iniciar la aplicación
      });
  
      return {
        fetchTodos,
        todos,
      };
    },
  };
  </script>
  