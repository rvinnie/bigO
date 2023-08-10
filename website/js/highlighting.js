import { CodeJar } from "https://medv.io/codejar/codejar.js";

const highlight = (editor) => {
    editor.textContent = editor.textContent;
    hljs.highlightElement(editor);
};

const editor = document.querySelector(".editor");
const jar = new CodeJar(editor, highlight);

let currentStyle = 0;
const styles = [
    "light",
    "dark",
].map((name) => {
    const styleLink = document.createElement("link");
    if (name == "light") {
        styleLink.href = `css/light.css`
    } else {
        styleLink.href = `css/dark.css`
    }
    styleLink.rel = "stylesheet";
    styleLink.disabled = "true";
    document.head.appendChild(styleLink);
    return styleLink;
});

styles[currentStyle].removeAttribute("disabled");

const switchStyleButton = document.querySelector(".switch-style");

switchStyleButton.addEventListener("click", (event) => {
    event.preventDefault();

    styles[currentStyle].setAttribute("disabled", "true");
    currentStyle = (currentStyle + 1) % styles.length;
    styles[currentStyle].removeAttribute("disabled");

    if (currentStyle == 0) {
        name = "light"
    } else {
        name = "dark"
    }

    switchStyleButton.textContent = name;
});

let currentLanguage = 0;
const languages = [
    function () {
        editor.className = "editor language-js";
        jar.updateCode(
`function binarySearch(arr, val) {
\tlet start = 0;
\tlet end = arr.length - 1;

\twhile (start <= end) {
\t\tlet mid = Math.floor((start + end) / 2);

\t\tif (arr[mid] === val) {
\t\t\treturn mid;
\t\t}

\t\tif (val < arr[mid]) {
\t\t\tend = mid - 1;
\t\t} else {
\t\t\tstart = mid + 1;
\t\t}
\t}
\treturn -1;
}`
        );
        jar.updateOptions({ tab: "\t" });
    },
    function () {
        editor.className = "editor language-go";
        jar.updateCode(
`func binarySearch(needle int, haystack []int) bool {
\tlow := 0
\thigh := len(haystack) - 1

\tfor low <= high{
\t\tmedian := (low + high) / 2

\t\tif haystack[median] < needle {
\t\t\tlow = median + 1
\t\t}else{
\t\t\thigh = median - 1
\t\t}
\t}

\tif low == len(haystack) || haystack[low] != needle {
\t\treturn false
\t}

\treturn true
}`
        );
        jar.updateOptions({ tab: "\t" });
    },
    function () {
        editor.className = "editor language-python";
        jar.updateCode(
`def binary_search(arr, low, high, x):
\tif high >= low:
 
\t\tmid = (high + low) // 2
 
\t\tif arr[mid] == x:
\t\t\treturn mid
\t\telif arr[mid] > x:
\t\t\treturn binary_search(arr, low, mid - 1, x)
\t\telse:
\t\t\treturn binary_search(arr, mid + 1, high, x)
\telse:
\t\t\treturn -1`
        );
        jar.updateOptions({tab: "\t"});
    }
];

// Supported languages - https://highlightjs.readthedocs.io/en/latest/supported-languages.html

languages[currentLanguage]();

const switchLanguageButton = document.querySelector(".switch-language");
switchLanguageButton.addEventListener("click", (event) => {
    event.preventDefault();
    currentLanguage = (currentLanguage + 1) % languages.length;
    languages[currentLanguage]();
    const [, name] = editor.className.match(/language-(\w+)/);
    switchLanguageButton.textContent = name;
});

