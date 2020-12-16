
const hdr_source: any = [
'                <ul>',
'                    <li><a href="">Статьи &emsp; &#9660;</a>',
'                        <ul class="submenu">',
'                            <li><a href="algorithm.html">Алгоритмы</a></li>',
'                            <li><a href="technology.html">Технология</a></li>',
'                            <li><a href="nix.html">Администрирование</a></li>',
'                            <li><a href="cpp.html">C / C++</a></li>',
'                            <li><a href="golang.html">GOLANG</a></li>',
'                            <li><a href="web.html">HTML + CSS + JAVASCRIPT / TYPESCRIPT</a></li>',
'                        </ul>',
'                    </li>',
'                    <li><a href="">Литература &emsp; &#9660;</a>',
'                        <ul class="submenu">',
'                            <li><a href="">Математика</a></li>',
'                            <li><a href="">Технология</a></li>',
'                            <li><a href="">Программирование</a></li>',
'                            <li><a href="">Статьи</a></li>',
'                        </ul>',
'                    </li>',
'                    <li><a href="">Разное &emsp; &#9660;</a>',
'                        <ul class="submenu">',
'                            <li><a href="blog/blog.html">Blog</a></li>',
'                            <li><a href="music.html">Музыка</a></li>',
'                            <li><a href="links.html">Ссылки</a></li>',
'                        </ul>',
'                    </li>',
'                </ul>', 
].join("\n");

function main() {
    let elem: HTMLElement = document.getElementById("header_menu");
    elem.innerHTML = hdr_source;
}

main()
