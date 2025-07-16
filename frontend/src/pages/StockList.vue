
<template>
  <!--
    StockList.vue
    Página principal que muestra el listado de stocks.
    Incluye Navbar, spinner de carga, manejo de errores y la tabla de stocks.
  -->
  <Navbar />
  <div class="flex justify-center items-center min-h-[80vh] bg-gradient-to-br from-indigo-50 to-white dark:from-gray-900 dark:to-gray-800 py-8">
    <div class="w-full max-w-9/10 bg-white dark:bg-gray-900 rounded-2xl shadow-2xl p-10 border border-gray-200 dark:border-gray-800">
      <h2 class="text-3xl font-extrabold text-center text-indigo-700 dark:text-indigo-300 tracking-tight">
        📈 Listado De Stock
      </h2>
      <!-- Spinner de carga -->
      <div v-if="stockStore.loading" class="flex flex-col items-center justify-center py-8">
        <div class="w-10 h-8 border-4 border-indigo-500 border-t-transparent rounded-full animate-spin"></div>
        <span class="mt-4 text-indigo-600 dark:text-indigo-400 text-lg font-semibold">Cargando datos...</span>
      </div>
      <!-- Contenido principal -->
      <div v-else>
        <!-- Mensaje de error si ocurre -->
        <div v-if="stockStore.error" class="text-center text-red-500 font-semibold mb-4">{{ stockStore.error }}</div>
        <!-- Tabla de stocks -->
        <div class="text-gray-500 text-center sm:text-left">
          <StockTable />
        </div>
      </div>
    </div>
  </div>
</template>


<script lang="ts" setup>
// --- StockList.vue ---
// Página principal: muestra listado de stocks y maneja carga inicial.

// --- Imports ordenados ---
import { onMounted } from 'vue'
import Navbar from '@/components/Navbar.vue'
import StockTable from '@/components/StockTable.vue'
import { useStockStore } from '@/store/stockStore'

// --- Store Pinia ---
const stockStore = useStockStore()

// Cargar los stocks al montar la página principal
onMounted(() => {
  stockStore.fetchStocks()
})
</script>


<style scoped>
/* Fondo degradado para la página de listado de stocks */
body {
  background: linear-gradient(135deg, #eef2ff 0%, #f8fafc 100%);
}
</style>
