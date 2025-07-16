
<template>
  <!-- Navbar principal, sticky en la parte superior -->
  <header class="bg-white dark:bg-gray-900 shadow-md sticky top-0 z-50">
    <div class="max-w-screen-xl mx-auto px-4 sm:px-6 lg:px-8 pb-4 md:pb-0">
      <div class="flex h-16 items-center justify-between">
        <!-- Logo de la aplicación -->
        <div class="flex items-center">
          <router-link to="/" class="text-teal-600 dark:text-teal-300">
            <span class="sr-only">Inicio</span>
            <img class="h-6 sm:h-7" src="@/assets/icons/vue.svg" alt="Vue Logo" />
          </router-link>
        </div>

        <!-- Navegación de escritorio -->
        <nav class="hidden md:flex items-center gap-8 text-sm font-medium">
          <!-- Enlace a la página de inicio -->
          <router-link
            to="/"
            class="text-gray-500 dark:text-white hover:text-blue-600 dark:hover:text-blue-400 transition"
          >
            Inicio
          </router-link>
          <!-- Enlace a la página de creación de stock -->
          <router-link
            to="/create"
            class="text-gray-500 dark:text-white hover:text-blue-600 dark:hover:text-blue-400 transition"
          >
            Crear Stock
          </router-link>
          <!-- Enlace a la información de stock (puede ser eliminado si la página no se usa) -->
          <router-link
            to="/StockInfo"
            class="text-gray-500 dark:text-white hover:text-blue-600 dark:hover:text-blue-400 transition"
          >
            Información de stock
          </router-link>
        </nav>

        <!-- Botón para abrir/cerrar el menú móvil -->
        <div class="md:hidden">
          <button
            @click="toggleMobileMenu"
            class="rounded-md bg-gray-100 dark:bg-gray-800 p-2 text-gray-600 dark:text-white hover:text-blue-600 dark:hover:text-blue-400 transition"
            aria-label="Abrir menú móvil"
          >
            <!-- Icono de menú hamburguesa -->
            <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24"
              stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M4 6h16M4 12h16M4 18h16" />
            </svg>
          </button>
        </div>
      </div>

      <!-- Menú móvil, visible solo en pantallas pequeñas -->
      <transition name="fade">
        <nav
          v-if="isMobileMenuOpen"
          class="md:hidden mt-3 rounded-md bg-gray-50 dark:bg-gray-800 shadow-lg border dark:border-gray-700 px-6 py-4"
        >
          <ul class="flex flex-col gap-4 text-base font-medium">
            <li>
              <router-link
                to="/"
                @click="toggleMobileMenu"
                class="block text-gray-700 dark:text-white hover:text-blue-600 dark:hover:text-blue-400"
              >
                Inicio
              </router-link>
            </li>
            <li>
              <router-link
                to="/create"
                @click="toggleMobileMenu"
                class="block text-gray-700 dark:text-white hover:text-blue-600 dark:hover:text-blue-400"
              >
                Crear Stock
              </router-link>
            </li>
            <li>
              <router-link
                to="/StockInfo"
                @click="toggleMobileMenu"
                class="block text-gray-700 dark:text-white hover:text-blue-600 dark:hover:text-blue-400"
              >
                Información de stock
              </router-link>
            </li>
          </ul>
        </nav>
      </transition>
    </div>
  </header>
</template>

<script setup lang="ts">
// Navbar.vue
// Componente de barra de navegación principal, incluye menú responsivo y enlaces de navegación.
import { ref } from "vue";

// Estado reactivo para controlar la visibilidad del menú móvil
const isMobileMenuOpen = ref(false);

// Alterna la visibilidad del menú móvil
function toggleMobileMenu() {
  isMobileMenuOpen.value = !isMobileMenuOpen.value;
}
</script>

<style scoped>
/* Transición de opacidad para el menú móvil */
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}
.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
