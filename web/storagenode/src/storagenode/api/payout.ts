// Copyright (C) 2020 Storj Labs, Inc.
// See LICENSE for copying information.

import { HeldInfo, PaymentInfoParameters, PayoutApi, TotalPayoutInfo } from '@/app/types/payout';
import { HttpClient } from '@/storagenode/utils/httpClient';

/**
 * NotificationsHttpApi is a http implementation of Notifications API.
 * Exposes all notifications-related functionality
 */
export class PayoutHttpApi implements PayoutApi {
    private readonly client: HttpClient = new HttpClient();
    private readonly ROOT_PATH: string = '/api/heldamount';

    /**
     * Fetch held amount information.
     *
     * @returns held amount information
     * @throws Error
     */
    public async getHeldInfo(paymentInfoParameters: PaymentInfoParameters): Promise<HeldInfo> {
        let path = `${this.ROOT_PATH}/paystubs/`;

        if (paymentInfoParameters.start) {
            path += paymentInfoParameters.start.period + '/';
        }

        path += paymentInfoParameters.end.period;

        if (paymentInfoParameters.satelliteId) {
            path += '?id=' + paymentInfoParameters.satelliteId;
        }

        const response = await this.client.get(path);

        if (!response.ok) {
            throw new Error('can not get held information');
        }

        const data: any[] = await response.json() || [];

        let usageAtRest: number = 0;
        let usageGet: number = 0;
        let usagePut: number = 0;
        let usageGetRepair: number = 0;
        let usagePutRepair: number = 0;
        let usageGetAudit: number = 0;
        let compAtRest: number = 0;
        let compGet: number = 0;
        let compPut: number = 0;
        let compGetRepair: number = 0;
        let compPutRepair: number = 0;
        let compGetAudit: number = 0;
        let held: number = 0;
        let owed: number = 0;
        let disposed: number = 0;
        let paid: number = 0;

        data.forEach((paystub: any) => {
            usageAtRest += paystub.usageAtRest;
            usageGet += paystub.usageGet;
            usagePut += paystub.usagePut;
            usageGetRepair += paystub.usageGetRepair;
            usagePutRepair += paystub.usagePutRepair;
            usageGetAudit += paystub.usageGetAudit;
            compAtRest += paystub.compAtRest;
            compGet += paystub.compGet;
            compPut += paystub.compPut;
            compGetRepair += paystub.compGetRepair;
            compPutRepair += paystub.compPutRepair;
            compGetAudit += paystub.compGetAudit;
            held += paystub.held;
            owed += paystub.owed;
            disposed += paystub.disposed;
            paid += paystub.paid;
        });

        return new HeldInfo(
            usageAtRest,
            usageGet,
            usagePut,
            usageGetRepair,
            usagePutRepair,
            usageGetAudit,
            compAtRest,
            compGet,
            compPut,
            compGetRepair,
            compPutRepair,
            compGetAudit,
            0,
            held,
            owed,
            disposed,
            paid,
        );
    }

    /**
     * Fetch total payout information.
     *
     * @returns total payout information
     * @throws Error
     */
    public async getTotal(paymentInfoParameters: PaymentInfoParameters): Promise<TotalPayoutInfo> {
        let path = `${this.ROOT_PATH}/paystubs/`;

        if (paymentInfoParameters.start) {
            path += paymentInfoParameters.start.period + '/';
        }

        path += paymentInfoParameters.end.period;

        if (paymentInfoParameters.satelliteId) {
            path += '?id=' + paymentInfoParameters.satelliteId;
        }

        const response = await this.client.get(path);

        if (!response.ok) {
            throw new Error('can not get total payout information');
        }

        const data: any[] = await response.json() || [];

        let held: number = 0;
        let paid: number = 0;

        data.forEach((paystub: any) => {
            held += paystub.held;
            paid += paystub.paid;
        });

        return new TotalPayoutInfo(
            held,
            paid,
        );
    }
}
