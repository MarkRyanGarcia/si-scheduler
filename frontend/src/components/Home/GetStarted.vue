<template>
    <section class="get-started">
        <h2>Ready to get started?</h2>
        <button @click="goToSchedule">View Schedules</button>
    </section>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'
const router = useRouter()

function goToSchedule() {
    router.push('/schedules')
}
</script>

<style scoped>
.get-started {
    text-align: center;
    padding: 3rem 1rem;
}

h2 {
    font-size: 2rem;
}
</style>