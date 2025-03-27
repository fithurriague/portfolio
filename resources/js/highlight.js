document.addEventListener("htmx:afterSwap", function () {
    const codeBlock = document.getElementById("code-block");
    codeBlock.style.tabSize = "4";
    codeBlock.removeAttribute("data-highlighted");
    hljs.highlightElement(codeBlock);
});
