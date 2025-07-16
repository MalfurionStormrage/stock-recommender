// --- Modelo de datos para un Stock ---
export interface Stock {
  id: string           // ID único del stock
  ticker: string       // Símbolo bursátil
  target_from: string  // Precio objetivo inicial
  target_to: string    // Precio objetivo final
  company: string      // Nombre de la compañía
  action: string       // Acción recomendada
  brokerage: string    // Casa de bolsa
  rating_from: string  // Calificación inicial
  rating_to: string    // Calificación final
  time: string         // Fecha y hora de la recomendación
  origin?: string      // (Opcional) Origen del dato
  next_page?: string   // (Opcional) Siguiente página (paginación)
}
