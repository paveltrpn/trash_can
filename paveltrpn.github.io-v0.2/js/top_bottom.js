var top_style = `
<style>
    .top_grid {
        display: grid;
        grid-template-rows: 1fr 1fr;
        grid-template-columns: 1fr 1fr 1fr;
        grid-gap: 32px;
    }

    .top_grid > div {
        display: flex;
        flex-direction: column;
        justify-content: center;
        color: white;
    }

    .top_grid > div > a {
        font-family: Roboto Mono, monospace;
        font-weight: bold;
        font-size: 40px;
        letter-spacing: 0.1em;
        text-decoration: none;
        color: white;
        border: 2px solid white;
        text-align: center;
        padding: 25px 5px;
    }

    .top_grid > div > a:hover { 
        color: red; /* Цвет ссылки */
        border: 2px solid red;
    }

</style>
`

var top_source = `
<div class="top_grid">
    <div>
        <a href="index.html">
            LOG
        </a>
    </div>

    <div>
        <a href="links.html">
            LINKS
        </a>
    </div>

    <div>
        <a href="articles.html">
            ARTICLES
        </a>
    </div>

    <div>
        <a href="javascript:void(0)">
        &#9746;
        </a>
    </div>

    <div>
        <a href="javascript:void(0)">
        &#9746;
        </a>
    </div>

    <div>
        <a href="javascript:void(0)">
        &#9746;
        </a>
    </div>
</div>
`

var top_article_source = `
<div class="top_grid">
    <div>
        <a href="../../index.html">
            LOG
        </a>
    </div>

    <div>
        <a href="../../links.html">
            LINKS
        </a>
    </div>

    <div>
        <a href="../../articles.html">
            ARTICLES
        </a>
    </div>

    <div>
        <a href="javascript:void(0)">
        &#9746;
        </a>
    </div>

    <div>
        <a href="javascript:void(0)">
        &#9746;
        </a>
    </div>

    <div>
        <a href="javascript:void(0)">
        &#9746;
        </a>
    </div>
</div>
`

function main() {
    let elem = document.getElementById("top");

    if (elem != null) {
        elem.insertAdjacentHTML("beforeend", top_style);
        elem.insertAdjacentHTML("beforeend", top_source);
    }

    elem = document.getElementById("top_article");

    if (elem != null) {
        elem.insertAdjacentHTML("beforeend", top_style);
        elem.insertAdjacentHTML("beforeend", top_article_source);
    }
}

main();
