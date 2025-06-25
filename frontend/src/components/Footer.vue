<template>
    <footer class="app-footer">
        <p>© {{ new Date().getFullYear() }} | Made with 💗 by <a href="https://www.markg.dev/" target="_blank">Mark
                Garcia</a> and <a href="https://chom.vercel.app/" target="_blank">Tomas Oh</a></p>
        <!-- <nav class="footer-links">
            <a href="#">Privacy</a>
            <a href="#">Terms</a>
            <a href="#">Contact</a>
        </nav> -->
    </footer>
</template>

<style scoped>
.app-footer {
    background-color: #1e2430;
    color: #eaf4ff;
    /* padding: 1rem 2rem; */
    text-align: center;
    border-top: 1px solid #3a455b;
    font-size: 0.9rem;
}

.footer-links {
    margin-top: 0.5rem;
    display: flex;
    justify-content: center;
    gap: 1rem;
    flex-wrap: wrap;
}

.footer-links a {
    color: #80bfff;
    text-decoration: none;
    transition: color 0.2s ease;
}

.footer-links a:hover {
    color: #a3cbff;
}
</style>