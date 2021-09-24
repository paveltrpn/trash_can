var top_source = `
<style>
    .top_grid {
        display: grid;
        grid-template-rows: 1fr;
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
        font-size: 32px;
        letter-spacing: 0.4em;
        text-decoration: none;
        color: white;
        border: 2px solid white;
        text-align: center;
        padding: 15px 15px;
    }

    .top_grid > div > a:hover { 
        color: red; /* Цвет ссылки */
        border: 2px solid red;
    }

</style>

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
        <a href="https://habr.com/ru/post/425837/">
            ARTICLES
        </a>
    </div>
</div>
`

function main() {
    var elem = document.getElementById("top");
    elem.innerHTML = top_source;
}

main();
