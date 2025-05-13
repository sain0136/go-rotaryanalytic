<script setup>
import { RouterView } from "vue-router";
import Toast from "primevue/toast";
import { useAuth } from "@/composables/useAuth";

</script>

<template>
  <Toast />
  <div>
    <RouterView />
  </div>
</template>

<style scoped></style>
