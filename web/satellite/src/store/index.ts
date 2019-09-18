// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

import Vue from 'vue';
import Vuex from 'vuex';

import { ApiKeysApiGql } from '@/api/apiKeys';
import { BucketsApiGql } from '@/api/buckets';
import { CreditsApiGql } from '@/api/credits';
import { ProjectMembersApiGql } from '@/api/projectMembers';
import { ProjectsApiGql } from '@/api/projects';
import { ProjectUsageApiGql } from '@/api/usage';
import { UsersApiGql } from '@/api/users';
import router from '@/router';
import { ApiKeysState, makeApiKeysModule } from '@/store/modules/apiKeys';
import { appStateModule } from '@/store/modules/appState';
import { BucketsState, makeBucketsModule } from '@/store/modules/buckets';
import { makeCreditsModule } from '@/store/modules/credits';
import { makeNotificationsModule, NotificationsState } from '@/store/modules/notifications';
import { projectPaymentsMethodsModule } from '@/store/modules/paymentMethods';
import { makeProjectMembersModule, ProjectMembersState } from '@/store/modules/projectMembers';
import { makeProjectsModule, PROJECTS_MUTATIONS, ProjectsState } from '@/store/modules/projects';
import { makeUsageModule, UsageState } from '@/store/modules/usage';
import { makeUsersModule } from '@/store/modules/users';
import { CreditUsage } from '@/types/credits';
import { User } from '@/types/users';

Vue.use(Vuex);

export class StoreModule<S> {
    public state: S;
    public mutations: any;
    public actions: any;
    public getters?: any;
}

// TODO: remove it after we will use modules as classes and use some DI framework
const usersApi = new UsersApiGql();
const apiKeysApi = new ApiKeysApiGql();
const creditsApi = new CreditsApiGql();
const bucketsApi = new BucketsApiGql();
const projectMembersApi = new ProjectMembersApiGql();
const projectsApi = new ProjectsApiGql();
const projectUsageApi = new ProjectUsageApiGql();

class ModulesState {
    public notificationsModule: NotificationsState;
    public apiKeysModule: ApiKeysState;
    public appStateModule;
    public creditsModule: CreditUsage;
    public projectMembersModule: ProjectMembersState;
    public projectPaymentsMethodsModule;
    public usersModule: User;
    public projectsModule: ProjectsState;
    public usageModule: UsageState;
    public bucketUsageModule: BucketsState;
}

// Satellite store (vuex)
const store = new Vuex.Store<ModulesState>({
    modules: {
        notificationsModule: makeNotificationsModule(),
        apiKeysModule: makeApiKeysModule(apiKeysApi),
        appStateModule,
        creditsModule: makeCreditsModule(creditsApi),
        projectMembersModule: makeProjectMembersModule(projectMembersApi),
        projectPaymentsMethodsModule,
        usersModule: makeUsersModule(usersApi),
        projectsModule: makeProjectsModule(projectsApi),
        usageModule: makeUsageModule(projectUsageApi),
        bucketUsageModule: makeBucketsModule(bucketsApi),
    },
});

store.subscribe((mutation, state) => {
    if (mutation.type === PROJECTS_MUTATIONS.SELECT_PROJECT) {
        document.title = `${state.projectsModule.selectedProject.name} | ${router.currentRoute.name} | us-central-1 - Tardigrade`;
    }
});

export default store;
