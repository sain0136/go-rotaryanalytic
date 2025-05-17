<script setup>
import { RouterView, useRoute, useRouter } from "vue-router";
import Toast from "primevue/toast";
import { useAuth } from "@/composables/useAuth";
import Menubar from "primevue/menubar";

const { isLoggedIn, logout } = useAuth();
const router = useRouter();

const menuItems = [
  {
    label: "Log Viewer",
    icon: "pi pi-list",
    command: () => router.push("/logs"),
  },
];
</script>

<template>
  <Toast />
  <div class="main-container">
    <template v-if="isLoggedIn">
      <Menubar :model="menuItems" class="app-menubar">
        <template #end>
          <button class="logout-btn" @click="logout">Logout</button>
        </template>
      </Menubar>
      <div class="content-container">
        <RouterView />
      </div>
    </template>
    <template v-else>
      <RouterView />
    </template>
  </div>
</template>

<style scoped>
.main-container {
  display: flex;
  flex-direction: column;
  height: 100vh;
  background-color: #f0f4f8;
}
.app-menubar {
  background: #14b8a6;
  border: none;
  box-shadow: 0 2px 8px rgba(20, 184, 166, 0.08);
  margin-bottom: 2rem;
  margin-top: 0;
  position: fixed;
  top: 0;
  left: 0;
  width: 100vw;
  z-index: 1000;
  border-radius: 0%;
}
.content-container {
  margin-top: 4rem;
  padding: 2rem;
}
.logout-btn {
  background: #14b8a6;
  color: #fff;
  border: none;
  border-radius: 0.5rem;
  padding: 0.5rem 1.5rem;
  font-size: 1rem;
  font-weight: 500;
  cursor: pointer;
  transition: background 0.2s;
  margin-left: 1rem;
}
.logout-btn:hover {
  background: #0d9488;
}
</style>
