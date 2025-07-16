
// --- routes.ts ---
// Definición de rutas principales de la aplicación

import StockList from '@/pages/StockList.vue'
import StockCreate from '@/pages/StockCreate.vue'
import StockInfo from '@/pages/StockInfo.vue'

/**
 * Rutas de la aplicación:
 * - Home: listado principal de stocks
 * - CreateStock: formulario para crear un nuevo stock
 * - StockInfo: información detallada de un stock (actualmente placeholder)
 *
 * Nota: La ruta StockInfo puede eliminarse si no se utiliza.
 */
export const routes = [
  {
    path: '/',
    name: 'Home',
    component: StockList, // Página principal
  },
  {
    path: '/create',
    name: 'CreateStock',
    component: StockCreate, // Formulario de creación
  },
  {
    path: '/StockInfo',
    name: 'StockInfo',
    component: StockInfo, // Detalle de stock (placeholder, puede eliminarse)
    props: true,
  },
]
