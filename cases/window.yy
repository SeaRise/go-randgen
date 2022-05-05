/* test window funtion */

query:
    select

select:
    SELECT
           window_function,
           fieldA,
           fieldB
    FROM (
	SELECT _field AS fieldA, _field AS fieldB
	FROM _table
    ) as t
    WINDOW window_name AS (window_spec)

window_function:
           DENSE_RANK() over_clause AS 'dense_rank'
           | RANK() over_clause AS 'rank'
           | ROW_NUMBER() over_clause AS 'row_number'


window_clause:
	window_name AS (window_spec)

window_name:
	w1

window_spec:
	partition_clause order_clause
	| partition_clause order_clause
	| partition_clause order_clause
	| window_name partition_clause order_clause

partition_clause:
	PARTITION BY LOWER(fieldB)

order_clause:
	#ORDER BY LOWER(fieldA) order_clause_indication
	#| ORDER BY LOWER(fieldB) order_clause_indication
	ORDER BY LOWER(fieldA) order_clause_indication , LOWER(fieldB) order_clause_indication
	| ORDER BY LOWER(fieldB) order_clause_indication , LOWER(fieldA) order_clause_indication

order_clause_indication:
	ASC | DESC

over_clause:
	OVER (window_spec) | OVER window_name

