export enum ObjectType {
    FILE = 0,
    DIRECTORY = 1,
    CAR_FILE_SHARDS = 2
}

export type CID = string;

export type DID = `did:key:{$Key}`;

export type StorachaUserInfo = {
    account: string;
    did: DID;
}