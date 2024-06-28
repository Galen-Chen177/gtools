<template>
    <div class="redis-cli">
        <h1>Redis CLI</h1>
        <div class="input-area">
            <input type="text" ref="command">
            <button @click="executeCommand">Execute</button>
        </div>
        <div class="output-area" ref="outputArea">
            <pre v-for="(output, index) in outputs" :key="index">{{ output }}</pre>
        </div>
    </div>
</template>

<script setup name="redisshell" lang="ts">
import { ref, onMounted } from 'vue'

let command = ref()
let outputs = ref([""])
let outputArea = ref()

onMounted(() => {

});


function executeCommand() {
    if (command.value.value == "") {
        return
    }
    outputs.value.push(command.value.value)
    outputs.value.push("Output: Placeholder for command execution result")

    // 显示最新的值
    outputArea.value.scrollTop = outputArea.value.scrollHeight
}

</script>


<style scoped>
.redis-cli {
    max-width: 600px;
    margin: 50px auto;
    padding: 20px;
    border: 1px solid #ccc;
    border-radius: 5px;
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
}

.input-area input[type="text"] {
    width: calc(100% - 80px);
    padding: 10px;
    font-size: 16px;
    border: 1px solid #ccc;
    border-radius: 3px;
}

.input-area button {
    padding: 10px 20px;
    font-size: 16px;
    border: none;
    background-color: #007bff;
    color: #fff;
    cursor: pointer;
    border-radius: 3px;
}

.output-area {
    margin-top: 10px;
    padding: 10px;
    background-color: #f0f0f0;
    border: 1px solid #ccc;
    border-radius: 3px;
    max-height: 300px;
    overflow-y: auto;
    white-space: pre-wrap;

}
</style>