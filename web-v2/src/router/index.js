import Vue from "vue";
import VueRouter from "vue-router";
import OrderList from '../components/OrderList.vue'

Vue.use(VueRouter);

const routes = [
    {
        path: '/orders',
        name: 'Orders',
        component: OrderList
    },
]

const router = new VueRouter({
    mode: 'history',
    routes // short for `routes: routes`
})

export default router;