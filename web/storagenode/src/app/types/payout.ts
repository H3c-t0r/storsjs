// Copyright (C) 2020 Storj Labs, Inc.
// See LICENSE for copying information.

// TODO: move comp division from api.
const PRICE_DIVIDER: number = 10000;

/**
 * Holds request arguments for payout information.
 */
export class PaymentInfoParameters {
    public constructor(
        public start: PayoutPeriod | null = null,
        public end: PayoutPeriod = new PayoutPeriod(),
        public satelliteId: string = '',
    ) {}
}

/**
 * Holds payout information.
 */
export class HeldInfo {
    public constructor(
        public usageAtRest: number = 0,
        public usageGet: number = 0,
        public usagePut: number = 0,
        public usageGetRepair: number = 0,
        public usagePutRepair: number = 0,
        public usageGetAudit: number = 0,
        public compAtRest: number = 0,
        public compGet: number = 0,
        public compPut: number = 0,
        public compGetRepair: number = 0,
        public compPutRepair: number = 0,
        public compGetAudit: number = 0,
        public surgePercent: number = 0,
        public held: number = 0,
        public owed: number = 0,
        public disposed: number = 0,
        public paid: number = 0,
        public paidWithoutSurge: number = 0,
    ) {}
}

/**
 * Represents payout period month and year.
 */
export class PayoutPeriod {
    public constructor(
        public year: number = new Date().getUTCFullYear(),
        public month: number = new Date().getUTCMonth(),
    ) {}

    public get period(): string {
        return this.month < 9 ? `${this.year}-0${this.month + 1}` : `${this.year}-${this.month + 1}`;
    }

    /**
     * Parses PayoutPeriod from string.
     * @param period string
     */
    public static fromString(period: string): PayoutPeriod {
        const periodArray = period.split('-');

        return new PayoutPeriod(parseInt(periodArray[0]), parseInt(periodArray[1]) - 1);
    }
}

/**
 * Holds 'start' and 'end' of payout period range.
 */
export class PayoutInfoRange {
    public constructor(
        public start: PayoutPeriod | null = null,
        public end: PayoutPeriod = new PayoutPeriod(),
    ) {}
}

/**
 * Holds accumulated held and earned payouts.
 */
export class TotalPayoutInfo {
    public constructor(
        public totalHeldAmount: number = 0,
        public totalEarnings: number = 0,
        public currentMonthEarnings: number = 0,
    ) {}
}

/**
 * Holds all payout module state.
 */
export class PayoutState {
    public constructor (
        public heldInfo: HeldInfo = new HeldInfo(),
        public periodRange: PayoutInfoRange = new PayoutInfoRange(),
        public totalHeldAmount: number = 0,
        public totalEarnings: number = 0,
        public currentMonthEarnings: number = 0,
        public heldPercentage: number = 0,
        public payoutPeriods: PayoutPeriod[] = [],
        public heldHistory: HeldHistory = new HeldHistory(),
        public payoutHistory: PayoutHistoryItem[] = [],
        public payoutHistoryPeriod: string = '',
        public estimation: EstimatedPayout = new EstimatedPayout(),
    ) {}
}

/**
 * Exposes all payout-related functionality.
 */
export interface PayoutApi {
    /**
     * Fetches held amount information by selected period.
     * @throws Error
     */
    getHeldInfoByPeriod(paymentInfoParameters: PaymentInfoParameters): Promise<HeldInfo>;

    /**
     * Fetches held amount information by selected month.
     * @throws Error
     */
    getHeldInfoByMonth(paymentInfoParameters: PaymentInfoParameters): Promise<HeldInfo>;

    /**
     * Fetches available payout periods.
     * @throws Error
     */
    getPayoutPeriods(id: string): Promise<PayoutPeriod[]>;

    /**
     * Fetches total payout information.
     * @throws Error
     */
    getTotal(paymentInfoParameters: PaymentInfoParameters): Promise<TotalPayoutInfo>;

    /**
     * Fetches held history for all satellites.
     * @throws Error
     */
    getHeldHistory(): Promise<HeldHistory>;

    /**
     * Fetch estimated payout information.
     * @throws Error
     */
    getEstimatedInfo(satelliteId: string): Promise<EstimatedPayout>;

    /**
     * Fetches payout history for all satellites.
     * @throws Error
     */
    getPayoutHistory(period: string): Promise<PayoutHistoryItem[]>;
}

/**
 * Holds held history information for all satellites.
 */
export class HeldHistory {
    public constructor(
        public monthlyBreakdown: HeldHistoryMonthlyBreakdownItem[] = [],
        public allStats: HeldHistoryAllStatItem[] = [],
    ) {}
}

/**
 * Contains held amounts of satellite grouped by periods.
 */
export class HeldHistoryMonthlyBreakdownItem {
    public constructor(
        public satelliteID: string = '',
        public satelliteName: string = '',
        public age: number = 1,
        public firstPeriod: number = 0,
        public secondPeriod: number = 0,
        public thirdPeriod: number = 0,
    ) {
        this.firstPeriod = this.firstPeriod / PRICE_DIVIDER;
        this.secondPeriod = this.secondPeriod / PRICE_DIVIDER;
        this.thirdPeriod = this.thirdPeriod / PRICE_DIVIDER;
    }
}

/**
 * Contains held information summary of satellite grouped by periods.
 */
export class HeldHistoryAllStatItem {
    public constructor(
        public satelliteID: string = '',
        public satelliteName: string = '',
        public age: number = 1,
        public totalHeld: number = 0,
        public totalDisposed: number = 0,
        public joinedAt: Date = new Date(),
    ) {
        this.totalHeld = this.totalHeld / PRICE_DIVIDER;
        this.totalDisposed = this.totalDisposed / PRICE_DIVIDER;
    }
}

/**
 * Contains estimated payout information for current and last periods.
 */
export class EstimatedPayout {
    public constructor(
        public currentMonth: PreviousMonthEstimatedPayout = new PreviousMonthEstimatedPayout(),
        public previousMonth: PreviousMonthEstimatedPayout = new PreviousMonthEstimatedPayout(),
    ) {}
}

/**
 * Contains last month estimated payout information.
 */
export class PreviousMonthEstimatedPayout {
    public constructor(
        public egressBandwidth: number = 0,
        public egressBandwidthPayout: number = 0,
        public egressRepairAudit: number = 0,
        public egressRepairAuditPayout: number = 0,
        public diskSpace: number = 0,
        public diskSpacePayout: number = 0,
        public heldRate: number = 0,
        public payout: number = 0,
        public held: number = 0,
    ) {}
}

/**
 * Contains payout information for payout history table.
 */
export class PayoutHistoryItem {
    public constructor(
        public satelliteID: string = '',
        public satelliteName: string = '',
        public age: number = 1,
        public earned: number = 0,
        public surge: number = 0,
        public surgePercent: number = 0,
        public held: number = 0,
        public afterHeld: number = 0,
        public disposed: number = 0,
        public paid: number = 0,
        public receipt: string = '',
        public isExitComplete: boolean = false,
        public heldPercent: number = 0,
    ) {
        this.earned = this.earned / PRICE_DIVIDER;
        this.surge = this.surge / PRICE_DIVIDER;
        this.held = this.held / PRICE_DIVIDER;
        this.afterHeld = this.afterHeld / PRICE_DIVIDER;
        this.disposed = this.disposed / PRICE_DIVIDER;
        this.paid = this.paid / PRICE_DIVIDER;
    }
}

export interface StoredMonthsByYear {
    [key: number]: MonthButton[];
}

/**
 * Holds all months names.
 */
export const monthNames = [
    'January', 'February', 'March', 'April',
    'May', 'June', 'July',	'August',
    'September', 'October', 'November',	'December',
];

/**
 * Describes month button entity for calendar.
 */
export class MonthButton {
    public constructor(
        public year: number = 0,
        public index: number = 0,
        public active: boolean = false,
        public selected: boolean = false,
    ) {}

    /**
     * Returns month label depends on index.
     */
    public get name(): string {
        return monthNames[this.index] ? monthNames[this.index].slice(0, 3) : '';
    }
}
