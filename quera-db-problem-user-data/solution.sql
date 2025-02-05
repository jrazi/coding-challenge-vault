-- Section1
SELECT 
    id,
    first_name,
    last_name,
    username
FROM "user"
WHERE 
    lower(first_name) LIKE 's%' 
    AND lower(last_name) LIKE '%e'
ORDER BY id;

-- Section2
SELECT 
    c.name AS company_name,
    array_agg(p.name ORDER BY p.name) AS products
FROM company c
JOIN product_company pc ON c.id = pc.company_id
JOIN product p ON pc.product_id = p.id
GROUP BY c.id
ORDER BY 
    COUNT(*) DESC,
    c.id;

-- Section3
SELECT 
    u.username,
    CONCAT(
        COALESCE(ca.country, ua.country, 'N/A'), ', ',
        COALESCE(ca.city, ua.city, 'N/A'), ', ',
        COALESCE(ca.zip_code::text, ua.zip_code::text, 'N/A')
    ) AS full_address
FROM "user" u
LEFT JOIN company comp ON u.company_id = comp.id
LEFT JOIN address ca ON comp.address_id = ca.id
LEFT JOIN address ua ON u.address_id = ua.id
ORDER BY u.id;

-- Section4
WITH marketplace_stats AS (
    SELECT AVG(price) AS avg_price FROM product
)
SELECT 
    p.name AS product,
    p.price,
    p.category,
    c.name AS manufacturer,
    a.city AS production_city,
    CASE 
        WHEN p.price > (SELECT avg_price FROM marketplace_stats) 
        THEN 'above_average'
        ELSE 'average_or_below' 
    END AS price_category
FROM product p
JOIN product_company pc ON p.id = pc.product_id
JOIN company c ON pc.company_id = c.id
JOIN address a ON c.address_id = a.id
ORDER BY p.id;