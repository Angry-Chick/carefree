const REFRESH_TOKEN_STORAGE_KEY = "CAREFREE_HOME_REFRESH_TOKEN";

export class CredsProvider {
  private _storage: Storage;
  private _refreshToken?: string;

  constructor(storage = localStorage) {
    this._storage = storage;
    try {
      const tkStr = this._storage.getItem(REFRESH_TOKEN_STORAGE_KEY);
      if (!!tkStr) {
        this._refreshToken = tkStr;
      }
    } catch (err) {
      console.error(`Failed to retrieve refresh token from storage: ${err}`);
    }
  }

  hasValidRefreshToken(): boolean {
    return !this._isExpired(this._refreshToken);
  }
  getRefreshToken(): string {
    if (!!this._refreshToken) {
      return this._refreshToken;
    }
    return "";
  }
  setRefreshToken(tkStr: string) {
    this._refreshToken = tkStr;
    try {
      this._storage.setItem(REFRESH_TOKEN_STORAGE_KEY, tkStr);
    } catch (err) {
      console.error(`Failed to store refresh token: ${err}`);
    }
  }
  clearRefreshToken() {
    this._refreshToken = undefined;
    try {
      this._storage.removeItem(REFRESH_TOKEN_STORAGE_KEY);
    } catch (err) {
      console.error(`Failed to clear refresh token: ${err}`);
    }
  }
  private _isExpired(tkStr?: string): boolean {
    if (!tkStr) {
      return true;
    }
    const tk = JSON.parse(tkStr);
    if (!tk.expiry) {
      return true;
    }
    if (Date.now() > tk.expiry / 1e6) {
      return true
    }
    return false;
  }
}

export const defaultCredsProvider = new CredsProvider();
