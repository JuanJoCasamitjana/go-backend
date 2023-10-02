<script>
export default {
    name: 'TodoItem',
    props: {
        todo: {
            type: Object,
            required: true
        }
    }
}
</script>

<template>
    <h3>{{ todo.title }}</h3>
    <p>{{ todo.description }}</p>
</template>

<style>

</style>