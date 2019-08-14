// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

import { projectMembersModule } from '@/store/modules/projectMembers';
import { createLocalVue } from '@vue/test-utils';
import { PROJECT_MEMBER_MUTATIONS } from '@/store/mutationConstants';
import Vuex from 'vuex';
import { ProjectMember, ProjectMemberCursor, ProjectMembersPage } from '@/types/projectMembers';

const mutations = projectMembersModule.mutations;

describe('mutations', () => {
    beforeEach(() => {
        createLocalVue().use(Vuex);
    });

    it('success delete project members', () => {
        const state = {
            cursor: new ProjectMemberCursor(),
            page: {
                projectMembers: [{user: {email: '1'}}, {user: {email: '2'}}]
            } as ProjectMembersPage,
        };
        const store = new Vuex.Store({state, mutations});

        const membersToDelete = ['1', '2'];

        store.commit(PROJECT_MEMBER_MUTATIONS.DELETE, membersToDelete);

        expect(state.page.projectMembers.length).toBe(0);
    });

    it('error delete project members', () => {
        const state = {
            cursor: new ProjectMemberCursor(),
            page: {
                projectMembers: [{user: {email: '1'}}, {user: {email: '2'}}]
            } as ProjectMembersPage,
        };
        const store = new Vuex.Store({state, mutations});

        const membersToDelete = ['3', '4'];

        store.commit(PROJECT_MEMBER_MUTATIONS.DELETE, membersToDelete);

        expect(state.page.projectMembers.length).toBe(2);
    });

    it('toggle selection', () => {
        const state = {
            cursor: new ProjectMemberCursor(),
            page: {
                projectMembers: [{
                    user: {id: '1'},
                    isSelected: false
                }, {
                    user: {id: '2'},
                    isSelected: false
                }]
            } as ProjectMembersPage,
        };
        const store = new Vuex.Store({state, mutations});

        const memberId = '1';

        store.commit(PROJECT_MEMBER_MUTATIONS.TOGGLE_SELECTION, memberId);

        expect(state.page.projectMembers[0].isSelected).toBeTruthy();
        expect(state.page.projectMembers[1].isSelected).toBeFalsy();
    });

    it('clear selection', () => {
        const state = {
            cursor: new ProjectMemberCursor(),
            page: {
                projectMembers: [{
                    user: {id: '1'},
                    isSelected: true
                }, {
                    user: {id: '2'},
                    isSelected: true
                }]
            } as ProjectMembersPage,
        };
        const store = new Vuex.Store({state, mutations});

        store.commit(PROJECT_MEMBER_MUTATIONS.CLEAR_SELECTION);

        expect(state.page.projectMembers[0].isSelected).toBeFalsy();
        expect(state.page.projectMembers[1].isSelected).toBeFalsy();
    });

    it('fetch team members', () => {
        const state = {
            cursor: new ProjectMemberCursor(),
            page: new ProjectMembersPage(),
        };
        const store = new Vuex.Store({state, mutations});

        const teamMembers = new ProjectMembersPage();
        teamMembers.projectMembers = [
            new ProjectMember('', '', '', '', '1'),
        ];

        store.commit(PROJECT_MEMBER_MUTATIONS.FETCH, teamMembers);

        expect(state.page.projectMembers.length).toBe(1);
    });
});

// describe('actions', () => {
//     beforeEach(() => {
//         jest.resetAllMocks();
//     });
//
//     it('success add project members', async function () {
//         const rootGetters = {
//             selectedProject: {
//                 id: '1'
//             },
//             searchParameters: {},
//             pagination: {limit: 20, offset: 0}
//         };
//         jest.spyOn(api, 'addProjectMembersRequest').mockReturnValue(Promise.resolve(<RequestResponse<null>>{isSuccess: true}));
//
//         const emails = ['1', '2'];
//
//         const dispatchResponse = await projectMembersModule.actions.addProjectMembers({rootGetters}, emails);
//
//         expect(dispatchResponse.isSuccess).toBeTruthy();
//     });
//
//     it('error add project members', async function () {
//         const rootGetters = {
//             selectedProject: {
//                 id: '1'
//             }
//         };
//         jest.spyOn(api, 'addProjectMembersRequest').mockReturnValue(Promise.resolve(<RequestResponse<null>>{isSuccess: false}));
//
//         const emails = ['1', '2'];
//
//         const dispatchResponse = await projectMembersModule.actions.addProjectMembers({rootGetters}, emails);
//
//         expect(dispatchResponse.isSuccess).toBeFalsy();
//     });
//
//     it('success delete project members', async () => {
//         const rootGetters = {
//             selectedProject: {
//                 id: '1'
//             }
//         };
//         jest.spyOn(api, 'deleteProjectMembersRequest').mockReturnValue(Promise.resolve(<RequestResponse<null>>{isSuccess: true}));
//
//         const commit = jest.fn();
//         const emails = ['1', '2'];
//
//         const dispatchResponse = await projectMembersModule.actions.deleteProjectMembers({commit, rootGetters}, emails);
//
//         expect(dispatchResponse.isSuccess).toBeTruthy();
//         expect(commit).toHaveBeenCalledWith(PROJECT_MEMBER_MUTATIONS.DELETE, emails);
//     });
//
//     it('error delete project members', async () => {
//         const rootGetters = {
//             selectedProject: {
//                 id: '1'
//             }
//         };
//         jest.spyOn(api, 'deleteProjectMembersRequest').mockReturnValue(Promise.resolve(<RequestResponse<null>>{isSuccess: false}));
//
//         const commit = jest.fn();
//         const emails = ['1', '2'];
//
//         const dispatchResponse = await projectMembersModule.actions.deleteProjectMembers({commit, rootGetters}, emails);
//
//         expect(dispatchResponse.isSuccess).toBeFalsy();
//         expect(commit).toHaveBeenCalledTimes(0);
//     });
//
//     it('toggle selection', function () {
//         const commit = jest.fn();
//         const projectMemberId = '1';
//
//         projectMembersModule.actions.toggleProjectMemberSelection({commit}, projectMemberId);
//
//         expect(commit).toHaveBeenCalledWith(PROJECT_MEMBER_MUTATIONS.TOGGLE_SELECTION, projectMemberId);
//     });
//
//     it('clear selection', function () {
//         const commit = jest.fn();
//
//         projectMembersModule.actions.clearProjectMemberSelection({commit});
//
//         expect(commit).toHaveBeenCalledTimes(1);
//     });
//
//     it('success fetch project members', async function () {
//         const rootGetters = {
//             selectedProject: {
//                 id: '1'
//             }
//         };
//         const state = {
//             pagination: {
//                 limit: 20,
//                 offset: 0
//             },
//             searchParameters: {
//                 searchQuery: ''
//             }
//         };
//         const commit = jest.fn();
//         const projectMemberMockModel: ProjectMember = new ProjectMember('1', '1', '1', '1', '1');
//         jest.spyOn(api, 'fetchProjectMembersRequest').mockReturnValue(
//             Promise.resolve(<RequestResponse<ProjectMember[]>>{
//                 isSuccess: true,
//                 data: [projectMemberMockModel]
//             }));
//
//         const dispatchResponse = await projectMembersModule.actions.fetchProjectMembers({
//             state,
//             commit,
//             rootGetters
//         });
//
//         expect(dispatchResponse.isSuccess).toBeTruthy();
//         expect(commit).toHaveBeenCalledWith(PROJECT_MEMBER_MUTATIONS.FETCH, [projectMemberMockModel]);
//     });
//
//     it('error fetch project members', async function () {
//         const rootGetters = {
//             selectedProject: {
//                 id: '1'
//             }
//         };
//         const state = {
//             pagination: {
//                 limit: 20,
//                 offset: 0
//             },
//             searchParameters: {
//                 searchQuery: ''
//             }
//         };
//         const commit = jest.fn();
//         jest.spyOn(api, 'fetchProjectMembersRequest').mockReturnValue(
//             Promise.resolve(<RequestResponse<ProjectMember[]>>{
//                 isSuccess: false,
//             })
//         );
//
//         const dispatchResponse = await projectMembersModule.actions.fetchProjectMembers({
//             state,
//             commit,
//             rootGetters
//         });
//
//         expect(dispatchResponse.isSuccess).toBeFalsy();
//         expect(commit).toHaveBeenCalledTimes(0);
//     });
// });

// describe('getters', () => {
//     it('project members', function () {
//         const state = {
//             projectMembers: [{user: {email: '1'}}]
//         };
//         const retrievedProjectMembers = projectMembersModule.getters.projectMembers(state);
//
//         expect(retrievedProjectMembers.length).toBe(1);
//     });
//
//     it('selected project members', function () {
//         const state = {
//             projectMembers: [
//                 {isSelected: false},
//                 {isSelected: true},
//                 {isSelected: true},
//                 {isSelected: true},
//                 {isSelected: false},
//             ]
//         };
//         const retrievedSelectedProjectMembers = projectMembersModule.getters.selectedProjectMembers(state);
//         expect(retrievedSelectedProjectMembers.length).toBe(3);
//     });
// });
