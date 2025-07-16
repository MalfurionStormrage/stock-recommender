
<template>
  <!--
    StockCreate.vue
    Página para crear un nuevo stock.
    Incluye formulario validado y botones de acción.
  -->
  <Navbar />
  <div class="flex justify-center items-center min-h-[80vh] bg-gradient-to-br from-indigo-50 to-white dark:from-gray-900 dark:to-gray-800 py-8">
    <div class="w-full max-w-5xl bg-white dark:bg-gray-900 rounded-2xl shadow-2xl p-10 border border-gray-200 dark:border-gray-800">
      <h2 class="text-3xl font-extrabold mb-8 text-center text-indigo-700 dark:text-indigo-300 tracking-tight">Crear nuevo Stock</h2>
      <form @submit.prevent="onSubmit" class="grid grid-cols-1 md:grid-cols-3 gap-x-8 gap-y-6">
        <!-- Campo: Ticker -->
        <div>
          <label class="block mb-2 text-base font-semibold text-gray-700 dark:text-gray-200">Ticker</label>
          <input v-model="form.ticker" required type="text" placeholder="Ej: AKBA" class="w-full rounded-lg border border-gray-300 dark:border-gray-700 bg-white dark:bg-gray-800 px-4 py-2 text-base text-gray-800 dark:text-gray-100 placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 transition shadow-sm" />
        </div>
        <!-- Campo: Target From -->
        <div>
          <label class="block mb-2 text-base font-semibold text-gray-700 dark:text-gray-200">Target From</label>
          <input v-model="form.target_from" required type="number" min="0" step="0.01" placeholder="Ej: 8.00" class="w-full rounded-lg border border-gray-300 dark:border-gray-700 bg-white dark:bg-gray-800 px-4 py-2 text-base text-gray-800 dark:text-gray-100 placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 transition shadow-sm" />
        </div>
        <!-- Campo: Target To -->
        <div>
          <label class="block mb-2 text-base font-semibold text-gray-700 dark:text-gray-200">Target To</label>
          <input v-model="form.target_to" required type="number" min="0" step="0.01" placeholder="Ej: 8.00" class="w-full rounded-lg border border-gray-300 dark:border-gray-700 bg-white dark:bg-gray-800 px-4 py-2 text-base text-gray-800 dark:text-gray-100 placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 transition shadow-sm" />
        </div>
        <!-- Campo: Company -->
        <div>
          <label class="block mb-2 text-base font-semibold text-gray-700 dark:text-gray-200">Company</label>
          <input v-model="form.company" required type="text" placeholder="Ej: Akebia Therapeutics" class="w-full rounded-lg border border-gray-300 dark:border-gray-700 bg-white dark:bg-gray-800 px-4 py-2 text-base text-gray-800 dark:text-gray-100 placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 transition shadow-sm" />
        </div>
        <!-- Campo: Action -->
        <div>
          <label class="block mb-2 text-base font-semibold text-gray-700 dark:text-gray-200">Action</label>
          <input v-model="form.action" required type="text" placeholder="Ej: initiated by" class="w-full rounded-lg border border-gray-300 dark:border-gray-700 bg-white dark:bg-gray-800 px-4 py-2 text-base text-gray-800 dark:text-gray-100 placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 transition shadow-sm" />
        </div>
        <!-- Campo: Brokerage -->
        <div>
          <label class="block mb-2 text-base font-semibold text-gray-700 dark:text-gray-200">Brokerage</label>
          <input v-model="form.brokerage" required type="text" placeholder="Ej: HC Wainwright" class="w-full rounded-lg border border-gray-300 dark:border-gray-700 bg-white dark:bg-gray-800 px-4 py-2 text-base text-gray-800 dark:text-gray-100 placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 transition shadow-sm" />
        </div>
        <!-- Campo: Rating From -->
        <div>
          <label class="block mb-2 text-base font-semibold text-gray-700 dark:text-gray-200">Rating From</label>
          <input v-model="form.rating_from" required type="text" placeholder="Ej: Buy" class="w-full rounded-lg border border-gray-300 dark:border-gray-700 bg-white dark:bg-gray-800 px-4 py-2 text-base text-gray-800 dark:text-gray-100 placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 transition shadow-sm" />
        </div>
        <!-- Campo: Rating To -->
        <div>
          <label class="block mb-2 text-base font-semibold text-gray-700 dark:text-gray-200">Rating To</label>
          <input v-model="form.rating_to" required type="text" placeholder="Ej: Buy" class="w-full rounded-lg border border-gray-300 dark:border-gray-700 bg-white dark:bg-gray-800 px-4 py-2 text-base text-gray-800 dark:text-gray-100 placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 transition shadow-sm" />
        </div>
        <!-- Campo: Time -->
        <div>
          <label class="block mb-2 text-base font-semibold text-gray-700 dark:text-gray-200">Time</label>
          <input v-model="form.time" required type="datetime-local" placeholder="Ej: 2025-05-01T12:05" class="w-full rounded-lg border border-gray-300 dark:border-gray-700 bg-white dark:bg-gray-800 px-4 py-2 text-base text-gray-800 dark:text-gray-100 placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 transition shadow-sm" />
        </div>
        <!--
        <div>
          <label class="block mb-2 text-base font-semibold text-gray-700 dark:text-gray-200">Origin</label>
          <input v-model="form.origin" required type="text" placeholder="Ej: " class="w-full rounded-lg border border-gray-300 dark:border-gray-700 bg-white dark:bg-gray-800 px-4 py-2 text-base text-gray-800 dark:text-gray-100 placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 transition shadow-sm" />
        </div>
        <div>
          <label class="block mb-2 text-base font-semibold text-gray-700 dark:text-gray-200">Next Page</label>
          <input v-model="form.next_page" required type="text" placeholder="Ej: ETR" class="w-full rounded-lg border border-gray-300 dark:border-gray-700 bg-white dark:bg-gray-800 px-4 py-2 text-base text-gray-800 dark:text-gray-100 placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 transition shadow-sm" />
        </div>
        -->
      </form>
      <!-- Botones de acción -->
      <div class="mt-10 flex justify-end gap-4">
        <button type="button" @click="onReset" class="px-6 py-2 bg-gray-200 hover:bg-gray-300 text-gray-800 rounded-lg text-lg font-semibold shadow-md transition">Limpiar</button>
        <button @click="onSubmit" class="px-6 py-2 bg-indigo-600 hover:bg-indigo-700 text-white rounded-lg text-lg font-semibold shadow-md transition">Guardar</button>
      </div>
    </div>
  </div>
</template>


<script setup lang="ts">
// --- StockCreate.vue ---
// Página para crear un nuevo stock. Incluye validación y acciones de formulario.

// --- Imports ordenados ---
import { reactive } from 'vue'
import { useRouter } from 'vue-router'
import type { Stock } from '@/types/Stock'
import { useStockStore } from '@/store/stockStore'
import Navbar from '@/components/Navbar.vue'

// --- Router y store ---
const router = useRouter()
const stockStore = useStockStore()

// --- Estado reactivo del formulario ---
// Solo los campos usados en la creación
const form = reactive<Partial<Stock>>({
  ticker: '',
  target_from: '',
  target_to: '',
  company: '',
  action: '',
  brokerage: '',
  rating_from: '',
  rating_to: '',
  time: ''
})

// Estado inicial para resetear el formulario
const initialForm = { ...form }

// Limpia el formulario y lo deja en estado inicial
function onReset() {
  Object.assign(form, initialForm)
}

// Envía el formulario para crear un nuevo stock
// Valida que todos los campos estén completos
async function onSubmit() {
  for (const key in form) {
    if (!form[key as keyof Stock] || String(form[key as keyof Stock]).trim() === '') {
      alert('Todos los campos son obligatorios. Por favor, completa: ' + key.replace('_', ' '));
      return;
    }
  }
  // Llama al store para crear el stock
  const res = await stockStore.createStock(form as Stock)
  if (res) {
    router.push('/')
  }
}
</script>


<style scoped>
/* Fondo degradado para la página de creación de stock */
body {
  background: linear-gradient(135deg, #eef2ff 0%, #f8fafc 100%);
}
</style>