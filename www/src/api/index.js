import {get, post} from './comms'

export function getUserList() {
    const uri = '/api/user/list'
    return get(uri)
}

export function getUserDetail(userid) {
    const uri = '/api/user/'+userid
    return get(uri)
}

export function loginUser(params) {
    const uri = '/api/login'
    return post(uri, params)
}
