import {request} from "@/utils/axios";
import type {User} from "@/utils/tables";
import {defineStore} from "pinia";
import {ref} from "vue";
import MarkdownIt from 'markdown-it';

export const useSessionStore = defineStore('session', () => {

  const token = ref('')
  const user = ref<User | null>(null)
  // watchEffect(async () => user.value = await request('GET', '/my-info'))

  interface Enum {
    id: number
    name: string
    color: string
  }

  const enums = ref<Record<string, Enum[]>>({})
  request<Record<string, Enum[]>>('GET', '/enums').then(res => enums.value = res)

  const md = ref<MarkdownIt>(new MarkdownIt())

  //watchEffect(async () => enums.value = await request('GET', '/enums'))

  return { token, user, enums, md }
})
