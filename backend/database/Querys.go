package database

//-----------------------------------------//
//-- QUERIES UTILIZADOS EN LA APLICACION --//
//----------------------------------------//

// -- QUERY PARA CREAR TABLA EN CASO DE NO EXISTIR --//
const CreateTableStock = `
CREATE TABLE IF NOT EXISTS stock (
	id SERIAL PRIMARY KEY,
	ticker TEXT NOT NULL,
	target_from NUMERIC(12, 2),
	target_to NUMERIC(12, 2),
	company TEXT,
	action TEXT,
	brokerage TEXT,
	rating_from TEXT,
	rating_to TEXT,
	time TIMESTAMP,
	origin TEXT,
	next_page TEXT,
	UNIQUE (ticker, time)
);`

// -- OBTENER TODOS LOS REGISTROS --//
const Get_all_stocks = `
SELECT id, ticker, target_from, target_to, company, action, brokerage,
       rating_from, rating_to, TO_CHAR(time, 'DD/MM/YYYY HH:MM') as time,origin, next_page
FROM stock;
`

// -- FILTRAR REGISTROS POR TICKER O NEXT_PAGE --//
const FilterStocks = `
SELECT id, ticker, target_from, target_to, company, action, brokerage,
       rating_from, rating_to, TO_CHAR(time, 'DD/MM/YYYY HH:MM') as time, origin,next_page
FROM stock
WHERE ticker ILIKE '%' || $1 || '%' OR next_page ILIKE '%' || $2 || '%';
`

// -- INSERTAR NUEVO REGISTRO --//
const Create_stock = `
INSERT INTO stock (
	ticker, target_from, target_to, company, action, brokerage,
	rating_from, rating_to, time,origin, next_page
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10 , $11)
ON CONFLICT DO NOTHING;
`

// -- ACTUALIZAR REGISTRO EXISTENTE --//
const Update_stock = `
UPDATE stock
SET target_from = $1,
    target_to = $2,
    company = $3,
    action = $4,
    brokerage = $5,
    rating_from = $6,
    rating_to = $7,
    time = $8,
	origin = $9,
    next_page = $10
WHERE id = $11;
`

// -- ELIMINAR REGISTRO POR ID --//
const Delete_stock = `DELETE FROM stock WHERE id = $1;`

// -- OBTENER MEJOR RECOMENDACION DE STOCKS --//
// Esta consulta obtiene las acciones con mejor potencial de ganancia y ordena por porcentaje de ganancia
// y fecha, limitando los resultados a los proximos 6 meses.
// Se filtran las acciones con rating_to positivo y se calcula el porcentaje de ganancia.
// Se ordena por porcentaje de ganancia y fecha, devolviendo la mejor recomendación.
const GetBestStockRecommendation = `SELECT stock.ticker,
       stock.target_from,
       stock.target_to,
       stock.company,
       stock.action,
       stock.brokerage,
       stock.rating_from,
       stock.rating_to,
       stock.TIME,
       stock.origin,
       stock.next_page,
       round((((target_to - target_from) / target_from)*100) , 2 ) AS profit_percentage,
       CONCAT(rating_from,' → ',rating_to) AS change_rating
FROM stock
WHERE stock.rating_to IN (
  'Buy', 'Strong Buy', 'Top Pick', 'Outperform', 'Overweight', 'Market Outperform',
  'Mkt Outperform', 'Sector Outperform', 'Positive', 'Speculative Buy', 'Action List Buy', 'Moderate Buy'
)
AND stock.time >= (NOW() - INTERVAL '6 months')
ORDER BY profit_percentage DESC, stock.time DESC`
