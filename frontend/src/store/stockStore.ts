import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import type { Stock } from '@/types/Stock'
import * as stockService from '@/services/stockService'

// --- Store centralizado de stocks ---
export const useStockStore = defineStore('stock', () => {
  // --- Estado principal ---
  const stocks = ref<Stock[]>([])
  const loading = ref(false)
  const error = ref<string | null>(null)

  // --- CRUD y fetch ---
  async function fetchStocks() {
    loading.value = true
    try {
      stocks.value = await stockService.svc_getAllStocks()
      error.value = null
    } catch (e: any) {
      error.value = e.message
    } finally {
      loading.value = false
    }
  }

  // Crea un nuevo stock
  async function createStock(stock: Stock) {
    loading.value = true
    try {
      await stockService.svc_createStock(stock)
      await fetchStocks()
      error.value = null
      return true
    } catch (e: any) {
      error.value = e.message
      return false
    } finally {
      loading.value = false
    }
  }

  // Elimina un stock por id
  const deleteStock = async (id: string) => {
    const success = await stockService.svc_deleteStock(id)
    if (success) {
      stocks.value = stocks.value.filter(s => s.id !== id)
    }
  }

  // Actualiza un stock
  const updateStock = async (stock: Stock) => {
    const success = await stockService.svc_updateStock(stock.id, stock)
    if (success) await fetchStocks()
  }


  // --- Paginación, búsqueda y utilidades ---

  // Término de búsqueda para filtrar stocks
  const search = ref('')
  // Página actual del paginador
  const currentPage = ref(1)
  // Tamaño de página (cuántos stocks por página)
  const pageSize = ref(10)

  /**
   * Devuelve los stocks filtrados por el término de búsqueda.
   * Busca en todos los campos del stock.
   */
  const filteredStocks = computed(() => {
    const term = search.value.toLowerCase().trim()
    if (!term) return stocks.value
    return stocks.value.filter(stock =>
      Object.values(stock).some(val =>
        String(val).toLowerCase().includes(term)
      )
    )
  })

  /**
   * Devuelve los stocks de la página actual.
   */
  const paginatedStocks = computed(() => {
    const start = (currentPage.value - 1) * pageSize.value
    return filteredStocks.value.slice(start, start + pageSize.value)
  })

  /**
   * Total de páginas según el filtro y el tamaño de página.
   */
  const totalPages = computed(() =>
    Math.ceil(filteredStocks.value.length / pageSize.value)
  )

  /**
   * Devuelve las páginas visibles para el paginador (actual, anterior y siguiente).
   */
  const visiblePages = computed(() => {
    const pages: number[] = []
    const start = Math.max(1, currentPage.value - 1)
    const end = Math.min(totalPages.value, currentPage.value + 1)
    for (let i = start; i <= end; i++) {
      pages.push(i)
    }
    return pages
  })

  /**
   * Cambia la página actual si está dentro del rango válido.
   */
  const goToPage = (page: number) => {
    if (page >= 1 && page <= totalPages.value) {
      currentPage.value = page
    }
  }

  // --- Exponer API del store ---
  return {
    stocks,
    loading,
    error,
    fetchStocks,
    createStock,
    updateStock,
    deleteStock,
    search,
    currentPage,
    pageSize,
    filteredStocks,
    paginatedStocks,
    totalPages,
    visiblePages,
    goToPage
  }
})
