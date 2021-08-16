tables = {
    rows = { 1, 10, 20, 25, 50, 75, 100 },
    charsets = { 'utf8mb4 collate=utf8mb4_general_ci', 'utf8mb4 collate = utf8mb4_bin', 'latin1', 'binary' },
    partitions = { 'undef' },
};

fields = {
    -- The four BLOB types are TINYBLOB, BLOB, MEDIUMBLOB, and LONGBLOB.
    -- The four TEXT types are TINYTEXT, TEXT, MEDIUMTEXT, and LONGTEXT.
    types = { 'int',
              'char(20)', 'varchar(20)', 'binary(20)', 'varbinary(20)',
        --              'enum', 'set',
        -- todo key blob and text prefix
        --      'tinyblob','blob','mediumblob','longblob',
        --      'tinytext','text','mediumtext','longtext',
    },
    sign = { 'signed', 'unsigned' },
    keys = { 'key', 'undef' }
}

data = {
    --   enum = { '"y"', '"b"', '"Abc"', '"null"' },
    numbers = { 'null', 'tinyint', '0', '-1', '1', '10', '20' },
    strings = { 'null', 'english', 'english', 'english', 'english', 'english', 'english', 'english' },
}