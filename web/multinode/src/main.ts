// Copyright (C) 2020 Storj Labs, Inc.
// See LICENSE for copying information.

import Vue from 'vue';
import Router from 'vue-router';

import App from '@/app/App.vue';
import { router } from '@/app/router';

Vue.config.productionTip = false;

Vue.use(Router);

const app = new Vue({
    router,
    render: (h) => h(App),
});

app.$mount('#app');
