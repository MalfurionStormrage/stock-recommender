// --- Tipo genérico para respuestas de la API de stocks ---
export interface StockApiResponse<T> {
  status: string        // Estado de la respuesta ("success", "error", etc)
  data: T               // Datos devueltos por la API
  message?: string      // Mensaje opcional (usado en errores o confirmaciones)
}