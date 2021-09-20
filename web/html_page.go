package web
var mainPageHTML = `<html>
<style>
    button {
        margin-left :5px
    }
    .grid {
        display: grid;
        grid-template-columns: repeat(4, 1fr);
    }
    .grid > span {
        padding: 8px 4px;
    }
</style>
<body onload="loadConfig()">
<script>
    async function loadConfig() {
        // Read configuration
        let response = await fetch('/config');
        let data = await response.text();
        console.log(data);
        let config = JSON.parse(data);

        // Read device status
        response = await fetch('/usb_device');
        data = await response.text();
        let device = JSON.parse(data);
        console.log(device);

        // Read parameters
        response = await fetch('/parameter');
        data = await response.text();
        let parameterResponse = JSON.parse(data);
        console.log(device);

        document.body.innerHTML = document.body.innerHTML + "<h1>" + config.name + " (" + device.status + ")</h1>";
        document.body.innerHTML = document.body.innerHTML + '<hr>'
        let buttonConfig = config.button_config;
        let buttonValue = parseInt(buttonConfig.start_value);
        let buttonLabels = buttonConfig.labels;
        let buttonsString = "";
        buttonLabels.forEach(label => {
           buttonsString = buttonsString + "<button onmousedown=" + upDownCommandString(buttonValue, "down") +
               " onmouseup=" + upDownCommandString(buttonValue, "up") + ">" + label + "</button>";
            buttonValue++;
        })
        document.body.innerHTML = document.body.innerHTML + "<p>" + buttonsString + "</p>";
        document.body.innerHTML = document.body.innerHTML + '<div class="block"><label>velocity</label></div><input id="velocity" oninput="handleVelocityChange()" type="range" min="0" max="127" value="64"></div>'
        document.body.innerHTML = document.body.innerHTML + "<hr>";
        let knobConfig = config.knob_config;
        let knobLabels = knobConfig.labels
        let knobsString = "";
        let controlValue = 0;
        knobLabels.forEach(label => {
            let currentKnobValue = parameterResponse.Data[controlValue]
            knobsString = knobsString + '<span class="block"><label>' + label + '</label><input id="' + label + '" oninput="controlCommandString(\'' + label + '\',\'' + controlValue + '\')" type="range" min="0" max="127" value="' + currentKnobValue + '"></span>'
            controlValue++;
        })
        document.body.innerHTML = document.body.innerHTML + "<p><div class='grid'>" + knobsString + "</div></p>";
    }

    function controlCommandString(id, control_value) {
        let sliderValue = document.getElementById(id).value;
        const value = {control_value:control_value, value: sliderValue};
        const myJSON = JSON.stringify(value);
        console.log(myJSON);
        sendCommand('/control', myJSON);
    }

    function handleVelocityChange() {
        let sliderValue = document.getElementById("velocity").value;
        const value = {value: sliderValue};
        const myJSON = JSON.stringify(value);
        console.log(myJSON);
        sendCommand('/velocity', myJSON);
    }


    function upDownCommandString(key, upDown) {
        let jsonValue = '{\"key\":\"' + key + '\",\"value\":\"' + upDown + '\"}';
        let command = "sendKeyCommand('" + jsonValue + "')"
        return command
    }

    async function sendCommand(url, data) {
        const response = await fetch(url, {
            method: 'PUT',
            headers: {
                'content-type': 'application/json'
            },
            body: data
        })

        // Awaiting response.json()
        const resData = await response.json();

        console.log(resData)


    }

    async function sendKeyCommand(data) {
        const response = await fetch("/key", {
            method: 'PUT',
            headers: {
                'Content-type': 'application/json'
            },
            body: data
        })

        // Awaiting response.json()
        const resData = await response.json();

        console.log(resData)

    }

</script>
</body>
</html>`
