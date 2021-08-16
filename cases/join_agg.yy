# 希望能够尽量覆盖足够多的 sql select 模式
# 1. 简单查询
# select xxx from single_table where expression order by xxx limit xxx
# 2. join
# 常规 join，子查询，多层嵌套 join
# 3. aggregation
# 简单查询 / join aggregation group by on expression

# 一些问题
# 在多层嵌套的时候如何在 expression 中指定对应的表名
# select * from _table as t1 join _table as t2 on expression join _table t3 where expression

# 使用非关键字的时候，像调用函数一样可以加上参数传递 -- 支持该功能
# select * from _table as t1 join _table as t2 on expression(t1,t2) join _table t3 where expression(t2,t3)


expression:{a = 1}
    CREATE TABLE
    {print(string.format("t%d", a))} (a INT)



expression:(tx,ty)
    tx. _field > ty. _field
