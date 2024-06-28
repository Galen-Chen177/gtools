import { createRouter, createWebHistory } from 'vue-router'

import json2struct from '../components/json2struct.vue'
import jsonfile2struct from '../components/jsonfile2struct.vue'
import urlEscap from '../components/urlEscap.vue'
import redisshell from '../components/redisshell.vue'


const router = createRouter({
    history: createWebHistory(),
    routes: [
        {
            path: '/json2struct',
            component: json2struct
        },
        {
            path: '/jsonfile2struct',
            component: jsonfile2struct
        },
        {
            path: '/urlescap',
            component: urlEscap
        },
        {
            path: '/redisshell',
            component: redisshell
        }
    ]
})

export default router