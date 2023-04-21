// AUTOGENERATED BY private/apigen
// DO NOT EDIT.

import { HttpClient } from '@/utils/httpClient';
import { MemorySize, Time, UUID } from '@/types/common';

export class APIKeyInfo {
    id: UUID;
    projectId: UUID;
    userAgent: string;
    name: string;
    createdAt: Time;
}

export class APIKeyPage {
    apiKeys: APIKeyInfo[];
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
    keyInfo: APIKeyInfo;
}

export class Project {
    id: UUID;
    publicId: UUID;
    name: string;
    description: string;
    userAgent: string;
    ownerId: UUID;
    rateLimit: number;
    burstLimit: number;
    maxBuckets: number;
    createdAt: Time;
    memberCount: number;
    storageLimit: MemorySize;
    bandwidthLimit: MemorySize;
    userSpecifiedStorageLimit: MemorySize;
    userSpecifiedBandwidthLimit: MemorySize;
    segmentLimit: number;
}

export class ProjectInfo {
    name: string;
    description: string;
    storageLimit: MemorySize;
    bandwidthLimit: MemorySize;
    createdAt: Time;
}

export class ResponseUser {
    id: UUID;
    fullName: string;
    shortName: string;
    email: string;
    userAgent: string;
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

export class projectsHttpApiV0 {
    private readonly http: HttpClient = new HttpClient();
    private readonly ROOT_PATH: string = '/api/v0/projects';

    public async createProject(request: ProjectInfo): Promise<Project> {
        const path = `${this.ROOT_PATH}/create`;
        const response = await this.http.post(path, JSON.stringify(request));
        if (response.ok) {
            return response.json().then((body) => body as Project);
        }
        const err = await response.json();
        throw new Error(err.error);
    }

    public async updateProject(request: ProjectInfo, id: UUID): Promise<Project> {
        const path = `${this.ROOT_PATH}/update/${id}`;
        const response = await this.http.patch(path, JSON.stringify(request));
        if (response.ok) {
            return response.json().then((body) => body as Project);
        }
        const err = await response.json();
        throw new Error(err.error);
    }

    public async deleteProject(id: UUID): Promise<void> {
        const path = `${this.ROOT_PATH}/delete/${id}`;
        const response = await this.http.delete(path);
        if (response.ok) {
            return;
        }
        const err = await response.json();
        throw new Error(err.error);
    }

    public async getProjects(): Promise<Array<Project>> {
        const path = `${this.ROOT_PATH}/`;
        const response = await this.http.get(path);
        if (response.ok) {
            return response.json().then((body) => body as Array<Project>);
        }
        const err = await response.json();
        throw new Error(err.error);
    }

    public async getBucketRollup(projectID: UUID, bucket: string, since: Time, before: Time): Promise<BucketUsageRollup> {
        const path = `${this.ROOT_PATH}/bucket-rollup?projectID=${projectID}&bucket=${bucket}&since=${since}&before=${before}`;
        const response = await this.http.get(path);
        if (response.ok) {
            return response.json().then((body) => body as BucketUsageRollup);
        }
        const err = await response.json();
        throw new Error(err.error);
    }

    public async getBucketRollups(projectID: UUID, since: Time, before: Time): Promise<Array<BucketUsageRollup>> {
        const path = `${this.ROOT_PATH}/bucket-rollups?projectID=${projectID}&since=${since}&before=${before}`;
        const response = await this.http.get(path);
        if (response.ok) {
            return response.json().then((body) => body as Array<BucketUsageRollup>);
        }
        const err = await response.json();
        throw new Error(err.error);
    }

    public async getAPIKeys(projectID: UUID, search: string, limit: number, page: number, order: number, orderDirection: number): Promise<APIKeyPage> {
        const path = `${this.ROOT_PATH}/apikeys/${projectID}?search=${search}&limit=${limit}&page=${page}&order=${order}&orderDirection=${orderDirection}`;
        const response = await this.http.get(path);
        if (response.ok) {
            return response.json().then((body) => body as APIKeyPage);
        }
        const err = await response.json();
        throw new Error(err.error);
    }
}

export class apikeysHttpApiV0 {
    private readonly http: HttpClient = new HttpClient();
    private readonly ROOT_PATH: string = '/api/v0/apikeys';

    public async createAPIKey(request: CreateAPIKeyRequest): Promise<CreateAPIKeyResponse> {
        const path = `${this.ROOT_PATH}/create`;
        const response = await this.http.post(path, JSON.stringify(request));
        if (response.ok) {
            return response.json().then((body) => body as CreateAPIKeyResponse);
        }
        const err = await response.json();
        throw new Error(err.error);
    }

    public async deleteAPIKey(id: UUID): Promise<void> {
        const path = `${this.ROOT_PATH}/delete/${id}`;
        const response = await this.http.delete(path);
        if (response.ok) {
            return;
        }
        const err = await response.json();
        throw new Error(err.error);
    }
}

export class usersHttpApiV0 {
    private readonly http: HttpClient = new HttpClient();
    private readonly ROOT_PATH: string = '/api/v0/users';

    public async getUser(): Promise<ResponseUser> {
        const path = `${this.ROOT_PATH}/`;
        const response = await this.http.get(path);
        if (response.ok) {
            return response.json().then((body) => body as ResponseUser);
        }
        const err = await response.json();
        throw new Error(err.error);
    }
}
