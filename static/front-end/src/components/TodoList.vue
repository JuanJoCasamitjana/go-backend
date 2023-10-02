<script>
import { watchEffect } from 'vue';
import TodoItem from './TodoItem.vue';
export default {
    name: 'TodoList',
    props: {
        API_URL: {
            type: String,
            required: true
        }
    },
    components: {
        TodoItem
    },
    data() {
        return {
            todos: []
        }
    },
    setup() {
        watchEffect(async () => {
        this.todos = await fetch(API_URL).then(res => res.json())
    })
    }
}


</script>

<template>
    <ul>
        <TodoItem v-for="todo in todos" :key="todo.id" :todo="todo"/>
    </ul>
</template>

<style scoped>

</style>