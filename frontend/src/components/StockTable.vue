<template>
  <!--
    StockTable.vue
    Tabla principal para la gestión de stocks.
    Incluye búsqueda, paginación, edición y eliminación.
  -->
  <section class="p-4 bg-white dark:bg-gray-900 rounded-2xl shadow-lg">
    <!-- Encabezado con título y buscador -->
    <div class="flex flex-col md:flex-row md:items-center md:justify-between gap-4 mb-6">
      <h2 class="text-2xl font-bold text-gray-800 dark:text-white"></h2>
      <input
        v-model="search"
        type="text"
        placeholder="Buscar en la tabla ..."
        class="px-4 pr-20 py-2 border border-gray-300 dark:border-gray-700 rounded-lg bg-white dark:bg-gray-800 text-sm text-gray-800 dark:text-white placeholder:text-gray-400 focus:outline-none focus:ring-2 focus:ring-indigo-500 transition"
      />
    </div>

    <!-- Spinner de carga -->
    <div v-if="loading" class="flex justify-center items-center py-6">
      <div class="w-6 h-6 border-4 border-indigo-500 border-t-transparent rounded-full animate-spin"></div>
      <span class="ml-3 text-indigo-600 dark:text-indigo-400">Cargando datos...</span>
    </div>

    <!-- Tabla de stocks -->
    <div v-else class="overflow-auto max-h-[400px] border rounded-lg shadow-sm">
      <table class="w-full text-sm text-left border-collapse">
        <thead>
          <tr class="bg-gray-100 dark:bg-gray-700 text-gray-700 dark:text-white sticky top-0">
            <!-- <th class="px-4 py-2">ID</th> -->
            <th class="px-4 py-2 text-center">Ticker</th>
            <th class="px-4 py-2 text-center">Target From</th>
            <th class="px-4 py-2 text-center">Target To</th>
            <th class="px-4 py-2 text-center">Company</th>
            <th class="px-4 py-2 text-center">Action</th>
            <th class="px-4 py-2 text-center">Brokerage</th>
            <th class="px-4 py-2 text-center">Rating From</th>
            <th class="px-4 py-2 text-center">Rating To</th>
            <th class="px-4 py-2 text-center">Time</th>
            <th class="px-4 py-2 text-center">Acciones</th>
          </tr>
        </thead>
        <tbody>
          <!-- Fila por cada stock paginado -->
          <tr
            v-for="stock in paginatedStocks"
            :key="stock.id"
            class="hover:bg-gray-50 dark:hover:bg-gray-800 transition"
          >
            <!-- <td class="px-4 py-2">{{ stock.id }}</td> -->
            <td class="px-4 py-2 text-center">{{ stock.ticker }}</td>
            <td class="px-4 py-2 text-center">${{ stock.target_from }}</td>
            <td class="px-4 py-2 text-center">${{ stock.target_to }}</td>
            <td class="px-4 py-2 text-center">{{ stock.company }}</td>
            <td class="px-4 py-2 text-center">{{ stock.action }}</td>
            <td class="px-4 py-2 text-center">{{ stock.brokerage }}</td>
            <td class="px-4 py-2 text-center">{{ stock.rating_from }}</td>
            <td class="px-4 py-2 text-center">{{ stock.rating_to }}</td>
            <td class="px-4 py-2 text-center">{{ stock.time }}</td>
            <td class="px-4 py-2 text-center">
              <div class="flex justify-center gap-2">
                <!-- Botón para editar -->
                <button
                  @click="openModal(stock)"
                  class="text-black bg-yellow-500 hover:bg-yellow-600 px-3 py-1 rounded text-sm"
                >
                  Editar
                </button>
                <!-- Botón para eliminar -->
                <button
                  @click="deleteStock(stock.id)"
                  class="text-white bg-red-500 hover:bg-red-600 px-3 py-1 rounded text-sm"
                >
                  Eliminar
                </button>
              </div>
            </td>
          </tr>

          <!-- Mensaje si no hay resultados -->
          <tr v-if="filteredStocks.length === 0">
            <td colspan="11" class="text-center py-4 text-gray-500 dark:text-gray-400">
              No se encontraron resultados.
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- Paginación -->
    <div class="flex flex-wrap justify-center gap-2 mt-6">
      <!-- Botón para ir a la primera página -->
      <button
        @click="goToPage(1)"
        :disabled="currentPage === 1"
        class="px-3 py-1 bg-gray-200 dark:bg-gray-700 rounded disabled:opacity-50"
      >
        Inicio
      </button>

      <!-- Puntos suspensivos si hay muchas páginas antes -->
      <span v-if="currentPage > 3" class="px-2 py-1 text-gray-500 dark:text-gray-400">...</span>

      <!-- Botones de páginas visibles -->
      <button
        v-for="page in visiblePages"
        :key="page"
        @click="goToPage(page)"
        class="px-3 py-1 rounded transition"
        :class="{
          'bg-indigo-600 text-white': currentPage === page,
          'bg-gray-200 dark:bg-gray-700 text-gray-700 dark:text-white': currentPage !== page,
        }"
      >
        {{ page }}
      </button>

      <!-- Puntos suspensivos si hay muchas páginas después -->
      <span v-if="currentPage < totalPages - 2" class="px-2 py-1 text-gray-500 dark:text-gray-400">...</span>

      <!-- Botón para ir a la última página -->
      <button
        @click="goToPage(totalPages)"
        :disabled="currentPage === totalPages"
        class="px-3 py-1 bg-gray-200 dark:bg-gray-700 rounded disabled:opacity-50"
      >
        Final
      </button>
    </div>

    <!-- Modal de edición -->
    <Modal
      v-if="selectedStock"
      :stock="selectedStock"
      @close="selectedStock = null"
      @save="updateStock"
    />
  </section>
</template>

<script setup lang="ts">
// --- StockTable.vue ---
// Componente de tabla para mostrar, buscar, editar y eliminar stocks con paginación.

// --- Imports ordenados ---
import { ref } from "vue";
import { storeToRefs } from "pinia";
import Modal from "./Modal.vue";
import type { Stock } from "@/types/Stock";
import { useStockStore } from "@/store/stockStore";

// --- Store y referencias reactivas ---
const stockStore = useStockStore();
const {
  search,            // Búsqueda reactiva
  paginatedStocks,   // Stocks paginados
  filteredStocks,    // Stocks filtrados por búsqueda
  currentPage,       // Página actual
  totalPages,        // Total de páginas
  visiblePages,      // Páginas visibles en el paginador
  loading            // Estado de carga
} = storeToRefs(stockStore);

// Métodos del store para paginación y CRUD
const goToPage = stockStore.goToPage;           // Cambiar de página
const deleteStock = stockStore.deleteStock;     // Eliminar stock
const updateStock = stockStore.updateStock;     // Actualizar stock

// --- Modal de edición ---
const selectedStock = ref<Stock | null>(null);  // Stock seleccionado para editar
function openModal(stock: Stock) {
  selectedStock.value = stock;
}
</script>

<style scoped>
/* Estilos para la tabla de stocks */
table {
  border-spacing: 0;
}
th,
td {
  border: 1px solid #e5e7eb;
}
</style>
