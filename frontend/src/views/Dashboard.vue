<template>
	<main id="Home-page">
		<h1>Dashboard</h1>
		<p>This is the dashboard page</p>
		<br/>
		<div class="data-table">
			<span>search by year: </span>
			<input type="text" v-model="searchValue">
			<EasyDataTable
				:headers="headers"
				:items="items"
				:search-field="searchField"
				:search-value="searchValue"
				:rows-per-page="10"
			/>
		</div>
	</main>
</template>
<script setup>
import { ref, onMounted } from 'vue';
import axios from 'axios';
import EasyDataTable from 'vue3-easy-data-table';
import 'vue3-easy-data-table/dist/style.css';

const searchField = ref("year");
const searchValue = ref("");

const headers = ref([
  { text: 'Vault ID', value: 'vault_id' },
  { text: 'Year', value: 'year' },
  { text: 'University', value: 'university' },
  { text: 'School', value: 'school' },
  { text: 'Degree', value: 'degree' },
  { text: 'Employment Rate Overall', value: 'employment_rate_overall' },
  { text: 'Employment Rate Full-time Permanent', value: 'employment_rate_ft_perm' },
  { text: 'Basic Monthly Mean', value: 'basic_monthly_mean' },
  { text: 'Basic Monthly Median', value: 'basic_monthly_median' },
  { text: 'Gross Monthly Mean', value: 'gross_monthly_mean' },
  { text: 'Gross Monthly Median', value: 'gross_monthly_median' },
  { text: 'Gross Monthly 25th Percentile', value: 'gross_mthly_25_percentile' },
  { text: 'Gross Monthly 75th Percentile', value: 'gross_mthly_75_percentile' }
]);

const items = ref([]);

const fetchData = async () => {
  try {
    const response = await axios.get('https://api-production.data.gov.sg/v2/internal/api/datasets/d_3c55210de27fcccda2ed0c63fdd2b352/rows?limit=1000');
    items.value = response.data.data.rows;
  } catch (error) {
    console.error('Error fetching data:', error);
  }
};

onMounted(() => {
  fetchData();
});
</script>