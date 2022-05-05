-- 表相关定义
tables = {
    -- 生成的表的记录数
    rows = {10, 20, 30, 90},
    -- 表的字符编码
    charsets = {'utf8', 'latin1', 'binary'},
    -- 表的分区数, 'undef' 表示不分区
    partitions = {'undef'},
}

-- 字段相关定义
fields = {
    -- 需要测试的数据类型
    types = {'bigint', 'float', 'double', 'decimal(40, 20)',
        'char(20)', 'varchar(20)'},
    -- 所有的上面的数字类型都要测试带符合和不带符号两种
    sign = {'signed', 'unsigned'}
}

-- 数据初始化相关定义
data = {
    -- 数字字段的生成方案
    numbers = {'null', 'tinyint', 'smallint',
        '12.991', '1.009', '-9.183',
        'decimal',
    },
    -- 字符串字段的生成方案
    strings = {'null', 'letter', 'english'},
}
