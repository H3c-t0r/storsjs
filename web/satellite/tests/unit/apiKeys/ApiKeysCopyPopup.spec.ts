// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

import { createLocalVue, mount } from '@vue/test-utils';
import Vuex from 'vuex';
import ApiKeysCopyPopup from '@/components/apiKeys/ApiKeysCopyPopup.vue';
import { makeApiKeysModule } from '@/store/modules/apiKeys';

const localVue = createLocalVue();
localVue.use(Vuex);
const apiKeysModule = makeApiKeysModule();
const store = new Vuex.Store(apiKeysModule);

describe('ApiKeysCopyPopup', () => {
    it('renders correctly', () => {
        const wrapper = mount(ApiKeysCopyPopup, {
            store,
            localVue
        });

        expect(wrapper).toMatchSnapshot();
    });


    it('function onCloseClick works correctly', () => {
        const wrapper = mount(ApiKeysCopyPopup, {
            store,
            localVue,
        });

        wrapper.vm.onCloseClick();

        expect(wrapper.vm.$data.isCopiedButtonShown).toBe(false);
        expect(wrapper.emitted()).toEqual({'closePopup': [[]]});
    });

    it('function onCopyClick works correctly', () => {
        const wrapper = mount(ApiKeysCopyPopup, {
            store,
            localVue,
        });

        wrapper.vm.onCopyClick();

        expect(wrapper.vm.$data.isCopiedButtonShown).toBe(true);
    });
});
