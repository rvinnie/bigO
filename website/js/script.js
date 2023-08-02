import { CodeJar } from "https://medv.io/codejar/codejar.js";

const highlight = (editor) => {
    // highlight.js does not trims old tags,
    // let's do it by this hack.
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
    // const themeLink = document.createElement("link");
    const styleLink = document.createElement("link");
    if (name == "light") {
        // themeLink.href = `https://cdnjs.cloudflare.com/ajax/libs/highlight.js/9.18.1/styles/atom-one-light.min.css`;
        styleLink.href = `css/light.css`
    } else {
        // themeLink.href = `https://cdnjs.cloudflare.com/ajax/libs/highlight.js/9.18.1/styles/atom-one-dark.min.css`;
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
        jar.updateCode(`import {CodeJar} from '@medv/codejar';
import Prism from 'prismjs';

const editor = document.querySelector('#editor');
const jar = new CodeJar(editor, Prism.highlightElement, {tab: '\\t'});

// Update code
jar.updateCode('let foo = bar');

// Get code
let code = jar.toString();

// Listen to updates
jar.onUpdate(code => {
  console.log(code);
});
`);
        jar.updateOptions({ tab: "  " });
    },
    function () {
        editor.className = "editor language-go";
        jar.updateCode(`package main

import (
\t"fmt"
\t"github.com/antonmedv/expr"
)

func main() {
\tfmt.Println("Hello, CodeJar")

\toutput, err := expr.Eval("1+2")
\tif err != nil {
\t\tpanic(err)
\t}
}
`);
        jar.updateOptions({ tab: "\t" });
    },
    function () {
        editor.className = "editor language-ts";
        jar.updateCode(`interface Person {
    firstName: string;
    lastName: string;
}

function greeter(person: Person) {
    return "Hello, " + person.firstName + " " + person.lastName;
}

let user = {
    firstName: "Jane",
    lastName: "User"
};

document.body.textContent = greeter(user);`);
        jar.updateOptions({ tab: "    " });
    },
    function () {
        editor.className = "editor language-rust";
        jar.updateCode(`#[derive(Debug)]
struct Rectangle {
    width: u32,
    height: u32,
}

impl Rectangle {
    fn area(&self) -> u32 {
        self.width * self.height
    }
}

fn main() {
    let rect1 = Rectangle { width: 30, height: 50 };

    println!(
        "The area of the rectangle is {} square pixels.",
        rect1.area()
    );
}`);
        jar.updateOptions({ tab: "    " });
    },
    function () {
        editor.className = "editor language-kotlin";
        jar.updateCode(`suspend fun main() = coroutineScope {
    for (i in 0 until 10) {
        launch {
            delay(1000L - i * 10)
            print("❤️$i ")
        }
    }
}

val positiveNumbers = list.filter { it > 0 }

fun calculateTotal(obj: Any) {
    if (obj is Invoice)
        obj.calculateTotal()
}`);
        jar.updateOptions({ tab: "    " });
    }
];

languages[currentLanguage]();

const switchLanguageButton = document.querySelector(".switch-language");
switchLanguageButton.addEventListener("click", (event) => {
    event.preventDefault();
    currentLanguage = (currentLanguage + 1) % languages.length;
    languages[currentLanguage]();
    const [, name] = editor.className.match(/language-(\w+)/);
    switchLanguageButton.textContent = name;
});


// Notifications
function pushNotify(title) {
    let myNotify = new Notify({
        status: 'warning',
        title: title,
        effect: 'slide',
        autoclose: true,
        autotimeout: 3000,
        type: 3
    })
}

const codeInput = document.getElementById("codeInput");
const countComplexityBttn = document.getElementById("countComplexity");
const languageInput = document.getElementById("languageInput");
const shortDescriptionBox = document.getElementById("shortDescriptionBox");
const fullDescriptionBox = document.getElementById("fullDescriptionBox");


const url = 'http://localhost:8080'
const maxBodySize = 4096

async function countAlgorithmComplexity() {
    const uri = `${url}/api/complexity/count`

    const algorithmBodyJson = { code: codeInput.textContent, language: languageInput.textContent };

    if (algorithmBodyJson.algorithm.length > maxBodySize) {
        pushNotify(`Body is too large!`)
        return
    }

    shortDescriptionBox.textContent = ""
    fullDescriptionBox.textContent = ""

    const options = {
        method: 'POST',
        headers: {
            'Content-Type': codeInput.type,
        },
        body: JSON.stringify(algorithmBodyJson),
    };

    const response = await fetch(uri, options)
    const body = await response.json()


    switch (response.status) {
        case 200:
            shortDescriptionBox.textContent = body['shortDescriptionBox']
            fullDescriptionBox.textContent = body['fullDescriptionBox']
            break
        case 400:
            window.location.href = "https://previews.123rf.com/images/leremy/leremy1702/leremy170200035/71368169-website-error-400-bad-request.jpg"
            break
        case 404:
            window.location.href = "https://cdn.vectorstock.com/i/1000x1000/60/19/404-error-the-page-not-found-with-ghost-vector-20356019.webp";
            break
        case 500:
            window.location.href = "https://www.crazydomains.com.au/help/500-internal-server-error/?image_id=2981"
            break
    }
}

countComplexityBttn.onclick = countAlgorithmComplexity;