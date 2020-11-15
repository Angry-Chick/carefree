import React, {useState} from 'react'

const REFRESH_TOKEN_STORAGE_KEY = 'CAREFREE_HOME_REFRESH_TOKEN';

export class CredsProvider {
    private _storage: Storage;
    private _refreshToken?: tokenProps;

    constructor(storage = localStorage) {
        this._storage = storage;
    }

    setRefreshToken(tk: tokenProps) {
        this._refreshToken = tk;
        try {
            this._storage.setItem(REFRESH_TOKEN_STORAGE_KEY, JSON.stringify(tk));
        } catch (err) {
            console.error(`Failed to store refresh token: ${err}`);
        }
    }
}

export const defaultCredsProvider = new CredsProvider();

interface tokenProps {
    Opaque: string,
    Expiry: number,
    Type: string,
}