CREATE DATABASE IF NOT EXISTS {{.Database}};

CREATE TABLE IF NOT EXISTS {{.Database}}.{{.Table}}
(

    {{.Fields}}

    );