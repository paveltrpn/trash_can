"use strict";
const express = require('express');
const app = express();
const host = 'localhost';
const port = 8080;
app.use((req, res, next) => {
    res.header('Access-Control-Allow-Origin', '*');
    res.header('Access-Control-Allow-Headers', 'Origin, X-Requested-With, Content-Type, Accept');
    next();
    app.options('*', (req, res) => {
        res.header('Access-Control-Allow-Methods', 'GET, PATCH, PUT, POST, DELETE, OPTIONS');
        res.send();
    });
});
app.use(express.static(`${__dirname}`));
app.listen(port, host, function () {
    console.log(`Server listens http://${host}:${port}`);
});
