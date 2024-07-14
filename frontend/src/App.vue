<template>
	<div class="app">
		<!-- Sidebar -->
		<Sidebar />

		<!-- Content -->
		<router-view />

		<!-- Footer -->
		<footer class="footer">
			<p>&copy; 2024 Dhoni. All rights reserved.</p>
		</footer>
	</div>
</template>

<script setup>
import Sidebar from './components/Sidebar.vue'
</script>

<style lang="scss">
:root {
	--primary: #4ade80;
	--primary-alt: #22c55e;
	--grey: #64748b;
	--dark: #1e293b;
	--dark-alt: #334155;
	--light: #f1f5f9;
	--sidebar-width: 300px;
}

* {
	margin: 0;
	padding: 0;
	box-sizing: border-box;
	font-family: 'Fira sans', sans-serif;
}

body {
	background: var(--light);
}

button {
	cursor: pointer;
	appearance: none;
	border: none;
	outline: none;
	background: none;
}

.app {
	display: flex;

	main {
		flex: 1 1 0;
		padding: 4rem;

		@media (max-width: 1024px) {
			padding-left: 6rem;
		}
	}
}

.footer {
  text-align: center;
  padding: 1rem;
  background-color: var(--dark);
  color: var(--light);
  position: fixed;
  width: 100%;
  bottom: 0;
}
</style>