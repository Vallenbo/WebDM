import {createRouter, createWebHistory} from 'vue-router'
import HomeView from '../components/layout/layout.vue'

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL), routes: [{
        path: '/', name: 'home', component: HomeView
    }, {
        path: '/image', name: 'about', redirect: "", // route level code-splitting
        // this generates a separate chunk (About.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        component: HomeView, children: [{
            path: "", component: import("../views/imagelist.vue")
        }, {
            path: "container/:id", component: import("../views/createcontainer.vue")
        }, {
            path: "", component: import("../views/containerlist.vue")
        }]
    }, {
        path: "/container", component: HomeView, redirect: "", children: [{
            path: "", component: import("../views/containerlist.vue")
        }]
    }, {
        path: "/servermonitor", component: HomeView, redirect: "", children: [{
            path: "", component: import("../views/servermonitor.vue")
        }]
    }, {
        path: "/createimage", component: HomeView, redirect: "", children: [{
            path: "", component: import("../views/createimage.vue")
        }]
    }, {
        path: "/version", component: HomeView, redirect: "", children: [{
            path: "", component: import("../views/DockerVersion.vue")
        }]
    }, {
        path: "/monitor", component: HomeView, redirect: "", children: [{
            path: "", component: import("../views/serverinfo.vue")
        }]
    }, {
        path: "/pullImage", component: HomeView, redirect: "", children: [{
            path: "", component: import("../views/pullImage.vue")
        }]
    }, {
        path: "/searchimage", component: HomeView, redirect: "", children: [{
            path: "", component: import("../views/searchimage.vue")
        }]
    }, {
        path: "/test", component: HomeView, redirect: "", children: [{
            path: "", component: import("../views/popups.vue")
        }]
    },]
})

export default router