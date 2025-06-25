<template>
    <section class="welcome">
        <h1>Generate the Perfect SI Schedule – Instantly</h1>
        <p>Based on SI leader availability and lecture times, this tool helps you schedule sessions in minutes.</p>
        <button @click="navigateToAbout">About</button>
    </section>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'

const router = useRouter()
function navigateToAbout() {
    router.push('/about')
}
</script>

<style scoped>
.welcome {
    text-align: center;
    padding: 4rem 2rem;
}

button {
    margin-top: 1rem;
    padding: 0.75rem 1.5rem;
    font-size: 1rem;
    cursor: pointer;
}
</style>