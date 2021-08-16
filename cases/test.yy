{
 function field(column_prefix)
    print(column_prefix)
 end

}
query:
     {column_prefix1="t1.";reset_value1="t1.";column_prefix2="t2.";reset_value2="t2."} select expression from _table as t1 left join _table as t2 on expression
     |  select agg_expression ,"debug" from _table as t1 where t1. _field_char > any(select t2. _field_char from _table as t2 where expression)

#select_list:
#   _field AS { num=num+1; field(num) } |
#   count(*) AS { num=num+1; field(num) } , select_list

column_str:
    { print(column_prefix1)} _field_char
    | { print(column_prefix2) }  _field_char

agg_expression: {column_prefix1="t1.";column_prefix2="";}
    count(expression2)

expression:
    column_str > column_str

expression2:
    column_str > column_str
