// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

import Vue, { VNode } from 'vue';
import VueClipboard from 'vue-clipboard2';
import { DirectiveBinding } from 'vue/types/options';

import { NotificatorPlugin } from '@/utils/plugins/notificator';
import { SegmentioPlugin } from '@/utils/plugins/segment';

import App from './App.vue';
import { router } from './router';
import { store } from './store';

Vue.config.devtools = true;
Vue.config.performance = true;
Vue.config.productionTip = false;

const notificator = new NotificatorPlugin();
const segment = new SegmentioPlugin();

Vue.use(notificator);
Vue.use(segment);
Vue.use(VueClipboard);

let clickOutsideEvent: EventListener;

/**
 * Binds closing action to outside popups area.
 */
Vue.directive('click-outside', {
    bind: function (el: HTMLElement, binding: DirectiveBinding, vnode: VNode) {
        clickOutsideEvent = function(event: Event): void {
            if (el === event.target || el.contains((event.target as Node))) {
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

/**
 * centsToDollars is a Vue filter that converts amount of cents in dollars string.
 */
Vue.filter('centsToDollars', (cents: number): string => {
    return `$${(cents / 100).toFixed(2)}`;
});

new Vue({
    router,
    store,
    render: (h) => h(App),
}).$mount('#app');
