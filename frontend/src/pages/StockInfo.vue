
<template>
  <!-- Página de recomendaciones de acciones destacadas -->
  <div>
    <!-- Barra de navegación principal -->
    <Navbar />
    <!-- Contenedor principal, ancho máximo casi total -->
    <div class="container mx-auto max-w-[95vw] py-10 px-4">
      <!-- Loader mientras se buscan recomendaciones -->
      <div v-if="loading" class="flex flex-col items-center justify-center py-12">
        <div class="w-10 h-10 border-4 border-indigo-500 border-t-transparent rounded-full animate-spin"></div>
        <span class="mt-4 text-indigo-600 dark:text-indigo-400 text-lg font-semibold">
          Buscando las mejores recomendaciones...
        </span>
      </div>

      <!-- Si hay recomendaciones -->
      <div v-else-if="stocks.length">
        <!-- Selector de cantidad de recomendaciones a mostrar -->
        <div class="flex items-center gap-4 mb-8">
          <label for="cantidad" class="text-base font-semibold text-gray-700 dark:text-gray-200">
            Ver mejores:
          </label>
          <select
            id="cantidad"
            v-model="cantidad"
            class="rounded-lg border border-gray-300 dark:border-gray-700 bg-white dark:bg-gray-800 px-3 py-2 text-base text-gray-800 dark:text-gray-100 focus:outline-none focus:ring-2 focus:ring-indigo-500"
          >
            <option v-for="n in [1,2,3,5,10,15,20,25,30,40,50,60,70,80,90,100]" :key="n" :value="n">Top {{ n }}</option>
          </select>
        </div>

        <!-- Grid de tarjetas de recomendaciones -->
        <div class="grid gap-8 grid-cols-1 md:grid-cols-2 xl:grid-cols-3">
          <div
            v-for="(stock, idx) in stocks.slice(0, cantidad)"
            :key="stock.id"
            class="bg-white dark:bg-gray-900 rounded-2xl shadow-2xl p-8 border border-gray-200 dark:border-gray-800 relative flex flex-col w-full max-w-[90%] mx-auto cursor-pointer hover:scale-[1.02] transition-transform"
            @click="openProfile(stock.ticker)"
            title="Ver información externa de la empresa"
          >
            <!-- Etiqueta de rating actual -->
            <div class="absolute top-0 right-0 mt-4 mr-4">
              <span class="inline-block px-3 py-1 rounded-full text-xs font-bold bg-indigo-100 text-indigo-700 dark:bg-indigo-800 dark:text-indigo-200">
                {{ stock.rating_to }}
              </span>
            </div>

            <!-- Título con ticker y empresa -->
            <h1 class="text-2xl font-extrabold text-indigo-700 dark:text-indigo-300 mb-2 flex items-center gap-2">
              #{{ idx+1 }} {{ stock.ticker }}
              <span class="text-lg font-semibold text-gray-500 dark:text-gray-400">
                ({{ stock.company }})
              </span>
            </h1>

            <!-- Datos principales de la recomendación -->
            <div class="flex flex-wrap gap-6 mt-4 mb-6">
              <!-- Precio actual -->
              <div class="flex flex-col items-center">
                <span class="text-sm text-gray-500 dark:text-gray-400">Precio actual</span>
                <span class="text-2xl font-bold text-green-600 dark:text-green-400">
                  ${{ stock.target_from }}
                </span>
              </div>
              <!-- Precio objetivo -->
              <div class="flex flex-col items-center">
                <span class="text-sm text-gray-500 dark:text-gray-400">Precio objetivo</span>
                <span class="text-2xl font-bold text-blue-600 dark:text-blue-400">
                  ${{ stock.target_to }}
                </span>
              </div>
              <!-- Potencial de ganancia -->
              <div class="flex flex-col items-center">
                <span class="text-sm text-gray-500 dark:text-gray-400">Potencial de ganancia</span>
                <span class="text-2xl font-bold text-indigo-700 dark:text-indigo-300">
                  {{ winPercent(stock) }}%
                </span>
              </div>
              <!-- Fecha de recomendación -->
              <div class="flex flex-col items-center">
                <span class="text-sm text-gray-500 dark:text-gray-400">Fecha recomendación</span>
                <span class="text-lg font-semibold text-gray-700 dark:text-gray-200">
                  {{ formatDate(stock.time) }}
                </span>
              </div>
            </div>

            <!-- Detalles adicionales -->
            <div class="flex flex-col md:flex-row gap-4 mt-6">
              <!-- Casa de bolsa -->
              <div class="flex-1">
                <div class="mb-2 text-sm text-gray-500 dark:text-gray-400">Casa de bolsa</div>
                <div class="font-semibold text-gray-800 dark:text-gray-100">{{ stock.brokerage }}</div>
              </div>
              <!-- Acción recomendada -->
              <div class="flex-1">
                <div class="mb-2 text-sm text-gray-500 dark:text-gray-400">Acción recomendada</div>
                <div class="font-semibold text-gray-800 dark:text-gray-100">{{ stock.action }}</div>
              </div>
              <!-- Cambio de rating -->
              <div class="flex-1">
                <div class="mb-2 text-sm text-gray-500 dark:text-gray-400">Cambio de rating</div>
                <div class="font-semibold text-gray-800 dark:text-gray-100">{{ stock.rating_from }} → {{ stock.rating_to }}</div>
              </div>
            </div>

            <!-- Mensaje de oportunidad -->
            <div class="mt-8 p-4 rounded-lg bg-indigo-50 dark:bg-indigo-900/30 text-indigo-800 dark:text-indigo-200 text-center text-lg font-medium">
              <span>Oportunidad de inversión según los datos actuales y las recomendaciones de los analistas.</span>
            </div>
          </div>
        </div>
      </div>

      <!-- Modal de perfil externo -->
      <div v-if="showProfile" class="fixed inset-0 z-50 flex items-center justify-center bg-black/40">
        <div class="bg-white dark:bg-gray-900 rounded-2xl shadow-2xl p-8 max-w-lg w-full relative">
          <button @click="showProfile = false" class="absolute top-2 right-2 text-gray-400 hover:text-red-500 text-2xl">&times;</button>
          <div v-if="profileLoading" class="flex flex-col items-center justify-center py-8">
            <div class="w-8 h-8 border-4 border-indigo-500 border-t-transparent rounded-full animate-spin"></div>
            <span class="mt-4 text-indigo-600 dark:text-indigo-400 text-base font-semibold">Cargando información...</span>
          </div>
          <div v-else-if="companyProfile">
            <div class="flex items-center gap-4 mb-4">
              <img v-if="companyProfile.logo" :src="companyProfile.logo" alt="Logo" class="w-14 h-14 rounded-full border" />
              <div>
                <div class="text-xl font-bold text-indigo-700 dark:text-indigo-300">{{ companyProfile.name }}</div>
                <div class="text-sm text-gray-500 dark:text-gray-400">{{ companyProfile.ticker }}</div>
              </div>
            </div>
            <div class="mb-2 text-gray-700 dark:text-gray-200"><b>Industria:</b> {{ companyProfile.finnhubIndustry || '-' }}</div>
            <div class="mb-2 text-gray-700 dark:text-gray-200"><b>País:</b> {{ companyProfile.country || '-' }}</div>
            <div class="mb-2 text-gray-700 dark:text-gray-200"><b>Web:</b> <a :href="companyProfile.weburl" target="_blank" class="text-blue-600 underline">{{ companyProfile.weburl }}</a></div>
            <div class="mb-2 text-gray-700 dark:text-gray-200"><b>Capitalización de mercado:</b> {{ companyProfile.marketCapitalization ? '$' + companyProfile.marketCapitalization + 'M' : '-' }}</div>
            <div class="mb-2 text-gray-700 dark:text-gray-200"><b>Acciones en circulación:</b> {{ companyProfile.shareOutstanding ? companyProfile.shareOutstanding : '-' }}</div>
          </div>
          <div v-else class="text-center text-gray-500 dark:text-gray-400 py-8">No se encontró información externa para este ticker.</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
// Importación de componentes y utilidades
import Navbar from '@/components/Navbar.vue'
import { ref, onMounted } from 'vue'
import { svc_getBestStockRecommendation } from '@/services/stockService'
import type { Stock } from '@/types/Stock'
import { getCompanyProfile } from '@/services/externalStockService'

// Estado reactivo para las mejores recomendaciones
const stocks = ref<Stock[]>([])
const loading = ref(true)
const cantidad = ref(3) // Cantidad de recomendaciones a mostrar

// Estado para el modal de perfil externo
const showProfile = ref(false)
const companyProfile = ref<any>(null)
const profileLoading = ref(false)

/**
 * Calcula el porcentaje de ganancia potencial para un stock
 * @param stock Objeto stock
 * @returns {string} Porcentaje de ganancia o '--' si no es válido
 */
function winPercent(stock: any): string {
  if (!stock) return '--'
  const from = parseFloat(stock.target_from)
  const to = parseFloat(stock.target_to)
  if (isNaN(from) || isNaN(to)) return '--'
  return ((to - from) / from * 100).toFixed(2)
}

/**
 * Formatea la fecha a formato legible en español
 * @param dateStr Fecha en string
 * @returns {string} Fecha formateada o '--' si no es válida
 */
function formatDate(dateStr: string): string {
  if (!dateStr) return '--'
  const d = new Date(dateStr)
  return d.toLocaleDateString('es-ES', { year: 'numeric', month: 'long', day: 'numeric' })
}

/**
 * Abre el modal y carga el perfil externo de la empresa
 * @param ticker Ticker de la acción
 */
async function openProfile(ticker: string) {
  showProfile.value = true
  companyProfile.value = null
  profileLoading.value = true
  try {
    const profile = await getCompanyProfile(ticker)
    companyProfile.value = profile && profile.ticker ? profile : null
  } catch (e) {
    companyProfile.value = null
  }
  profileLoading.value = false
}

// Carga las mejores recomendaciones al montar el componente
onMounted(async () => {
  loading.value = true
  try {
    const result = await svc_getBestStockRecommendation()
    // Si el resultado es un array, lo asignamos; si es un solo objeto, lo envolvemos en array
    if (Array.isArray(result)) {
      stocks.value = result
    } else if (result) {
      stocks.value = [result]
    } else {
      stocks.value = []
    }
  } catch (e) {
    stocks.value = []
  }
  loading.value = false
})
</script>

<style scoped>
/* Contenedor principal con altura mínima */
.container {
  min-height: 60vh;
}
/* Sombra personalizada para tarjetas */
.shadow-2xl {
  box-shadow: 0 10px 32px 0 rgba(60, 72, 88, 0.15), 0 2px 4px 0 rgba(60, 72, 88, 0.10);
}
</style>