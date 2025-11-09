import router from '@/router';
import {usePersistedStore} from '@/stores/persisted';
import {useSessionStore} from '@/stores/session';
import axios from 'axios';
import {ElMessage} from 'element-plus';

export const api = axios.create({
  baseURL: 'http://axogc.net:8091',
  timeout: 10000,
})

export function request<T>(method: 'GET' | 'POST' | 'PUT' | 'PATCH' | 'DELETE', url: string, data?: any): Promise<T> {
  return api.request({
    method, url,
    params: method === 'GET' ? data : undefined,
    data: ['POST', 'PUT', 'PATCH'].includes(method) ? data : undefined,
  })
}

api.interceptors.request.use(config => {
  config.withCredentials = true
  const session = useSessionStore()
  const persisted = usePersistedStore()
  config.baseURL = persisted.serverAddr
  if (session.token) {
    config.headers.set("Authorization", `Bearer ${session.token}`)
  }
  return config
}, err => {
  return Promise.reject(err)
})

api.interceptors.response.use(res => {
  if (res.data.message) {
    ElMessage({type: 'info', message: res.data.message})
  }
  const token = res.headers['x-access-token']
  const session = useSessionStore()
  if (token) {
    session.token = token
  }
  return res.data.data
}, async err => {
  const res = err.response
  if (res.status === 401) {
    router.push('/login')
  }
  ElMessage({type: 'error', message: res.data.message})
  return Promise.reject(err)
})
