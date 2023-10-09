"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.testfreadSync = void 0;
const fs = require("fs");
const smnp = require("./string_mnp");
function testfreadSync() {
    console.log("main(): fs test.");
    let fileContent = fs.readFileSync("../assets/raven.txt", "utf8");
    let words = smnp.splitBySpace(fileContent);
    for (let i = 0; i < 5; i++) {
        console.log(words[i]);
    }
    console.log(smnp.removePunctuation(fileContent));
}
exports.testfreadSync = testfreadSync;
(() => {
    testfreadSync();
})();
