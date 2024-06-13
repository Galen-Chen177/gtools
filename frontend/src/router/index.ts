import { createRouter, createWebHistory } from 'vue-router'

import json2struct from '../components/json2struct.vue'


const router = createRouter({
    history: createWebHistory(),
    routes: [
        {
            path: '/json2struct',
            component: json2struct
        }
    ]
})

export default router