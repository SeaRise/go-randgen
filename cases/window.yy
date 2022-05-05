/* test window funtion */

query:
    select

select:
    SELECT
           window_function_with_window_name,
           fieldA,
           fieldB
    FROM (
	SELECT _field AS fieldA, _field AS fieldB
	FROM _table
    ) as t
    WINDOW window_name AS (window_spec)
    | SELECT
           window_function,
           fieldA,
           fieldB
    FROM (
	SELECT _field AS fieldA, _field AS fieldB
	FROM _table
    ) as t

window_function:
           DENSE_RANK() OVER (window_spec) AS 'dense_rank'
           | RANK() OVER (window_spec) AS 'rank'
           | ROW_NUMBER() OVER (window_spec) AS 'row_number'

window_function_with_window_name:
                            DENSE_RANK() OVER window_name AS 'dense_rank'
                            | RANK() OVER window_name AS 'rank'
                            | ROW_NUMBER() OVER window_name AS 'row_number'

window_clause:
	window_name AS (window_spec)

window_name:
	w1

window_spec:
	partition_clause order_clause
	| partition_clause order_clause
	| partition_clause order_clause

partition_clause:
	PARTITION BY LOWER(fieldB)

order_clause:
	#ORDER BY LOWER(fieldA) order_clause_indication
	#| ORDER BY LOWER(fieldB) order_clause_indication
	ORDER BY LOWER(fieldA) order_clause_indication , LOWER(fieldB) order_clause_indication
	| ORDER BY LOWER(fieldB) order_clause_indication , LOWER(fieldA) order_clause_indication

order_clause_indication:
	ASC | DESC
