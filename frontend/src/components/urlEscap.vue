<template>
    <div class="main_div">
        <div class="input_div">
            <textarea cols="50" rows="20" placeholder="请输入url" ref="input"></textarea>
            <textarea readonly cols="50" rows="20" ref="output"></textarea>
        </div>
        <br />
        <div class="button_div">
            <button @click="noIncludeStruct">解析</button>
        </div>
    </div>
</template>

<script setup lang="ts" name="urlEscap">
import { ref } from 'vue'
import { decodeApiData } from '../utils/utils'
import { UrlParse } from '../../wailsjs/go/jsonfunc/JsonFunc'

let input = ref()
let output = ref()



function noIncludeStruct() {
    UrlParse(input.value.value).then(result => {
        let resp = decodeApiData(result)
        if (resp.err != "") {
            output.value.value = "err:\n" + resp.err
            return
        }
        output.value.value = resp.data
    })
}

</script>



<style scoped></style>
