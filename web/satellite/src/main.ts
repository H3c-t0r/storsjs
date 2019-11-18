// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

import Vue, { VNode } from 'vue';
import { DirectiveBinding } from 'vue/types/options';

import { NotificatorPlugin } from '@/utils/plugins/notificator';
import { AnalyticsPlugin } from '@/utils/plugins/analytics';
import { SegmentioPlugin } from '@/utils/plugins/segment';

import App from './App.vue';
import { router } from './router';
import { store } from './store';

Vue.config.devtools = true;
Vue.config.performance = true;
Vue.config.productionTip = false;



const notificator = new NotificatorPlugin();
const analytics = new AnalyticsPlugin();
const segment = new SegmentioPlugin();

Vue.use(notificator);
Vue.use(segment);

// TODO: needs to have access to sgement ID served from the backend
Vue.use(analytics, {
    id: process.env.VUE_APP_SEGMENTID,
    router,
});

let clickOutsideEvent: EventListener;

Vue.directive('click-outside', {
    bind: function (el: HTMLElement, binding: DirectiveBinding, vnode: VNode) {
        clickOutsideEvent = function(event: Event): void {
            if (el === event.target) {
                return;
            }

            if (vnode.context) {
                vnode.context[binding.expression](event);
            }
        };

        document.body.addEventListener('click', clickOutsideEvent);
    },
    unbind: function(): void {
        document.body.removeEventListener('click', clickOutsideEvent);
    },
});

new Vue({
    router,
    store,
    render: (h) => h(App),
}).$mount('#app');
