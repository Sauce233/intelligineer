import {useColorMode} from "@vueuse/core";
import {defineStore} from "pinia";
import {ref} from "vue";

export const usePersistedStore = defineStore('persisted', () => {
  const colorMode = useColorMode()
  const serverAddr = ref('http://axogc.net:8091')

  return { colorMode, serverAddr }
}, { persist: true })
