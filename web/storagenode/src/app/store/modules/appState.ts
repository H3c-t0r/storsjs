// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

export const APPSTATE_MUTATIONS = {
    TOGGLE_SATELLITE_SELECTION: 'TOGGLE_SATELLITE_SELECTION',
    TOGGLE_BANDWIDTH_CHART: 'TOGGLE_BANDWIDTH_CHART',
    TOGGLE_EGRESS_CHART: 'TOGGLE_EGRESS_CHART',
    TOGGLE_INGRESS_CHART: 'TOGGLE_INGRESS_CHART,',
    CLOSE_ADDITIONAL_CHARTS: 'CLOSE_ADDITIONAL_CHARTS',
    CLOSE_ALL_POPUPS: 'CLOSE_ALL_POPUPS',
};

export const APPSTATE_ACTIONS = {
    TOGGLE_SATELLITE_SELECTION: 'TOGGLE_SATELLITE_SELECTION',
    TOGGLE_BANDWIDTH_CHART: 'TOGGLE_BANDWIDTH_CHART',
    TOGGLE_EGRESS_CHART: 'TOGGLE_EGRESS_CHART',
    TOGGLE_INGRESS_CHART: 'TOGGLE_INGRESS_CHART',
    CLOSE_ADDITIONAL_CHARTS: 'CLOSE_ADDITIONAL_CHARTS',
    CLOSE_ALL_POPUPS: 'CLOSE_ALL_POPUPS',
};

const {
    TOGGLE_SATELLITE_SELECTION,
    TOGGLE_BANDWIDTH_CHART,
    TOGGLE_EGRESS_CHART,
    TOGGLE_INGRESS_CHART,
    CLOSE_ADDITIONAL_CHARTS,
    CLOSE_ALL_POPUPS,
} = APPSTATE_MUTATIONS;

export const appStateModule = {
    state: {
        isSatelliteSelectionShown: false,
        isBandwidthChartShown: true,
        isEgressChartShown: false,
        isIngressChartShown: false,
    },
    mutations: {
        [TOGGLE_SATELLITE_SELECTION](state: any): void {
            state.isSatelliteSelectionShown = !state.isSatelliteSelectionShown;
        },
        [TOGGLE_BANDWIDTH_CHART](state: any): void {
            state.isBandwidthChartShown = !state.isBandwidthChartShown;
        },
        [TOGGLE_EGRESS_CHART](state: any): void {
            state.isEgressChartShown = !state.isEgressChartShown;
        },
        [TOGGLE_INGRESS_CHART](state: any): void {
            state.isIngressChartShown = !state.isIngressChartShown;
        },
        [CLOSE_ADDITIONAL_CHARTS](state: any): void {
            state.isBandwidthChartShown = true;
            state.isIngressChartShown = false;
            state.isEgressChartShown = false;
        },
        [CLOSE_ALL_POPUPS](state: any): void {
            state.isSatelliteSelectionShown = false;
        },
    },
    actions: {
        [APPSTATE_ACTIONS.TOGGLE_SATELLITE_SELECTION]: function ({commit, state}: any): void {
            if (!state.isSatelliteSelectionShown) {
                commit(APPSTATE_MUTATIONS.TOGGLE_SATELLITE_SELECTION);

                return;
            }

            commit(APPSTATE_MUTATIONS.CLOSE_ALL_POPUPS);
        },
        [APPSTATE_ACTIONS.TOGGLE_EGRESS_CHART]: function ({commit, state}: any): void {
            if (!state.isBandwidthChartShown) {
                commit(APPSTATE_MUTATIONS.TOGGLE_EGRESS_CHART);
                commit(APPSTATE_MUTATIONS.TOGGLE_INGRESS_CHART);

                return;
            }

            commit(APPSTATE_MUTATIONS.TOGGLE_BANDWIDTH_CHART);
            commit(APPSTATE_MUTATIONS.TOGGLE_EGRESS_CHART);
        },
        [APPSTATE_ACTIONS.TOGGLE_INGRESS_CHART]: function ({commit, state}: any): void {
            if (!state.isBandwidthChartShown) {
                commit(APPSTATE_MUTATIONS.TOGGLE_INGRESS_CHART);
                commit(APPSTATE_MUTATIONS.TOGGLE_EGRESS_CHART);

                return;
            }

            commit(APPSTATE_MUTATIONS.TOGGLE_BANDWIDTH_CHART);
            commit(APPSTATE_MUTATIONS.TOGGLE_INGRESS_CHART);
        },

        [APPSTATE_ACTIONS.CLOSE_ADDITIONAL_CHARTS]: function ({commit}: any): void {
            commit(APPSTATE_MUTATIONS.CLOSE_ADDITIONAL_CHARTS);
        },
        [APPSTATE_ACTIONS.CLOSE_ALL_POPUPS]: function ({commit}: any): void {
            commit(APPSTATE_MUTATIONS.CLOSE_ALL_POPUPS);
        }
    },
};
