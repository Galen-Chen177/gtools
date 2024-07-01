export function decodeApiData(resp: string) {
    console.log(resp)
    return JSON.parse(resp);
}

import { Log } from "../../wailsjs/go/log/Log";
export function log(level: string, inlog: string) {
    Log(level, inlog)
}