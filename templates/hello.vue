<template>
    <div class="hello">
        <h1> This is your first component in Vue </h1>
        <h3> {{ webpack }} </h3>
    </div>
</template>

<script>
    export default {
        name: "hello",
        data: () => ({
            webpack: 'Powered by webpack!',
        }),
    }
</script>

<style scoped>
    .hello {
        margin: 0 auto;
        font-family: sans-serif;
        background-color: #ccf7e2;
        padding: 10px;
        border-radius: 5px;
        max-width: 500px;
    }
</style>