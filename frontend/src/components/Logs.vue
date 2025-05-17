<script setup lang="ts">
import { ref, onMounted } from "vue";
import { logApi } from "../lib/api";
import DataTable from "primevue/datatable";
import Column from "primevue/column";
import Paginator from "primevue/paginator";

const logs = ref([]);
const loading = ref(false);
const page = ref(1);
const pageSize = ref(50);
const totalRecords = ref(0);

const fetchLogs = async () => {
  loading.value = true;
  try {
    const response = await logApi.getLogs({
      page: page.value,
      pageSize: pageSize.value,
    });
    logs.value = response.logs;
    totalRecords.value = response.total;
  } catch (err) {
    // Optionally show a toast or error message
    logs.value = [];
    totalRecords.value = 0;
  } finally {
    loading.value = false;
  }
};

onMounted(fetchLogs);

const onPage = (event: any) => {
  page.value = event.page + 1; // PrimeVue paginator is 0-based
  pageSize.value = event.rows;
  fetchLogs();
};
</script>

<template>
  <DataTable
    :value="logs"
    :loading="loading"
    responsiveLayout="scroll"
    stripedRows
    showGridlines
  >
    <Column
      field="timestamp"
      header="Timestamp"
      sortable
      style="min-width: 160px"
    />
    <Column field="type" header="Type" sortable style="min-width: 100px" />
    <Column field="route" header="Route" style="min-width: 120px" />
    <Column field="requestId" header="Request ID" style="min-width: 120px" />
    <Column field="message" header="Message" style="min-width: 200px" />
    <Column
      field="customMessage"
      header="Custom Message"
      style="min-width: 200px"
    />
    <Column field="detailsString" header="Details" style="min-width: 200px" />
  </DataTable>
  <Paginator
    :rows="pageSize"
    :totalRecords="totalRecords"
    :first="(page - 1) * pageSize"
    :rowsPerPageOptions="[10, 25, 50, 100]"
    @page="onPage"
    class="mt-3"
    :pt="{
      root: {
        class: 'p-pp',
      },
    }"
  />
</template>

<style scoped lang="css">
:deep(.p-datatable) {
  font-size: 0.98rem;
}
</style>
