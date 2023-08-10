const url = 'http://localhost:8080'
const maxBodySize = 4096

const descriptionBody = document.getElementById("descriptionBody")
const shortDescriptionBox = document.getElementById("shortDescriptionBox");
const fullDescriptionBox = document.getElementById("fullDescriptionBox");
const codeInput = document.getElementById("codeInput");
const countComplexityBttn = document.getElementById("countComplexity");
const languageInput = document.getElementById("languageInput");

async function countAlgorithmComplexity() {
    const uri = `${url}/api/complexity/count`

    countComplexityBttn.style.opacity = '0.6'
    countComplexityBttn.style.cursor = 'wait'
    countComplexityBttn.disabled = true;
    descriptionBody.style.visibility = 'hidden'

    const algorithmBodyJson = { code: codeInput.textContent, language: languageInput.textContent };

    if (algorithmBodyJson.code.length > maxBodySize) {
        pushNotify(`Body is too large!`)
        return
    } else if (algorithmBodyJson.code.length == 0) {
        pushNotify(`Body is empty!`)
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

    countComplexityBttn.style.opacity = '1'
    countComplexityBttn.style.cursor = ''
    countComplexityBttn.disabled = false;


    switch (response.status) {
        case 200:
            const body = await response.json()
            shortDescriptionBox.textContent = body['shortDescription']
            fullDescriptionBox.textContent = body['fullDescription']
            descriptionBody.style.visibility = 'visible'
            break
        case 400:
            pushNotify(`Wrong code body!`)
            break
        case 404:
            pushNotify(`Not Found!`)
            break
        case 500:
            pushNotify(`Something went wrong :(`)
            break
        default:
            pushNotify(`Unexpected Error :(`)
            break
    }
}

countComplexityBttn.onclick = countAlgorithmComplexity;

// Notifications
function pushNotify(title) {
    let myNotify = new Notify({
        status: 'warning',
        title: title,
        effect: 'slide',
        autoclose: true,
        autotimeout: 3000,
        type: 2
    })
}


