<template>
    <div class="main_div">
        <div class="input_div">
            <text>{{ fileName }}</text>
            <button @click="getOneFile">选择文件</button>
            <br />
            <textarea readonly cols="70" rows="20" ref="output"></textarea>
        </div>
        <br />
        <div class="button_div">
            <button @click="noIncludeStruct">无嵌套转换</button>
            <button @click="IncludeStruct">有嵌套转换</button>
        </div>
    </div>
</template>

<script setup lang="ts" name="jsonfile2struct">
import { ref } from 'vue'
import { decodeApiData } from '../utils/utils'
import { JsonFile2Struct } from '../../wailsjs/go/jsonfunc/JsonFunc'
import { OneFileDialog } from '../../wailsjs/go/dialog/BackendDialog'

let output = ref()
let fileName = ref('')

function getOneFile() {
    OneFileDialog("Json 文件 (*.json, *.txt)", "*.json;*.txt").then(result => {
        let resp = decodeApiData(result)
        if (resp.err != "") {
            fileName.value = "err:\n" + resp.err
            return
        }
        fileName.value = resp.data
    })
}

function noIncludeStruct() {
    if (fileName.value == "") {
        return
    }
    alert("未实现")
}

function IncludeStruct() {
    if (fileName.value == "") {
        return
    }
    JsonFile2Struct(fileName.value).then(result => {
        let resp = decodeApiData(result)
        if (resp.err != "") {
            output.value.value = "err:\n" + resp.err
            return
        }
        output.value.value = resp.data
    })
}
</script>

<style scoped>
.main_div {
    text-align: center;
}
</style>