<template>
    <div class="redis-config">
        <h2>Redis 配置</h2>
        <form @submit.prevent="submitConfig">
            <div>
                <label for="host">主机:</label>
                <input type="text" ref="redisHost" id="host" required />
            </div>
            <div>
                <label for="port">端口:</label>
                <input type="number" ref="redisPort" id="port" required />
            </div>
            <div>
                <label for="dbname">db:</label>
                <input type="number" ref="redisDbname" id="dbname" />
            </div>
            <div>
                <label for="password">密码:</label>
                <input type="password" ref="redisPassword" id="password" />
            </div>
            <label>{{ redisErr }}</label>
            <button type="submit">提交配置</button>
        </form>
    </div>
</template>


<script setup lang="ts" name="configF">
import { ref, onMounted } from 'vue'
import { log, decodeApiData } from '../utils/utils'
import { Read, Write } from '../../wailsjs/go/config/ConfigFrontend'


let redisHost = ref()
let redisPort = ref()
let redisDbname = ref()
let redisPassword = ref()
let redisErr = ref()
let config = ref()

// 打开界面时，将后端保存的配置填入到界面上
onMounted(() => {
    Read().then(result => {
        log("info", result)
        console.log(result)
        let resp = decodeApiData(result)
        if (resp.err != "") {
            redisErr.value = "err:" + resp.err
            return
        }
        redisErr.value = ""
        config.value = resp.data
        redisHost.value.value = config.value.redishost
        redisPort.value.value = config.value.redisport
        redisDbname.value.value = config.value.redisdb
        redisPassword.value.value = config.value.redispwd
    })
});

// 点击保存时，调用后端，保存到文件中，并且重新读取一次配置文件
function submitConfig() {
    let req = {
        "redis": {
            "host": redisHost.value.value,
            "port": parseInt(redisPort.value.value, 10),
            "db": parseInt(redisDbname.value.value, 10),
            "password": redisPassword.value.value
        }
    }
    let reqStr = JSON.stringify(req, null, 2);
    Write(reqStr).then(result => {
        console.log(result)
        let resp = decodeApiData(result)
        if (resp.err != "") {
            redisErr.value = "err:" + resp.err
            return
        }
        redisErr.value = ""
    })
}



</script>


<style scoped>
.redis-config {
    max-width: 400px;
    margin: 0 auto;
    padding: 20px;
    border: 1px solid #ccc;
    border-radius: 4px;
}

.redis-config div {
    margin-bottom: 10px;
}

.redis-config label {
    display: block;
    margin-bottom: 5px;
}

.redis-config input {
    width: 100%;
    padding: 8px;
    box-sizing: border-box;
}
</style>