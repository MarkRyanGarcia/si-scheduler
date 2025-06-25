<template>
    <section class="how-it-works">
        <h2>How it Works</h2>
        <div class="card-container">
            <div class="card" v-for="(step, index) in steps" :key="index">
                <div class="step-number">Step {{ index + 1 }}</div>
                <h3 class="step-title">{{ step.title }}</h3>
                <p class="step-description">{{ step.description }}</p>
            </div>
        </div>
    </section>
</template>

<script lang="ts" setup>
const steps = [
    {
        title: "Collect Availability",
        description:
            "SI leaders submit their weekly availability and course assignments through our simple form.",
    },
    {
        title: "Smart Algorithm",
        description:
            "Our algorithm automatically generates optimal schedules considering lecture times, availability, and spacing requirements.",
    },
    {
        title: "Review & Export",
        description:
            "View generated schedules in calendar or table format, make adjustments, and export for distribution.",
    },
];
</script>

<style scoped>
.how-it-works {
    padding: 2rem;
    text-align: center;
}

h2 {
    font-size: 2rem;
    margin-bottom: 1.5rem;
}

.card-container {
    display: flex;
    flex-wrap: wrap;
    gap: 1.5rem;
    justify-content: center;
}

.card {
    background-color: #263043;
    color: #eaf4ff;
    border-radius: 12px;
    padding: 1.5rem;
    width: 280px;
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.2);
    transition: transform 0.2s ease;
}

.card:hover {
    transform: translateY(-4px);
}

.step-number {
    font-weight: bold;
    color: #80bfff;
    margin-bottom: 0.5rem;
    font-size: 1.1rem;
}

.step-title {
    font-size: 1.3rem;
    margin-bottom: 0.5rem;
}

.step-description {
    font-size: 0.95rem;
    line-height: 1.4;
}
</style>