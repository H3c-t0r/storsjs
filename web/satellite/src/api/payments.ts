// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

import { ErrorUnauthorized } from '@/api/errors/ErrorUnauthorized';
import { BillingHistoryItem, CreditCard, PaymentsApi, ProjectCharge } from '@/types/payments';
import { HttpClient } from '@/utils/httpClient';
import { Time } from '@/utils/time';

/**
 * PaymentsHttpApi is a http implementation of Payments API.
 * Exposes all payments-related functionality
 */
export class PaymentsHttpApi implements PaymentsApi {
    private readonly client: HttpClient = new HttpClient();
    private readonly ROOT_PATH: string = '/api/v0/payments';

    /**
     * Get account balance.
     *
     * @returns balance in cents
     * @throws Error
     */
    public async getBalance(): Promise<AccountBalance> {
        const path = `${this.ROOT_PATH}/account/balance`;
        const response = await this.client.get(path);

        if (!response.ok) {
            if (response.status === 401) {
                throw new ErrorUnauthorized();
            }

            throw new Error('Can not get account balance');
        }

        const balance = await response.json();
        if (balance) {
            return new AccountBalance(balance.freeCredits, balance.coins);
        }

        return new AccountBalance();
    }

    /**
     * Try to set up a payment account.
     *
     * @throws Error
     */
    public async setupAccount(): Promise<void> {
        const path = `${this.ROOT_PATH}/account`;
        const response = await this.client.post(path, null);

        if (response.ok) {
            return;
        }

        if (response.status === 401) {
            throw new ErrorUnauthorized();
        }

        throw new Error('can not setup account');
    }

    /**
     * projectsCharges returns how much money current user will be charged for each project which he owns.
     */
    public async projectsCharges(): Promise<ProjectCharge[]> {
        const path = `${this.ROOT_PATH}/account/charges`;
        const response = await this.client.get(path);

        if (!response.ok) {
            if (response.status === 401) {
                throw new ErrorUnauthorized();
            }

            throw new Error('can not get projects charges');
        }

        const charges = await response.json();
        if (charges) {
            return charges.map(charge =>
                new ProjectCharge(
                    charge.projectId,
                    charge.storage,
                    charge.egress,
                    charge.objectCount),
            );
        }

        return [];
    }

    /**
     * projectsUsageAndCharges returns usage and how much money current user will be charged for each project which he owns.
     */
    public async projectsUsageAndCharges(start: Date, end: Date): Promise<ProjectUsageAndCharges[]> {
        const since = Time.toUnixTimestamp(start).toString();
        const before = Time.toUnixTimestamp(end).toString();
        const path = `${this.ROOT_PATH}/account/charges?from=${since}&to=${before}`;
        const response = await this.client.get(path);

        if (!response.ok) {
            if (response.status === 401) {
                throw new ErrorUnauthorized();
            }

            throw new Error('can not get projects charges');
        }

        const charges = await response.json();
        if (charges) {
            return charges.map(charge =>
                new ProjectUsageAndCharges(
                    new Date(charge.since),
                    new Date(charge.before),
                    charge.egress,
                    charge.storage,
                    charge.objectCount,
                    charge.projectId,
                    charge.storagePrice,
                    charge.egressPrice,
                    charge.objectPrice,
                ),
            );
        }

        return [];
    }

    /**
     * Add credit card.
     *
     * @param token - stripe token used to add a credit card as a payment method
     * @throws Error
     */
    public async addCreditCard(token: string): Promise<void> {
        const path = `${this.ROOT_PATH}/cards`;
        const response = await this.client.post(path, token);

        if (response.ok) {
            return;
        }

        if (response.status === 401) {
            throw new ErrorUnauthorized();
        }

        throw new Error('can not add credit card');
    }

    /**
     * Detach credit card from payment account.
     *
     * @param cardId
     * @throws Error
     */
    public async removeCreditCard(cardId: string): Promise<void> {
        const path = `${this.ROOT_PATH}/cards/${cardId}`;
        const response = await this.client.delete(path);

        if (response.ok) {
            return;
        }

        if (response.status === 401) {
            throw new ErrorUnauthorized();
        }

        throw new Error('can not remove credit card');
    }

    /**
     * Get list of user`s credit cards.
     *
     * @returns list of credit cards
     * @throws Error
     */
    public async listCreditCards(): Promise<CreditCard[]> {
        const path = `${this.ROOT_PATH}/cards`;
        const response = await this.client.get(path);

        if (!response.ok) {
            if (response.status === 401) {
                throw new ErrorUnauthorized();
            }
            throw new Error('can not list credit cards');
        }

        const creditCards = await response.json();

        if (creditCards) {
            return creditCards.map(card => new CreditCard(card.id, card.expMonth, card.expYear, card.brand, card.last4, card.isDefault));
        }

        return [];
    }

    /**
     * Make credit card default.
     *
     * @param cardId
     * @throws Error
     */
    public async makeCreditCardDefault(cardId: string): Promise<void> {
        const path = `${this.ROOT_PATH}/cards`;
        const response = await this.client.patch(path, cardId);

        if (response.ok) {
            return;
        }

        if (response.status === 401) {
            throw new ErrorUnauthorized();
        }

        throw new Error('can not make credit card default');
    }

    /**
     * Returns a list of invoices, transactions and all others payments history items for payment account.
     *
     * @returns list of payments history items
     * @throws Error
     */
    public async paymentsHistory(): Promise<PaymentsHistoryItem[]> {
        const path = `${this.ROOT_PATH}/billing-history`;
        const response = await this.client.get(path);

        if (!response.ok) {
            if (response.status === 401) {
                throw new ErrorUnauthorized();
            }
            throw new Error('can not list billing history');
        }

        const paymentsHistoryItems = await response.json();
        if (paymentsHistoryItems) {
            return paymentsHistoryItems.map(item =>
                new PaymentsHistoryItem(
                    item.id,
                    item.description,
                    item.amount,
                    item.received,
                    item.status,
                    item.link,
                    new Date(item.start),
                    new Date(item.end),
                    item.type,
                    item.remaining,
                ),
            );
        }

        return [];
    }

    /**
     * makeTokenDeposit process coin payments.
     *
     * @param amount
     * @throws Error
     */
    public async makeTokenDeposit(amount: number): Promise<TokenDeposit> {
        const path = `${this.ROOT_PATH}/tokens/deposit`;
        const response = await this.client.post(path, JSON.stringify({ amount }));

        if (!response.ok) {
            if (response.status === 401) {
                throw new ErrorUnauthorized();
            }

            throw new Error('can not process coin payment');
        }

        const result = await response.json();

        return new TokenDeposit(result.amount, result.address, result.link);
    }

    /**
     * Indicates if paywall is enabled.
     *
     * @param userId
     * @throws Error
     */
    public async getPaywallStatus(userId: string): Promise<boolean> {
        const path = `${this.ROOT_PATH}/paywall-enabled/${userId}`;
        const response = await this.client.get(path);

        if (!response.ok) {
            if (response.status === 401) {
                throw new ErrorUnauthorized();
            }

            throw new Error('can not get paywall status');
        }

        return await response.json();
    }
}
