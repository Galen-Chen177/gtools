<template>
    <div class="main_div">
        <div class="input_div">
            <textarea cols="50" rows="20" placeholder="请输入json数据..." ref="input"></textarea>
            <textarea readonly cols="50" rows="20" ref="output"></textarea>
        </div>
        <br />
        <div class="button_div">
            <button @click="noIncludeStruct">无嵌套转换</button>
            <button @click="IncludeStruct">有嵌套转换</button>
        </div>
    </div>
</template>


<script setup lang="ts" name="json2struct">
import { ref } from 'vue'
import { decodeApiData } from '../utils/utils'
import { Json2Struct } from '../../wailsjs/go/jsonfunc/JsonFunc'

let input = ref()
let output = ref()



function noIncludeStruct() {
    Json2Struct(input.value.value).then(result => {
        let resp = decodeApiData(result)
        if (resp.err != "") {
            output.value.value = "err:\n" + resp.err
            return
        }
        output.value.value = resp.data
    })
}
function IncludeStruct() {
    alert("未实现")
}
</script>

<style scoped>
.main_div {
    height: 100%;
    width: 100%;
}

textarea {
    margin-left: 3%;
    width: 45%;
}

.button_div {
    text-align: center;
}
</style>