// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

import { PROJECT_MEMBER_MUTATIONS } from '../mutationConstants';
import {
    addProjectMembersRequest,
    deleteProjectMembersRequest,
    fetchProjectMembersRequest
} from '@/api/projectMembers';
import { ProjectMember, ProjectMemberCursor, ProjectMemberOrderBy, ProjectMembersPage } from '@/types/projectMembers';
import { RequestResponse } from '@/types/response';
import { PM_ACTIONS } from '@/utils/constants/actionNames';
import { SortDirection } from '@/types/common';

const projectMembersLimit = 8;
const firstPage = 1;

export const projectMembersModule = {
    state: {
        cursor: new ProjectMemberCursor(),
        page: new ProjectMembersPage(),
        selectedProjectMembers: [],
    },
    mutations: {
        [PROJECT_MEMBER_MUTATIONS.DELETE](state: any, projectMemberEmails: string[]) {
            const emailsCount = projectMemberEmails.length;

            for (let j = 0; j < emailsCount; j++) {
                state.page.projectMembers = state.page.projectMembers.filter((element: any) => {
                    return element.user.email !== projectMemberEmails[j];
                });
            }
        },
        [PROJECT_MEMBER_MUTATIONS.FETCH](state: any, page: ProjectMembersPage) {
            // todo expand this assignment
            state.page = page;
        },
        [PROJECT_MEMBER_MUTATIONS.SET_PAGE](state: any, page: number) {
            state.cursor.page = page;
        },
        [PROJECT_MEMBER_MUTATIONS.SET_SEARCH_QUERY](state: any, search: string) {
            state.cursor.search = search;
        },
        [PROJECT_MEMBER_MUTATIONS.CHANGE_SORT_ORDER](state: any, order: ProjectMemberOrderBy) {
            state.cursor.order = order;
        },
        [PROJECT_MEMBER_MUTATIONS.CHANGE_SORT_ORDER_DIRECTION](state: any, direction: SortDirection) {
            state.cursor.orderDirection = direction;
        },
        [PROJECT_MEMBER_MUTATIONS.CLEAR](state: any) {
            state.cursor = {limit: projectMembersLimit, search: '', page: firstPage} as ProjectMemberCursor;
            state.page = {projectMembers: [] as ProjectMember[]} as ProjectMembersPage;
        },
        [PROJECT_MEMBER_MUTATIONS.TOGGLE_SELECTION](state: any, projectMemberId: string) {
            state.page.projectMembers = state.page.projectMembers.map((projectMember: any) => {
                if (projectMember.user.id === projectMemberId) {
                    projectMember.isSelected = !projectMember.isSelected;
                }

                return projectMember;
            });
        },
        [PROJECT_MEMBER_MUTATIONS.CLEAR_SELECTION](state: any) {
            state.page.projectMembers = state.page.projectMembers.map((projectMember: any) => {
                projectMember.isSelected = false;

                return projectMember;
            });
        },
    },
    actions: {
        [PM_ACTIONS.ADD]: async function ({rootGetters}: any, emails: string[]): Promise<RequestResponse<null>> {
            const projectId = rootGetters.selectedProject.id;

            return await addProjectMembersRequest(projectId, emails);
        },
        [PM_ACTIONS.DELETE]: async function ({commit, rootGetters}: any, projectMemberEmails: string[]): Promise<RequestResponse<null>> {
            const projectId = rootGetters.selectedProject.id;

            const response = await deleteProjectMembersRequest(projectId, projectMemberEmails);

            if (response.isSuccess) {
                commit(PROJECT_MEMBER_MUTATIONS.DELETE, projectMemberEmails);
            }

            return response;
        },
        [PM_ACTIONS.FETCH]: async function ({commit, rootGetters, state}: any, page: number): Promise<RequestResponse<ProjectMembersPage>> {
            const projectID = rootGetters.selectedProject.id;
            state.cursor.page = page;

            commit(PROJECT_MEMBER_MUTATIONS.SET_PAGE, page);

            let result = await fetchProjectMembersRequest(projectID, state.cursor);
            if (result.isSuccess) {
                commit(PROJECT_MEMBER_MUTATIONS.FETCH, result.data);
            }

            return result;
        },

        [PM_ACTIONS.SET_SEARCH_QUERY]: function ({commit}, search: string) {
            commit(PROJECT_MEMBER_MUTATIONS.SET_SEARCH_QUERY, search);
        },
        [PM_ACTIONS.SET_SORT_BY]: function ({commit}, order: ProjectMemberOrderBy) {
            commit(PROJECT_MEMBER_MUTATIONS.CHANGE_SORT_ORDER, order);
        },
        [PM_ACTIONS.SET_SORT_DIRECTION]: function ({commit}, direction: SortDirection) {
            commit(PROJECT_MEMBER_MUTATIONS.CHANGE_SORT_ORDER_DIRECTION, direction);
        },
        [PM_ACTIONS.CLEAR]: function ({commit}) {
            commit(PROJECT_MEMBER_MUTATIONS.CLEAR);
        },
        [PM_ACTIONS.TOGGLE_SELECTION]: function ({commit}: any, projectMemberId: string) {
            commit(PROJECT_MEMBER_MUTATIONS.TOGGLE_SELECTION, projectMemberId);
        },
        [PM_ACTIONS.CLEAR_SELECTION]: function ({commit}: any) {
            commit(PROJECT_MEMBER_MUTATIONS.CLEAR_SELECTION);
        },
    },
    getters: {
        selectedProjectMembers: (state: any) => state.page.projectMembers.filter((member: any) => member.isSelected),
    }
};
