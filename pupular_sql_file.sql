select * from stock_headers sh ;
select * from "StockFinancialStatements" sfs ;
select * from "StockSimllarAssets" ssa ;
select * from "StockStats" ss ;
select * from "StockShareHoldingPatterns" sshp ;
select * from "StockPriceDatas" spd ;
select * from "StockNews" sn ;

--  PROFIT PER YEAR
SELECT sfs.stock_header_id,
       sh.search_id,
       sh.display_name, 
       (profit_yearly::jsonb)->>'2024' AS _2024, 
       (profit_yearly::jsonb)->>'2023' AS _2023, 
       (profit_yearly::jsonb)->>'2022' AS _2022, 
       (profit_yearly::jsonb)->>'2021' AS _2021, 
       (profit_yearly::jsonb)->>'2020' AS _2020  
FROM stock_financial_statements sfs
JOIN stock_headers sh 
ON sh.id = sfs.stock_header_id 
WHERE (profit_yearly::jsonb)->>'2024' >= (profit_yearly::jsonb)->>'2023' 
AND (profit_yearly::jsonb)->>'2023' >= (profit_yearly::jsonb)->>'2022'
AND (profit_yearly::jsonb)->>'2022' >= (profit_yearly::jsonb)->>'2021'
AND (profit_yearly::jsonb)->>'2021' >= (profit_yearly::jsonb)->>'2020';



-- PROFIT INCREASED BY PERCENTAGE
SELECT sfs.stock_header_id,
       sh.search_id,
       sh.display_name, 
       (profit_yearly::jsonb)->>'2024' AS _2024, 
       (profit_yearly::jsonb)->>'2023' AS _2023, 
       (profit_yearly::jsonb)->>'2022' AS _2022, 
       (profit_yearly::jsonb)->>'2021' AS _2021, 
       (profit_yearly::jsonb)->>'2020' AS _2020,
       CASE 
           WHEN (profit_yearly::jsonb)->>'2023' IS NULL 
                OR (profit_yearly::jsonb)->>'2023' = '0' 
           THEN NULL 
           ELSE (( (profit_yearly::jsonb)->>'2024')::numeric - ( (profit_yearly::jsonb)->>'2023')::numeric ) 
                / ( (profit_yearly::jsonb)->>'2023')::numeric * 100
       END AS pct_inc_2023_2024,
       CASE 
           WHEN (profit_yearly::jsonb)->>'2022' IS NULL 
                OR (profit_yearly::jsonb)->>'2022' = '0' 
           THEN NULL 
           ELSE (( (profit_yearly::jsonb)->>'2023')::numeric - ( (profit_yearly::jsonb)->>'2022')::numeric ) 
                / ( (profit_yearly::jsonb)->>'2022')::numeric * 100
       END AS pct_inc_2022_2023,
       CASE 
           WHEN (profit_yearly::jsonb)->>'2021' IS NULL 
                OR (profit_yearly::jsonb)->>'2021' = '0' 
           THEN NULL 
           ELSE (( (profit_yearly::jsonb)->>'2022')::numeric - ( (profit_yearly::jsonb)->>'2021')::numeric ) 
                / ( (profit_yearly::jsonb)->>'2021')::numeric * 100
       END AS pct_inc_2021_2022,
       CASE 
           WHEN (profit_yearly::jsonb)->>'2020' IS NULL 
                OR (profit_yearly::jsonb)->>'2020' = '0' 
           THEN NULL 
           ELSE (( (profit_yearly::jsonb)->>'2021')::numeric - ( (profit_yearly::jsonb)->>'2020')::numeric ) 
                / ( (profit_yearly::jsonb)->>'2020')::numeric * 100
       END AS pct_inc_2020_2021
FROM stock_financial_statements sfs
JOIN stock_headers sh 
ON sh.id = sfs.stock_header_id 
WHERE (profit_yearly::jsonb)->>'2024' >= (profit_yearly::jsonb)->>'2023' 
AND (profit_yearly::jsonb)->>'2023' >= (profit_yearly::jsonb)->>'2022'
AND (profit_yearly::jsonb)->>'2022' >= (profit_yearly::jsonb)->>'2021'
AND (profit_yearly::jsonb)->>'2021' >= (profit_yearly::jsonb)->>'2020';

-- PROFIT INCREASED BY PERCENTAGE ONLY POSITIVE YEAR WISE
SELECT 
    sfs.stock_header_id,
    sh.search_id,
    sh.display_name, 
    (profit_yearly::jsonb)->>'2024' AS profit_2024, 
    (profit_yearly::jsonb)->>'2023' AS profit_2023, 
    (profit_yearly::jsonb)->>'2022' AS profit_2022, 
    (profit_yearly::jsonb)->>'2021' AS profit_2021, 
    (profit_yearly::jsonb)->>'2020' AS profit_2020,
    CASE 
        WHEN ((profit_yearly::jsonb)->>'2023') IS NULL 
            OR ((profit_yearly::jsonb)->>'2023')::numeric = 0 
        THEN NULL 
        ELSE (((profit_yearly::jsonb)->>'2024')::numeric - ((profit_yearly::jsonb)->>'2023')::numeric) 
            / ((profit_yearly::jsonb)->>'2023')::numeric * 100
    END AS pct_inc_2023_2024,
    CASE 
        WHEN ((profit_yearly::jsonb)->>'2022') IS NULL 
            OR ((profit_yearly::jsonb)->>'2022')::numeric = 0 
        THEN NULL 
        ELSE (((profit_yearly::jsonb)->>'2023')::numeric - ((profit_yearly::jsonb)->>'2022')::numeric) 
            / ((profit_yearly::jsonb)->>'2022')::numeric * 100
    END AS pct_inc_2022_2023,
    CASE 
        WHEN ((profit_yearly::jsonb)->>'2021') IS NULL 
            OR ((profit_yearly::jsonb)->>'2021')::numeric = 0 
        THEN NULL 
        ELSE (((profit_yearly::jsonb)->>'2022')::numeric - ((profit_yearly::jsonb)->>'2021')::numeric) 
            / ((profit_yearly::jsonb)->>'2021')::numeric * 100
    END AS pct_inc_2021_2022,
    CASE 
        WHEN ((profit_yearly::jsonb)->>'2020') IS NULL 
            OR ((profit_yearly::jsonb)->>'2020')::numeric = 0 
        THEN NULL 
        ELSE (((profit_yearly::jsonb)->>'2021')::numeric - ((profit_yearly::jsonb)->>'2020')::numeric) 
            / ((profit_yearly::jsonb)->>'2020')::numeric * 100
    END AS pct_inc_2020_2021
FROM stock_financial_statements sfs
JOIN stock_headers sh 
ON sh.id = sfs.stock_header_id 
WHERE 
    (
        ((profit_yearly::jsonb)->>'2023') IS NOT NULL 
        AND ((profit_yearly::jsonb)->>'2023')::numeric > 0 
        AND ((profit_yearly::jsonb)->>'2022') IS NOT NULL  
        AND ((profit_yearly::jsonb)->>'2022')::numeric > 0 
        AND (((profit_yearly::jsonb)->>'2024')::numeric - ((profit_yearly::jsonb)->>'2023')::numeric) 
            / ((profit_yearly::jsonb)->>'2023')::numeric * 100 > 
        	(((profit_yearly::jsonb)->>'2023')::numeric - ((profit_yearly::jsonb)->>'2022')::numeric) 
            / ((profit_yearly::jsonb)->>'2022')::numeric * 100
    )
    AND 
    (
        ((profit_yearly::jsonb)->>'2022') IS NOT NULL 
        AND ((profit_yearly::jsonb)->>'2022')::numeric > 0 
        AND ((profit_yearly::jsonb)->>'2021') IS NOT NULL  
        AND ((profit_yearly::jsonb)->>'2021')::numeric > 0 
        AND (((profit_yearly::jsonb)->>'2023')::numeric - ((profit_yearly::jsonb)->>'2022')::numeric) 
            / ((profit_yearly::jsonb)->>'2022')::numeric * 100 > 
        	(((profit_yearly::jsonb)->>'2022')::numeric - ((profit_yearly::jsonb)->>'2021')::numeric) 
            / ((profit_yearly::jsonb)->>'2021')::numeric * 100
    )
    AND 
    (
        ((profit_yearly::jsonb)->>'2021') IS NOT NULL 
        AND ((profit_yearly::jsonb)->>'2021')::numeric > 0 
        AND ((profit_yearly::jsonb)->>'2020') IS NOT NULL  
        AND ((profit_yearly::jsonb)->>'2020')::numeric > 0 
        AND (((profit_yearly::jsonb)->>'2022')::numeric - ((profit_yearly::jsonb)->>'2021')::numeric) 
            / ((profit_yearly::jsonb)->>'2021')::numeric * 100 > 
        	(((profit_yearly::jsonb)->>'2021')::numeric - ((profit_yearly::jsonb)->>'2020')::numeric) 
            / ((profit_yearly::jsonb)->>'2020')::numeric * 100
    );


-- PROFIT INCREASED BY PERCENTAGE ONLY POSITIVE QUARTER WISE
SELECT 
    sfs.stock_header_id,
    sh.search_id,
    sh.display_name, 
    (profit_quarterly::jsonb)->>'Dec 24' AS dec_24, 
    (profit_quarterly::jsonb)->>'Sep 24' AS sep_24, 
    (profit_quarterly::jsonb)->>'Jun 24' AS jun_24, 
    (profit_quarterly::jsonb)->>'Mar 24' AS mar_24, 
    (profit_quarterly::jsonb)->>'Dec 23' AS dec_23,
    CASE 
        WHEN ((profit_quarterly::jsonb)->>'Dec 23') IS NULL 
            OR ((profit_quarterly::jsonb)->>'Dec 23')::numeric = 0 
        THEN NULL 
        ELSE (((profit_quarterly::jsonb)->>'Dec 24')::numeric - ((profit_quarterly::jsonb)->>'Dec 23')::numeric) 
            / ((profit_quarterly::jsonb)->>'Dec 23')::numeric * 100
    END AS pct_inc_dec23_dec24,
    CASE 
        WHEN ((profit_quarterly::jsonb)->>'Sep 24') IS NULL 
            OR ((profit_quarterly::jsonb)->>'Sep 24')::numeric = 0 
        THEN NULL 
        ELSE (((profit_quarterly::jsonb)->>'Dec 24')::numeric - ((profit_quarterly::jsonb)->>'Sep 24')::numeric) 
            / ((profit_quarterly::jsonb)->>'Sep 24')::numeric * 100
    END AS pct_inc_sep24_dec24,
    CASE 
        WHEN ((profit_quarterly::jsonb)->>'Jun 24') IS NULL 
            OR ((profit_quarterly::jsonb)->>'Jun 24')::numeric = 0 
        THEN NULL 
        ELSE (((profit_quarterly::jsonb)->>'Sep 24')::numeric - ((profit_quarterly::jsonb)->>'Jun 24')::numeric) 
            / ((profit_quarterly::jsonb)->>'Jun 24')::numeric * 100
    END AS pct_inc_jun24_sep24,
    CASE 
        WHEN ((profit_quarterly::jsonb)->>'Mar 24') IS NULL 
            OR ((profit_quarterly::jsonb)->>'Mar 24')::numeric = 0 
        THEN NULL 
        ELSE (((profit_quarterly::jsonb)->>'Jun 24')::numeric - ((profit_quarterly::jsonb)->>'Mar 24')::numeric) 
            / ((profit_quarterly::jsonb)->>'Mar 24')::numeric * 100
    END AS pct_inc_mar24_jun24
FROM stock_financial_statements sfs
JOIN stock_headers sh 
ON sh.id = sfs.stock_header_id 
WHERE 
    (
        ((profit_quarterly::jsonb)->>'Dec 23') IS NOT NULL 
        AND ((profit_quarterly::jsonb)->>'Dec 23')::numeric > 0 
        AND ((profit_quarterly::jsonb)->>'Sep 24') IS NOT NULL  
        AND ((profit_quarterly::jsonb)->>'Sep 24')::numeric > 0 
        AND (((profit_quarterly::jsonb)->>'Dec 24')::numeric - ((profit_quarterly::jsonb)->>'Dec 23')::numeric) 
            / ((profit_quarterly::jsonb)->>'Dec 23')::numeric * 100 > 
        	(((profit_quarterly::jsonb)->>'Sep 24')::numeric - ((profit_quarterly::jsonb)->>'Jun 24')::numeric) 
            / ((profit_quarterly::jsonb)->>'Jun 24')::numeric * 100
    )
    AND 
    (
        ((profit_quarterly::jsonb)->>'Sep 24') IS NOT NULL 
        AND ((profit_quarterly::jsonb)->>'Sep 24')::numeric > 0 
        AND ((profit_quarterly::jsonb)->>'Jun 24') IS NOT NULL  
        AND ((profit_quarterly::jsonb)->>'Jun 24')::numeric > 0 
        AND (((profit_quarterly::jsonb)->>'Sep 24')::numeric - ((profit_quarterly::jsonb)->>'Jun 24')::numeric) 
            / ((profit_quarterly::jsonb)->>'Jun 24')::numeric * 100 > 
        	(((profit_quarterly::jsonb)->>'Jun 24')::numeric - ((profit_quarterly::jsonb)->>'Mar 24')::numeric) 
            / ((profit_quarterly::jsonb)->>'Mar 24')::numeric * 100
    )
    AND 
    (
        ((profit_quarterly::jsonb)->>'Jun 24') IS NOT NULL 
        AND ((profit_quarterly::jsonb)->>'Jun 24')::numeric > 0 
        AND ((profit_quarterly::jsonb)->>'Mar 24') IS NOT NULL  
        AND ((profit_quarterly::jsonb)->>'Mar 24')::numeric > 0 
        AND (((profit_quarterly::jsonb)->>'Jun 24')::numeric - ((profit_quarterly::jsonb)->>'Mar 24')::numeric) 
            / ((profit_quarterly::jsonb)->>'Mar 24')::numeric * 100 > 
        	(((profit_quarterly::jsonb)->>'Mar 24')::numeric - ((profit_quarterly::jsonb)->>'Dec 23')::numeric) 
            / ((profit_quarterly::jsonb)->>'Dec 23')::numeric * 100
    );

   
   
-- Foreign percentage increased with first and last partation of period   
WITH cta AS (
    SELECT 
        stock_header_id,
        period,
        FIRST_VALUE(foreign_institutions) OVER (
            PARTITION BY stock_header_id ORDER BY id 
            ROWS BETWEEN UNBOUNDED PRECEDING AND UNBOUNDED FOLLOWING
        ) AS first_fii,
        LAST_VALUE(foreign_institutions) OVER (
            PARTITION BY stock_header_id ORDER BY id
            ROWS BETWEEN UNBOUNDED PRECEDING AND UNBOUNDED FOLLOWING
        ) AS last_fii,
        LAST_VALUE(foreign_institutions) OVER (
            PARTITION BY stock_header_id ORDER BY id
            ROWS BETWEEN UNBOUNDED PRECEDING AND UNBOUNDED FOLLOWING
        ) - FIRST_VALUE(foreign_institutions) OVER (
            PARTITION BY stock_header_id ORDER BY id
            ROWS BETWEEN UNBOUNDED PRECEDING AND UNBOUNDED FOLLOWING
        ) AS difference,
        ARRAY_AGG(foreign_institutions) OVER (
            PARTITION BY stock_header_id ORDER BY id
        ) AS fii_trend
    FROM stock_share_holding_patterns
    GROUP BY stock_header_id, id, foreign_institutions
    ORDER BY stock_header_id, id
)
SELECT 
    cta.*, 
    s.company_name, 
    s.search_id, 
    cta.difference 
FROM cta 
JOIN stocks s ON s.id = cta.stock_header_id
WHERE cta.period = 'Dec ''24'
ORDER BY difference DESC;

    
select id,stock_header_id ,period,foreign_institutions from stock_share_holding_patterns sshp ;
select stock_header_id ,count(id) from stock_share_holding_patterns sshp group by stock_header_id order by count(id) desc;

-- THIS IS THE MAIN QUERY TO GET THE DIFFERENCE IS STAKE HOLDER PER QUARTER RELEASE
WITH cta AS (
    SELECT 
        stock_header_id,
        period,
        NTH_VALUE(period, 2) OVER (
            PARTITION BY stock_header_id ORDER BY id desc
            ROWS BETWEEN UNBOUNDED PRECEDING AND UNBOUNDED FOLLOWING
        ) AS period1,
        NTH_VALUE(foreign_institutions,1) OVER (
            PARTITION BY stock_header_id ORDER BY id desc
            ROWS BETWEEN UNBOUNDED PRECEDING AND UNBOUNDED FOLLOWING
        ) AS first_fii,
        NTH_VALUE(foreign_institutions, 2) OVER (
            PARTITION BY stock_header_id ORDER BY id desc
            ROWS BETWEEN UNBOUNDED PRECEDING AND UNBOUNDED FOLLOWING
        ) AS second_fii,
        NTH_VALUE(foreign_institutions, 3) OVER (
            PARTITION BY stock_header_id ORDER BY id desc
            ROWS BETWEEN UNBOUNDED PRECEDING AND UNBOUNDED FOLLOWING
        ) AS third_fii,
        NTH_VALUE(foreign_institutions,4) OVER (
            PARTITION BY stock_header_id ORDER BY id desc
            ROWS BETWEEN UNBOUNDED PRECEDING AND UNBOUNDED FOLLOWING
        ) AS forth_fii,
        NTH_VALUE(foreign_institutions,5) OVER (
            PARTITION BY stock_header_id ORDER BY id desc
            ROWS BETWEEN UNBOUNDED PRECEDING AND UNBOUNDED FOLLOWING
        ) AS fifth_fii,
        ROW_NUMBER() OVER (PARTITION BY stock_header_id ORDER BY id DESC) AS rn
    FROM stock_share_holding_patterns
)
SELECT cta.*,sh.search_id ,first_fii-second_fii as diff,sh.display_name FROM cta 
join stock_headers sh on cta.stock_header_id = sh.id
WHERE rn = 1 and cta.stock_header_id <> 161 and cta.second_fii <> 0 and cta.stock_header_id not in (1775, 1874, 2632, 4080)
order by diff desc;

select * from 	 sshp join stock_headers sh on sshp.stock_header_id = sh.id  where search_id ilike '%icici%'


-- DELETE AND DROP QUERY
--DROP TABLE IF EXISTS 
--    "StockHeaders",
--    "StockFinancialStatements",
--    "StockSimllarAssets",
--    "StockStats",
--    "StockShareHoldingPatterns",
--    "StockPriceDatas",
--    "StockNews",
--    "Stocks",
--    "LivePriceDtos" CASCADE;
--
--DO $$ 
--DECLARE 
--    r RECORD;
--BEGIN
--    FOR r IN (SELECT tablename FROM pg_tables WHERE schemaname = 'public') 
--    LOOP
--        EXECUTE 'DROP TABLE IF EXISTS "' || r.tablename || '" CASCADE';
--    END LOOP;
--END $$;


select * from stock_news sn order by id desc;

SELECT stock_id,count(*),array_agg(concat(pub_date,title)) 
FROM public.stock_news sn
WHERE pub_date::date >= '2025-04-12'::date
group by stock_id
order by count(*) desc;