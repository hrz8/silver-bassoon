-- name: GetOrders :many
SELECT * FROM orders WHERE 1 = 1;

-- name: GetCustomerOrders :many
SELECT
    o.order_name AS "order_name",
    cc.company_name AS "customer_company_name",
    c.name AS "customer_name",
    TO_CHAR(o.created_at AT TIME ZONE @time_zone::text, 'Mon DDth, HH:MI AM') AS "order_date",
    CASE WHEN POSITION('.' IN TO_CHAR(SUM(d.delivered_quantity), 'FM999999.99')) > 0
        THEN COALESCE('$' || TRIM(TRAILING '.' FROM TO_CHAR(SUM(d.delivered_quantity), 'FM999999.99')), '-')
        ELSE COALESCE('$' || TO_CHAR(SUM(d.delivered_quantity), 'FM999999'), '-')
    END AS "delivered_amount",
    CASE WHEN POSITION('.' IN TO_CHAR(SUM(oi.price_per_unit * oi.quantity), 'FM999999.99')) > 0
        THEN COALESCE('$' || TRIM(TRAILING '.' FROM TO_CHAR(SUM(oi.price_per_unit * oi.quantity), 'FM999999.99')), '-')
        ELSE COALESCE('$' || TO_CHAR(SUM(oi.price_per_unit * oi.quantity), 'FM999999'), '-')
    END AS "total_amount"
FROM
    orders o
JOIN customers c ON o.customer_id = c.user_id
JOIN customer_companies cc ON c.company_id = cc.company_id
LEFT JOIN order_items oi ON o.id = oi.order_id
LEFT JOIN deliveries d ON oi.id = d.order_item_id
WHERE
    1=1 AND
    (
        CASE WHEN @is_search_term::bool THEN
            o.order_name ILIKE @search_term OR
            oi.product ILIKE @search_term
        ELSE
            TRUE
        END
    )
    AND (
        CASE WHEN @using_date_filter::bool THEN
            o.created_at AT TIME ZONE @time_zone::text >= @start_date AND
            o.created_at AT TIME ZONE @time_zone::text <= @end_date
        ELSE
            TRUE
        END
    )
GROUP BY
    o.order_name,
    cc.company_name,
    c.name,
    o.created_at
ORDER BY
    o.created_at DESC
LIMIT CASE WHEN @using_pagination::bool THEN @page_size::int END
OFFSET CASE WHEN @using_pagination::bool THEN (@page_number - 1) * @page_size::int END
;