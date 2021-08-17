query:
     select col1 from tables group by col1;
     | select (((col1 + 1))) from tables group by col1;
     | select cast(col1 as DECIMAL) from tables group by col1;
     | select aggregation(col1) from tables group by col1;
     | select (select col1) aa, count(distinct col2) from tables group by aa;
     | select col1 from tables group by col1 order by col1;
     | select HEX(col1) from tables group by HEX(col1);
     | select length(col1) from tables group by col1;
     | select col1 from tables group by col1 having col1 compare_operator col1_?;
     | select col1 from tables group by col1 having aggregation(col1) compare_operator  col1_?;
     | select col1, GROUP_CONCAT(col2, col3 order by col2 desc) from tables group by col1;
     | select col1 from tables where col1 is null group by col1;
     | select col1 from tables group by col1 having col1 is null;
     | select col1, col2 from tables group by col1, col2;
     | select col1, col2 from tables group by 1, 2;
     | select col1, aggregation(col2) from tables group by col1;
     | select col1, col2, col3 from tables group by col1, col2, col3;
     | select aggregation(col1), col2, col3 from tables group by col2, col3;
     | select col1 from tables where col1 where_condition group by col1;
     | select aggregation(col1) from tables where col1 where_condition group by col1;
     | select col1, col2 from tables where col1 where_condition  group by col1, col2;
     | select col1, aggregation(col2) from tables where col1 where_condition group by col1;
     | select max(col1), min(col1), sum(col1), count(col1), avg(col1), max(col2), min(col2), sum(col2), count(col2), avg(col2), count(1), count(*) from tables where col1 where_condition group by col1, col2;
     | SELECT * FROM tables WHERE COL1 IN (SELECT COL1 FROM tables GROUP BY COL1 HAVING COUNT(*)<>1);
     | select join_field from tables as t1 cjoin tables as t2 on t1.col1 join_compare t2.col1 group by t1.col1, t2.col1;
     | select hint_join_field from tables as t1 cjoin tables as t2 on t1.col1 join_compare t2.col1 group by 1, 2;
     | select hint_join_field from tables as t1 cjoin tables as t2 on t1.col1 join_compare t2.col1 group by t1.col1, t2.col1;
     | select join_field from tables as t1 join tables as t2 on t1.col1 join_cpmpare t2.col1 group by t1.col1, t2.col1;
     | select join_field from tables as t1 cjoin tables as t2 on t1.col1 join_compare t2.col1 group by t1.col1, t2.col1 having t1.col1 where_condition;
     | select join_field from tables as t1 cjoin tables as t2 on t1.col1 join_compare t2.col1 group by t1.col1, t2.col1 having t1.col1 join_compare t2.col1;
     | select join_field from tables as t1 cjoin tables as t2 on t1.col1 join_compare t2.col1 group by t1.col1, t2.col1 having aggregation(t1.col1) join_compare col1_?;
     | select hint_join_field from tables as t1 cjoin tables as t2 on t1.col1 join_compare t2.col1 group by t1.col1, t2.col1 having aggregation(t1.col1) join_compare col1_?;
     | select join_field from tables as t1 cjoin tables as t2 on t1.col1 join_compare t2.col1 where t1.col1 where_condition group by t1.col1, t2.col1;
     | select join_field from tables as t1 cjoin tables as t2 on t1.col1 join_compare t2.col1 where t1.col1 where_condition group by t1.col1, t2.col1 having t1.col1 where_condition;
     | select join_field from tables as t1 cjoin tables as t2 on t1.col1 join_compare t2.col1 where t1.col1 where_condition group by t1.col1, t2.col1  having t1.col1 join_compare t2.col1;
     | select join_field from tables as t1 cjoin tables as t2 on t1.col1 join_compare t2.col1 where t1.col1 where_condition group by t1.col1, t2.col1 having aggregation(t1.col1) join_compare col1_?;
     | select join_field from tables as t1 cjoin tables as t2 on t1.col1 != t2.col1 group by t1.col1, t2.col1;
     | select join_field from tables as t1 cjoin tables as t2 on t1.col1 = t2.col1 where t1.col1 >= t2.col1 group by t1.col1, t2.col1;
     | select join_field from tables as t1 cjoin tables as t2 on t1.col1 = t2.col1 where t1.col1 <= t2.col1 group by t1.col1, t2.col1;
     | select year(col1), aggregation(col2) from tables group by year(col1);
     | select group_concat(col1,col2,col3 order by col1) from tables group by col1;
     | select group_concat(col1,col2,col3 order by col1 desc) from tables group by col1;

where_condition:
     compare_operator col1_?
     | cbetween col1_? and col1_?
     | cin (col1_?, col1_?, col1_?)
     | is not null

cbetween:
     between
     | not between

cin:
    in
    | not in

arithmetic_operators:
     +
     | -
     | *
     | /
     | %
     | MOD
     | DIV

compare_operator:
     >
     | <
     | =
     | >=
     | <=
     | !=
     | <>
     | <=>

cjoin:
     inner join
     | right join
     | left join

join_field:
     t1.col1, t2.col1
     | t1.col1, aggregation(t2.col1)
     | max(col1), min(col1), sum(col1), count(col1), avg(col1), max(col2), min(col2), sum(col2), count(col2), avg(col2), count(1), count(*)

hint_join_field:
     | /*HASH_JOIN*/ t1.col1, aggregation(t2.col1)
     | /*MERGE_JOIN*/ t1.col1, aggregation(t2.col1)
     | /*INL_JOIN*/ t1.col1, aggregation(t2.col1)
     | /*INL_MERGE_JOIN*/ t1.col1, aggregation(t2.col1)
     | /*INL_HASH_JOIN*/ t1.col1, aggregation(t2.col1)

join_compare:
     =

aggregation:
     max
     | min
     | avg
     | count
     | sum