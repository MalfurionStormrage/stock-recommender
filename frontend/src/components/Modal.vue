<template>
  <!-- Fondo oscuro desenfocado -->
  <div class="fixed inset-0 z-50 flex items-center justify-center bg-black/40 backdrop-blur-sm">
    <div class="bg-white dark:bg-gray-900 text-gray-800 dark:text-gray-100 p-8 rounded-2xl shadow-xl max-w-3xl w-full overflow-y-auto max-h-[90vh]">
      <!-- Encabezado -->
      <div class="flex justify-between items-center mb-6">
        <h2 class="text-2xl font-semibold">Actualizar datos del registro</h2>
        <button @click="$emit('close')" class="text-gray-400 hover:text-gray-600 text-2xl font-bold">
          &times;
        </button>
      </div>

      <!-- Formulario -->
      <form class="grid grid-cols-1 sm:grid-cols-2 gap-6">
        <div>
          <label class="block mb-1 text-sm font-medium">Ticker</label>
          <input v-model="localStock.ticker" disabled type="text" placeholder="Ej: AKBA" class="w-full rounded-lg border border-gray-300 dark:border-gray-700 bg-white dark:bg-gray-800 px-3 py-2 text-sm text-gray-800 dark:text-gray-100 placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 transition" />
        </div>

        <div>
          <label class="block mb-1 text-sm font-medium">Target From</label>
          <input v-model="localStock.target_from" required type="number" min="0" step="0.01" placeholder="Ej: 8.00" class="w-full rounded-lg border border-gray-300 dark:border-gray-700 bg-white dark:bg-gray-800 px-3 py-2 text-sm text-gray-800 dark:text-gray-100 placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 transition" />
        </div>

        <div>
          <label class="block mb-1 text-sm font-medium">Target To</label>
          <input v-model="localStock.target_to" required type="number" min="0" step="0.01" placeholder="Ej: 8.00" class="w-full rounded-lg border border-gray-300 dark:border-gray-700 bg-white dark:bg-gray-800 px-3 py-2 text-sm text-gray-800 dark:text-gray-100 placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 transition" />
        </div>

        <div>
          <label class="block mb-1 text-sm font-medium">Company</label>
          <input v-model="localStock.company" required type="text" placeholder="Ej: Akebia Therapeutics" class="w-full rounded-lg border border-gray-300 dark:border-gray-700 bg-white dark:bg-gray-800 px-3 py-2 text-sm text-gray-800 dark:text-gray-100 placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 transition" />
        </div>

        <div>
          <label class="block mb-1 text-sm font-medium">Action</label>
          <input v-model="localStock.action" required type="text" placeholder="Ej: initiated by" class="w-full rounded-lg border border-gray-300 dark:border-gray-700 bg-white dark:bg-gray-800 px-3 py-2 text-sm text-gray-800 dark:text-gray-100 placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 transition" />
        </div>

        <div>
          <label class="block mb-1 text-sm font-medium">Brokerage</label>
          <input v-model="localStock.brokerage" required type="text" placeholder="Ej: HC Wainwright" class="w-full rounded-lg border border-gray-300 dark:border-gray-700 bg-white dark:bg-gray-800 px-3 py-2 text-sm text-gray-800 dark:text-gray-100 placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 transition" />
        </div>

        <div>
          <label class="block mb-1 text-sm font-medium">Rating From</label>
          <input v-model="localStock.rating_from" required type="text" placeholder="Ej: Buy" class="w-full rounded-lg border border-gray-300 dark:border-gray-700 bg-white dark:bg-gray-800 px-3 py-2 text-sm text-gray-800 dark:text-gray-100 placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 transition" />
        </div>

        <div>
          <label class="block mb-1 text-sm font-medium">Rating To</label>
          <input v-model="localStock.rating_to" required type="text" placeholder="Ej: Buy" class="w-full rounded-lg border border-gray-300 dark:border-gray-700 bg-white dark:bg-gray-800 px-3 py-2 text-sm text-gray-800 dark:text-gray-100 placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 transition" />
        </div>

        <div>
          <label class="block mb-1 text-sm font-medium">Time</label>
          <input v-model="localStock.time" required type="datetime-local" placeholder="Ej: 2025-05-01T12:05" class="w-full rounded-lg border border-gray-300 dark:border-gray-700 bg-white dark:bg-gray-800 px-3 py-2 text-sm text-gray-800 dark:text-gray-100 placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 transition" />
        </div>

        <!--
        <div>
          <label class="block mb-1 text-sm font-medium">Origin</label>
          <input v-model="localStock.origin" type="text" placeholder="Ej: " class="w-full rounded-lg border border-gray-300 dark:border-gray-700 bg-white dark:bg-gray-800 px-3 py-2 text-sm text-gray-800 dark:text-gray-100 placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 transition" />
        </div>

        <div>
          <label class="block mb-1 text-sm font-medium">Next Page</label>
          <input v-model="localStock.next_page" type="text" placeholder="Ej: ETR" class="w-full rounded-lg border border-gray-300 dark:border-gray-700 bg-white dark:bg-gray-800 px-3 py-2 text-sm text-gray-800 dark:text-gray-100 placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 transition" />
        </div>
        -->
      </form>

      <!-- Footer -->
      <div class="mt-8 flex justify-end gap-4">
        <button @click="$emit('close')" class="px-4 py-2 bg-gray-200 hover:bg-gray-300 text-gray-800 rounded-lg">
          Cancelar
        </button>
        <button @click="guardar" class="px-4 py-2 bg-indigo-600 hover:bg-indigo-700 text-white rounded-lg">
          Guardar
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
// --- Imports y tipado ---
import { reactive } from 'vue'
import type { Stock } from '@/types/Stock'

// Props y emits del modal
const props = defineProps<{ stock: Stock }>()
const emit = defineEmits<{
  (e: 'close'): void
  (e: 'save', value: Stock): void
}>()

// Convierte fecha de "01/05/2025 12:05" a "2025-05-01T12:05" para el input
function toDatetimeLocal(str: string): string {
  if (!str) return '';
  const [d, m, yhm] = str.split('/');
  if (!d || !m || !yhm) return '';
  const [y, hm] = yhm.split(' ');
  if (!y || !hm) return '';
  return `${y}-${m.padStart(2, '0')}-${d.padStart(2, '0')}T${hm}`;
}

// Convierte fecha de "2025-05-01T12:05" a "01/05/2025 12:05" para guardar
function fromDatetimeLocal(str: string): string {
  if (!str) return '';
  const [date, time] = str.split('T');
  if (!date || !time) return '';
  const [y, m, d] = date.split('-');
  if (!y || !m || !d) return '';
  return `${d}/${m}/${y} ${time}`;
}

// Estado local del stock a editar
const localStock = reactive({
  ...props.stock,
  time: toDatetimeLocal(props.stock.time)
})

// Guarda los cambios y emite eventos
function guardar() {
  // Validar campos requeridos
  const requiredFields = [
    'ticker', 'target_from', 'target_to', 'company', 'action', 'brokerage',
    'rating_from', 'rating_to', 'time'
  ];
  for (const field of requiredFields) {
    if (!(localStock as any)[field] || String((localStock as any)[field]).trim() === '') {
      alert('Todos los campos son obligatorios. Por favor, completa: ' + field.replace('_', ' '));
      return;
    }
  }
  emit('save', { ...localStock, time: fromDatetimeLocal(localStock.time) })
  emit('close')
}
</script>
