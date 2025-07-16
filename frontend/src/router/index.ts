
// --- index.ts ---
// Configuración principal del router de Vue Router 4

import { createRouter, createWebHistory } from 'vue-router'
import { routes } from './routes'

// Exporta la instancia del router para ser usada en main.ts
export const router = createRouter({
  history: createWebHistory(), // Usa el modo history para URLs limpias
  routes,                      // Rutas definidas en routes.ts
})