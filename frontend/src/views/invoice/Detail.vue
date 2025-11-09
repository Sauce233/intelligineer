<script setup lang="ts">
import {request} from '@/utils/axios';
import type {Invoice} from '@/utils/tables';
import {ref} from 'vue';
import {useRoute} from 'vue-router';

const route = useRoute()

const invoice = ref<Invoice | null>(null)

request<Invoice>('GET', `/invoices/${route.params.id}`).then(res => invoice.value = res)
</script>

<template>
  <div class="max-w-5xl mx-auto flex flex-col gap-4 p-4">
    {{invoice}}
  </div>
</template>
