const express = require('express');
const app = express();

const host = 'localhost';
const port = 8080;

app.use((req: any, res: any, next: any) => {
    res.header('Access-Control-Allow-Origin', '*');

    // authorized headers for preflight requests
    // https://developer.mozilla.org/en-US/docs/Glossary/preflight_request
    res.header('Access-Control-Allow-Headers', 'Origin, X-Requested-With, Content-Type, Accept');
    next();

    app.options('*', (req: any, res: any) => {
        // allowed XHR methods  
        res.header('Access-Control-Allow-Methods', 'GET, PATCH, PUT, POST, DELETE, OPTIONS');
        res.send();
    });
});

app.use(express.static(`${__dirname}`));

app.listen(port, host, function () {
    console.log(`Server listens http://${host}:${port}`);
});