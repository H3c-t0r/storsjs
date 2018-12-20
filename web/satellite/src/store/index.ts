// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

import Vue from 'vue';
import Vuex from 'vuex';


import { usersModule } from '@/store/modules/users';
import { projectsModule } from '@/store/modules/projects';
import { projectMembersModule } from '@/store/modules/projectMembers';
import { notificationsModule } from '@/store/modules/notifications';

Vue.use(Vuex);

// Satellite store (vuex)
const store = new Vuex.Store({
	modules: {
		usersModule,
		projectsModule,
		projectMembersModule,
		notificationsModule,
	}
});

export default store;
