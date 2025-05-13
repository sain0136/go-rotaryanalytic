<script setup>
import { useAuth } from "../composables/useAuth";
import { onMounted } from "vue";
const { logout } = useAuth();
import { authApi } from "@/lib/api";

onMounted(async () => {
  try {
    const res = await authApi.test();
    console.log(res);
  } catch (error) {
    console.error("Error during component mount:", error);
  }
});
</script>

<template>
  <div class="min-h-screen bg-gray-100">
    <nav class="bg-white shadow-md p-4">
      <div class="max-w-7xl mx-auto flex justify-between items-center">
        <h1 class="text-2xl font-bold text-teal-600">Rotary Analytics</h1>
        <button
          @click="logout"
          class="px-4 py-2 bg-teal-500 text-white rounded-lg hover:bg-teal-600 transition-colors"
        >
          Logout
        </button>
      </div>
    </nav>

    <main class="max-w-7xl mx-auto p-4">
      <!-- Your main content here -->
      <h2 class="text-xl font-semibold mb-4">Welcome to Rotary Analytics</h2>
    </main>
  </div>
</template>

<style scoped>
h1 {
  color: #42b983;
  text-align: center;
}
</style>
