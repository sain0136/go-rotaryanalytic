<script setup>
import { ref, watch } from "vue";
import { useRouter } from "vue-router";
import { useAuth } from "../composables/useAuth";
import { authApi } from "../lib/api";
import { useToast } from "primevue/usetoast";

const toast = useToast();
const router = useRouter();
const { login, loggingout, setLoggingOut } = useAuth();

const username = ref("");
const password = ref("");
const error = ref("");
const isLoading = ref(false);

if (loggingout.value) {
  setLoggingOut(false);
  toast.add({
    severity: "info",
    summary: "Logged out",
    detail: "You have been logged out successfully.",
    life: 3000,
  });
}

async function handleSubmit(e) {
  e.preventDefault();
  error.value = "";
  isLoading.value = true;

  try {
    const response = await authApi.login({
      username: username.value,
      password: password.value,
    });
    if (response.status === "success") {
      login();
      router.push("/");
    }
  } catch (err) {
    error.value = err.message || "Failed to login";
  } finally {
    isLoading.value = false;
  }
}
</script>

<template>
  <section class="w-full min-h-screen flex items-center justify-center p-4">
    <div
      class="w-full max-w-md bg-white rounded-2xl shadow-lg p-8 transform transition-all duration-300 hover:shadow-xl"
    >
      <div class="flex flex-col gap-2 text-center !mb-10">
        <h1 class="text-3xl font-bold text-teal-600 mb-2">Rotary Analytics</h1>
        <div class="h-1 w-full bg-teal-500 mx-auto rounded-full mb-6"></div>
        <h2 class="text-xl text-gray-600">Sign in to your admin account</h2>
      </div>

      <form @submit="handleSubmit" class="space-y-8 flex flex-col gap-4">
        <div class="space-y-3">
          <label
            for="username"
            class="block text-sm font-semibold text-gray-700"
          >
            Username
          </label>
          <input
            v-model="username"
            type="text"
            id="username"
            class="block w-full px-4 py-3 rounded-lg bg-gray-50 border border-gray-200 focus:border-teal-500 focus:ring-2 focus:ring-teal-200 transition-all duration-200"
            placeholder="Enter your username"
            required
          />
        </div>

        <div class="space-y-3">
          <label
            for="password"
            class="block text-sm font-semibold text-gray-700"
          >
            Password
          </label>
          <input
            v-model="password"
            type="password"
            id="password"
            class="block w-full px-4 py-3 rounded-lg bg-gray-50 border border-gray-200 focus:border-teal-500 focus:ring-2 focus:ring-teal-200 transition-all duration-200"
            placeholder="••••••••"
            required
          />
        </div>
        <div v-if="error" class="p-3 text-sm text-red-500 bg-red-50 rounded-lg">
          {{ error }}
        </div>
        <button
          type="submit"
          :disabled="isLoading"
          class="w-full py-4 px-4 bg-teal-500 hover:bg-teal-600 text-white font-medium rounded-lg text-sm transition-colors duration-200 focus:ring-4 focus:ring-teal-200 focus:outline-none shadow-sm hover:shadow-md mt-4 disabled:opacity-50 disabled:cursor-not-allowed"
        >
          {{ isLoading ? "Signing in..." : "Sign in" }}
        </button>
      </form>
    </div>
  </section>
</template>

<style scoped>
input::placeholder {
  color: black;
}

input:focus {
  outline: none;
}

button:active {
  transform: scale(0.98);
}
</style>
