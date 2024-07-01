import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
    history: createWebHistory(),
    routes: [
        {
            path: '/json2struct',
            component: () => import('../components/json2struct.vue'),
            meta: {
                refreshPage: true
            },
        },
        {
            path: '/jsonfile2struct',
            component: () => import('../components/jsonfile2struct.vue'),
            meta: {
                refreshPage: true
            },
        },
        {
            path: '/urlescap',
            component: () => import('../components/urlEscap.vue'),
            meta: {
                refreshPage: true
            },
        },
        {
            path: '/redisshell',
            component: () => import('../components/redisshell.vue'),
            meta: {
                refreshPage: true
            },
        },
        {
            path: '/configF',
            component: () => import('../components/config.vue'),
            meta: {
                refreshPage: true
            },
        }
        // Always leave this as last one,
        // but you can also remove it
        // {
        //     path: '/:catchAll(.*)*',
        //     component: () => import('pages/ErrorNotFound.vue'),
        // },
    ]
})

export default router