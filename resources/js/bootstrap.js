import hljs from "highlight.js/lib/core";
import go from "highlight.js/lib/languages/go";
import htmx from "htmx.org";

hljs.registerLanguage("go", go);

window.hljs = hljs;
window.htmx = htmx;
