// Servicio para obtener información adicional de acciones desde Finnhub
import axios from 'axios'

// URL base de la API de Finnhub
const BASE_URL = 'https://finnhub.io/api/v1'
// API_KEY: Tu clave pública de Finnhub
const API_KEY = import.meta.env.VITE_EXTERNALKEY

/**
 * Obtiene el perfil de la empresa por símbolo (ticker)
 * @param symbol Ticker de la acción (ej: 'AAPL', 'MSFT')
 * @returns Perfil de la empresa (nombre, industria, país, web, etc.)
 */
export async function getCompanyProfile(symbol: string) {
  const res = await axios.get(`${BASE_URL}/stock/profile2`, {
    params: { symbol, token: API_KEY },
    headers: { 'Content-Type': 'application/json' }
  })
  return res.data
}

/**
 * Obtiene cotización en tiempo real por símbolo
 * @param symbol Ticker de la acción (ej: 'AAPL', 'MSFT')
 * @returns Cotización actual, apertura, cierre anterior, etc.
 */
export async function getQuote(symbol: string) {
  const res = await axios.get(`${BASE_URL}/quote`, {
    params: { symbol, token: API_KEY },
    headers: { 'Content-Type': 'application/json' }
  })
  return res.data
}
