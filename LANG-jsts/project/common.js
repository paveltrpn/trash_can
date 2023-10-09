class cat_c {
    constructor() {
        this.mew = "mew mew mew!";
    }
    say_mew(times) {
        let i = 0;
        for (i = 0; i < times; i++) {
            console.log("cat_c::say_mew() at " + i + "-th time!");
        }
    }
}
export { cat_c };
