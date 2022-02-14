import axios from 'axios'
axios.defaults.withCredentials = true
const csrfHeader = {'X-Token': 'abcdefghigklmn'}
const defaultOptions = {method: 'GET'}

export function getHeaders(headers = {}) {
    return {
        Accept: 'application/json',
        'Content-Type': 'application/json',
        ...headers
    }
}

export function getAPIRoot() {
    const {host, protocol} = window.location;
    let baseURL = process.env.VUE_APP_APIURL ? process.env.VUE_APP_APIURL:`${protocol}//${host}`
    if (baseURL.endsWith('/')) {
        baseURL = baseURL.slice(0, -1)
    }
    return baseURL;
}

export async function request(url, options = defaultOptions) {
    const headers = {...options.headers, ...csrfHeader}
    let reqConfig = {url, ...defaultOptions, ...options, headers,
        baseURL:getAPIRoot(),withCredentials: true}
    return axios.request(reqConfig).catch(function (error) {
        console.log(error.response)
        switch (error.response.status) {
            case 401:
                localStorage.clear()
                location.href = "/login.html?next="+location.href
                return
            case 403:
                alert(error.response.data)
                return
            default:
                console.log("请求异常", error.response)
                return error.response
        }
    })
}

export function get(uri, headers, options = {}) {
    return request(uri, {method: 'GET', headers: getHeaders(headers)}, options.stream)
}

export function post(uri, body) {
    return request(uri, {method: 'POST', headers: getHeaders(), data: JSON.stringify(body)})
}

export function getAPI(type, queryParams) {
    return ["", type, queryParams ? `?${new URLSearchParams(queryParams).toString()}` : ''].join('')
}