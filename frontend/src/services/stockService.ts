
// --- stockService.ts ---
// Servicios para consumir la API de stocks y mostrar notificaciones.

import axios, { AxiosError } from 'axios'
import Swal from 'sweetalert2'
import type { Stock } from '@/types/Stock'
import type { StockApiResponse } from '@/types/StockApiResponse'

// const API_URL = 'http://localhost:8080/api/v1/stock' // Descomentar esta línea para pruebas locales
const API_URL = 'http://ec2-18-205-28-183.compute-1.amazonaws.com:8080/api/v1/stock' // URL de la API desplegada en AWS

/**
 * Obtiene todos los stocks desde la API.
 * Formatea los campos numéricos y el id como string.
 */
export async function svc_getAllStocks(): Promise<Stock[]> {
  const res = await axios.get<StockApiResponse<Stock[]>>(API_URL)
  if (res.data.status !== 'success') throw new Error('Error en la API')

  return res.data.data.map((item: any) => ({
    ...item,
    id: String(item.id),
    target_from: Number(item.target_from).toFixed(2),
    target_to: Number(item.target_to).toFixed(2),
  }))
}

/**
 * Crea un nuevo stock en la API.
 * @param stock - Objeto stock sin id
 * @returns true si fue exitoso, false si hubo error
 */
export async function svc_createStock(stock: Omit<Stock, 'id'>): Promise<boolean> {
  if (stock.time) {
    stock.time = parseToISOString(stock.time)
  }

  stock.target_from = stock.target_from.toString().replace(/[$,]/g, '')
  stock.target_to = stock.target_to.toString().replace(/[$,]/g, '')

  try {
    const res = await axios.post(API_URL, stock, {
      headers: { 'Content-Type': 'application/json' }
    })

    if (res.data.status === 'success') {
      await Swal.fire("¡Creado!", res.data.message, "success")
      return true
    }

  } catch (err) {
    const error = err as AxiosError<{ message: string }>
    await Swal.fire("Error", error.response?.data?.message || "No se pudo crear.", "error")
    return false
  }

  return false
}

/**
 * Actualiza un stock existente en la API.
 * @param id - ID del stock
 * @param stock - Objeto stock sin id
 * @returns true si fue exitoso, false si hubo error
 */
export async function svc_updateStock(id: string, stock: Omit<Stock, 'id'>): Promise<boolean> {
  if (stock.time) {
    stock.time = parseToISOString(stock.time)
  }

  stock.target_from = stock.target_from.toString().replace(/[$,]/g, '')
  stock.target_to = stock.target_to.toString().replace(/[$,]/g, '')

  try {
    const res = await axios.put(`${API_URL}/${id}`, stock, {
      headers: { 'Content-Type': 'application/json' }
    })

    if (res.data.status === 'success') {
      await Swal.fire("¡Actualizado!", res.data.message, "success")
      return true
    }

  } catch (err) {
    const error = err as AxiosError<{ message: string }>
    await Swal.fire("Error", error.response?.data?.message || "No se pudo actualizar.", "error")
  }

  return false
}

/**
 * Elimina un stock de la API tras confirmación del usuario.
 * @param id - ID del stock
 * @returns true si fue exitoso, false si hubo error
 */
export async function svc_deleteStock(id: string): Promise<boolean> {
  const result = await Swal.fire({
    title: "¿Estás seguro?",
    text: "¡Esta acción no se puede deshacer!",
    icon: "warning",
    showCancelButton: true,
    confirmButtonText: "Sí, eliminar",
    confirmButtonColor: "#d33",
  })

  if (result.isConfirmed) {
    try {
      const res = await axios.delete(`${API_URL}/${id}`)
      await Swal.fire("¡Eliminado!", res.data.message, "success")
      return true
    } catch (err) {
      const error = err as AxiosError<{ message: string }>
      await Swal.fire("Error", error.response?.data?.message || "No se pudo eliminar.", "error")
    }
  }

  return false
}

export async function svc_getBestStockRecommendation(): Promise<Stock[] | null> {
  try {
    // Llamada a la API para obtener las mejores recomendación
    const res = await axios.get<StockApiResponse<Stock[]>>(`${API_URL}/best-recommendation`)
    if (res.data.status !== 'success') throw new Error('Error al obtener la mejor recomendación')

    return res.data.data.length > 0 ? res.data.data : null
  } catch (err) {
    const error = err as AxiosError<{ message: string }>
    await Swal.fire("Error", error.response?.data?.message || "No se pudo obtener la mejor recomendación.", "error")
    return null
  }
}

// --- Utilidad para parsear fechas a ISO ---
/**
 * Convierte una fecha string a formato ISO 8601.
 * Soporta formatos ISO y "dd/mm/yyyy hh:mm".
 * Lanza error si el formato es inválido.
 */
function parseToISOString(input: string): string {
  if (!input) return '';

  // Si ya es un formato ISO, lo intentamos parsear directo
  const isoDate = new Date(input);
  if (!isNaN(isoDate.getTime())) {
    return isoDate.toISOString();
  }

  // Si es formato tipo "16/07/2025 12:07"
  const parts = input.split('/');
  if (parts.length === 3) {
    const [day, month, yearTime] = parts;
    const [year, time] = yearTime.split(' ');
    if (year && time) {
      const formatted = `${year}-${month.padStart(2, '0')}-${day.padStart(2, '0')}T${time}:00`;
      const parsed = new Date(formatted);
      if (!isNaN(parsed.getTime())) {
        return parsed.toISOString();
      }
    }
  }

  // Si nada funcionó, lanza error
  throw new Error(`Formato de fecha inválido: "${input}"`);
}