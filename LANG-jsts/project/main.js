import { cat_c } from "./common.js";
import { divider, concat_str } from "./utils.js";
function main() {
    let cat = new cat_c();
    cat.say_mew(5);
    console.log("");
    let c = divider(19, 5);
    console.log("divider(19,5) say - " + c);
    let sum = concat_str("left ", "right");
    console.log("concat_str() say - " + sum);
}
main();
