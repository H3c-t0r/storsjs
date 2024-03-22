// AUTOGENERATED BY private/apigen
// DO NOT EDIT.

import { HttpClient } from '@/utils/httpClient';
import { MemorySize, Time, UUID } from '@/types/common';

export class APIKeyInfo {
    id: UUID;
    projectId: UUID;
    projectPublicId: UUID;
    createdBy: UUID;
    userAgent: string | null;
    name: string;
    createdAt: Time;
    version: number;
}

export class APIKeyPage {
    apiKeys: APIKeyInfo[] | null;
    search: string;
    limit: number;
    order: number;
    orderDirection: number;
    offset: number;
    pageCount: number;
    currentPage: number;
    totalCount: number;
}

export class BucketUsageRollup {
    projectID: UUID;
    bucketName: string;
    totalStoredData: number;
    totalSegments: number;
    objectCount: number;
    metadataSize: number;
    repairEgress: number;
    getEgress: number;
    auditEgress: number;
    since: Time;
    before: Time;
}

export class CreateAPIKeyRequest {
    projectID: string;
    name: string;
}

export class CreateAPIKeyResponse {
    key: string;
    keyInfo: APIKeyInfo | null;
}

export class Project {
    id: UUID;
    publicId: UUID;
    name: string;
    description: string;
    userAgent: string | null;
    ownerId: UUID;
    rateLimit: number | null;
    burstLimit: number | null;
    maxBuckets: number | null;
    createdAt: Time;
    memberCount: number;
    storageLimit: MemorySize | null;
    bandwidthLimit: MemorySize | null;
    userSpecifiedStorageLimit: MemorySize | null;
    userSpecifiedBandwidthLimit: MemorySize | null;
    segmentLimit: number | null;
    defaultPlacement: number;
    defaultVersioning: number;
}

export class ResponseUser {
    id: UUID;
    fullName: string;
    shortName: string;
    email: string;
    userAgent: string | null;
    projectLimit: number;
    isProfessional: boolean;
    position: string;
    companyName: string;
    employeeCount: string;
    haveSalesContact: boolean;
    paidTier: boolean;
    isMFAEnabled: boolean;
    mfaRecoveryCodeCount: number;
}

export class UpsertProjectInfo {
    name: string;
    description: string;
    storageLimit: MemorySize;
    bandwidthLimit: MemorySize;
    createdAt: Time;
}

class APIError extends Error {
    constructor(
        public readonly msg: string,
        public readonly responseStatusCode?: number,
    ) {
        super(msg);
    }
}

export class ProjectManagementHttpApiV1 {
    private readonly http: HttpClient = new HttpClient();
    private readonly ROOT_PATH: string = '/public/v1/projects';

    public async createProject(request: UpsertProjectInfo): Promise<Project> {
        const fullPath = `${this.ROOT_PATH}/create`;
        const response = await this.http.post(fullPath, JSON.stringify(request));
        if (response.ok) {
            return response.json().then((body) => body as Project);
        }
        const err = await response.json();
        throw new APIError(err.error, response.status);
    }

    public async updateProject(request: UpsertProjectInfo, id: UUID): Promise<Project> {
        const fullPath = `${this.ROOT_PATH}/update/${id}`;
        const response = await this.http.patch(fullPath, JSON.stringify(request));
        if (response.ok) {
            return response.json().then((body) => body as Project);
        }
        const err = await response.json();
        throw new APIError(err.error, response.status);
    }

    public async deleteProject(id: UUID): Promise<void> {
        const fullPath = `${this.ROOT_PATH}/delete/${id}`;
        const response = await this.http.delete(fullPath);
        if (response.ok) {
            return;
        }
        const err = await response.json();
        throw new APIError(err.error, response.status);
    }

    public async getProjects(): Promise<Project[]> {
        const fullPath = `${this.ROOT_PATH}/`;
        const response = await this.http.get(fullPath);
        if (response.ok) {
            return response.json().then((body) => body as Project[]);
        }
        const err = await response.json();
        throw new APIError(err.error, response.status);
    }

    public async getBucketRollup(projectID: UUID, bucket: string, since: Time, before: Time): Promise<BucketUsageRollup> {
        const u = new URL(`${this.ROOT_PATH}/bucket-rollup`, window.location.href);
        u.searchParams.set('projectID', projectID);
        u.searchParams.set('bucket', bucket);
        u.searchParams.set('since', since);
        u.searchParams.set('before', before);
        const fullPath = u.toString();
        const response = await this.http.get(fullPath);
        if (response.ok) {
            return response.json().then((body) => body as BucketUsageRollup);
        }
        const err = await response.json();
        throw new APIError(err.error, response.status);
    }

    public async getBucketRollups(projectID: UUID, since: Time, before: Time): Promise<BucketUsageRollup[]> {
        const u = new URL(`${this.ROOT_PATH}/bucket-rollups`, window.location.href);
        u.searchParams.set('projectID', projectID);
        u.searchParams.set('since', since);
        u.searchParams.set('before', before);
        const fullPath = u.toString();
        const response = await this.http.get(fullPath);
        if (response.ok) {
            return response.json().then((body) => body as BucketUsageRollup[]);
        }
        const err = await response.json();
        throw new APIError(err.error, response.status);
    }

    public async getAPIKeys(projectID: UUID, search: string, limit: number, page: number, order: number, orderDirection: number): Promise<APIKeyPage> {
        const u = new URL(`${this.ROOT_PATH}/apikeys/${projectID}`, window.location.href);
        u.searchParams.set('search', search);
        u.searchParams.set('limit', limit);
        u.searchParams.set('page', page);
        u.searchParams.set('order', order);
        u.searchParams.set('orderDirection', orderDirection);
        const fullPath = u.toString();
        const response = await this.http.get(fullPath);
        if (response.ok) {
            return response.json().then((body) => body as APIKeyPage);
        }
        const err = await response.json();
        throw new APIError(err.error, response.status);
    }
}

export class APIKeyManagementHttpApiV1 {
    private readonly http: HttpClient = new HttpClient();
    private readonly ROOT_PATH: string = '/public/v1/apikeys';

    public async createAPIKey(request: CreateAPIKeyRequest): Promise<CreateAPIKeyResponse> {
        const fullPath = `${this.ROOT_PATH}/create`;
        const response = await this.http.post(fullPath, JSON.stringify(request));
        if (response.ok) {
            return response.json().then((body) => body as CreateAPIKeyResponse);
        }
        const err = await response.json();
        throw new APIError(err.error, response.status);
    }

    public async deleteAPIKey(id: UUID): Promise<void> {
        const fullPath = `${this.ROOT_PATH}/delete/${id}`;
        const response = await this.http.delete(fullPath);
        if (response.ok) {
            return;
        }
        const err = await response.json();
        throw new APIError(err.error, response.status);
    }
}

export class UserManagementHttpApiV1 {
    private readonly http: HttpClient = new HttpClient();
    private readonly ROOT_PATH: string = '/public/v1/users';

    public async getUser(): Promise<ResponseUser> {
        const fullPath = `${this.ROOT_PATH}/`;
        const response = await this.http.get(fullPath);
        if (response.ok) {
            return response.json().then((body) => body as ResponseUser);
        }
        const err = await response.json();
        throw new APIError(err.error, response.status);
    }
}
