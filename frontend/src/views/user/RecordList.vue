<script setup lang="ts">
import {request} from '@/utils/axios';
import type {Record} from '@/utils/tables';
import {useUrlSearchParams} from '@vueuse/core';
import {ref, watchEffect} from 'vue';
import {formatDate, type List} from '@/utils/utils';
import {useRouteParams} from '@vueuse/router';

const id = useRouteParams('id')
const records = ref<List<Record> | null>(null);

const params = useUrlSearchParams<any>('history');
watchEffect(async () => records.value = await request('GET', `/users/${id.value}/records`, params));
</script>

<template>

  <div class="text-xl">{{$t('userDetail.recordsTitle')}}</div>

  <div v-if="records" class="grid grid-cols-[repeat(auto-fill,minmax(250px,1fr))] gap-4">
    <div v-for="r in records.data" :key="r.id" class="card space-y-2">
      <div>{{r.title}}</div>
      <div class="text-subtle line-clamp-2">{{r.content}}</div>
      <div class="text-subtle">
        {{formatDate(r.createdAt)}} {{r.minutes}}
      </div>
    </div>
  </div>

  <el-pagination
    v-if="records"
    layout="prev, pager, next, total"
    :total="records.total"
    :current-page="Number(params.page) || 1"
    :page-size="Number(params.size) || 12"
    @update:current-page="v => params.page = v"
    @update:page-size="v => params.size = v"
  />

</template>
